<template>
  <view class="container">
    <view class="header">
      <text class="title">我的短链接</text>
    </view>
    <scroll-view
      class="list"
      scroll-y
      @scrolltolower="loadMore"
      :style="{ height: 'calc(100vh - 200px)' }"
    >
      <view v-for="url in urls" :key="url.short_code" class="item">
        <text class="original">{{ url.original_url }}</text>
        <text class="short">短链接：{{ host }}/{{ url.short_code }}</text>
        <view class="actions">
          <button @click="copyLink(url.short_code)">复制</button>
          <button @click="extendExpiry(url.short_code)">更新有效期</button>
          <button @click="deleteUrl(url.short_code)">删除</button>
        </view>
      </view>
      <view v-if="loading" class="loading">加载中...</view>
      <view v-if="!hasMore" class="no-more">没有更多了</view>
    </scroll-view>
    <navigator url="/pages/url/create" open-type="navigate" class="btn">
      创建新短链接
    </navigator>
  </view>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import api, { type UrlItem } from "@/services/api";

const host = "http://localhost:8080";
const urls = ref<UrlItem[]>([]);
const currentPage = ref(1);
const pageSize = 10;
const total = ref(0);
const loading = ref(false);
const hasMore = ref(true);

const loadMore = async () => {
  if (loading.value || !hasMore.value) return;

  loading.value = true;
  try {
    currentPage.value++;
    const { urls: newUrls, total: totalCount } = await api.getUrls(
      currentPage.value,
      pageSize
    );

    if (newUrls.length > 0) {
      urls.value = [...urls.value, ...newUrls];
    } else {
      hasMore.value = false;
    }

    if (urls.value.length >= totalCount) {
      hasMore.value = false;
    }
  } catch (error) {
    console.error("加载更多失败:", error);
    uni.showToast({
      title: "加载更多失败",
      icon: "none",
    });
  } finally {
    loading.value = false;
  }
};

const loadUrls = async (page: number = 1) => {
  try {
    const { urls: urlList, total: totalCount } = await api.getUrls(
      page,
      pageSize
    );
    urls.value = urlList;
    total.value = totalCount;
  } catch (error) {
    console.error("获取短链接失败:", error);
    uni.showToast({
      title: "获取短链接失败",
      icon: "error",
      duration: 2000,
    });
  }
};

onMounted(() => {
  loadUrls();
});

const copyLink = (shortCode: string) => {
  uni.setClipboardData({
    data: `${host}/${shortCode}`,
    success: () => {
      uni.showToast({
        title: "已复制到剪贴板",
        icon: "success",
        duration: 2000,
      });
    },
  });
};

const extendExpiry = async (shortCode: string) => {
  uni.showModal({
    title: "更新有效期",
    editable: true,
    placeholderText: "请输入新的有效期（0-168小时，0表示立即过期）",
    success: async (res) => {
      if (res.confirm) {
        const hours = parseInt(res.content || "0");
        if (isNaN(hours) || hours < 0 || hours > 168) {
          uni.showToast({
            title: "请输入0-168之间的有效数字",
            icon: "none",
          });
          return;
        }

        try {
          await api.updateUrl(shortCode, hours);
          uni.showToast({
            title: `有效期已更新为${hours}小时`,
            icon: "success",
            duration: 2000,
          });
          loadUrls(); // 刷新列表
        } catch (error) {
          console.error("更新有效期失败:", error);
          uni.showToast({
            title: "更新有效期失败",
            icon: "error",
            duration: 2000,
          });
        }
      }
    },
  });
};

const deleteUrl = async (shortCode: string) => {
  uni.showModal({
    title: "确认删除",
    content: "确定要删除这个短链接吗？",
    success: async (res) => {
      if (res.confirm) {
        try {
          await api.deleteUrl(shortCode);
          urls.value = urls.value.filter((url) => url.short_code !== shortCode);
          uni.showToast({
            title: "删除成功",
            icon: "success",
            duration: 2000,
          });
        } catch (error) {
          console.error("删除失败:", error);
          uni.showToast({
            title: "删除失败",
            icon: "error",
            duration: 2000,
          });
        }
      }
    },
  });
};
</script>

<style lang="scss">
.container {
  padding: 20px;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  min-height: 100vh;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 30px;
  padding: 20px;
  background: white;
  border-radius: 15px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
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

.list {
  margin-bottom: 20px;
}

.item {
  padding: 20px;
  background: white;
  border-radius: 15px;
  margin-bottom: 15px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  transition: transform 0.2s ease;
}

.item:hover {
  transform: translateY(-2px);
}

.original {
  display: block;
  color: #666;
  margin-bottom: 10px;
  font-size: 14px;
  word-break: break-all;
}

.short {
  display: block;
  color: #3498db;
  margin-bottom: 15px;
  font-size: 16px;
  font-weight: 500;
}

.actions {
  display: flex;
  gap: 10px;
}

.actions button {
  flex: 1;
  padding: 8px;
  font-size: 14px;
  border: none;
  border-radius: 8px;
  background: #3498db;
  color: white;
  cursor: pointer;
  transition: all 0.2s ease;
}

.actions button:hover {
  opacity: 0.9;
  transform: translateY(-1px);
}

.actions button:nth-child(2) {
  background: #2ecc71;
}

.actions button:nth-child(3) {
  background: #e74c3c;
}

.btn {
  position: fixed;
  bottom: 20px;
  left: 50%;
  transform: translateX(-50%);
  padding: 14px 24px;
  background: #3498db;
  color: white;
  border-radius: 8px;
  text-align: center;
  font-size: 16px;
  text-decoration: none;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  transition: all 0.2s ease;
  width: 90%;
  max-width: 320px;
  z-index: 100;
}

.btn:hover {
  transform: translateY(-1px);
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.15);
}

.loading,
.no-more {
  text-align: center;
  color: #666;
  padding: 20px 0;
  margin-bottom: 80px; /* 增加底部间距避免被按钮遮挡 */
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

.item {
  animation: fadeIn 0.5s ease forwards;
  opacity: 0;
}

.item:nth-child(1) {
  animation-delay: 0.1s;
}
.item:nth-child(2) {
  animation-delay: 0.2s;
}
.item:nth-child(3) {
  animation-delay: 0.3s;
}
.item:nth-child(4) {
  animation-delay: 0.4s;
}
.item:nth-child(5) {
  animation-delay: 0.5s;
}
</style>
