import { z } from "zod";

export const formSchema = z.object({
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
