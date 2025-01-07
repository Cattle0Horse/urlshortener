import "./globals.css";
import { ThemeProvider } from "../components/theme-provider";
import { Navbar } from "../components/nav/navbar";
import { Toaster } from "sonner";
import { AuthProvider } from "../components/auth/auth-provider";

export default function RootLayout({
	children,
}: {
	children: React.ReactNode;
}) {
	return (
		<html lang="zh-CN" suppressHydrationWarning>
			<body className="min-h-screen bg-background antialiased">
				<AuthProvider>
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
						<Toaster expand={true} position="top-center" richColors />
					</ThemeProvider>
				</AuthProvider>
			</body>
		</html>
	);
}
