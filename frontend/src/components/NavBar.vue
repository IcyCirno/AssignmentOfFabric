<template>
  <nav class="navbar">
    <div class="nav-left">
      <router-link to="/">首页</router-link>
      <router-link to="/market">市场</router-link>
      <router-link v-if="isLoggedIn" to="/mine">副本</router-link>
    </div>
  </nav>
</template>

<script>
export default {
  name: 'NavBar',
  data() {
    return {
      token: localStorage.getItem('token') || ''
    }
  },
  computed: {
    isLoggedIn() {
      return !!this.token
    }
  },
  mounted() {
    // 监听登录状态变化事件
    window.addEventListener('token-changed', this.refreshToken)
  },
  beforeUnmount() {
    window.removeEventListener('token-changed', this.refreshToken)
  },
  methods: {
    refreshToken() {
      this.token = localStorage.getItem('token') || ''
    }
  }
}
</script>

<style scoped>
.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.8rem 2rem;
  background-color: #4588a0;
  color: white;
}

.navbar a {
  color: white;
  text-decoration: none;
  margin-right: 1.5rem;
  font-weight: bold;
}

.navbar a:hover {
  color: #d1f5e9;
}

.nav-left, .nav-right {
  display: flex;
  align-items: center;
}
</style>
