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
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { Icons } from "@/components/ui/icons"
import { Card } from "@/components/ui/card"
import { isValidUrl } from "@/lib/utils"
import { api } from "@/lib/api"
import { getCookie } from "@/lib/cookies"

const formSchema = z.object({
  url: z.string().url({
    message: "请输入有效的URL地址",
  }),
  customSlug: z
    .string()
    .min(3, {
      message: "自定义短链至少需要3个字符",
    })
    .max(20, {
      message: "自定义短链不能超过20个字符",
    })
    .regex(/^[a-zA-Z0-9-_]+$/, {
      message: "只能使用字母、数字、横线和下划线",
    })
    .optional(),
})

export function CreateLinkForm() {
  const router = useRouter()
  const [isLoading, setIsLoading] = React.useState<boolean>(false)
  const [shortUrl, setShortUrl] = React.useState<string>("")

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      url: "",
      customSlug: "",
    },
  })

  async function onSubmit(values: z.infer<typeof formSchema>) {
    setIsLoading(true)

    try {
      const token = getCookie('token')
      const data = await api.links.create({
        url: values.url,
        customSlug: values.customSlug,
      }, token)
      setShortUrl(data.shortUrl)
      toast.success('创建成功')
    } catch (error) {
      toast.error('创建失败，请稍后重试')
    } finally {
      setIsLoading(false)
    }
  }

  const copyToClipboard = async () => {
    try {
      await navigator.clipboard.writeText(shortUrl)
      toast.success("已复制到剪贴板")
    } catch (err) {
      toast.error("复制失败，请手动复制")
    }
  }

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
                  <Input
                    placeholder="https://example.com/very-long-url"
                    {...field}
                    disabled={isLoading}
                  />
                </FormControl>
                <FormDescription>
                  输入您想要缩短的完整URL地址
                </FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="customSlug"
            render={({ field }) => (
              <FormItem>
                <FormLabel>自定义短链（可选）</FormLabel>
                <FormControl>
                  <Input
                    placeholder="my-custom-url"
                    {...field}
                    disabled={isLoading}
                  />
                </FormControl>
                <FormDescription>
                  自定义您的短链接后缀，如果不填写将自动生成
                </FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
          <Button type="submit" disabled={isLoading}>
            {isLoading && (
              <Icons.spinner className="mr-2 h-4 w-4 animate-spin" />
            )}
            创建短链接
          </Button>
        </form>
      </Form>

      {shortUrl && (
        <Card className="p-6">
          <div className="flex items-center justify-between space-x-4">
            <div className="space-y-1">
              <h3 className="font-semibold">您的短链接已生成</h3>
              <p className="text-sm text-muted-foreground break-all">
                {shortUrl}
              </p>
            </div>
            <Button
              variant="secondary"
              size="sm"
              onClick={copyToClipboard}
              className="shrink-0"
            >
              复制链接
            </Button>
          </div>
        </Card>
      )}
    </div>
  )
} 