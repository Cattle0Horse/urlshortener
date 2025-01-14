import axios from "axios";

export interface ApiResponse<T = any> {
  code: number;
  msg: string;
  data?: T;
}

export interface LoginResponse {
  access_token: string;
  user_id: number;
  email: string;
}

export interface CreateUrlResponse {
  short_code: string;
}

export interface UrlItem {
  id: number;
  short_code: string;
  original_url: string;
  expiry_time: string;
  created_at: string;
}

export interface UrlListResponse {
  total: number;
  urls: UrlItem[];
}

const api = axios.create({
  baseURL: "http://localhost:8080/api",
  headers: {
    "Content-Type": "application/json",
  },
});

// 添加请求拦截器
api.interceptors.request.use((config) => {
  const token = localStorage.getItem("token");
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export default {
  // 用户认证
  async register(email: string, password: string) {
    try {
      const response = await api.post("/auth/register", { email, password });
      return response.data;
    } catch (error: any) {
      if (error.response?.data?.code === 40003) {
        throw new Error("该邮箱已被注册");
      }
      throw error;
    }
  },

  async login(email: string, password: string): Promise<LoginResponse> {
    try {
      const { data: apidata } = await api.post<ApiResponse<LoginResponse>>(
        "/auth/login",
        { email, password }
      );
      const data: LoginResponse = apidata.data!;
      localStorage.setItem("token", data.access_token);
      return data;
    } catch (error: any) {
      if (error.response?.data?.code === 40004) {
        throw new Error("密码错误");
      } else if (error.response?.data?.code === 40002) {
        throw new Error("用户不存在");
      }
      throw error;
    }
  },

  // 短链接管理
  async createUrl(
    originalUrl: string,
    duration: number
  ): Promise<CreateUrlResponse> {
    try {
      const { data: apidata } = await api.post<ApiResponse<CreateUrlResponse>>(
        "/url",
        { original_url: originalUrl, duration }
      );
      return apidata.data!;
    } catch (error: any) {
      if (error.response?.data?.code === 40003) {
        throw new Error("短链接已存在");
      }
      throw error;
    }
  },

  async getUrls(page: number = 1, size: number = 10): Promise<UrlListResponse> {
    const { data: apidata } = await api.get<ApiResponse<UrlListResponse>>(
      `/urls?page=${page}&size=${size}`
    );
    return apidata.data!;
  },

  async updateUrl(shortCode: string, duration: number) {
    try {
      await api.patch(`/url/${shortCode}`, { duration });
    } catch (error: any) {
      if (error.response?.data?.code === 40002) {
        throw new Error("短链接不存在");
      }
      throw error;
    }
  },

  async deleteUrl(shortCode: string) {
    try {
      await api.delete(`/url/${shortCode}`);
    } catch (error: any) {
      if (error.response?.data?.code === 40002) {
        throw new Error("短链接不存在");
      }
      throw error;
    }
  },
};
