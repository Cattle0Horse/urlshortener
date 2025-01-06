"use client"

import { useState, useEffect } from 'react'
import { motion, AnimatePresence } from 'framer-motion'
import { Copy, Trash2, Clock, ExternalLink } from 'lucide-react'
import { Button } from '@/components/ui/button'
import { useToast } from '@/components/ui/use-toast'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import { timeAgo, copyToClipboard } from '@/lib/utils'

interface Url {
  id: string
  originalUrl: string
  shortCode: string
  createdAt: number
  expiresAt: number
  clicks: number
}

export function UrlList() {
  const [urls, setUrls] = useState<Url[]>([])
  const [isLoading, setIsLoading] = useState(true)
  const { toast } = useToast()

  useEffect(() => {
    fetchUrls()
  }, [])

  const fetchUrls = async () => {
    try {
      const response = await fetch('/api/urls')
      if (!response.ok) throw new Error('获取数据失败')
      const data = await response.json()
      setUrls(data)
    } catch (error) {
      console.error('Failed to fetch URLs:', error)
      toast({
        variant: 'destructive',
        title: '获取数据失败',
        description: '请稍后重试',
      })
    } finally {
      setIsLoading(false)
    }
  }

  const handleCopy = async (shortUrl: string) => {
    const success = await copyToClipboard(shortUrl)
    if (success) {
      toast({
        title: '复制成功',
        description: '短链接已复制到剪贴板',
      })
    } else {
      toast({
        variant: 'destructive',
        title: '复制失败',
        description: '请手动复制链接',
      })
    }
  }

  const handleDelete = async (id: string) => {
    try {
      const response = await fetch(`/api/urls/${id}`, {
        method: 'DELETE',
      })
      if (!response.ok) throw new Error('删除失败')
      
      setUrls((prev) => prev.filter((url) => url.id !== id))
      toast({
        title: '删除成功',
        description: '短链接已删除',
      })
    } catch (error) {
      console.error('Failed to delete URL:', error)
      toast({
        variant: 'destructive',
        title: '删除失败',
        description: '请稍后重试',
      })
    }
  }

  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-64">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary" />
      </div>
    )
  }

  if (urls.length === 0) {
    return (
      <motion.div
        initial={{ opacity: 0 }}
        animate={{ opacity: 1 }}
        className="flex flex-col items-center justify-center h-64 text-center"
      >
        <h3 className="text-xl font-semibold mb-2">暂无短链接</h3>
        <p className="text-gray-600 dark:text-gray-300">
          点击上方按钮创建您的第一个短链接
        </p>
      </motion.div>
    )
  }

  return (
    <div className="rounded-lg border bg-card">
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>原始链接</TableHead>
            <TableHead>短链接</TableHead>
            <TableHead>创建时间</TableHead>
            <TableHead>过期时间</TableHead>
            <TableHead>点击次数</TableHead>
            <TableHead className="text-right">操作</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          <AnimatePresence>
            {urls.map((url) => (
              <motion.tr
                key={url.id}
                initial={{ opacity: 0, y: 20 }}
                animate={{ opacity: 1, y: 0 }}
                exit={{ opacity: 0, y: -20 }}
                className="group hover:bg-muted/50"
              >
                <TableCell className="font-medium max-w-[200px] truncate">
                  <a
                    href={url.originalUrl}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="flex items-center hover:text-primary"
                  >
                    {url.originalUrl}
                    <ExternalLink className="ml-2 h-4 w-4 opacity-0 group-hover:opacity-100 transition-opacity" />
                  </a>
                </TableCell>
                <TableCell>
                  <div className="flex items-center space-x-2">
                    <span className="font-mono">{url.shortCode}</span>
                    <Button
                      variant="ghost"
                      size="icon"
                      onClick={() => handleCopy(url.shortCode)}
                    >
                      <Copy className="h-4 w-4" />
                    </Button>
                  </div>
                </TableCell>
                <TableCell>{timeAgo(url.createdAt)}</TableCell>
                <TableCell>
                  <div className="flex items-center">
                    <Clock className="mr-2 h-4 w-4 text-gray-500" />
                    {timeAgo(url.expiresAt)}
                  </div>
                </TableCell>
                <TableCell>{url.clicks}</TableCell>
                <TableCell className="text-right">
                  <Button
                    variant="ghost"
                    size="icon"
                    onClick={() => handleDelete(url.id)}
                  >
                    <Trash2 className="h-4 w-4 text-red-500" />
                  </Button>
                </TableCell>
              </motion.tr>
            ))}
          </AnimatePresence>
        </TableBody>
      </Table>
    </div>
  )
} 