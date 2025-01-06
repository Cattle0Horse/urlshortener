#!/usr/bin/env nu

def launch_mysql [] {
    (
        docker run
            -d	
            --name urlshortener_mysql
            -p 3306:3306
            -e MYSQL_ROOT_PASSWORD=123456
            -e MYSQL_DATABASE=urlshortener
            mysql:latest
    )
}

def launch_redis [] {
    (
        docker run
            -d
            --name urlshortener_redis
            -p 6379:6379
            redis/redis-stack-server:latest
    )
}

def gen [] {
    go run cmd/gen/gen.go
}

def run [] {
    go run main.go
}

def enter_mysql [] {
    docker exec -it urlshortener_mysql mysql -u root -p
}

def enter_redis [] {
    docker exec -it urlshortener_redis /bin/sh
}

def gen_config_struct_tag [] {
    let files = ["config" "cache", "jwt", "mysql", "server", "tddl", "url"]
    $files | each { |name| (
        gomodifytags -all -w
            -add-tags yaml,mapstructure
            -transform snakecase
            -file $"config/($name).go"
    )}
}
