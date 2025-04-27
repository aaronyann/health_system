<template>
  <div class="register-container">
    <el-card class="register-card">
      <h2>用户注册</h2>
      <el-form :model="registerForm" ref="registerFormRef" label-width="80px">
        <el-form-item label="用户名">
          <el-input v-model="registerForm.username" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="邮箱">
          <el-input v-model="registerForm.email" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="密码">
          <el-input v-model="registerForm.password" type="password" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="确认密码">
          <el-input
            v-model="registerForm.confirmPassword"
            type="password"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleRegister">注册</el-button>
        </el-form-item>
      </el-form>
      <p>已有账号？ <router-link to="/login">立即登录</router-link></p>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'IndexRegister',
  data() {
    return {
      registerForm: {
        username: '',
        email: '',
        password: '',
        confirmPassword: '',
      },
    }
  },
  methods: {
    handleRegister() {
      console.log('注册信息：', this.registerForm)
      // 在此调用 API 执行注册逻辑
    },
  },
}
</script>

<style scoped>
.register-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}
.register-card {
  width: 400px;
  padding: 20px;
}
</style>
