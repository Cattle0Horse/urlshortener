#!/usr/bin/env nu

# nextjs standalone 打包至 ./build/目录，并移动相关文件，运行只需 node ./build/server.js
# 让本应由CDN处理的public及static文件交由本机处理
def next_build_and_move [] {
	mkdir -v build
	pnpm build
	cp -r public ./build/
	mv -u ./.next/standalone/* ./build/
    rm ./.next/standalone
	cp -r .next/static ./build/.next
}

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

def start_mysql [] {
    docker start urlshortener_mysql
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

def start_redis [] {
    docker start urlshortener_redis
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

def run_env [] {
    docker compose -f docker-compose.env.yaml up
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
