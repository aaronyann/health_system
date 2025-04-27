<template>
  <div class="login-container">
    <el-card class="login-card">
      <h2>用户登录</h2>
      <el-form :model="loginForm" ref="loginFormRef" label-width="80px">
        <el-form-item label="用户名">
          <el-input v-model="loginForm.username" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="密码">
          <el-input v-model="loginForm.password" type="password" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleLogin">登录</el-button>
        </el-form-item>
      </el-form>
      <p>没有账号？ <router-link to="/register">立即注册</router-link></p>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'LoginPage',
  data() {
    return {
      loginForm: {
        username: '',
        password: '',
      },
    }
  },
  methods: {
    handleLogin() {
      // 此处仅做模拟，实际登录应调用 API
      console.log('登录信息：', this.loginForm)
      // 模拟登录成功后将状态保存在 localStorage
      localStorage.setItem('isAuthenticated', 'true')
      // 登录成功后跳转到首页
      this.$router.push({ name: 'Home' })
    },
  },
}
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}
.login-card {
  width: 400px;
  padding: 20px;
}
</style>
