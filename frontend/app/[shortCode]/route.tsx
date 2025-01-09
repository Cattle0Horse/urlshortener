import { notFound } from "next/navigation";
import { NextResponse } from "next/server";

// todo: 将notFound()重定向为自定义的页面
export async function GET(
	request: Request,
	{ params }: { params: { shortCode: string } }
) {
	const url = `${process.env.NEXT_PUBLIC_API_URL}/api/url/${params.shortCode}`;
	// try {
	const response = await fetch(url, {
		headers: request.headers,
		redirect: "manual",
	});

	// 处理重定向
	if (response.status === 307) {
		const location = response.headers.get("location");
		if (!location) {
			console.error("重定向响应缺少location头");
			notFound();
		}
		return NextResponse.redirect(location);
	}
	
	// 处理其他非2xx响应（await json前需要检查ok，否则会抛出异常）
	if (!response.ok) {
		const data = await response.json();
		// todo: 更优雅的响应处理
		if (data?.code == 40002) {
			// note: 重定向至not-found页面
			notFound();
		}
		return NextResponse.json(
			{ error: data.message || "请求失败", status: response.status },
			{ status: response.status }
		);
	}

	// 返回正常响应，这是不应该的
	return new NextResponse(response.body, {
		status: response.status,
		headers: response.headers,
	});
	// } catch (error) {
	// 	console.error("请求失败:", error);
	// 	return NextResponse.json({ error: "服务器内部错误" }, { status: 500 });
	// }
}
