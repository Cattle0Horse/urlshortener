<template>
  <view class="container">
    <text class="title">注册</text>
    <view class="form">
      <input v-model="email" placeholder="邮箱" type="text" />
      <input v-model="password" placeholder="密码" type="password" />
      <input 
        v-model="confirmPassword" 
        placeholder="确认密码" 
        type="password"
      />
      <view class="btn" @click="handleRegister">提交注册</view>
      <navigator url="/pages/login/login" open-type="redirect" class="link">
        已有账号？去登录
      </navigator>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useAuthStore } from '@/stores/auth'

const email = ref('')
const password = ref('')
const confirmPassword = ref('')
const authStore = useAuthStore()

const validatePassword = () => {
  if (password.value !== confirmPassword.value) {
    uni.showToast({
      title: '两次输入的密码不一致',
      icon: 'error',
      duration: 2000
    })
    return false
  }
  return true
}

const handleRegister = async () => {
  if (!validatePassword()) return
  
  try {
    await authStore.register(email.value, password.value)
    uni.showToast({
      title: '注册成功，请登录',
      icon: 'success',
      duration: 2000
    })
    uni.navigateTo({ url: '/pages/login/login' })
  } catch (error: any) {
    uni.showToast({
      title: error.message,
      icon: 'error',
      duration: 2000
    })
  }
}
</script>

<style lang="scss" scoped>
.container {
  padding: 40px 20px;
  min-height: 100vh;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
}

.title {
  font-size: 32px;
  font-weight: bold;
  color: #2c3e50;
  text-align: center;
  margin-bottom: 40px;
  position: relative;
}

.title::after {
  content: "";
  position: absolute;
  bottom: -10px;
  left: 50%;
  transform: translateX(-50%);
  width: 60px;
  height: 4px;
  background: #3498db;
}

.form {
  max-width: 400px;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  gap: 20px;
  background: white;
  padding: 30px;
  border-radius: 15px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

input {
  padding: 12px 16px;
  border: 2px solid #e0e0e0;
  border-radius: 8px;
  font-size: 16px;
  transition: all 0.2s ease;
}

input:focus {
  border-color: #3498db;
  box-shadow: 0 0 0 3px rgba(52, 152, 219, 0.2);
}

.btn {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 14px 24px;
  background: #3498db;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 16px;
  transition: all 0.2s ease;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 240px;
  margin: 0 auto;
}

.link {
  margin-top: 20px;
  text-align: center;
  color: #3498db;
  text-decoration: none;
  font-size: 14px;
  transition: color 0.2s ease;
}

.link:hover {
  color: #2980b9;
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

.form {
  animation: fadeIn 0.5s ease forwards;
  opacity: 0;
}
</style>
