import { redirect } from 'next/navigation'
import { headers } from 'next/headers'
import { NextResponse } from 'next/server'

export async function GET(
  request: Request,
  { params }: { params: { shortCode: string } }
) {
  const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/url/${params.shortCode}`, {
    headers: headers(),
    redirect: 'manual', // 不要自动跟随重定向
  })
  
  if (!response.ok && response.status !== 302) {
    redirect('/')
  }

  // 如果是重定向响应，获取Location头并返回重定向
  if (response.status === 302) {
    const location = response.headers.get('location')
    if (!location) {
      redirect('/')
    }
    return NextResponse.redirect(location!)
  }

  return response
}
