"use client";
import { useRouter } from "next/navigation";
import { useAuth } from "@/components/auth/auth-provider";
import { useEffect } from "react";

export default function AuthLayout({
	children,
}: Readonly<{
	children: React.ReactNode;
}>) {
	const router = useRouter();
	const { isAuth } = useAuth();

	/*
	// 根据React的规则，不能在渲染过程中执行导航操作，因为这会导致状态更新。
	if (isAuth) {
		router.push("/dashboard");
	}
	*/

	useEffect(() => {
		if (isAuth) {
			router.push("/dashboard");
		}
	}, [isAuth, router]);

	return <>{children}</>;
}
