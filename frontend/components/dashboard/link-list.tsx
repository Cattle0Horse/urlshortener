"use client";


import { useEffect, useState } from "react";
import { LinkCard } from "@/components/dashboard/link-card";
import { api } from "@/lib/api";
import { getCookie } from "@/lib/cookies";
import { toast } from "sonner";
import { Loading } from "@/components/ui/loading";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button"
import { CustomPagination } from "@/components/ui/custom-pagination";


interface Link {
	id: number;
	short_code: string;
	original_url: string;
	expiry_time: string;
	created_at: string;
}

interface PaginationState {
	page: number;
	pageSize: number;
	total: number;
}

export function LinkList() {
	const router = useRouter();
	const [links, setLinks] = useState<Link[]>([]);
	const [isLoading, setIsLoading] = useState(true);
	const [pagination, setPagination] = useState<PaginationState>({
		page: 1,
		pageSize: 9, // 3x3 grid layout
		total: 0,
	});

	const fetchLinks = async (page = pagination.page) => {
		try {
			const token = getCookie("token");
			if (!token) {
				toast.error("登录已过期");
				router.push("/login");
				return;
			}

			const response = await api.links.list(token, page, pagination.pageSize);
			if (response) {
				setLinks(response.urls);
				setPagination((prev) => ({
					...prev,
					page,
					total: response.total,
				}));
			}
		} catch (error) {
			console.error("Failed to fetch links:", error);
			toast.error("获取链接列表失败");
		} finally {
			setIsLoading(false);
		}
	};

	useEffect(() => {
		fetchLinks();
	}, [router]);

	const handlePageChange = (newPage: number) => {
		setIsLoading(true);
		fetchLinks(newPage);
		// Scroll to top when page changes
		window.scrollTo({ top: 0, behavior: "smooth" });
	};

	if (isLoading) {
		return <Loading />;
	}

	if (links.length === 0 && pagination.page === 1) {
		return (
			<div className="text-center py-12">
				<p className="text-lg text-muted-foreground">暂无短链接</p>
			</div>
		);
	}

	return (
		<div className="space-y-6">
			<div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3 animate-in fade-in-50 duration-500 min-h-[500px]">
				{links.map((link) => (
					<LinkCard
						key={link.id}
						link={link}
						onUpdate={() => fetchLinks(pagination.page)}
					/>
				))}
			</div>
			{pagination.total > pagination.pageSize && (
				<div className="flex justify-center mt-6">
					<CustomPagination
						currentPage={pagination.page}
						pageSize={pagination.pageSize}
						total={pagination.total}
						onPageChange={handlePageChange}
					/>
				</div>
			)}
		</div>
	);
}
