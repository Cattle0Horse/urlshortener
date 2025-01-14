<template>
  <view class="container">
    <view class="header">
      <text class="title">创建短链接</text>
      <text class="subtitle">将长链接转换为短链接</text>
    </view>
    <view class="form">
      <view class="input-group">
        <text class="label">原始链接（不超过255个字符）</text>
        <input
          v-model="originalUrl"
          placeholder="https://example.com/very-long-url"
          type="url"
          class="input"
        />
      </view>
      <view class="input-group">
        <text class="label">有效期（1-168小时）</text>
        <view class="duration-input">
          <input
            v-model.number="durationHours"
            type="number"
            min="1"
            max="168"
            placeholder="1-168"
            class="input"
          />
          <text class="unit">小时</text>
        </view>
      </view>
      <button
        class="create-btn"
        :class="{ disabled: !originalUrl }"
        @click="createUrl"
      >
        <text>创建短链接</text>
        <uni-icons type="forward" size="20" color="#fff"></uni-icons>
      </button>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref } from "vue";
import api from "@/services/api";

const originalUrl = ref("");
const durationHours = ref(24);
const urlPattern =
  /^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?(\?[^\s]*)?$/;

const createUrl = async () => {
  if (!originalUrl.value) {
    uni.showToast({ title: "请输入原始链接", icon: "none" });
    return;
  }

  // URL格式验证
  if (originalUrl.value.length <= 255 && !urlPattern.test(originalUrl.value)) {
    uni.showToast({ title: "请输入有效的URL地址", icon: "none" });
    return;
  }

  // 有效期验证
  if (durationHours.value < 1 || durationHours.value > 168) {
    uni.showToast({ title: "有效期需在1-168小时之间", icon: "none" });
    return;
  }

  try {
    await api.createUrl(originalUrl.value, durationHours.value);
    uni.showToast({ title: "创建成功" });
    uni.navigateBack();
  } catch (error) {
    console.error("创建失败:", error);
    uni.showToast({ title: "创建失败", icon: "none" });
  }
};
</script>

<style lang="scss">
.container {
  padding: 20px;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  min-height: 100vh;
}

.header {
  margin-bottom: 30px;
}

.title {
  font-size: 28px;
  font-weight: bold;
  color: #2c3e50;
  position: relative;
}

.title::after {
  content: "";
  position: absolute;
  bottom: -5px;
  left: 0;
  width: 50px;
  height: 3px;
  background: #3498db;
}

.subtitle {
  display: block;
  margin-top: 8px;
  color: #666;
  font-size: 14px;
}

.form {
  display: flex;
  flex-direction: column;
  gap: 20px;
  background: white;
  padding: 24px;
  border-radius: 15px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.input-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.label {
  font-size: 14px;
  color: #666;
  font-weight: 500;
}

.input {
  padding: 12px;
  border: 1px solid #ddd;
  border-radius: 8px;
  font-size: 14px;
  transition: all 0.2s ease;
  width: 100%;
}

.input:focus {
  border-color: #3498db;
  box-shadow: 0 0 0 3px rgba(52, 152, 219, 0.2);
}

.duration-input {
  display: flex;
  align-items: center;
  gap: 10px;
}

.duration-input .input {
  flex: 1;
  max-width: 120px;
}

.unit {
  color: #666;
  font-size: 14px;
}

.create-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  padding: 14px;
  background: #3498db;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 16px;
  cursor: pointer;
  transition: all 0.2s ease;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  margin-top: 10px;
}

.create-btn.disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.create-btn:hover:not(.disabled) {
  transform: translateY(-1px);
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.15);
}

/* 添加动画效果 */
@keyframes slideIn {
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
  animation: slideIn 0.5s ease forwards;
}
</style>
