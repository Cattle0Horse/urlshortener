"use client";

import { useCallback, useEffect, useState } from "react";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import {
  Pagination,
  PaginationContent,
  PaginationEllipsis,
  PaginationItem,
  PaginationLink,
  PaginationNext,
  PaginationPrevious,
} from "@/components/ui/pagination";
import { base_url } from "@/components/env";
import { useAuth } from "@/components/context";
import Link from "next/link";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { toast } from "sonner";

interface UrlData {
  views: number;
  original_url: string;
  expired_at: string;
  short_url: string;
  id: number;
}

export default function MyUrls() {
  const [urls, setUrls] = useState<UrlData[]>([]);
  const [currentPage, setCurrentPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);
  const [selectedUrl, setSelectedUrl] = useState<UrlData | null>(null);
  const [newExpiryDate, setNewExpiryDate] = useState("");
  const [isUpdateDialogOpen, setIsUpdateDialogOpen] = useState(false);
  const pageSize = 10;
  const { token } = useAuth();

  const fetchUrls = useCallback(
    async (page: number) => {
      try {
        const response = await fetch(
          `${base_url}/api/url?page=${page}&size=${pageSize}`,
          {
            method: "GET",
            headers: {
              Authorization: `Bearer ${token}`,
            },
          }
        );

        const data = await response.json();
        setUrls(data.items || []);
        setTotalPages(Math.ceil((data.total || 0) / pageSize));
      } catch (error) {
        console.error("Failed to fetch URLs:", error);
      }
    },
    [token, pageSize]
  );

  const handleDelete = async (shortUrl: string) => {
    const code = shortUrl.split("/").pop();
    try {
      const response = await fetch(`${base_url}/api/url/${code}`, {
        method: "DELETE",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (response.ok) {
        toast.success("短链接已删除");
        fetchUrls(currentPage);
      } else {
        toast.error("删除失败");
      }
    } catch (error) {
      toast.error("删除失败");
      console.error("Failed to delete URL:", error);
    }
  };

  const handleUpdate = async (shortUrl: string, newExpiredAt: string) => {
    const code = shortUrl.split("/").pop();
    try {
      const response = await fetch(`${base_url}/api/url/${code}`, {
        method: "PATCH",
        headers: {
          Authorization: `Bearer ${token}`,
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          expired_at: new Date(newExpiredAt).toISOString(),
        }),
      });

      if (response.ok) {
        toast.success("过期时间已更新");
        fetchUrls(currentPage);
        setIsUpdateDialogOpen(false);
      } else {
        toast.error("更新失败");
      }
    } catch (error) {
      toast.error("更新失败");
      console.error("Failed to update URL:", error);
    }
  };

  useEffect(() => {
    fetchUrls(currentPage);
  }, [currentPage, fetchUrls]);

  return (
    <div className="container mx-auto py-10">
      <h1 className="text-2xl font-bold mb-6">我的短链接</h1>
      <div className="rounded-md border">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>ID</TableHead>
              <TableHead>短链接</TableHead>
              <TableHead>原始链接</TableHead>
              <TableHead>访问次数</TableHead>
              <TableHead>过期时间</TableHead>
              <TableHead>操作</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {urls.map((url) => (
              <TableRow key={url.id}>
                <TableCell>{url.id}</TableCell>
                <TableCell className="truncate text-sky-500">
                  <Link href={url.short_url} target="_blank">
                    {url.short_url}
                  </Link>
                </TableCell>
                <TableCell className="max-w-[200px] truncate text-sky-500">
                  <Link href={url.original_url} target="_blank">
                    {url.original_url}
                  </Link>
                </TableCell>
                <TableCell>{url.views}</TableCell>
                <TableCell>
                  {new Date(url.expired_at).toLocaleString("zh-CN", {
                    year: "numeric",
                    month: "2-digit",
                    day: "2-digit",
                    hour: "2-digit",
                    minute: "2-digit",
                    second: "2-digit",
                    hour12: false,
                  })}
                </TableCell>
                <TableCell>
                  <div className="flex space-x-2">
                    <Dialog
                      open={isUpdateDialogOpen && selectedUrl?.id === url.id}
                      onOpenChange={(open) => {
                        setIsUpdateDialogOpen(open);
                        if (!open) {
                          setSelectedUrl(null);
                          setNewExpiryDate("");
                        }
                      }}
                    >
                      <DialogTrigger asChild>
                        <Button
                          variant="outline"
                          size="sm"
                          onClick={() => {
                            setSelectedUrl(url);
                            const date = new Date(url.expired_at);
                            const year = date.getFullYear();
                            const month = String(date.getMonth() + 1).padStart(
                              2,
                              "0"
                            );
                            const day = String(date.getDate()).padStart(2, "0");
                            const hours = String(date.getHours()).padStart(
                              2,
                              "0"
                            );
                            const minutes = String(date.getMinutes()).padStart(
                              2,
                              "0"
                            );
                            setNewExpiryDate(
                              `${year}-${month}-${day}T${hours}:${minutes}`
                            );
                          }}
                        >
                          更新
                        </Button>
                      </DialogTrigger>
                      <DialogContent>
                        <DialogHeader>
                          <DialogTitle>更新过期时间</DialogTitle>
                        </DialogHeader>
                        <div className="space-y-4 py-4">
                          <Input
                            type="datetime-local"
                            value={newExpiryDate}
                            onChange={(e) => setNewExpiryDate(e.target.value)}
                          />
                          <Button
                            onClick={() =>
                              handleUpdate(url.short_url, newExpiryDate)
                            }
                          >
                            确认
                          </Button>
                        </div>
                      </DialogContent>
                    </Dialog>
                    <Button
                      variant="destructive"
                      size="sm"
                      onClick={() => handleDelete(url.short_url)}
                    >
                      删除
                    </Button>
                  </div>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>
      <div className="mt-4">
        <Pagination>
          <PaginationContent>
            <PaginationItem>
              <PaginationPrevious
                onClick={() => setCurrentPage((p) => Math.max(1, p - 1))}
              />
            </PaginationItem>

            {/* First page */}
            {totalPages > 0 && (
              <PaginationItem>
                <PaginationLink
                  onClick={() => setCurrentPage(1)}
                  isActive={currentPage === 1}
                >
                  1
                </PaginationLink>
              </PaginationItem>
            )}

            {/* Left ellipsis */}
            {currentPage > 3 && (
              <PaginationItem>
                <PaginationEllipsis />
              </PaginationItem>
            )}

            {/* Middle pages */}
            {Array.from({ length: totalPages }, (_, i) => i + 1)
              .filter((page) => {
                if (totalPages <= 7) return true;
                if (page === 1 || page === totalPages) return false;
                return Math.abs(currentPage - page) <= 1;
              })
              .map((page) => (
                <PaginationItem key={page}>
                  <PaginationLink
                    onClick={() => setCurrentPage(page)}
                    isActive={currentPage === page}
                  >
                    {page}
                  </PaginationLink>
                </PaginationItem>
              ))}

            {/* Right ellipsis */}
            {currentPage < totalPages - 2 && (
              <PaginationItem>
                <PaginationEllipsis />
              </PaginationItem>
            )}

            {/* Last page */}
            {totalPages > 1 && (
              <PaginationItem>
                <PaginationLink
                  onClick={() => setCurrentPage(totalPages)}
                  isActive={currentPage === totalPages}
                >
                  {totalPages}
                </PaginationLink>
              </PaginationItem>
            )}

            <PaginationItem>
              <PaginationNext
                onClick={() =>
                  setCurrentPage((p) => Math.min(totalPages, p + 1))
                }
              />
            </PaginationItem>
          </PaginationContent>
        </Pagination>
      </div>
    </div>
  );
}
