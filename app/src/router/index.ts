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
        path: 'exercise/list',
        component: () => import('@/views/tabs/exercise_types/ExerciseListPage.vue')
      },
      {
        path: 'exercise/add',
        component: () => import('@/views/tabs/exercise_types/ExercisePage.vue')
      },
      {
        path: 'workouts/list',
        component: () => import('@/views/tabs/workouts/WorkoutListPage.vue')
      },
      {
        path: 'workouts/add',
        component: () => import('@/views/tabs/workouts/WorkoutEditPage.vue')
      },
      {
        path: 'analytics',
        component: () => import('@/views/tabs/analytics/AnalyticsPage.vue')
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
