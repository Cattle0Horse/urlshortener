import { Metadata } from "next"
import { LinkList } from "@/components/dashboard/link-list"

export const metadata: Metadata = {
  title: "仪表板 | ShortLink",
  description: "管理您的短链接",
}

export default function DashboardPage() {
  return (
    <div className="container py-6 lg:py-10">
      <div className="flex flex-col items-start gap-4 md:flex-row md:justify-between md:gap-8">
        <div className="flex-1 space-y-4">
          <h1 className="inline-block font-heading text-4xl tracking-tight lg:text-5xl">
            我的短链接
          </h1>
          <p className="text-xl text-muted-foreground">
            管理和追踪您创建的所有短链接。
          </p>
        </div>
      </div>
      <hr className="my-8" />
      <LinkList />
    </div>
  )
} 