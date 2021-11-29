import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/components/Login'
import Workouts from '@/components/Workouts'
import HelloWorld from '@/components/HelloWorld'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'HelloWorld',
      component: HelloWorld
    },
    {
      path: '/login',
      name: 'Login',
      component: Login
    },
    {
      path: '/workouts',
      name: 'Workouts',
      component: Workouts
    }
  ]
})
