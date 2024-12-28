"use client";

import { useState } from "react";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import * as z from "zod";
import { Copy } from "lucide-react";

import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";

import { toast } from "sonner";
import { useAuth } from "@/components/context";
import { useRouter } from "next/navigation";

const formSchema = z.object({
  originalUrl: z.string().url({
    message: "请输入一个有效的URL",
  }),
  customCode: z.string().optional(),
  duration: z.string().optional(),
});

export default function Home() {
  const { token, userID, isAuth } = useAuth();
  const [shortUrl, setShortUrl] = useState<string>("");
  const router = useRouter();

  if (!isAuth) {
    router.push("/auth/login");
  }

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      originalUrl: "",
      customCode: "",
    },
  });

  async function onSubmit(values: z.infer<typeof formSchema>) {
    try {
      const response = await fetch("http://localhost:8080/api/url", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          original_url: values.originalUrl,
          custom_code: values.customCode || undefined,
          duration: values.duration ? Number(values.duration) : undefined,
          user_id: userID,
        }),
      });

      if (!response.ok) {
        throw new Error("创建失败");
      }

      const data = await response.json();
      setShortUrl(data.short_url);
      toast.success("短链接生成完成");
    } catch {
      toast.error("出现错误,请重试");
    }
  }

  return (
    <div className="min-h-screen min-w-full bg-background p-8 flex items-center justify-center">
      <div className="container max-w-md items-center justify-center mx-auto ">
        <div className="border-2 px-8 py-4 rounded-lg shadow-lg flex flex-col items-center justify-center">
          <h1 className="scroll-m-20 text-xl text-center font-extrabold tracking-tight lg:text-3xl mb-8">
            短链接生成器
          </h1>

          <Form {...form}>
            <form
              onSubmit={form.handleSubmit(onSubmit)}
              className="space-y-4 w-full"
            >
              <FormField
                control={form.control}
                name="originalUrl"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>长链接</FormLabel>
                    <FormControl>
                      <Input placeholder="https://example.com" {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <div className="flex w-full space-x-10">
                <FormField
                  control={form.control}
                  name="customCode"
                  render={({ field }) => (
                    <FormItem className="w-2/3">
                      <FormLabel>自定义别名(可选)</FormLabel>
                      <FormControl>
                        <Input placeholder="custom-code" {...field} />
                      </FormControl>
                      <FormDescription>
                        输入一个4-10个字符的自定义别名
                      </FormDescription>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <FormField
                  control={form.control}
                  name="duration"
                  render={({ field }) => (
                    <FormItem className="w-1/3">
                      <FormLabel>有效时长(可选)</FormLabel>
                      <FormControl>
                        <Input placeholder="10" type="number" {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>

              <div className="flex pt-4">
                <Button type="submit" className="mx-auto">
                  生成短链接
                </Button>
              </div>
            </form>
          </Form>
        </div>
        {shortUrl && (
          <div className="mt-8 p-4 border rounded-lg">
            <h2 className="text-lg font-semibold mb-2">Your Short URL:</h2>
            <div className="flex items-center gap-2">
              <a
                href={shortUrl}
                target="_blank"
                className="hover:underline text-sky-600"
              >
                <h2>{shortUrl}</h2>
              </a>

              <Copy
                onClick={() => {
                  navigator.clipboard.writeText(shortUrl);
                  toast.success("复制到粘贴板");
                }}
                size={15}
                className="hover:cursor-pointer hover:shadow-lg hover:scale-105"
              />
            </div>
          </div>
        )}
      </div>
    </div>
  );
}
