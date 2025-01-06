package config

import "time"

type Url struct {
	DefaultDuration time.Duration `yaml:"default_duration" mapstructure:"default_duration"`
	// BloomFilterSize is the size of bloom filter
	BloomFilterSize uint64 `validate:"required,gt=0" yaml:"bloom_filter_size" mapstructure:"bloom_filter_size"`
	// BloomFilterFalsePositiveRate is the false positive rate of bloom filter
	BloomFilterFalsePositiveRate float64 `validate:"required,gt=0,lt=1" yaml:"bloom_filter_false_positive_rate" mapstructure:"bloom_filter_false_positive_rate"`
}
