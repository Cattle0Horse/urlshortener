"use client";
import { LogOutIcon } from "lucide-react";
import { useAuth } from "./context";
import { Button } from "./ui/button";
import { useRouter } from "next/navigation";

export default function Logout() {
  const { setAuth } = useAuth();
  const router = useRouter();

  return (
    <Button
      variant={"outline"}
      onClick={() => {
        setAuth("", "", 0);
        router.push("/auth/login");
      }}
      className="h-8 w-8"
    >
      <LogOutIcon className="" />
    </Button>
  );
}
