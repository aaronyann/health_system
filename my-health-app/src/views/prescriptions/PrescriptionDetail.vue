<template>
  <div class="prescription-detail">
    <el-card>
      <h2>处方详情</h2>
      <el-descriptions title="处方信息" :column="1">
        <el-descriptions-item label="处方ID">{{ prescription.id }}</el-descriptions-item>
        <el-descriptions-item label="患者姓名">{{ prescription.patientName }}</el-descriptions-item>
        <el-descriptions-item label="医生姓名">{{ prescription.doctorName }}</el-descriptions-item>
        <el-descriptions-item label="处方日期">{{ prescription.date }}</el-descriptions-item>
        <el-descriptions-item label="药物信息">
          <ul>
            <li v-for="(med, index) in prescription.medications" :key="index">
              {{ med.name }} - {{ med.dosage }} - {{ med.frequency }}
            </li>
          </ul>
        </el-descriptions-item>
        <el-descriptions-item label="用药说明">{{
          prescription.instructions
        }}</el-descriptions-item>
      </el-descriptions>
      <div style="margin-top: 20px">
        <el-button type="primary" @click="goBack">返回列表</el-button>
      </div>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'PrescriptionDetail',
  data() {
    return {
      prescription: {},
    }
  },
  methods: {
    fetchPrescription(id) {
      // 模拟获取处方详情数据，实际开发中请替换为 API 调用
      const prescriptions = [
        {
          id: 1,
          patientName: 'Alice',
          doctorName: 'Dr. Smith',
          date: '2025-04-10',
          medications: [
            { name: 'Paracetamol', dosage: '500mg', frequency: '3次/天' },
            { name: 'Ibuprofen', dosage: '200mg', frequency: '2次/天' },
          ],
          instructions: '饭后服用',
        },
        {
          id: 2,
          patientName: 'Bob',
          doctorName: 'Dr. Johnson',
          date: '2025-04-11',
          medications: [{ name: 'Amoxicillin', dosage: '250mg', frequency: '3次/天' }],
          instructions: '按时服药',
        },
      ]
      this.prescription = prescriptions.find((p) => p.id == id) || {}
    },
    goBack() {
      this.$router.push({ name: 'PrescriptionList' })
    },
  },
  mounted() {
    const prescriptionId = this.$route.params.id
    this.fetchPrescription(prescriptionId)
  },
}
</script>

<style scoped>
.prescription-detail {
  padding: 20px;
}
</style>
