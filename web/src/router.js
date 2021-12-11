import { createWebHistory, createRouter } from "vue-router";
// lazy-loaded
const Home = () => import("./components/Home.vue")
const Login = () => import("./components/Login.vue")
const Register = () => import("./components/Register.vue")
const Profile = () => import("./components/Profile.vue")
const ExerciseTypeList = () => import("./components/ExerciseTypeList.vue")
const ExerciseType = () => import("./components/ExerciseType.vue")

const routes = [
  {
    path: "/",
    redirect : "/home",
  },
  {
    path: "/home",
    name: "home",
    component: Home,
  },
  {
    path: "/login",
    component: Login,
  },
  {
    path: "/register",
    component: Register,
  },
  {
    path: "/profile",
    name: "profile",
    component: Profile,
  },
  {
    path: "/exercise_type_list",
    name: "exercise_type_list",
    component: ExerciseTypeList,
  },
  {
    path: "/exercise_type",
    name: "exercise_type",
    component: ExerciseType,
    props: true
  } ,
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  const publicPages = ['/login', '/register', '/home', '/'];
  const authRequired = !publicPages.includes(to.path);
  const loggedIn = localStorage.getItem('user');

  // trying to access a restricted page + not logged in
  // redirect to login page
  if (authRequired && !loggedIn) {
    next('/login');
  } else {
    next();
  }
});

export default router;