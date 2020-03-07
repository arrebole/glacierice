import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Content from '../views/Content.vue'
import Refererence from '../views/Refererence.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/data-resource',
    name: 'DataResource',
    component: Content
  },
  {
    path: '/open-data-api',
    name: 'OpenDataAPI',
    component: Content
  },
  {
    path: '/contributor',
    name: 'Contributor',
    component: Content
  },
  {
    path: '/crash-course',
    name: 'CrashCourse',
    component: Refererence
  },
  {
    path: '/papers',
    name: 'Papers',
    component: Refererence
  },
  {
    path: '/models',
    name: 'Models',
    component: Refererence
  },
  {
    path: '/thanks',
    name: "Thanks",
    component: () => import(/* webpackChunkName: "thanks-page" */ '../views/Thanks.vue')
  }, {
    path: '*',
    redirect: { name: 'Home' }
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
  scrollBehavior() {
    return { x: 0, y: 0 }
  }
})

export default router
