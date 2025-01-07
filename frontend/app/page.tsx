import { Link2, BarChart3, Shield, Share2, ShoppingBag, MessageSquare, Presentation } from "lucide-react"

export default function Home() {
	return (
		<div className="min-h-[calc(100vh-var(--navbar-height))] flex flex-col">
			<main className="flex-1">
				{/* Hero Section */}
				<section className="w-full py-12 md:py-24 lg:py-32">
					<div className="container flex flex-col items-center justify-center space-y-8 px-4 md:px-6 text-center">
						{/* 标题区域 */}
						<div className="flex max-w-[980px] flex-col items-center gap-4">
							<h1 className="text-4xl font-bold tracking-tighter sm:text-5xl md:text-6xl bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
								更智能的链接管理方式
							</h1>
							<p className="mx-auto max-w-[700px] text-muted-foreground md:text-xl">
								简化您的链接，提升分享体验。快速生成短链接，轻松管理和追踪。
							</p>
						</div>

					</div>
				</section>

				{/* Features Section */}
				<section className="w-full py-12 md:py-24 bg-muted/50">
					<div className="container px-4 md:px-6">
						<div className="grid gap-8 md:gap-12 lg:grid-cols-3">
							<div className="space-y-4">
								<div className="inline-block rounded-lg bg-gradient-to-r from-primary to-secondary p-3 shadow-lg">
									<Link2 className="h-6 w-6 text-white" />
								</div>
								<h2 className="text-2xl font-bold tracking-tight bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
									简单快捷
								</h2>
								<p className="text-muted-foreground">
									只需一键即可生成短链接，无需复杂设置。支持自定义链接，让您的链接更具个性。提供批量处理功能，高效处理多个链接。
								</p>
							</div>
							<div className="space-y-4">
								<div className="inline-block rounded-lg bg-gradient-to-r from-primary to-secondary p-3 shadow-lg">
									<BarChart3 className="h-6 w-6 text-white" />
								</div>
								<h2 className="text-2xl font-bold tracking-tight bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
									数据分析
								</h2>
								<p className="text-muted-foreground">
									实时追踪链接点击量，了解访问来源。支持地理位置分析、设备统计、访问时段分析，助您做出更明智的决策。
								</p>
							</div>
							<div className="space-y-4">
								<div className="inline-block rounded-lg bg-gradient-to-r from-primary to-secondary p-3 shadow-lg">
									<Shield className="h-6 w-6 text-white" />
								</div>
								<h2 className="text-2xl font-bold tracking-tight bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
									安全可靠
								</h2>
								<p className="text-muted-foreground">
									提供链接有效期设置，支持访问密码保护。系统自动检测恶意链接，确保分享安全。多重备份确保数据安全。
								</p>
							</div>
						</div>
					</div>
				</section>

				{/* Stats Section */}
				<section className="w-full py-12 md:py-24">
					<div className="container px-4 md:px-6">
						<div className="grid gap-8 md:gap-12 lg:grid-cols-4 text-center">
							<div className="space-y-2">
								<h3 className="text-3xl font-bold bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">1M+</h3>
								<p className="text-muted-foreground">活跃用户</p>
							</div>
							<div className="space-y-2">
								<h3 className="text-3xl font-bold bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">100M+</h3>
								<p className="text-muted-foreground">生成的短链接</p>
							</div>
							<div className="space-y-2">
								<h3 className="text-3xl font-bold bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">99.9%</h3>
								<p className="text-muted-foreground">服务可用性</p>
							</div>
							<div className="space-y-2">
								<h3 className="text-3xl font-bold bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">24/7</h3>
								<p className="text-muted-foreground">全天候支持</p>
							</div>
						</div>
					</div>
				</section>

				{/* Use Cases Section */}
				<section className="w-full py-12 md:py-24 bg-muted/50">
					<div className="container px-4 md:px-6">
						<h2 className="text-3xl font-bold tracking-tight text-center mb-12 bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
							适用场景
						</h2>
						<div className="grid gap-8 md:gap-12 lg:grid-cols-4">
							<div className="space-y-4">
								<div className="inline-block rounded-lg bg-gradient-to-r from-primary to-secondary p-3 shadow-lg">
									<Share2 className="h-6 w-6 text-white" />
								</div>
								<h3 className="text-xl font-semibold">社交媒体</h3>
								<p className="text-muted-foreground">
									优化社交媒体分享链接，提升用户点击率
								</p>
							</div>
							<div className="space-y-4">
								<div className="inline-block rounded-lg bg-gradient-to-r from-primary to-secondary p-3 shadow-lg">
									<ShoppingBag className="h-6 w-6 text-white" />
								</div>
								<h3 className="text-xl font-semibold">电商营销</h3>
								<p className="text-muted-foreground">
									追踪营销活动效果，优化推广策略
								</p>
							</div>
							<div className="space-y-4">
								<div className="inline-block rounded-lg bg-gradient-to-r from-primary to-secondary p-3 shadow-lg">
									<MessageSquare className="h-6 w-6 text-white" />
								</div>
								<h3 className="text-xl font-semibold">即时通讯</h3>
								<p className="text-muted-foreground">
									简化长链接分享，提升沟通效率
								</p>
							</div>
							<div className="space-y-4">
								<div className="inline-block rounded-lg bg-gradient-to-r from-primary to-secondary p-3 shadow-lg">
									<Presentation className="h-6 w-6 text-white" />
								</div>
								<h3 className="text-xl font-semibold">数据分析</h3>
								<p className="text-muted-foreground">
									深入了解用户行为，优化业务决策
								</p>
							</div>
						</div>
					</div>
				</section>
			</main>
		</div>
	)
}
