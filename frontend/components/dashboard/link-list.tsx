"use client"

import * as React from "react"
import { toast } from "sonner"
import { LinkCard } from "@/components/dashboard/link-card"
import { Loading } from "@/components/ui/loading"

interface Link {
  id: string
  url: string
  shortUrl: string
  slug: string
  clicks: number
  createdAt: string
}

export function LinkList() {
  const [links, setLinks] = React.useState<Link[]>([])
  const [isLoading, setIsLoading] = React.useState(true)

  React.useEffect(() => {
    async function fetchLinks() {
      try {
        const response = await fetch("/api/links")
        if (!response.ok) throw new Error("获取失败")
        const data = await response.json()
        setLinks(data)
      } catch (error) {
        toast.error("获取链接列表失败")
      } finally {
        setIsLoading(false)
      }
    }

    fetchLinks()
  }, [])

  if (isLoading) return <Loading />

  if (links.length === 0) {
    return (
      <div className="flex flex-col items-center justify-center space-y-4">
        <p className="text-lg text-muted-foreground">
          您还没有创建任何短链接
        </p>
      </div>
    )
  }

  return (
    <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
      {links.map((link) => (
        <LinkCard key={link.id} link={link} />
      ))}
    </div>
  )
} 