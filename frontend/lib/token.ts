import { jwtDecode } from "jwt-decode"

interface JWTPayload {
  exp: number
  iat: number
  user_id: number
}

export function isTokenExpired(token: string): boolean {
  if (!token) return true

  try {
    const decoded = jwtDecode<JWTPayload>(token)
    const currentTime = Date.now() / 1000

    return decoded.exp < currentTime
  } catch {
    return true
  }
} 