import Vue from 'vue'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import App from './App.vue'
import Login from './views/Login.vue'
import Admin from './views/Admin.vue'
import Manage from './views/Manage.vue'
import Init from './views/Init.vue'
import Register from './views/Register.vue'
import Auction from './views/Auction.vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)
Vue.use(ElementUI)

const routes = [
	{path: '/init', component: Init},  
	{path: '/register', component: Register},  
  { path: '/login', component: Login },
  { path: '/admin', component: Admin },
  { path: '/manage', component: Manage },
  {path: '/auction', component: Auction},  
]

const router = new VueRouter({
  routes // （缩写）相当于 routes: routes
})

new Vue({
  el: '#app',
  router,
  render: h => h(App)
})
