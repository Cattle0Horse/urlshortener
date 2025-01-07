import { Metadata } from "next"
import Link from "next/link"
import { Button } from "@/components/ui/button"
import { LinkList } from "@/components/dashboard/link-list"
import { Plus } from "lucide-react"

export const metadata: Metadata = {
  title: "仪表板 | URLify",
  description: "管理您的短链接",
}

export default function DashboardPage() {
  return (
    <div className="container py-8">
      <div className="flex flex-col gap-8">
        <div className="flex items-center justify-between">
          <div>
            <h1 className="text-3xl font-bold tracking-tight">我的短链接</h1>
            <p className="text-muted-foreground">
              查看和管理您创建的所有短链接
            </p>
          </div>
          <Button asChild>
            <Link href="/create" className="bg-gradient-to-r from-primary to-secondary text-white">
              <Plus className="mr-2 h-4 w-4" />
              创建短链接
            </Link>
          </Button>
        </div>
        <LinkList />
      </div>
    </div>
  )
} 