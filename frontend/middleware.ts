import { NextResponse } from 'next/server'
import type { NextRequest } from 'next/server'

// 需要认证的路由
const protectedRoutes = [
  "/dashboard",
  "/create",
]

export function middleware(request: NextRequest) {
  const token = request.cookies.get('token')
  const { pathname } = request.nextUrl

  // 检查是否是受保护的路由
  if (protectedRoutes.some(route => pathname.startsWith(route))) {
    // 如果没有 token，重定向到登录页
    if (!token) {
      const url = new URL("/login", request.url)
      url.searchParams.set("from", pathname)
      return NextResponse.redirect(url)
    }
  }

  // 如果已登录，访问登录/注册页面时重定向到仪表板
  if ((pathname === "/login" || pathname === "/register") && token) {
    return NextResponse.redirect(new URL("/dashboard", request.url))
  }

  return NextResponse.next()
}

export const config = {
  matcher: [
    /*
     * 匹配所有需要认证的路由:
     * - /dashboard
     * - /create
     * - /login
     * - /register
     */
    "/dashboard/:path*",
    "/create/:path*",
    "/login",
    "/register",
  ],
} 