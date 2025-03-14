"use client";

import * as React from "react";
import { useRouter } from "next/navigation";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import * as z from "zod";
import { toast } from "sonner";
import { Button } from "@/components/ui/button";
import {
	Form,
	FormControl,
	FormField,
	FormItem,
	FormLabel,
	FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Loader2 } from "lucide-react";
import { loginSchema } from "@/lib/validations/auth";
import { api } from "@/lib/api";
import { useAuth } from "./auth-provider";

type LoginFormData = z.infer<typeof loginSchema>;

export function LoginForm() {
	const router = useRouter();
	const [isLoading, setIsLoading] = React.useState<boolean>(false);
	const { setAuth } = useAuth();

	const form = useForm<LoginFormData>({
		resolver: zodResolver(loginSchema),
		defaultValues: {
			email: "",
			password: "",
		},
	});

	async function onSubmit(values: LoginFormData) {
		try {
			setIsLoading(true);
			const response = await api.auth.login(values.email, values.password);
			console.log(response);

			setAuth(response?.access_token, response?.email, response?.user_id);

			// 显示成功消息并跳转
			toast.success("登录成功", {
				description: "欢迎回来！即将跳转到仪表盘...",
			});
			router.push("/dashboard");
		} catch (error) {
			console.error("Login failed:", error);
			// 显示错误消息
			if (error instanceof Error) {
				toast.error("登录失败", {
					description: error.message,
				});
			} else {
				toast.error("登录失败", {
					description: "邮箱或密码错误",
				});
			}
		} finally {
			setIsLoading(false);
		}
	}

	return (
		<Form {...form}>
			<form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
				<FormField
					control={form.control}
					name="email"
					render={({ field }) => (
						<FormItem>
							<FormLabel>邮箱</FormLabel>
							<FormControl>
								<Input
									type="email"
									placeholder="example@example.com"
									{...field}
								/>
							</FormControl>
							<FormMessage />
						</FormItem>
					)}
				/>
				<FormField
					control={form.control}
					name="password"
					render={({ field }) => (
						<FormItem>
							<FormLabel>密码</FormLabel>
							<FormControl>
								<Input type="password" placeholder="••••••••" {...field} />
							</FormControl>
							<FormMessage />
						</FormItem>
					)}
				/>
				<Button type="submit" className="w-full" disabled={isLoading}>
					{isLoading ? (
						<>
							<Loader2 className="mr-2 h-4 w-4 animate-spin" />
							登录中...
						</>
					) : (
						"登录"
					)}
				</Button>
			</form>
		</Form>
	);
}
