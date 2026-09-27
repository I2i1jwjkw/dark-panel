import { createRouter, createWebHistory } from 'vue-router'
import Login from '../views/Login.vue'
import Dashboard from '../views/Dashboard.vue'
const router=createRouter({history:createWebHistory(),routes:[{path:'/',redirect:'/login'},{path:'/login',component:Login},{path:'/dashboard',component:Dashboard}]})
router.beforeEach(to=>{const token=localStorage.getItem('token');if(to.path==='/dashboard'&&!token)return '/login';if(to.path==='/login'&&token)return '/dashboard'})
export default router