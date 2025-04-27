import { createRouter, createWebHistory } from 'vue-router'

// 引入页面组件
import Home from '../views/Home.vue'
import Login from '../views/Login.vue'
import Register from '../views/Register.vue'
import PasswordReset from '../views/PasswordReset.vue'
import Profile from '../views/Profile.vue'
import PatientDashboard from '../views/dashboard/PatientDashboard.vue'
import DoctorDashboard from '../views/dashboard/DoctorDashboard.vue'
import AdminDashboard from '../views/dashboard/AdminDashboard.vue'
import HealthRecords from '../views/health/HealthRecords.vue'
import HealthRecordDetail from '../views/health/HealthRecordDetail.vue'
import HealthRecordCreate from '../views/health/HealthRecordCreate.vue'
import AppointmentList from '../views/appointments/AppointmentList.vue'
import AppointmentDetail from '../views/appointments/AppointmentDetail.vue'
import AppointmentCreate from '../views/appointments/AppointmentCreate.vue'
import PrescriptionList from '../views/prescriptions/PrescriptionList.vue'
import PrescriptionDetail from '../views/prescriptions/PrescriptionDetail.vue'
// import LabReport from '../views/lab/LabReport.vue'
import MedicalImage from '../views/lab/MedicalImage.vue'
import InsuranceInfo from '../views/insurance/InsuranceInfo.vue'
import NotificationList from '../views/notifications/NotificationList.vue'
import LabReportList from '../views/lab/LabReportList.vue'
import LabReportDetail from '../views/lab/LabReportDetail.vue'

// 引入布局组件
import MainLayout from '../components/layout/MainLayout.vue'
import AuthLayout from '../components/layout/AuthLayout.vue'

// 定义路由规则（通过 meta.requiresAuth 标识需要登录的页面）
const routes = [
  {
    path: '/',
    component: MainLayout,
    meta: { requiresAuth: true },
    children: [
      { path: '', name: 'Home', component: Home },
      { path: 'profile', name: 'Profile', component: Profile },
      { path: 'dashboard/patient', name: 'PatientDashboard', component: PatientDashboard },
      { path: 'dashboard/doctor', name: 'DoctorDashboard', component: DoctorDashboard },
      { path: 'dashboard/admin', name: 'AdminDashboard', component: AdminDashboard },
      { path: 'health/records', name: 'HealthRecords', component: HealthRecords },
      { path: 'health/records/create', name: 'HealthRecordCreate', component: HealthRecordCreate },
      { path: 'health/records/:id', name: 'HealthRecordDetail', component: HealthRecordDetail },
      { path: 'appointments', name: 'AppointmentList', component: AppointmentList },
      { path: 'appointments/create', name: 'AppointmentCreate', component: AppointmentCreate },
      { path: 'appointments/:id', name: 'AppointmentDetail', component: AppointmentDetail },
      // 新增处方相关路由
      { path: 'prescriptions', name: 'PrescriptionList', component: PrescriptionList },
      { path: 'prescriptions/:id', name: 'PrescriptionDetail', component: PrescriptionDetail },
      // { path: 'lab/reports', name: 'LabReport', component: LabReport },
      // { path: 'lab/images', name: 'MedicalImage', component: MedicalImage },
      { path: 'lab/reports', name: 'LabReportList', component: LabReportList },
      { path: 'lab/reports/:id', name: 'LabReportDetail', component: LabReportDetail },
      { path: 'lab/images', name: 'MedicalImage', component: MedicalImage },
      { path: 'insurance', name: 'InsuranceInfo', component: InsuranceInfo },
      { path: 'notifications', name: 'NotificationList', component: NotificationList },
    ],
  },
  {
    path: '/',
    component: AuthLayout,
    children: [
      { path: 'login', name: 'Login', component: Login },
      { path: 'register', name: 'Register', component: Register },
      { path: 'password-reset', name: 'PasswordReset', component: PasswordReset },
    ],
  },
  // 404 兜底路由
  { path: '/:catchAll(.*)', redirect: '/' },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

// 全局路由守卫（示例：使用 localStorage 模拟登录状态）
router.beforeEach((to, from, next) => {
  const isAuthenticated = localStorage.getItem('isAuthenticated') === 'true'

  if (to.meta.requiresAuth && !isAuthenticated) {
    next({ name: 'Login' })
  } else if (
    (to.name === 'Login' || to.name === 'Register' || to.name === 'PasswordReset') &&
    isAuthenticated
  ) {
    next({ name: 'Home' })
  } else {
    next()
  }
})

export default router
