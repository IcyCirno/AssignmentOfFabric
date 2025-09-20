<template>
  <div class="profile-info">
    <h3>个人信息</h3>
    <div v-if="loading">加载中...</div>
    <div v-else-if="error" class="error">{{ error }}</div>
    <div v-else>
      <p><strong>用户名：</strong>{{ user.name }}</p>
      <p><strong>余额：</strong>{{ user.gocoin }} GoCoin</p>
      <button @click="$emit('logout')">退出登录</button>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    user: Object,
    loading: Boolean,
    error: String,
  },
};
</script>

<style scoped>
.profile-info {
  flex: 0 0 auto;
}
</style>
