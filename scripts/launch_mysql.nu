(
	docker run -d	
		--name urlshortener_mysql
		-p 3306:3306
		-e MYSQL_ROOT_PASSWORD=123456
		-e MYSQL_DATABASE=urlshortener
		-e MYSQL_USER=urlshortener
		-e MYSQL_PASSWORD=123456
		mysql:latest
)