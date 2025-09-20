export async function request(url, options = {}) {
  // 基础配置
  const baseUrl = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api'

  // 从 localStorage 取 token
  const token = localStorage.getItem('token')

  // 默认 headers
  const headers = {
    'Content-Type': 'application/json',
    ...(options.headers || {})
  }

  // 如果有 token，就加到请求头
  if (token) {
    headers['Authorization'] = `Bearer: ${token}`
  }

  // 发送请求
  const response = await fetch(baseUrl + url, {
    ...options,
    headers
  })
  
  // 如果未授权，清除 token 并跳转登录
  if (response.status === 401) {
  localStorage.removeItem('token')
  window.location.href = '/login'
  return
  }


  // 错误处理
  if (!response.ok) {
    const errorText = await response.text()
    throw new Error(errorText || '请求失败')
  }



  // 尝试解析 JSON
  return response.json()
}

export const getCards = async () => {
  return request("/auth/card/query", { method: "POST" });
};

// 销毁卡牌
export const destroyCard = async (hashid) => {
  return request("/auth/card/destroy", {
    method: "POST",
    body: JSON.stringify({ hashid }),
  });
};

// 上架卡牌
export const sellCard = async (hashid, cost) => {
  return request("/auth/card/sell", {
    method: "POST",
    body: JSON.stringify({ hashid, cost }),
  });
};

// 下架卡牌
export const cancelCard = async (orderid) => {
  return request("/auth/card/cancel", {
    method: "POST",
    body: JSON.stringify({ orderid }),
  });
};