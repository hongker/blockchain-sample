import Vue from 'vue'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import App from './App.vue'
import Login from './views/Login.vue'
import Admin from './views/Admin.vue'
import Manage from './views/Manage.vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)
Vue.use(ElementUI)

const routes = [
	{path: '/', redirect: Login},  
  { path: '/login', component: Login },
  { path: '/admin', component: Admin },
  { path: '/manage', component: Manage },
]

const router = new VueRouter({
  routes // （缩写）相当于 routes: routes
})

new Vue({
  el: '#app',
  router,
  render: h => h(App)
})
