module.exports = {
	assetPrefix: "", // 确保没有多余的前缀
	basePath: "", // 确保没有多余的基础路径
	reactStrictMode: true,
	swcMinify: true,
	output: "standalone",
	async headers() {
		return [
			{
				source: "/:path*",
				headers: [
					{
						key: "Access-Control-Allow-Origin",
						value: process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080",
					},
					{
						key: "Access-Control-Allow-Credentials",
						value: "true",
					},
					{
						key: "Access-Control-Allow-Methods",
						value: "GET, POST, PUT, DELETE, OPTIONS",
					},
					{
						key: "Access-Control-Allow-Headers",
						value: "Content-Type, Authorization",
					},
				],
			},
		];
	},
};
