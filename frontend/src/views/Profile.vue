<template>
  <div class="profile-page">
    <h2>个人中心</h2>

    <div v-if="loading" class="loading">加载中...</div>
    <div v-else-if="error" class="error">{{ error }}</div>

    <div v-else class="profile-card">
      <p><strong>用户名：</strong>{{ user.name }}</p>
      <p><strong>账户金额：</strong>{{ user.gocoin }} GoCoin</p>
      <button @click="handleLogout">退出登录</button>
    </div>
  </div>
</template>

<script>
import { getProfile, logout } from '../api/auth'

export default {
  name: 'Profile',
  data() {
    return {
      user: {},
      loading: true,
      error: ''
    }
  },
  async mounted() {
    await this.fetchProfile()
  },
  methods: {
    async fetchProfile() {
      try {
        const token = localStorage.getItem('token')
        if (!token) {
          // 未登录自动跳转
          window.location.href = '/auth'
          return
        }

        const res = await getProfile()
        if (res.code === 200) {
          this.user = res.data
        } else {
          this.error = res.msg || '加载失败'
        }
      } catch (err) {
        this.error = '服务器请求失败'
        console.error(err)
      } finally {
        this.loading = false
      }
    },
    handleLogout() {
      logout(this.$router)
    }
  }
}
</script>

<style scoped>
.profile-page {
  max-width: 600px;
  margin: 2rem auto;
  padding: 2rem;
  background: #f9fafb;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.05);
  text-align: center;
}

h2 {
  margin-bottom: 1.5rem;
  color: #1e3a8a;
}

.profile-card {
  font-size: 1.1rem;
  line-height: 2rem;
}

button {
  margin-top: 1rem;
  padding: 0.6rem 1.2rem;
  background-color: #42b983;
  border: none;
  border-radius: 5px;
  color: white;
  cursor: pointer;
}

button:hover {
  background-color: #369870;
}

.loading {
  color: #666;
}

.error {
  color: red;
  font-weight: bold;
}
</style>
