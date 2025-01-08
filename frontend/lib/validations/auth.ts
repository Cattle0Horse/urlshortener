import * as z from "zod";

export const loginSchema = z.object({
	email: z.string().email({
		message: "请输入有效的邮箱地址",
	}),
	password: z
		.string()
		.min(8, {
			message: "密码长度需在8-20个字符之间",
		})
		.max(20, {
			message: "密码长度需在8-20个字符之间",
		}),
});

export const registerSchema = z
	.object({
		email: z.string().email({
			message: "请输入有效的邮箱地址",
		}),
		password: z
			.string()
			.min(8, {
				message: "密码长度需在8-20个字符之间",
			})
			.max(20, {
				message: "密码长度需在8-20个字符之间",
			}),
		confirmPassword: z.string(),
	})
	.refine((data) => data.password === data.confirmPassword, {
		message: "两次输入的密码不一致",
		path: ["confirmPassword"],
	});

export type LoginInput = z.infer<typeof loginSchema>;
export type RegisterInput = z.infer<typeof registerSchema>;
