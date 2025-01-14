import { createSSRApp } from "vue";
import { createPinia } from 'pinia';
import App from "./App.vue";
import { useAuthStore } from '@/stores/auth'

export function createApp() {
  const app = createSSRApp(App);
  
  // 配置Pinia
  const pinia = createPinia()
  app.use(pinia)

  // 添加路由守卫
  const authStore = useAuthStore()
  
  // 需要登录的页面路径
  const authPages = [
    '/pages/urls/urls',
    '/pages/url/create'
  ]

  // 页面跳转拦截
  uni.addInterceptor('navigateTo', {
    invoke(args) {
      const url = args.url.split('?')[0]
      if (authPages.includes(url) && !authStore.isAuthenticated) {
        uni.showToast({
          title: '请先登录',
          icon: 'none'
        })
        uni.navigateTo({
          url: '/pages/login/login'
        })
        return false
      }
      return true
    }
  })

  uni.addInterceptor('switchTab', {
    invoke(args) {
      const url = args.url.split('?')[0]
      if (authPages.includes(url) && !authStore.isAuthenticated) {
        uni.showToast({
          title: '请先登录',
          icon: 'none'
        })
        uni.navigateTo({
          url: '/pages/login/login'
        })
        return false
      }
      return true
    }
  })

  return {
    app,
  };
}
