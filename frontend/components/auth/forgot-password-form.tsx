"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import Link from "next/link";
import { toast } from "sonner";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { forgotPasswordSchema } from "@/lib/validations/auth";
import { base_url } from "../env";
import { useAuth } from "../context";
import { Loading } from "../loading";

type FormData = z.infer<typeof forgotPasswordSchema>;

export function ForgotPasswordForm() {
  const router = useRouter();
  const [isSendingCode, setIsSendingCode] = useState(false);
  const [countdown, setCountdown] = useState(0);
  const { setAuth } = useAuth();

  const form = useForm<FormData>({
    resolver: zodResolver(forgotPasswordSchema),
    defaultValues: {
      email: "",
      verificationCode: "",
      password: "",
      confirmPassword: "",
    },
  });

  async function sendVerificationCode(email: string) {
    if (!email) {
      toast.error("请先输入邮箱");
      return;
    }

    setIsSendingCode(true);
    try {
      const response = await fetch(`${base_url}/api/auth/register/${email}`, {
        method: "GET",
      });

      const payload = await response.json();

      if (!response.ok) {
        toast.error(payload?.message);
        return;
      }

      toast.success("验证码已发送到您的邮箱");
      setCountdown(60);
      const timer = setInterval(() => {
        setCountdown((prev) => {
          if (prev <= 1) {
            clearInterval(timer);
            return 0;
          }
          return prev - 1;
        });
      }, 1000);
    } catch {
      toast.error("发送验证码失败，请稍后重试");
    } finally {
      setIsSendingCode(false);
    }
  }

  async function onSubmit(data: FormData) {
    try {
      const response = await fetch(`${base_url}/api/auth/forget`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          email: data.email,
          password: data.password,
          email_code: data.verificationCode,
        }),
      });

      const payload = await response.json();

      if (!response.ok) {
        toast.error(payload?.message);
        return;
      }

      setAuth(payload?.access_token, payload?.email, payload?.user_id);

      router.push("/");
      toast.success("密码重置成功");
    } catch {
      toast.error("密码重置失败，请稍后重试");
    }
  }

  return (
    <div className="w-full max-w-md space-y-6 p-6 rounded-lg shadow-md bg-background">
      <div className="space-y-2 text-center">
        <h1 className="text-3xl font-bold">重置密码</h1>
        <p className="text-gray-500">
          请输入您的邮箱，我们将发送验证码帮助您重置密码
        </p>
      </div>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
          <FormField
            control={form.control}
            name="email"
            render={({ field }) => (
              <FormItem>
                <FormLabel>邮箱</FormLabel>
                <FormControl>
                  <Input placeholder="your@email.com" type="email" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="verificationCode"
            render={({ field }) => (
              <FormItem>
                <FormLabel>验证码</FormLabel>
                <div className="flex gap-2">
                  <FormControl>
                    <Input placeholder="请输入验证码" {...field} />
                  </FormControl>
                  <Button
                    type="button"
                    variant="outline"
                    className="whitespace-nowrap"
                    disabled={
                      isSendingCode || countdown > 0 || !form.getValues("email")
                    }
                    onClick={() =>
                      sendVerificationCode(form.getValues("email"))
                    }
                  >
                    {countdown > 0 ? (
                      `${countdown}秒后重试`
                    ) : isSendingCode ? (
                      <Loading />
                    ) : (
                      "发送验证码"
                    )}
                  </Button>
                </div>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="password"
            render={({ field }) => (
              <FormItem>
                <FormLabel>新密码</FormLabel>
                <FormControl>
                  <Input type="password" {...field} />
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
                <FormLabel>确认新密码</FormLabel>
                <FormControl>
                  <Input type="password" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <Button
            className="w-full"
            type="submit"
            disabled={form.formState.isSubmitting}
          >
            {form.formState.isSubmitting ? <Loading /> : "重置密码"}
          </Button>
        </form>
      </Form>
      <div className="text-center">
        <Link
          href="/auth/login"
          className="text-sm text-blue-600 hover:underline"
        >
          返回登录
        </Link>
      </div>
    </div>
  );
}
