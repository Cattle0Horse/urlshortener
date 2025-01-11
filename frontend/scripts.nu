# nextjs standalone 打包至 ./build/目录，并移动相关文件，运行只需 node ./build/server.js
# 让本应由CDN处理的public及static文件交由本机处理
# todo: 使用覆盖逻辑优化
def next_build_and_move [] {
	rm -rf build
	mkdir -v build
	pnpm build
	cp -ur public ./build/
	cp -ur ./.next/standalone/* ./build/
	cp -ur .next/static ./build/.next
}

def --env next_env [] {
	$env.HOSTNAME = "cattle_horse.com"
	$env.PORT = "80"
	echo "设置成功,请关闭代理"
}

def frontend_build [] {
  docker build -f ./Dockerfile -t cattlehorse/urlshortener-frontend:latest .
}