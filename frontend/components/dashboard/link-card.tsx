"use client"

import { useState } from "react"
import { toast } from "sonner"
import { formatDateTime } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
} from "@/components/ui/card"
import {
  Copy,
  Clock,
  ExternalLink,
  Eye,
  Trash2,
  Loader2,
} from "lucide-react"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
  DialogFooter,
} from "@/components/ui/dialog"
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip"
import { Slider } from "@/components/ui/slider"
import { api } from "@/lib/api"
import { getCookie } from "@/lib/cookies"

interface LinkCardProps {
  link: {
    id: number
    short_code: string
    original_url: string
    expiry_time: string
    created_at: string
  }
  onUpdate?: () => void
}

export function LinkCard({ link, onUpdate }: LinkCardProps) {
  const [isUpdating, setIsUpdating] = useState(false)
  const [isDeleting, setIsDeleting] = useState(false)
  const [duration, setDuration] = useState(24)
  const [open, setOpen] = useState(false)
  const [showDeleteDialog, setShowDeleteDialog] = useState(false)

  const shortUrl = `${window.location.origin}/${link.short_code}`

  const copyToClipboard = async () => {
    try {
      await navigator.clipboard.writeText(shortUrl)
      toast.success('已复制到剪贴板', {
        description: shortUrl,
        action: {
          label: '访问链接',
          onClick: () => window.open(shortUrl, '_blank', 'noopener,noreferrer')
        }
      })
    } catch (err) {
      console.error('Failed to copy:', err)
      // 如果 clipboard API 失败，尝试使用传统方法
      const textarea = document.createElement('textarea')
      textarea.value = shortUrl
      textarea.style.position = 'fixed'
      textarea.style.opacity = '0'
      document.body.appendChild(textarea)
      textarea.select()
      try {
        document.execCommand('copy')
        toast.success('已复制到剪贴板', {
          description: shortUrl,
          action: {
            label: '访问链接',
            onClick: () => window.open(shortUrl, '_blank', 'noopener,noreferrer')
          }
        })
      } catch (err) {
        toast.error('复制失败', {
          description: '请手动复制以下链接：',
          action: {
            label: '访问链接',
            onClick: () => window.open(shortUrl, '_blank', 'noopener,noreferrer')
          }
        })
      } finally {
        document.body.removeChild(textarea)
      }
    }
  }

  const openShortUrl = () => {
    window.open(shortUrl, '_blank')
  }

  const openOriginalUrl = () => {
    let url = link.original_url
    // Add https:// if no protocol is specified
    if (!url.startsWith('http://') && !url.startsWith('https://')) {
      url = 'https://' + url
    }
    window.open(url, '_blank', 'noopener,noreferrer')
  }

  const handleUpdateDuration = async () => {
    try {
      setIsUpdating(true)
      const token = getCookie('token')
      if (!token) {
        toast.error('登录已过期')
        return
      }

      await api.links.update(link.short_code, duration, token)
      toast.success('更新成功', {
        description: `短链接有效期已更新为 ${duration} 小时`
      })
      onUpdate?.()
      setOpen(false)
    } catch (error) {
      console.error('Failed to update expiry:', error)
      toast.error('更新失败', {
        description: '请稍后重试'
      })
    } finally {
      setIsUpdating(false)
    }
  }

  const handleDelete = async () => {
    try {
      setIsDeleting(true)
      const token = getCookie('token')
      if (!token) {
        toast.error('登录已过期', {
          description: '请重新登录后再试'
        })
        return
      }

      await api.links.delete(link.short_code, token)
      toast.success('删除成功')
      onUpdate?.()
      setShowDeleteDialog(false)
    } catch (error) {
      console.error('Failed to delete link:', error)
      if (error instanceof Error) {
        toast.error('删除失败', {
          description: error.message
        })
      } else {
        toast.error('删除失败', {
          description: '请稍后重试'
        })
      }
    } finally {
      setIsDeleting(false)
    }
  }

  return (
    <div>
      <Card className="group hover:shadow-md transition-all duration-200">
        <CardHeader>
          <div className="flex items-center justify-between">
            <CardTitle className="text-xl flex items-center gap-2">
              <span className="bg-gradient-to-r from-primary/80 to-primary text-white px-2 py-1 rounded text-sm font-normal">
                {link.short_code}
              </span>
            </CardTitle>
            <div className="flex items-center gap-2">
              <TooltipProvider>
                <Tooltip>
                  <TooltipTrigger asChild>
                    <Button
                      variant="ghost"
                      size="icon"
                      className="h-8 w-8 hover:bg-primary hover:text-white transition-colors"
                      onClick={copyToClipboard}
                    >
                      <Copy className="h-4 w-4" />
                      <span className="sr-only">复制链接</span>
                    </Button>
                  </TooltipTrigger>
                  <TooltipContent>
                    <p className="text-xs">复制短链接</p>
                  </TooltipContent>
                </Tooltip>
              </TooltipProvider>

              <TooltipProvider>
                <Tooltip>
                  <TooltipTrigger asChild>
                    <Button
                      variant="ghost"
                      size="icon"
                      className="h-8 w-8 hover:bg-primary hover:text-white transition-colors"
                      onClick={() => setOpen(true)}
                    >
                      <Clock className="h-4 w-4" />
                      <span className="sr-only">更新到期时间</span>
                    </Button>
                  </TooltipTrigger>
                  <TooltipContent>
                    <p className="text-xs">更新到期时间</p>
                  </TooltipContent>
                </Tooltip>
              </TooltipProvider>

              <TooltipProvider>
                <Tooltip>
                  <TooltipTrigger asChild>
                    <Button
                      variant="ghost"
                      size="icon"
                      className="h-8 w-8 hover:bg-destructive hover:text-destructive-foreground transition-colors"
                      onClick={() => setShowDeleteDialog(true)}
                    >
                      <Trash2 className="h-4 w-4" />
                      <span className="sr-only">删除短链接</span>
                    </Button>
                  </TooltipTrigger>
                  <TooltipContent>
                    <p className="text-xs">删除短链接</p>
                  </TooltipContent>
                </Tooltip>
              </TooltipProvider>
            </div>
          </div>
          <CardDescription className="mt-2 flex items-center gap-2">
            <span className="text-muted-foreground shrink-0">原始链接：</span>
            <TooltipProvider>
              <Tooltip>
                <TooltipTrigger asChild>
                  <span className="font-medium truncate flex-1 cursor-default">{link.original_url}</span>
                </TooltipTrigger>
                <TooltipContent side="bottom" align="start" className="max-w-[300px]">
                  <p className="break-all text-sm">{link.original_url}</p>
                </TooltipContent>
              </Tooltip>
            </TooltipProvider>
            <TooltipProvider>
              <Tooltip>
                <TooltipTrigger asChild>
                  <Button
                    variant="ghost"
                    size="icon"
                    className="h-6 w-6 shrink-0"
                    onClick={openOriginalUrl}
                  >
                    <ExternalLink className="h-4 w-4" />
                    <span className="sr-only">访问原始链接</span>
                  </Button>
                </TooltipTrigger>
                <TooltipContent side="bottom" align="end">
                  <p className="text-xs">访问原始链接</p>
                </TooltipContent>
              </Tooltip>
            </TooltipProvider>
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div className="space-y-3">
            <div className="flex items-center justify-between group/item hover:bg-muted/50 p-2 rounded-lg transition-colors">
              <span className="text-sm text-muted-foreground">短链接</span>
              <button
                onClick={openShortUrl}
                className="text-sm font-medium text-primary hover:underline cursor-pointer"
                title="点击访问短链接"
              >
                {shortUrl}
              </button>
            </div>
            <div className="flex items-center justify-between group/item hover:bg-muted/50 p-2 rounded-lg transition-colors">
              <span className="text-sm text-muted-foreground">过期时间</span>
              <span className="text-sm font-medium">{formatDateTime(link.expiry_time)}</span>
            </div>
            <div className="flex items-center justify-between group/item hover:bg-muted/50 p-2 rounded-lg transition-colors">
              <span className="text-sm text-muted-foreground">创建时间</span>
              <span className="text-sm font-medium">
                {formatDateTime(link.created_at)}
              </span>
            </div>
          </div>
        </CardContent>
      </Card>

      {/* 更新时间对话框 */}
      <Dialog open={open} onOpenChange={setOpen}>
        <DialogContent className="sm:max-w-md">
          <DialogHeader>
            <DialogTitle>更新到期时间</DialogTitle>
            <DialogDescription>
              调整滑块来设置新的到期时间
            </DialogDescription>
          </DialogHeader>
          <div className="py-6">
            <div className="space-y-4">
              <div className="px-1">
                <Slider
                  min={1}
                  max={168}
                  step={1}
                  value={[duration]}
                  onValueChange={([value]) => setDuration(value)}
                  disabled={isUpdating}
                />
                <div className="flex items-center justify-between mt-4 text-sm">
                  <span className="text-muted-foreground">1小时</span>
                  <span className="font-medium text-base text-primary">{duration}小时</span>
                  <span className="text-muted-foreground">168小时</span>
                </div>
              </div>
              <Button
                onClick={handleUpdateDuration}
                disabled={isUpdating}
                className="w-full bg-gradient-to-r from-primary to-primary/80 text-white hover:opacity-90"
              >
                {isUpdating ? (
                  <>
                    <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                    更新中...
                  </>
                ) : (
                  "确认更新"
                )}
              </Button>
            </div>
          </div>
        </DialogContent>
      </Dialog>

      {/* 删除确认对话框 */}
      <Dialog open={showDeleteDialog} onOpenChange={setShowDeleteDialog}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>删除短链接</DialogTitle>
            <DialogDescription>
              确定要删除这个短链接吗？此操作不可撤销。
            </DialogDescription>
          </DialogHeader>
          <DialogFooter className="gap-2 sm:gap-0">
            <Button
              variant="outline"
              onClick={() => setShowDeleteDialog(false)}
            >
              取消
            </Button>
            <Button
              variant="destructive"
              onClick={handleDelete}
              disabled={isDeleting}
            >
              {isDeleting ? (
                <>
                  <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                  删除中...
                </>
              ) : (
                "确认删除"
              )}
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  )
}