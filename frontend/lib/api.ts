import { API_URL } from "./config";
import { removeCookie } from "./cookies";

interface RequestOptions extends RequestInit {
	token?: string;
	params?: Record<string, string | number>;
}

interface APIResponse<T = any> {
	code: number;
	data: T;
	msg: string;
}

async function fetchAPI<T>(
	endpoint: string,
	options: RequestOptions = {}
): Promise<T> {
	const { token, params, ...restOptions } = options;
	const headers = {
		"Content-Type": "application/json",
		Accept: "application/json",
		...(token && { Authorization: `Bearer ${token}` }),
		...options.headers,
	};

	const url = new URL(`${API_URL}${endpoint}`);
	if (params) {
		Object.entries(params).forEach(([key, value]) => {
			url.searchParams.append(key, String(value));
		});
	}

	const response = await fetch(url.toString(), {
		...restOptions,
		headers,
		credentials: "include",
		mode: "cors",
	});

	const result = (await response.json()) as APIResponse<T>;

	if (result.code !== 200) {
		// 根据错误码显示不同的错误信息
		switch (result.code) {
			case 40001:
				throw new Error(result.msg || "无效的请求");
			case 40002:
				throw new Error(result.msg || "用户不存在");
			case 40004:
				throw new Error(result.msg || "密码错误");
			case 40005:
				// 清除无效的token
				removeCookie("token");
				throw new Error(result.msg || "无效的token");
			case 40007:
				// 清除未授权的token
				removeCookie("token");
				throw new Error(result.msg || "未授权");
			case 50002:
				throw new Error(result.msg || "数据库错误");
			default:
				throw new Error(result.msg || "请求失败");
		}
	}

	return result.data;
}

interface LoginResponse {
	token: string;
	user_id: number;
}

interface ListResponse {
	urls: {
		id: number;
		short_code: string;
		original_url: string;
		expiry_time: string;
		created_at: string;
	}[];
	total: number;
}

interface CreateResponse {
	short_code: string;
}

export const api = {
	auth: {
		register: (email: string, password: string) =>
			fetchAPI("/api/auth/register", {
				method: "POST",
				body: JSON.stringify({ email, password }),
			}),

		login: (email: string, password: string): Promise<LoginResponse> =>
			fetchAPI("/api/auth/login", {
				method: "POST",
				body: JSON.stringify({ email, password }),
			}),
	},

	links: {
		create: (
			data: { url: string; duration: number },
			token: string
		): Promise<CreateResponse> =>
			fetchAPI("/api/url", {
				method: "POST",
				body: JSON.stringify({
					original_url: data.url,
					duration: data.duration,
				}),
				token,
			}),

		list: (token: string, page = 1, size = 10): Promise<ListResponse> =>
			fetchAPI("/api/urls", {
				headers: {
					Authorization: `Bearer ${token}`,
				},
				params: { page, size },
			}),

		update: (id: string, duration: number, token: string) =>
			fetchAPI(`/api/url/${id}`, {
				method: "PATCH",
				body: JSON.stringify({ duration }),
				token,
			}),

		delete: (id: string, token: string) =>
			fetchAPI(`/api/url/${id}`, {
				method: "DELETE",
				token,
			}),

		resolve: (shortCode: string): Promise<{ original_url: string }> => {
			return fetchAPI(`/api/url/${shortCode}`);
		},
	},
};
