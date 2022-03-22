import { createRouter, createWebHistory } from '@ionic/vue-router';
import { RouteRecordRaw } from 'vue-router';
import TabsPage from '@/views/TabsPage.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    redirect: '/tabs/workouts'
  },
  {
    path: '/login',
    component: () => import('@/views/LoginPage.vue')
  },
  {
    path: '/tabs/',
    component: TabsPage,
    children: [
      {
        path: '',
        redirect: '/tabs/workouts'
      },
      {
        path: 'exercise_list',
        component: () => import('@/views/tabs/ExerciseListPage.vue')
      },
      {
        path: 'exercise',
        component: () => import('@/views/tabs/ExercisePage.vue')
      },
      {
        path: 'workouts',
        component: () => import('@/views/tabs/WorkoutsPage.vue')
      },
      {
        path: 'analytics',
        component: () => import('@/views/tabs/AnalyticsPage.vue')
      }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

router.beforeEach((to, from, next) => {
  const publicPages = ["/login", "/register", "/"];
  const authRequired = !publicPages.includes(to.path);
  const loggedIn = localStorage.getItem("user");

  // trying to access a restricted page + not logged in
  // redirect to login page
  (authRequired && !loggedIn) ? next("/login") : next();
});

export default router
