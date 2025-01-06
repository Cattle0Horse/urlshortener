"use client";

import { useEffect, useState } from 'react'
import Link from 'next/link'
import { useTheme } from 'next-themes'
import { motion } from 'framer-motion'
import { Moon, Sun, Link as LinkIcon, User, LogOut } from 'lucide-react'
import { Button } from '@/components/ui/button'
import { getCookie, removeCookie } from '@/lib/cookies'
import { useRouter } from 'next/navigation'
import { toast } from "sonner"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"

export function Navbar() {
  const { theme, setTheme } = useTheme()
  const router = useRouter()
  const [mounted, setMounted] = useState(false)
  const [isLoggedIn, setIsLoggedIn] = useState(false)
  const [userEmail, setUserEmail] = useState("")

  useEffect(() => {
    setMounted(true)
  }, [])

  useEffect(() => {
    const token = getCookie('token')
    const email = getCookie('email')
    setIsLoggedIn(!!token)
    setUserEmail(email || "")
  }, [])

  if (!mounted) {
    return (
      <header className="sticky top-0 z-50 w-full border-b bg-background/95 backdrop-blur h-[var(--navbar-height)]">
        <div className="container flex h-full items-center">
          <div className="flex items-center justify-between h-16 w-full">
            <div className="flex items-center">
              <Link href="/" className="flex items-center space-x-2">
                <LinkIcon className="h-6 w-6 text-blue-600 dark:text-blue-400" />
                <span className="text-xl font-bold bg-gradient-to-r from-blue-600 to-cyan-500 bg-clip-text text-transparent">
                  URLify
                </span>
              </Link>
            </div>
          </div>
        </div>
      </header>
    )
  }

  const handleLogout = () => {
    // 清除所有 cookies
    removeCookie('token')
    removeCookie('email')
    removeCookie('user_id')
    
    // 更新状态
    setIsLoggedIn(false)
    setUserEmail("")
    
    // 显示退出提示
    toast.success("已退出登录")
    
    // 跳转到首页
    router.push('/')
    router.refresh()
  }

  return (
    <header className="sticky top-0 z-50 w-full border-b bg-background/95 backdrop-blur h-[var(--navbar-height)]">
      <div className="container flex h-full items-center">
        <div className="flex items-center justify-between h-16 w-full">
          <div className="flex items-center">
            <Link href="/" className="flex items-center space-x-2">
              <LinkIcon className="h-6 w-6 text-blue-600 dark:text-blue-400" />
              <span className="text-xl font-bold bg-gradient-to-r from-blue-600 to-cyan-500 bg-clip-text text-transparent">
                URLify
              </span>
            </Link>
          </div>

          <div className="flex items-center space-x-4">
            <motion.button
              whileHover={{ scale: 1.05 }}
              whileTap={{ scale: 0.95 }}
              onClick={() => setTheme(theme === 'dark' ? 'light' : 'dark')}
              className="p-2 rounded-lg bg-gray-100 dark:bg-gray-800 hover:bg-gray-200 dark:hover:bg-gray-700 transition-colors"
            >
              {theme === 'dark' ? (
                <Sun className="h-5 w-5 text-yellow-500" />
              ) : (
                <Moon className="h-5 w-5 text-blue-600" />
              )}
            </motion.button>

            {isLoggedIn ? (
              <DropdownMenu>
                <DropdownMenuTrigger asChild>
                  <Button variant="ghost" className="relative">
                    <User className="h-5 w-5" />
                    <span className="ml-2 hidden md:inline-block">{userEmail}</span>
                  </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent align="end">
                  <DropdownMenuItem asChild>
                    <Link href="/dashboard">
                      仪表板
                    </Link>
                  </DropdownMenuItem>
                  <DropdownMenuItem onClick={handleLogout}>
                    <LogOut className="mr-2 h-4 w-4" />
                    退出登录
                  </DropdownMenuItem>
                </DropdownMenuContent>
              </DropdownMenu>
            ) : (
              <div className="hidden md:flex items-center space-x-2">
                <Button variant="ghost" asChild>
                  <Link href="/login">登录</Link>
                </Button>
                <Button asChild>
                  <Link href="/register">注册</Link>
                </Button>
              </div>
            )}
          </div>
        </div>
      </div>
    </header>
  )
}
