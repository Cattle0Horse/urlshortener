import { defineStore } from 'pinia'
import api from '@/services/api'

interface User {
  email: string
  token: string
  userId: number
}

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: JSON.parse(localStorage.getItem('user') || 'null') as User | null
  }),
  getters: {
    isAuthenticated: (state) => !!state.user?.token
  },
  actions: {
    async register(email: string, password: string) {
      try {
        await api.register(email, password)
      } catch (error) {
        throw error
      }
    },
    
    async login(email: string, password: string) {
      try {
        const { access_token, user_id, email: userEmail } = await api.login(email, password)
        const user = {
          email: userEmail,
          token: access_token,
          userId: user_id
        }
        this.user = user
        localStorage.setItem('user', JSON.stringify(user))
      } catch (error) {
        throw error
      }
    },

    logout() {
      this.user = null
      localStorage.removeItem('user')
    }
  }
})
