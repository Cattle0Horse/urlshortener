export const assetPrefix = "";
export const basePath = "";
export const reactStrictMode = true;
export const swcMinify = true;
export const output = "standalone";
export async function headers() {
	return [
		{
			source: "/:path*",
			headers: [
				{
					key: "Access-Control-Allow-Origin",
					value: "*",
				},
			],
		},
	];
}
