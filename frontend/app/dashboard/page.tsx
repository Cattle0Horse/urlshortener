"use client";

import Link from "next/link";
import { Button } from "@/components/ui/button";
import { LinkList } from "@/components/dashboard/link-list";
import { Plus } from "lucide-react";
import { useRouter } from "next/navigation";
import { useAuth } from "@/components/auth/auth-provider";
import { useEffect } from "react";

export default function DashboardPage() {
	const { isAuth } = useAuth();
	const router = useRouter();
	/*
	// 根据React的规则，不能在渲染过程中执行导航操作，因为这会导致状态更新。
	if (isAuth) {
		router.push("/dashboard");
	}
	*/
	useEffect(() => {
		if (!isAuth) {
			router.push("/auth/login");
		}
	}, [router, isAuth]);

	return (
		<div className="container py-8">
			<div className="flex flex-col gap-8">
				<div className="flex items-center justify-between">
					<div>
						<h1 className="text-3xl font-bold tracking-tight">我的短链接</h1>
						<p className="text-muted-foreground">
							查看和管理您创建的所有短链接
						</p>
					</div>
					<Button asChild>
						<Link
							href="/create"
							className="bg-gradient-to-r from-primary to-secondary text-white"
						>
							<Plus className="mr-2 h-4 w-4" />
							创建短链接
						</Link>
					</Button>
				</div>
				<LinkList />
			</div>
		</div>
	);
}
