import "./globals.css";
import { Inter } from "next/font/google";
import MainLayout from "@/components/layout/main-layout";
import { cn } from "@/lib/utils";
import { ThemeProvider } from "@/components/theme-provider";
import { Navbar } from "@/components/nav/navbar";
import { Toaster } from "sonner";

const inter = Inter({ subsets: ["latin"] });

export default function RootLayout({
	children,
}: {
	children: React.ReactNode;
}) {
	return (
		<html lang="zh-CN">
			<body
				className={cn(
					"min-h-screen bg-background antialiased",
					inter.className
				)}
			>
				<ThemeProvider
					attribute="class"
					defaultTheme="system"
					enableSystem
					disableTransitionOnChange
				>
					<div className="relative flex min-h-screen flex-col">
						<Navbar />
						{children}
					</div>
					<Toaster richColors closeButton position="top-right" />
				</ThemeProvider>
			</body>
		</html>
	);
}
