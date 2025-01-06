"use client"

import * as React from "react"
import { useRouter } from "next/navigation"
import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import * as z from "zod"
import { toast } from "sonner"
import { Button } from "@/components/ui/button"
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { Icons } from "@/components/ui/icons"
import { loginSchema } from "@/lib/validations/auth"
import { api } from "@/lib/api"
import { setCookie } from "@/lib/cookies"

export function LoginForm() {
  const router = useRouter()
  const [isLoading, setIsLoading] = React.useState<boolean>(false)

  const form = useForm<z.infer<typeof loginSchema>>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  })

  async function onSubmit(values: z.infer<typeof loginSchema>) {
    setIsLoading(true)

    try {
      const data = await api.auth.login(values.email, values.password)
      
      // 存储 token 和用户信息
      setCookie('token', data.access_token)
      setCookie('email', values.email)
      setCookie('user_id', String(data.user_id))

      // 显示成功消息
      toast.success("登录成功", {
        description: "欢迎回来！正在跳转到仪表板...",
        duration: 2000,  // 显示2秒
      })
      
      // 延迟跳转以显示成功消息
      setTimeout(() => {
        router.push("/dashboard")
        router.refresh()
      }, 1000)
    } catch (error) {
      // 显示错误消息
      toast.error("登录失败", {
        description: "邮箱或密码错误，请重试",
      })
    } finally {
      setIsLoading(false)
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
                  placeholder="输入您的密码"
                  autoComplete="current-password"
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
          {isLoading && (
            <Icons.spinner className="mr-2 h-4 w-4 animate-spin" />
          )}
          登录
        </Button>
      </form>
    </Form>
  )
}
