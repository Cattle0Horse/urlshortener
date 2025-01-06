import * as z from "zod";

export const loginSchema = z.object({
  email: z.string().email({
    message: "请输入有效的邮箱地址",
  }),
  password: z.string().min(6, {
    message: "密码至少需要6个字符",
  }),
});

export const registerSchema = z.object({
  email: z.string().email({
    message: "请输入有效的邮箱地址",
  }),
  password: z.string().min(6, {
    message: "密码至少需要6个字符",
  }),
  confirmPassword: z.string(),
}).refine((data) => data.password === data.confirmPassword, {
  message: "两次输入的密码不一致",
  path: ["confirmPassword"],
});

export const forgotPasswordSchema = z.object({
  email: z.string().email({
    message: "请输入有效的邮箱地址",
  }),
  verificationCode: z.string().min(4, {
    message: "验证码至少需要4个字符",
  }),
  password: z.string().min(8, {
    message: "密码至少需要8个字符",
  }),
  confirmPassword: z.string(),
}).refine((data) => data.password === data.confirmPassword, {
  message: "两次输入的密码不一致",
  path: ["confirmPassword"],
})

export type LoginInput = z.infer<typeof loginSchema>
export type RegisterInput = z.infer<typeof registerSchema>
