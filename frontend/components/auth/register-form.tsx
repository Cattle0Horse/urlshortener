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
import { api } from "@/lib/api";

const formSchema = z
	.object({
		email: z.string().email({ message: "请输入有效的邮箱地址" }),
		password: z.string().min(8, { message: "密码至少需要8个字符" }),
		confirmPassword: z.string(),
	})
	.refine((data) => data.password === data.confirmPassword, {
		message: "两次输入的密码不一致",
		path: ["confirmPassword"],
	});

export default function RegisterForm() {
	const router = useRouter();
	const [isLoading, setIsLoading] = React.useState<boolean>(false);

	const form = useForm<z.infer<typeof formSchema>>({
		resolver: zodResolver(formSchema),
		defaultValues: {
			email: "",
			password: "",
			confirmPassword: "",
		},
	});

	async function onSubmit(values: z.infer<typeof formSchema>) {
		setIsLoading(true);

		try {
			await api.auth.register(values.email, values.password);
			// 显示成功消息
			toast.success("注册成功", { description: "正在跳转到登录页面..." });
			router.push("/auth/login");
		} catch (error) {
			// 显示详细的错误信息
			if (error instanceof Error) {
				toast.error("注册失败", { description: error.message || "请稍后重试" });
			} else {
				toast.error("注册失败", {
					description: "该邮箱可能已被注册，请尝试直接登录",
				});
			}
		} finally {
			setIsLoading(false);
		}
	}

	return (
		<Form {...form}>
			<form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
				<FormField
					control={form.control}
					name="email"
					render={({ field }) => (
						<FormItem>
							<FormLabel>邮箱</FormLabel>
							<FormControl>
								<Input
									type="email"
									placeholder="hello@example.com"
									autoComplete="email"
									{...field}
									disabled={isLoading}
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
								<Input
									type="password"
									placeholder="输入密码"
									autoComplete="new-password"
									{...field}
									disabled={isLoading}
								/>
							</FormControl>
							<FormMessage />
						</FormItem>
					)}
				/>
				<FormField
					control={form.control}
					name="confirmPassword"
					render={({ field }) => (
						<FormItem>
							<FormLabel>确认密码</FormLabel>
							<FormControl>
								<Input
									type="password"
									placeholder="再次输入密码"
									autoComplete="new-password"
									{...field}
									disabled={isLoading}
								/>
							</FormControl>
							<FormMessage />
						</FormItem>
					)}
				/>
				<Button
					className="w-full bg-gradient-to-r from-blue-600 to-cyan-500 hover:from-blue-700 hover:to-cyan-600 text-white"
					type="submit"
					disabled={isLoading}
				>
					{isLoading && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
					注册
				</Button>
			</form>
		</Form>
	);
}
