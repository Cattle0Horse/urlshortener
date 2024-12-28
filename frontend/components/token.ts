"use client";
import { jwtDecode } from "jwt-decode";

export function isTokenExpired(token: string | null) {
  try {
    if (!token) {
      return true;
    }

    const pay = jwtDecode(token);

    if (!pay) {
      return true;
    }
    const currentTime = Math.floor(Date.now() / 1000); // s
    if (pay.exp && pay.exp > currentTime) {
      return false;
    }

    return true;
  } catch {
    return true;
  }
}
