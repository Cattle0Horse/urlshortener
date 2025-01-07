import { Metadata } from "next"
import { CreateLinkForm } from "@/components/create-link-form"

export const metadata: Metadata = {
  title: "创建短链接 | ShortLink",
  description: "创建一个新的短链接",
}

export default function CreatePage() {
  return (
    <div className="container max-w-2xl py-6 lg:py-10">
      <div className="flex flex-col items-start gap-4 md:flex-row md:justify-between md:gap-8">
        <div className="flex-1 space-y-4">
          <h1 className="inline-block font-heading text-4xl tracking-tight lg:text-5xl">
            创建短链接
          </h1>
          <p className="text-xl text-muted-foreground">
            输入您想要缩短的链接，获取一个简短的URL。
          </p>
        </div>
      </div>
      <hr className="my-8" />
      <CreateLinkForm />
    </div>
  )
} 