"use client"

import { toast } from "sonner"
import { formatDate } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
  CardFooter,
} from "@/components/ui/card"
import { Icons } from "@/components/ui/icons"

interface LinkCardProps {
  link: {
    id: string
    url: string
    shortUrl: string
    slug: string
    clicks: number
    createdAt: string
  }
}

export function LinkCard({ link }: LinkCardProps) {
  const copyToClipboard = async () => {
    try {
      await navigator.clipboard.writeText(link.shortUrl)
      toast.success("已复制到剪贴板")
    } catch (err) {
      toast.error("复制失败，请手动复制")
    }
  }

  return (
    <Card>
      <CardHeader>
        <CardTitle className="text-xl">{link.slug}</CardTitle>
        <CardDescription className="truncate">
          原始链接：{link.url}
        </CardDescription>
      </CardHeader>
      <CardContent>
        <div className="space-y-2">
          <div className="flex items-center justify-between">
            <span className="text-sm text-muted-foreground">短链接</span>
            <span className="text-sm font-medium">{link.shortUrl}</span>
          </div>
          <div className="flex items-center justify-between">
            <span className="text-sm text-muted-foreground">访问次数</span>
            <span className="text-sm font-medium">{link.clicks}</span>
          </div>
          <div className="flex items-center justify-between">
            <span className="text-sm text-muted-foreground">创建时间</span>
            <span className="text-sm font-medium">
              {formatDate(link.createdAt)}
            </span>
          </div>
        </div>
      </CardContent>
      <CardFooter>
        <Button
          variant="secondary"
          className="w-full"
          onClick={copyToClipboard}
        >
          <Icons.copy className="mr-2 h-4 w-4" />
          复制链接
        </Button>
      </CardFooter>
    </Card>
  )
} 