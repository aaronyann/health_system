<template>
  <div class="medical-image">
    <el-card>
      <h2>医疗影像展示</h2>
      <el-row :gutter="20">
        <el-col :span="8" v-for="(img, index) in images" :key="index">
          <el-image
            style="width: 100%; height: 150px"
            :src="img.url"
            :preview-src-list="previewList"
          >
          </el-image>
          <div style="text-align: center; margin-top: 5px">{{ img.description }}</div>
        </el-col>
      </el-row>
      <div style="margin-top: 20px">
        <el-button type="primary" @click="goBack">返回列表</el-button>
      </div>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'MedicalImage',
  data() {
    return {
      images: [
        { url: 'https://via.placeholder.com/300x150?text=Image1', description: 'X光图像' },
        { url: 'https://via.placeholder.com/300x150?text=Image2', description: 'CT扫描图像' },
        { url: 'https://via.placeholder.com/300x150?text=Image3', description: 'MRI图像' },
      ],
      previewList: [],
    }
  },
  methods: {
    goBack() {
      // 根据项目路由调整返回页面名称
      this.$router.push({ name: 'LabReport' })
    },
  },
  mounted() {
    this.previewList = this.images.map((img) => img.url)
  },
}
</script>

<style scoped>
.medical-image {
  padding: 20px;
}
</style>
