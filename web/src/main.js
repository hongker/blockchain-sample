import Vue from 'vue'
import 'element-ui/lib/theme-chalk/index.css'
import App from './views/App.vue'
import routes from './routes'

import {
  Menu,
  Submenu,
  MenuItem,
  MenuItemGroup,
  Carousel,
  CarouselItem,
} from 'element-ui'

Vue.component(Menu.name, Menu);
Vue.component(Submenu.name, Submenu);
Vue.component(MenuItem.name, MenuItem);
Vue.component(MenuItemGroup.name, MenuItemGroup);
Vue.component(Carousel.name, Carousel);
Vue.component(CarouselItem.name, CarouselItem);

import VueRouter from 'vue-router'

Vue.use(VueRouter)

const router = new VueRouter({
  routes
})

new Vue({
  router,
  // el: '#app',
  render: h => h(App)
}).$mount('#app')
