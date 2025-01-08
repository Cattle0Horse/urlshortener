"use client";

import { createContext, useContext, useEffect, useState } from "react";
import { getCookie } from "@/lib/cookies";
import { api } from "@/lib/api"
import { useRouter } from "next/navigation";


interface AuthContextType {
	isAuthenticated: boolean;
	isLoading: boolean;
	user: {
		id: string;
		email: string;
	} | null;
}

const AuthContext = createContext<AuthContextType>({
	isAuthenticated: false,
	isLoading: true,
	user: null,
});

export function AuthProvider({ children }: { children: React.ReactNode }) {
	const [authState, setAuthState] = useState<AuthContextType>({
		isAuthenticated: false,
		isLoading: true,
		user: null,
	});
	const router = useRouter();

	useEffect(() => {
		async function checkAuth() {
			const token = getCookie("token");
			const user_id = getCookie("user_id");
			const email = getCookie("email");

			// 严格检查 cookie 值
			if (
				typeof token !== "string" ||
				typeof user_id !== "string" ||
				typeof email !== "string" ||
				token.trim() === "" ||
				user_id.trim() === "" ||
				email.trim() === ""
			) {
				setAuthState({
					isAuthenticated: false,
					isLoading: false,
					user: null,
				});
				// 清除无效的 cookie
				document.cookie =
					"token=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;";
				document.cookie =
					"user_id=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;";
				document.cookie =
					"email=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;";
				return;
			}

			try {
				setAuthState({
					isAuthenticated: true,
					isLoading: false,
					user: {
						id: user_id,
						email: email,
					},
				});
			} catch (error) {
				// 清除无效的 cookie
				document.cookie =
					"token=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;";
				document.cookie =
					"user_id=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;";
				document.cookie =
					"email=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;";

				setAuthState({
					isAuthenticated: false,
					isLoading: false,
					user: null,
				});
				router.push("/login");
			}
		}

		checkAuth();
	}, [router, setAuthState]);

	return (
		<AuthContext.Provider value={authState}>{children}</AuthContext.Provider>
	);
}

export function useAuth() {
	return useContext(AuthContext);
}
