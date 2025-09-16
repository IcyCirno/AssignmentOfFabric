<template>
  <div class="home-page">

    <!-- 主体内容 -->
    <main class="content">
      <h1>欢迎来到区块链电子卡片交易平台</h1>
      <p>在这里你可以创建、购买和交易独特的数字卡片。</p>
    </main>
  </div>
</template>

<script>
export default {
  name: 'Home',
  data() {
    return {
      isLoggedIn: false
    }
  },
  mounted() {
    this.checkLogin()
  },
  methods: {
    checkLogin() {
      this.isLoggedIn = !!localStorage.getItem('token')
    }
  }
}
</script>

<style scoped>
/* 整体布局 */
.home-page {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

/* 顶部导航栏 */
.navbar {
  display: flex;
  justify-content: space-between; /* 左右分布 */
  align-items: center;
  background-color: #1e3a8a;
  padding: 1rem 2rem;
}

/* 左侧导航 */
.nav-left {
  display: flex;
  gap: 1.5rem;
}

/* 右侧导航 */
.nav-right {
  display: flex;
  gap: 1.5rem;
}

/* 导航项样式 */
.nav-item {
  color: white;
  text-decoration: none;
  font-weight: 500;
}

.nav-item:hover {
  text-decoration: underline;
}

/* 主体内容 */
.content {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  text-align: center;
  padding: 2rem;
}
</style>
