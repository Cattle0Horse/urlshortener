"use client";

import * as React from "react";
import { Moon, Sun } from "lucide-react";
import { useTheme } from "next-themes";
import { Button } from "@/components/ui/button";

export function ThemeToggle() {
	const { theme, setTheme } = useTheme();
	const [mounted, setMounted] = React.useState(false);

	// Handle mounting on client side
	React.useEffect(() => {
		setMounted(true);
	}, []);

	// 防止初始化时的水合不匹配
	React.useEffect(() => {
		if (mounted) {
			document.body.style.setProperty(
				"--initial-color-scheme",
				theme || "light"
			);
		}
	}, [mounted, theme]);

	// 在主题未加载完成前不渲染按钮
	if (!mounted) {
		return null;
	}

	return (
		<Button
			variant="ghost"
			size="icon"
			onClick={() => setTheme(theme === "dark" ? "light" : "dark")}
		>
			<Sun className="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0" />
			<Moon className="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100" />
			<span className="sr-only">切换主题</span>
		</Button>
	);
}
