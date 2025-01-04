package otel

import (
	"context"
	"fmt"
	"time"

	"github.com/Cattle0Horse/url-shortener/config"
	"github.com/Cattle0Horse/url-shortener/tools"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	oteltrace "go.opentelemetry.io/otel/sdk/trace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
)

var (
	meterProvider *sdkmetric.MeterProvider
	meter         metric.Meter
)

// List of supported exporters
// https://opentelemetry.io/docs/instrumentation/go/exporters/

// OTLP Trace Exporter
func newOTLPExporter(ctx context.Context) (oteltrace.SpanExporter, error) {
	// Change default HTTPS -> HTTP
	insecureOpt := otlptracehttp.WithInsecure()

	// Update default OTLP reciver endpoint
	endPoint := fmt.Sprintf("%s:%s", config.Get().OTel.AgentHost, config.Get().OTel.AgentPort)
	endpointOpt := otlptracehttp.WithEndpoint(endPoint)
	return otlptracehttp.New(ctx, insecureOpt, endpointOpt)
}

// OTLP Metric Exporter
func newOTLPMetricExporter(ctx context.Context) (sdkmetric.Exporter, error) {
	insecureOpt := otlpmetrichttp.WithInsecure()
	endPoint := fmt.Sprintf("%s:%s", config.Get().OTel.AgentHost, config.Get().OTel.AgentPort)
	endpointOpt := otlpmetrichttp.WithEndpoint(endPoint)
	return otlpmetrichttp.New(ctx, insecureOpt, endpointOpt)
}

func Init() {
	ctx := context.Background()

	// Initialize resource
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(config.Get().OTel.ServiceName),
		),
	)
	tools.PanicOnErr(err)

	// Initialize tracing
	traceProvider := sdktrace.NewTracerProvider(sdktrace.WithResource(r))
	otel.SetTracerProvider(traceProvider)
	traceExp, err := newOTLPExporter(ctx)
	tools.PanicOnErr(err)

	bsp := sdktrace.NewBatchSpanProcessor(traceExp)
	traceProvider.RegisterSpanProcessor(bsp)

	// Initialize metrics
	metricExp, err := newOTLPMetricExporter(ctx)
	if err != nil {
		meter = noop.NewMeterProvider().Meter("noop")
		return
	}

	meterProvider = sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(r),
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(metricExp)),
	)
	otel.SetMeterProvider(meterProvider)
	meter = meterProvider.Meter("url-shortener")
}

// RecordRequestMetrics 记录请求相关指标
func RecordRequestMetrics(method, path string, statusCode int, duration time.Duration) {
	if meter == nil {
		return
	}

	// 请求计数器
	requestCounter, _ := meter.Int64Counter(
		"http_requests_total",
		metric.WithDescription("Total number of HTTP requests"),
	)
	requestCounter.Add(context.Background(), 1,
		metric.WithAttributes(
			semconv.HTTPMethod(method),
			semconv.HTTPRoute(path),
			semconv.HTTPStatusCode(statusCode),
		),
	)

	// 请求耗时直方图
	requestDuration, _ := meter.Float64Histogram(
		"http_request_duration_seconds",
		metric.WithDescription("HTTP request duration in seconds"),
	)
	requestDuration.Record(context.Background(), duration.Seconds(),
		metric.WithAttributes(
			semconv.HTTPMethod(method),
			semconv.HTTPRoute(path),
			semconv.HTTPStatusCode(statusCode),
		),
	)

	// 错误计数器
	if statusCode >= 400 {
		errorCounter, _ := meter.Int64Counter(
			"http_errors_total",
			metric.WithDescription("Total number of HTTP errors"),
		)
		errorCounter.Add(context.Background(), 1,
			metric.WithAttributes(
				semconv.HTTPMethod(method),
				semconv.HTTPRoute(path),
				semconv.HTTPStatusCode(statusCode),
			),
		)
	}
}
