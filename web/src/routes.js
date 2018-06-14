import Index from './views/Index.vue'
import Home from './views/Home.vue'
import Foo from './views/Foo.vue'

const routes= [
    { path: '/', component: Index },
    { path: '/foo', component: Foo },
    { path: '/home', component: Home }
  ]

export default routes