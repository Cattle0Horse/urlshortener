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
	FormDescription,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Slider } from "@/components/ui/slider";
import { Loader2, Clock, ExternalLink, LayoutDashboard } from "lucide-react";
import { Card } from "@/components/ui/card";
import { api } from "@/lib/api";
import { getCookie } from "@/lib/cookies";

const formSchema = z.object({
	url: z
		.string()
		.url({
			message: "请输入有效的URL地址",
		})
		.max(255, {
			message: "URL长度不能超过255个字符",
		}),
	duration: z
		.number()
		.min(1, { message: "到期时长最少1小时" })
		.max(168, { message: "到期时长最多168小时" })
		.default(24),
});

export function CreateLinkForm() {
	const router = useRouter();
	const [isLoading, setIsLoading] = React.useState<boolean>(false);
	const [shortUrl, setShortUrl] = React.useState<string>("");

	const form = useForm<z.infer<typeof formSchema>>({
		resolver: zodResolver(formSchema),
		defaultValues: {
			url: "",
			duration: 24,
		},
	});

	async function onSubmit(values: z.infer<typeof formSchema>) {
		try {
			setIsLoading(true);
			const token = getCookie("token");
			if (!token) {
				toast.error("登录已过期", {
					description: "请重新登录后再试",
				});
				router.push("/login");
				return;
			}

			const response = await api.links.create(
				{
					url: values.url,
					duration: values.duration,
				},
				token
			);

			const newShortUrl = `${window.location.origin}/${response.short_code}`;
			setShortUrl(newShortUrl);

			toast.success("短链接创建成功！", {
				description: newShortUrl,
				action: {
					label: "前往仪表盘",
					onClick: () => router.push("/dashboard"),
				},
			});

			// 重置表单
			form.reset({
				url: "",
				duration: 24,
			});
		} catch (error) {
			console.error("Failed to create link:", error);
			if (error instanceof Error) {
				toast.error("创建失败", {
					description: error.message,
				});
			} else {
				toast.error("创建失败", {
					description: "请稍后重试",
				});
			}
		} finally {
			setIsLoading(false);
		}
	}

	const copyToClipboard = async () => {
		try {
			await navigator.clipboard.writeText(shortUrl);
			toast.success("已复制到剪贴板", {
				description: shortUrl,
				action: {
					label: "访问链接",
					onClick: () => window.open(shortUrl, "_blank", "noopener,noreferrer"),
				},
			});
		} catch (err) {
			console.error("Failed to copy:", err);
			toast.error("复制失败", {
				description: "请手动复制以下链接：" + shortUrl,
			});
		}
	};

	return (
		<div className="grid gap-8">
			<Form {...form}>
				<form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
					<FormField
						control={form.control}
						name="url"
						render={({ field }) => (
							<FormItem>
								<FormLabel>原始链接</FormLabel>
								<FormControl>
									<div className="relative">
										<Input
											placeholder="https://example.com/very-long-url"
											className="pr-4 transition-all duration-200 border-input focus:border-primary"
											{...field}
											disabled={isLoading}
										/>
										<div className="absolute right-3 top-1/2 -translate-y-1/2">
											<div
												className={`h-2 w-2 rounded-full transition-colors ${
													field.value && !form.formState.errors.url
														? "bg-green-500"
														: "bg-gray-300"
												}`}
											/>
										</div>
									</div>
								</FormControl>
								<FormDescription>输入您想要缩短的完整URL地址</FormDescription>
								<FormMessage />
							</FormItem>
						)}
					/>
					<FormField
						control={form.control}
						name="duration"
						render={({ field }) => (
							<FormItem className="space-y-4">
								<FormLabel className="flex items-center gap-2 text-base">
									<Clock className="h-4 w-4" />
									<span>到期时长</span>
								</FormLabel>
								<FormControl>
									<div className="px-1 space-y-4">
										<Slider
											min={1}
											max={168}
											step={1}
											value={[field.value]}
											onValueChange={([value]) => field.onChange(value)}
											disabled={isLoading}
										/>
										<div className="flex items-center justify-between text-sm">
											<span className="text-muted-foreground">1小时</span>
											<span className="font-medium text-base text-primary">
												{field.value}小时
											</span>
											<span className="text-muted-foreground">168小时</span>
										</div>
									</div>
								</FormControl>
								<FormDescription className="text-xs text-muted-foreground">
									设置短链接的有效期，默认24小时（1周 = 168小时）
								</FormDescription>
								<FormMessage />
							</FormItem>
						)}
					/>
					<Button
						type="submit"
						disabled={isLoading}
						className="w-full bg-gradient-to-r from-primary to-secondary text-white"
					>
						{isLoading && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
						创建短链接
					</Button>
				</form>
			</Form>

			{shortUrl && (
				<Card className="p-6">
					<h3 className="text-lg font-semibold mb-4">创建成功！</h3>
					<div className="space-y-4">
						<div className="flex items-center justify-between">
							<span className="text-sm">短链接：</span>
							<span className="text-sm font-medium">{shortUrl}</span>
						</div>
						<div className="flex gap-4">
							<Button
								variant="outline"
								className="w-full"
								onClick={copyToClipboard}
							>
								复制链接
							</Button>
							<Button
								variant="outline"
								className="w-full"
								onClick={() => window.open(shortUrl, "_blank")}
							>
								<ExternalLink className="mr-2 h-4 w-4" />
								访问链接
							</Button>
							<Button
								className="w-full"
								onClick={() => router.push("/dashboard")}
							>
								<LayoutDashboard className="mr-2 h-4 w-4" />
								前往仪表盘
							</Button>
						</div>
					</div>
				</Card>
			)}
		</div>
	);
}
