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