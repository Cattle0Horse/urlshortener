import Link from "next/link"
import { Button } from "@/components/ui/button"
import { Link2 } from "lucide-react"

export default function NotFound() {
  return (
    <div className="min-h-[calc(100vh-var(--navbar-height))] w-full bg-gradient-to-b from-background to-muted relative overflow-hidden">
      {/* 装饰性背景元素 */}
      <div className="absolute inset-0 w-full h-full">
        <div className="absolute top-0 -left-4 w-72 h-72 bg-primary/10 rounded-full filter blur-3xl animate-blob" />
        <div className="absolute -bottom-8 right-4 w-72 h-72 bg-secondary/10 rounded-full filter blur-3xl animate-blob animation-delay-2000" />
      </div>

      <div className="container relative flex flex-col items-center justify-center min-h-[calc(100vh-var(--navbar-height))] text-center">
        <div className="bg-gradient-to-r from-primary to-secondary p-3 rounded-2xl mb-8">
          <Link2 className="w-12 h-12 text-white" />
        </div>
        <h1 className="text-4xl font-bold tracking-tighter sm:text-5xl md:text-6xl lg:text-7xl bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent mb-4">
          404
        </h1>
        <p className="max-w-[600px] text-lg text-muted-foreground mb-8">
          抱歉，您访问的页面不存在。可能是链接已过期或输入的地址有误。
        </p>
        <Button asChild>
          <Link href="/" className="bg-gradient-to-r from-primary to-secondary text-white">
            返回首页
          </Link>
        </Button>
      </div>
    </div>
  )
} 