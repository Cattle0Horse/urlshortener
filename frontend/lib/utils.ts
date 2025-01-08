import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
	return twMerge(clsx(inputs));
}

export function formatDate(input: string | number): string {
	const date = new Date(input);
	return date.toLocaleDateString("zh-CN", {
		month: "long",
		day: "numeric",
		year: "numeric",
	});
}

export function formatDateTime(input: string | number): string {
	const date = new Date(input);
	return date.toLocaleString("zh-CN", {
		year: "numeric",
		month: "2-digit",
		day: "2-digit",
		hour: "2-digit",
		minute: "2-digit",
		second: "2-digit",
		hour12: false,
	});
}

// 将时间戳格式化为相对时间
export function timeAgo(timestamp: number): string {
	const seconds = Math.floor((Date.now() - timestamp) / 1000);

	const intervals = {
		年: 31536000,
		月: 2592000,
		周: 604800,
		天: 86400,
		小时: 3600,
		分钟: 60,
		秒: 1,
	};

	for (const [unit, secondsInUnit] of Object.entries(intervals)) {
		const interval = Math.floor(seconds / secondsInUnit);
		if (interval >= 1) {
			return `${interval}${unit}前`;
		}
	}

	return "刚刚";
}

// 复制文本到剪贴板
export async function copyToClipboard(text: string): Promise<boolean> {
	try {
		await navigator.clipboard.writeText(text);
		return true;
	} catch (err) {
		console.error("Failed to copy text: ", err);
		return false;
	}
}

// 格式化大数字
export function formatNumber(num: number): string {
	if (num >= 1000000) {
		return (num / 1000000).toFixed(1) + "M";
	}
	if (num >= 1000) {
		return (num / 1000).toFixed(1) + "K";
	}
	return num.toString();
}

export function isValidUrl(url: string): boolean {
	try {
		new URL(url);
		return true;
	} catch (_) {
		return false;
	}
}
