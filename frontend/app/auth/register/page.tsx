"use client";
import { RegisterForm } from "@/components/auth/register-form";

export default function RegisterPage() {
  return (
    <div className="min-h-screen flex items-center justify-center dark:bg-gray-900 bg-gray-50">
      <RegisterForm />
    </div>
  );
}
