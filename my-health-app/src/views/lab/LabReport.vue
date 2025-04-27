<template>
  <div class="lab-report">
    <el-card>
      <h2>检验报告详情</h2>
      <el-descriptions title="检验报告" :column="1" border>
        <el-descriptions-item label="报告ID">{{ report.id }}</el-descriptions-item>
        <el-descriptions-item label="预约ID">{{ report.appointmentId }}</el-descriptions-item>
        <el-descriptions-item label="患者姓名">{{ report.patientName }}</el-descriptions-item>
        <el-descriptions-item label="检验机构">{{ report.labName }}</el-descriptions-item>
        <el-descriptions-item label="检验项目">{{ report.testType }}</el-descriptions-item>
        <el-descriptions-item label="检验结果">{{ report.results }}</el-descriptions-item>
        <el-descriptions-item label="报告文件">
          <el-link type="primary" :underline="false" @click="downloadReport">点击下载</el-link>
        </el-descriptions-item>
        <el-descriptions-item label="报告文件哈希">{{ report.resultsHash }}</el-descriptions-item>
        <el-descriptions-item label="发布时间">{{ report.issuedTime }}</el-descriptions-item>
      </el-descriptions>
      <div style="margin-top: 20px; text-align: right">
        <el-button type="primary" @click="goBack">返回报告列表</el-button>
      </div>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'LabReport',
  data() {
    return {
      report: {},
    }
  },
  methods: {
    fetchReport(id) {
      // 模拟获取检验报告详情数据，实际项目中请替换为 API 调用
      const reports = [
        {
          id: 1,
          appointmentId: 101,
          patientName: 'Alice',
          labName: 'ABC 实验室',
          testType: '血常规',
          results: '所有指标正常',
          resultsHash: 'QmReportHash1',
          issuedTime: '2025-04-15 10:00:00',
          fileUrl: 'https://via.placeholder.com/500x700?text=LabReport1.pdf',
        },
        {
          id: 2,
          appointmentId: 102,
          patientName: 'Bob',
          labName: 'XYZ 实验室',
          testType: '尿常规',
          results: '发现异常指标，请复查',
          resultsHash: 'QmReportHash2',
          issuedTime: '2025-04-16 11:30:00',
          fileUrl: 'https://via.placeholder.com/500x700?text=LabReport2.pdf',
        },
      ]
      this.report = reports.find((r) => r.id == id) || {}
    },
    goBack() {
      // 假设报告列表页面的路由名称为 'LabReportList'
      this.$router.push({ name: 'LabReportList' })
    },
    downloadReport() {
      // 模拟下载报告文件逻辑
      window.open(this.report.fileUrl, '_blank')
    },
  },
  mounted() {
    const reportId = this.$route.params.id
    this.fetchReport(reportId)
  },
}
</script>

<style scoped>
.lab-report {
  padding: 20px;
}
</style>
