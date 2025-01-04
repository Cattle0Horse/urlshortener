#!/usr/bin/env nu

def launch_mysql [] {
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
}

def launch_redis [] {
    (
        docker run -d
            --name urlshortener_redis
            -p 6379:6379
            redis:latest
    )
}

def gen_config_struct_tag [] {
    (
        gomodifytags -all
            -add-tags json,yaml,mapstructure
            -w
            -file "config/config.go"
    )
}
