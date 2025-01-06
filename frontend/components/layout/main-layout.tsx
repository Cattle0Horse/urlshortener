"use client"

import { ThemeProvider } from "next-themes"
import { Toaster } from "sonner"
import Header from "./header"

export default function MainLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <ThemeProvider 
      attribute="class" 
      defaultTheme="system" 
      enableSystem
    >
      <div className="relative min-h-screen bg-background">
        <div className="absolute inset-0 bg-gradient-to-br from-background to-secondary opacity-50" />
        <div className="relative z-10">
          <Header />
          <main className="container mx-auto px-4 py-8">
            {children}
          </main>
          <Toaster 
            position="top-center"
            toastOptions={{
              style: {
                background: 'hsl(var(--background))',
                color: 'hsl(var(--foreground))',
                border: '1px solid hsl(var(--border))',
              },
            }}
          />
        </div>
      </div>
    </ThemeProvider>
  )
} 