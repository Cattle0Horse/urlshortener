"use client"

import { useEffect } from "react"
import { Button } from "@/components/ui/button"

export default function ErrorBoundary({
  error,
  reset,
}: {
  error: Error & { digest?: string }
  reset: () => void
}) {
  useEffect(() => {
    console.error(error)
  }, [error])

  return (
    <div className="flex h-[50vh] flex-col items-center justify-center space-y-4">
      <div className="text-center">
        <h2 className="text-2xl font-bold">出错了</h2>
        <p className="text-muted-foreground">抱歉，发生了一些错误</p>
      </div>
      <Button
        variant="outline"
        onClick={
          // 尝试恢复渲染
          () => reset()
        }
      >
        重试
      </Button>
    </div>
  )
} 