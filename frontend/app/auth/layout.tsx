"use client";

import { useRouter } from "next/navigation";
import { useAuth } from "@/components/context";

export default function AuthLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const router = useRouter();
  const { isAuth } = useAuth();

  if (isAuth) {
    router.push("/");
    return;
  }

  return <>{children}</>;
}
