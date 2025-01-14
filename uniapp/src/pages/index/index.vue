<template>
  <view class="container">
    <view class="header">
      <image class="logo" src="/static/logo.svg" />
      <text class="title">短链接生成器</text>
      <text class="subtitle">轻松创建和管理短链接</text>
    </view>
    <view class="nav-buttons">
      <view
        v-if="authStore.isAuthenticated"
        class="btn"
        @click="handleLogout"
      >
        <uni-icons type="logout" size="18" color="#fff"></uni-icons>
        <text>登出</text>
    </view>
      <navigator
        v-else
        url="/pages/login/login"
        open-type="navigate"
        class="btn"
      >
        <uni-icons type="person" size="18" color="#fff"></uni-icons>
        <text>登录/注册</text>
      </navigator>
      <navigator url="/pages/urls/urls" open-type="navigate" class="btn">
        <uni-icons type="list" size="18" color="#fff"></uni-icons>
        <text>短链接管理</text>
      </navigator>
    </view>
  </view>
</template>

<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import { onShow } from '@dcloudio/uni-app';

const authStore = useAuthStore();

onShow(() => {
  // 确保页面显示时样式正确
  document.body.style.overflow = 'hidden';
});

const handleLogout = () => {
  authStore.logout();
  uni.reLaunch({
    url: "/pages/login/login",
  });
};
</script>

<style lang="scss" scoped>
.container {
  padding: 40px 20px;
  min-height: 100vh;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  position: relative;
  z-index: 1;
}

.header {
  text-align: center;
  margin-bottom: 40px;
}

.logo {
  height: 120px;
  width: 120px;
  margin-bottom: 20px;
}

.title {
  font-size: 32px;
  font-weight: bold;
  color: #2c3e50;
  margin: 20px 0 10px 0;
  line-height: 1.2;
  display: block;
}

.subtitle {
  font-size: 16px;
  color: #666;
  line-height: 1.6;
  max-width: 280px;
  margin: 10px auto 0;
  display: block;
}

.nav-buttons {
  width: 100%;
  max-width: 320px;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.btn {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  padding: 14px 24px;
  background: #3498db;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 16px;
  cursor: pointer;
  transition: all 0.2s ease;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 280px;
  margin: 0 auto;
}

.btn:hover {
  transform: translateY(-1px);
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.15);
}

/* 添加动画效果 */
@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.nav-buttons {
  animation: fadeIn 0.5s ease forwards;
}
</style>
