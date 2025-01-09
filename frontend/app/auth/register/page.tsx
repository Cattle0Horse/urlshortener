import { Metadata } from "next";
import Link from "next/link";
import RegisterForm from "@/components/auth/register-form";
import { LinkIcon } from "lucide-react";

export const metadata: Metadata = {
	title: "注册 | URLify",
	description: "创建您的URLify账户",
};

export default function RegisterPage() {
	return (
		<div className="min-h-screen w-full bg-gradient-to-b from-background to-muted relative overflow-hidden">
			{/* 装饰性背景元素 */}
			<div className="absolute inset-0 w-full h-full">
				<div className="absolute top-0 -right-4 w-72 h-72 bg-primary/10 rounded-full filter blur-3xl animate-blob" />
				<div className="absolute -bottom-8 left-20 w-72 h-72 bg-secondary/10 rounded-full filter blur-3xl animate-blob animation-delay-2000" />
				<div className="absolute top-1/2 left-1/2 w-72 h-72 bg-accent/10 rounded-full filter blur-3xl animate-blob animation-delay-4000" />
			</div>

			<div className="container relative">
				{/* 主要内容区域 */}
				<div className="flex min-h-screen items-center justify-center py-6">
					<div className="relative w-full max-w-lg -mt-20">
						{/* Logo和标题区域 */}
						<div className="flex flex-col items-center space-y-4 mb-8">
							<LinkIcon className="h-12 w-12 text-blue-600 dark:text-blue-400" />
							<h1 className="text-4xl font-bold tracking-tight bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
								创建账户
							</h1>
							<p className="text-muted-foreground text-center max-w-sm">
								注册 URLify，体验专业的链接管理服务
							</p>
						</div>

						{/* 注册表单卡片 */}
						<div className="relative">
							<div className="absolute inset-0 bg-gradient-to-r from-secondary/20 to-primary/20 rounded-2xl blur" />
							<div className="relative bg-background/80 backdrop-blur-xl rounded-xl shadow-lg border p-8">
								<RegisterForm />
								<div className="mt-6 text-center">
									<p className="text-sm text-muted-foreground">
										已有账号？{" "}
										<Link
											href="/auth/login"
											className="text-primary hover:text-primary/80 font-medium transition-colors"
										>
											立即登录
										</Link>
									</p>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	);
}
