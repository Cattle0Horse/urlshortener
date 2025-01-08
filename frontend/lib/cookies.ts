import Cookies from "js-cookie";

interface CookieOptions {
	expires?: number | Date;
	path?: string;
	secure?: boolean;
	sameSite?: "strict" | "lax" | "none";
}

export function getCookie(name: string): string {
	return Cookies.get(name) || "";
}

export function setCookie(
	name: string,
	value: string,
	options: CookieOptions = {}
) {
	Cookies.set(name, value, {
		...options,
		path: "/",
		secure: process.env.NODE_ENV === "production",
		sameSite: "lax",
	});
}

export function removeCookie(name: string) {
	Cookies.remove(name);
}
