<template>
  <div class="login-container">
    <div class="login-box">
      <h1>注册</h1>
      <form @submit.prevent="handleRegister">
        

        <div class="form-group">
          <label for="name">昵称</label>
          <input v-model="name" type="text" id="name" required />
        </div>

        <div class="form-group">
          <label for="password">密码</label>
          <input v-model="password" type="password" id="password" required />
        </div>

        <div class="form-group">
          <label for="confirmPassword">确认密码</label>
          <input v-model="confirmPassword" type="password" id="confirmPassword" required />
        </div>

        <button type="submit" class="btn">注册</button>
      </form>

      <div class="other-option">
        <p>已有账号？ <router-link to="/login">去登录</router-link></p>
      </div>
    </div>
  </div>
</template>

<script>
import { register } from '../api/auth'

export default {
  name: 'Register',
  data() {
    return {
      name: '',
      password: '',
      confirmPassword: ''
    }
  },
  methods: {
    async handleRegister() {
      if (this.password !== this.confirmPassword) {
        alert('两次输入的密码不一致')
        return
      }

      try {
        await register(this.name, this.password)
        alert('注册成功，请登录')
        this.$router.push('/login')
      } catch (error) {
        alert('注册失败: ' + error.message)
      }
    }
  }
}
</script>


<style scoped>
/* 复用 Login.vue 的样式 */
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background: url('/assets/bg-login.jpg') no-repeat center center/cover;
}

.login-box {
  background: white;
  padding: 2rem;
  border-radius: 12px;
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.15);
  width: 350px;
  text-align: center;
}

.login-box h1 {
  margin-bottom: 1.5rem;
  font-size: 1.8rem;
  color: #333;
}

.form-group {
  display: flex;
  flex-direction: column;
  text-align: left;
  margin-bottom: 1rem;
}

.form-group label {
  margin-bottom: 0.3rem;
  font-size: 0.9rem;
  color: #555;
}

.form-group input {
  padding: 0.6rem;
  border: 1px solid #ccc;
  border-radius: 6px;
  outline: none;
}

.form-group input:focus {
  border-color: #42b983;
  box-shadow: 0 0 4px rgba(66, 185, 131, 0.5);
}

.btn {
  width: 100%;
  padding: 0.8rem;
  margin-top: 0.5rem;
  border: none;
  border-radius: 6px;
  background-color: #42b983;
  color: white;
  font-size: 1rem;
  cursor: pointer;
  transition: background 0.3s;
}

.btn:hover {
  background-color: #369870;
}

.other-option {
  margin-top: 1.5rem;
}

.other-option p {
  margin: 1rem 0;
  font-size: 0.9rem;
  color: #777;
}
</style>
