import { API_URL } from './config'

interface RequestOptions extends RequestInit {
  token?: string
}

async function fetchAPI(endpoint: string, options: RequestOptions = {}) {
  const { token, ...restOptions } = options
  const headers = {
    'Content-Type': 'application/json',
    ...(token && { Authorization: `Bearer ${token}` }),
    ...options.headers,
  }

  const response = await fetch(`${API_URL}${endpoint}`, {
    ...restOptions,
    headers,
  })

  if (!response.ok) {
    throw new Error(`API error: ${response.statusText}`)
  }

  return response.json()
}

interface LoginResponse {
  access_token: string
  user_id: number
  email: string
}

export const api = {
  auth: {
    register: (email: string, password: string) => 
      fetchAPI('/api/auth/register', {
        method: 'POST',
        body: JSON.stringify({ email, password }),
      }),

    login: (email: string, password: string): Promise<LoginResponse> => 
      fetchAPI('/api/auth/login', {
        method: 'POST',
        body: JSON.stringify({ email, password }),
      }),
  },

  links: {
    create: (data: { url: string; customSlug?: string }, token: string) =>
      fetchAPI('/api/url', {
        method: 'POST',
        body: JSON.stringify({
          original_url: data.url,
          duration: 24,
          ...(data.customSlug && { custom_code: data.customSlug }),
        }),
        token,
      }),

    list: (page = 1, size = 10, token: string) =>
      fetchAPI(`/api/urls?page=${page}&size=${size}`, { token }),

    update: (id: string, duration: number, token: string) =>
      fetchAPI(`/api/url/${id}`, {
        method: 'PATCH',
        body: JSON.stringify({ duration }),
        token,
      }),

    delete: (id: string, token: string) =>
      fetchAPI(`/api/url/${id}`, {
        method: 'DELETE',
        token,
      }),
  },
} 