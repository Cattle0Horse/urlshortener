"use client";

import { useState, useEffect } from "react";
import Link from "next/link";
import { Button } from "../ui/button";
import { Menu } from "lucide-react";
import { ThemeToggle } from "../theme-toggle";

export default function Header() {
	const [mounted, setMounted] = useState(false);
	const [isMenuOpen, setIsMenuOpen] = useState(false);

	useEffect(() => setMounted(true), []);

	if (!mounted) return null;

	return (
		<header className="sticky top-0 z-50 w-full">
			<nav className="glass-effect mx-auto flex h-16 items-center justify-between px-4 sm:px-6 lg:px-8">
				<Link href="/" className="flex items-center space-x-2">
					<span className="text-xl font-bold">ShortLink</span>
				</Link>

				{/* Desktop Navigation */}
				<div className="hidden md:flex items-center space-x-6">
					<Link
						href="/dashboard"
						className="text-sm font-medium hover:text-primary"
					>
						仪表盘
					</Link>
					<Link
						href="/create"
						className="text-sm font-medium hover:text-primary"
					>
						创建链接
					</Link>
					<div className="flex items-center space-x-2">
						<Button variant="ghost" size="sm" asChild>
							<Link href="/login">登录</Link>
						</Button>
						<Button size="sm" asChild>
							<Link href="/register">注册</Link>
						</Button>
					</div>
					<ThemeToggle />
				</div>

				{/* Mobile Navigation */}
				<div className="md:hidden flex items-center space-x-2">
					<ThemeToggle />
					<Button
						variant="ghost"
						size="icon"
						onClick={() => setIsMenuOpen(!isMenuOpen)}
					>
						<Menu className="h-5 w-5" />
					</Button>
				</div>
			</nav>

			{/* Mobile Menu */}
			{isMenuOpen && (
				<div className="glass-effect md:hidden">
					<div className="px-2 pt-2 pb-3 space-y-1">
						<Link
							href="/dashboard"
							className="block px-3 py-2 rounded-md text-base font-medium hover:text-primary"
						>
							仪表盘
						</Link>
						<Link
							href="/create"
							className="block px-3 py-2 rounded-md text-base font-medium hover:text-primary"
						>
							创建链接
						</Link>
						<Link
							href="/login"
							className="block px-3 py-2 rounded-md text-base font-medium hover:text-primary"
						>
							登录
						</Link>
						<Link
							href="/register"
							className="block px-3 py-2 rounded-md text-base font-medium hover:text-primary"
						>
							注册
						</Link>
					</div>
				</div>
			)}
		</header>
	);
}
