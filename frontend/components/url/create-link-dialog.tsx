'use client'

import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { motion, AnimatePresence } from 'framer-motion'
import { z } from 'zod'
import { Loader2 } from 'lucide-react'

import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { isValidUrl } from '@/lib/utils'

const createLinkSchema = z.object({
  url: z.string().min(1, '请输入URL').refine(isValidUrl, '请输入有效的URL'),
  expiry: z.string().min(1, '请选择有效期'),
})

type CreateLinkForm = z.infer<typeof createLinkSchema>

interface CreateLinkDialogProps {
  open: boolean
  onOpenChange: (open: boolean) => void
  onSuccess: () => void
}

export function CreateLinkDialog({
  open,
  onOpenChange,
  onSuccess,
}: CreateLinkDialogProps) {
  const [isLoading, setIsLoading] = useState(false)
  const [createdUrl, setCreatedUrl] = useState<string | null>(null)

  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
    setValue,
    watch,
  } = useForm<CreateLinkForm>({
    resolver: zodResolver(createLinkSchema),
    defaultValues: {
      expiry: '7d',
    },
  })

  const handleClose = () => {
    reset()
    setCreatedUrl(null)
    onOpenChange(false)
  }

  const onSubmit = async (data: CreateLinkForm) => {
    try {
      setIsLoading(true)
      const response = await fetch('/api/urls', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      })

      if (!response.ok) {
        throw new Error('创建失败')
      }

      const result = await response.json()
      setCreatedUrl(result.shortUrl)
      onSuccess()
    } catch (error) {
      console.error('Failed to create short URL:', error)
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <Dialog open={open} onOpenChange={handleClose}>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>创建短链接</DialogTitle>
          <DialogDescription>
            输入您想要缩短的URL地址，选择有效期
          </DialogDescription>
        </DialogHeader>

        <form onSubmit={handleSubmit(onSubmit)} className="space-y-4 mt-4">
          <div className="space-y-2">
            <Input
              {...register('url')}
              placeholder="请输入URL地址"
              className={errors.url ? 'border-red-500' : ''}
            />
            {errors.url && (
              <p className="text-sm text-red-500">{errors.url.message}</p>
            )}
          </div>

          <div className="space-y-2">
            <Select
              value={watch('expiry')}
              onValueChange={(value: string) => setValue('expiry', value)}
            >
              <SelectTrigger>
                <SelectValue placeholder="选择有效期" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="1d">1天</SelectItem>
                <SelectItem value="7d">7天</SelectItem>
                <SelectItem value="30d">30天</SelectItem>
                <SelectItem value="90d">90天</SelectItem>
                <SelectItem value="180d">180天</SelectItem>
                <SelectItem value="365d">365天</SelectItem>
              </SelectContent>
            </Select>
            {errors.expiry && (
              <p className="text-sm text-red-500">{errors.expiry.message}</p>
            )}
          </div>

          <div className="flex justify-end space-x-2">
            <Button variant="outline" onClick={handleClose}>
              取消
            </Button>
            <Button type="submit" disabled={isLoading}>
              {isLoading ? (
                <>
                  <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                  创建中...
                </>
              ) : (
                '创建'
              )}
            </Button>
          </div>
        </form>

        <AnimatePresence>
          {createdUrl && (
            <motion.div
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              exit={{ opacity: 0, y: 20 }}
              className="mt-4 p-4 bg-green-50 dark:bg-green-900/20 rounded-lg"
            >
              <p className="text-sm text-green-600 dark:text-green-400">
                短链接创建成功：
              </p>
              <p className="mt-1 font-mono text-sm">{createdUrl}</p>
              <Button
                variant="outline"
                size="sm"
                className="mt-2 w-full"
                onClick={() => {
                  navigator.clipboard.writeText(createdUrl)
                }}
              >
                复制链接
              </Button>
            </motion.div>
          )}
        </AnimatePresence>
      </DialogContent>
    </Dialog>
  )
} 