import Cookies from 'js-cookie'

export function getCookie(name: string): string {
  return Cookies.get(name) || ''
}

export function setCookie(name: string, value: string, options = {}) {
  Cookies.set(name, value, {
    ...options,
    secure: process.env.NODE_ENV === 'production',
    sameSite: 'lax',
  })
}

export function removeCookie(name: string) {
  Cookies.remove(name)
} 