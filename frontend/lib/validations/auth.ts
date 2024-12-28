import * as z from "zod";

const passwordSchema = z
  .string()
  .trim()
  .min(8, "密码至少8位")
  .max(32, "密码最多32位")
  .regex(
    /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,}$/,
    "密码必须包含大小写字母和数字"
  );

export const registerSchema = z
  .object({
    email: z.string().trim().email("请输入有效的邮箱地址"),
    verificationCode: z
      .string()
      .trim()
      .min(4, "验证码至少4位")
      .max(6, "验证码最多6位"),
    password: passwordSchema,
    confirmPassword: z.string().trim(),
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: "两次输入的密码不一致",
    path: ["confirmPassword"],
  });

export const loginSchema = z.object({
  email: z.string().trim().email("请输入有效的邮箱地址"),
  password: z.string().trim().min(1, "请输入密码"),
});

export const forgotPasswordSchema = z
  .object({
    email: z.string().trim().email("请输入有效的邮箱地址"),
    verificationCode: z
      .string()
      .trim()
      .min(4, "验证码至少4位")
      .max(6, "验证码最多6位"),
    password: passwordSchema,
    confirmPassword: z.string().trim(),
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: "两次输入的密码不一致",
    path: ["confirmPassword"],
  });
