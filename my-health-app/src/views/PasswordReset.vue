<template>
  <div class="password-reset-container">
    <el-card class="password-reset-card">
      <h2>重置密码</h2>
      <el-form :model="resetForm" ref="resetFormRef" label-width="100px">
        <el-form-item label="邮箱">
          <el-input v-model="resetForm.email" placeholder="请输入注册邮箱" autocomplete="off" />
        </el-form-item>
        <el-form-item label="新密码">
          <el-input
            v-model="resetForm.newPassword"
            type="password"
            placeholder="请输入新密码"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item label="确认密码">
          <el-input
            v-model="resetForm.confirmPassword"
            type="password"
            placeholder="请再次输入新密码"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item label="验证码">
          <el-input
            v-model="resetForm.verificationCode"
            placeholder="请输入验证码"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handlePasswordReset"> 重置密码 </el-button>
        </el-form-item>
      </el-form>
      <p>记起密码了？ <router-link to="/login">返回登录</router-link></p>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'PasswordReset',
  data() {
    return {
      resetForm: {
        email: '',
        newPassword: '',
        confirmPassword: '',
        verificationCode: '',
      },
    }
  },
  methods: {
    handlePasswordReset() {
      // 验证两次输入的密码是否一致
      if (this.resetForm.newPassword !== this.resetForm.confirmPassword) {
        this.$message.error('两次输入的密码不一致')
        return
      }
      // 此处模拟调用 API 重置密码（例如使用 axios 发送请求）
      console.log('重置密码信息：', this.resetForm)
      // 重置成功后显示提示并跳转登录页面
      this.$message.success('密码重置成功，请登录')
      this.$router.push({ name: 'Login' })
    },
  },
}
</script>

<style scoped>
.password-reset-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}
.password-reset-card {
  width: 400px;
  padding: 20px;
}
</style>
