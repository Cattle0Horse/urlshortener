"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { cn } from "@/lib/utils";
import { useAuth } from "../context";
import Logout from "../logout";
import { ModeToggle } from "../mode-toggle";

export function Navbar() {
  const pathname = usePathname();
  const { email } = useAuth();

  // Hide navbar on auth pages
  if (pathname?.startsWith("/auth")) {
    return null;
  }

  const navItems = [
    {
      name: "创建短链接",
      href: "/",
    },
    {
      name: "我的短链接",
      href: "/url",
    },
  ];

  return (
    <nav className="border-b">
      <div className="flex h-16 items-center px-4 w-full">
        <div className="space-x-6 text-sm font-medium flex mx-auto items-center justify-center w-full">
          {navItems.map((item) => (
            <Link
              key={item.href}
              href={item.href}
              className={cn(
                "transition-colors hover:text-primary",
                pathname === item.href
                  ? "text-foreground"
                  : "text-foreground/60"
              )}
            >
              {item.name}
            </Link>
          ))}

          <p>{email}</p>

          <Logout />
          <ModeToggle />
        </div>
      </div>
    </nav>
  );
}
