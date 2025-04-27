<template>
  <div class="lab-report-list">
    <el-card>
      <h2>检验报告列表</h2>
      <el-table :data="reports" style="width: 100%">
        <el-table-column prop="id" label="报告ID" width="80"></el-table-column>
        <el-table-column prop="patientName" label="患者姓名"></el-table-column>
        <el-table-column prop="labName" label="检验机构"></el-table-column>
        <el-table-column prop="testType" label="检验项目"></el-table-column>
        <el-table-column prop="issuedTime" label="发布时间"></el-table-column>
        <el-table-column label="操作" width="120">
          <template #default="scope">
            <el-button type="primary" size="small" @click="viewDetail(scope.row.id)"
              >详情</el-button
            >
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px; text-align: right">
        <el-pagination
          background
          layout="prev, pager, next"
          :total="reports.length"
          @current-change="handlePageChange"
        >
        </el-pagination>
      </div>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'LabReportList',
  data() {
    return {
      reports: [
        {
          id: 1,
          appointmentId: 101,
          patientName: 'Alice',
          labName: 'ABC 实验室',
          testType: '血常规',
          issuedTime: '2025-04-15 10:00:00',
          fileUrl: 'https://via.placeholder.com/500x700?text=LabReport1.pdf',
          resultsHash: 'QmReportHash1',
        },
        {
          id: 2,
          appointmentId: 102,
          patientName: 'Bob',
          labName: 'XYZ 实验室',
          testType: '尿常规',
          issuedTime: '2025-04-16 11:30:00',
          fileUrl: 'https://via.placeholder.com/500x700?text=LabReport2.pdf',
          resultsHash: 'QmReportHash2',
        },
      ],
      currentPage: 1,
      pageSize: 10,
    }
  },
  methods: {
    viewDetail(id) {
      this.$router.push({ name: 'LabReportDetail', params: { id } })
    },
    handlePageChange(page) {
      this.currentPage = page
      // 根据需要调用 API 获取对应页数据
    },
  },
}
</script>

<style scoped>
.lab-report-list {
  padding: 20px;
}
</style>
