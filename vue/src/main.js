import Vue from 'vue'
import VueRouter from 'vue-router'

import App from './App.vue'
import Admin from './Admin.vue'
import Analytics from './components/Analytics.vue'

Vue.config.productionTip = false

// vueRouterでルーティング
const routes=[
  {
    path:'/',
    name:'App',
    component:App
  },
  {
    path:'/admin',
    component:Admin,
    children:[
      {
        path:'log',
        component:Analytics
      },
      {
        path:'*',
        redirect:'/admin'
      }
    ]
  },
  {
    path:'*',
    redirect:'/'
  }
]
const router = new VueRouter({
  routes,
  mode: 'history',
  base: process.env.BASE_URL
})

// グローバル変数の設定
let GlobalData = new Vue({
  data: {
    // apiのホストアドレス
    // host: 'http://localhost:8081',
    host: 'http://172.16.6.4:8081',
  }
});
Vue.mixin({
  computed:{
    host:{
      get:()=>{ return GlobalData.$data.host },
      set:(e)=>{GlobalData.$data.host = e;}
    }
  }
})

Vue.use(VueRouter)
new Vue({
  router,
  render: h => h('router-view')
}).$mount('#app')
