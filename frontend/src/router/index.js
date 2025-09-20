import { createRouter, createWebHistory } from 'vue-router'

// 路由页面组件（懒加载，提升性能）
const Home = () => import('../views/Home.vue')
const Market = () => import('../views/Market.vue')
//const CardDetail = () => import('../views/CardDetail.vue')
//const CreateCard = () => import('../views/CreateCard.vue')
const Collection = () => import('../views/Collection.vue')

const Profile = () => import('../views/Profile.vue')
//const NotFound = () => import('../views/NotFound.vue')
const Mint = () => import('../views/Mint.vue')
const Auth = () => import('../views/Auth.vue')
const Mine = () => import('../views/Mine.vue')



const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path:'/auth',
    name:'Auth',
    component: Auth
  },
  { path: '/profile',
    name: 'Profile',
    component: Profile
  },
  { path: '/mint',
    name: 'Mint',
    component: Mint
  },
  {
    path: '/mine',
    name: 'Mine',
    component: Mine
  },
  { path: '/collection',
    name: 'Collection',
    component: Collection
  },
  { path: '/market',
    name: 'Market',
    component: Market
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// 路由守卫（简单占位，后续可接入钱包验证）
router.beforeEach((to, from, next) => {
  if (to.meta.requiresAuth) {
    // TODO: 判断是否已连接钱包/已登录
    const isAuthenticated = false // 占位
    if (!isAuthenticated) {
      alert('请先连接钱包或登录！')
      return next('/')
    }
  }
  next()
})

export default router