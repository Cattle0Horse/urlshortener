(
	docker run -d
		--name urlshortener_redis
		-p 6379:6379
		redis:latest
)