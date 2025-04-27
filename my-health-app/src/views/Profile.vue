<template>
  <div class="profile-container">
    <el-card class="profile-card">
      <h2>个人资料</h2>
      <el-form :model="profileForm" ref="profileFormRef" label-width="100px">
        <el-form-item label="用户名">
          <el-input v-model="profileForm.username" disabled />
        </el-form-item>
        <el-form-item label="邮箱">
          <el-input v-model="profileForm.email" />
        </el-form-item>
        <el-form-item label="手机号">
          <el-input v-model="profileForm.phone" />
        </el-form-item>
        <el-form-item label="头像">
          <el-upload
            class="avatar-uploader"
            action="https://jsonplaceholder.typicode.com/posts/"
            :show-file-list="false"
            :on-success="handleAvatarSuccess"
            :before-upload="beforeAvatarUpload"
          >
            <img v-if="profileForm.avatar" :src="profileForm.avatar" class="avatar" />
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
          </el-upload>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleProfileUpdate">保存修改</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'ProfilePage',
  data() {
    return {
      profileForm: {
        username: 'johndoe', // 模拟初始数据
        email: 'johndoe@example.com',
        phone: '1234567890',
        avatar: '',
      },
    }
  },
  methods: {
    handleProfileUpdate() {
      // 模拟调用 API 更新用户信息
      console.log('更新后的个人资料：', this.profileForm)
      this.$message.success('个人资料已更新')
    },
    handleAvatarSuccess(response, file) {
      // 模拟上传成功后的处理，实际开发中根据返回结果设置图片 URL
      this.profileForm.avatar = URL.createObjectURL(file.raw)
    },
    beforeAvatarUpload(file) {
      const isJPGOrPNG = file.type === 'image/jpeg' || file.type === 'image/png'
      const isLt2M = file.size / 1024 / 1024 < 2

      if (!isJPGOrPNG) {
        this.$message.error('头像上传图片只能是 JPG/PNG 格式!')
      }
      if (!isLt2M) {
        this.$message.error('上传头像图片大小不能超过 2MB!')
      }
      return isJPGOrPNG && isLt2M
    },
  },
}
</script>

<style scoped>
.profile-container {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 50px 0;
}
.profile-card {
  width: 500px;
  padding: 20px;
}
.avatar-uploader {
  display: inline-block;
  width: 100px;
  height: 100px;
  border: 1px dashed #dcdfe6;
  border-radius: 6px;
  cursor: pointer;
  overflow: hidden;
}
.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 100%;
  height: 100%;
  line-height: 100px;
  text-align: center;
}
.avatar {
  width: 100%;
  height: 100%;
  display: block;
}
</style>
