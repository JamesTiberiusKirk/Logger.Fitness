import { createWebHistory, createRouter } from "vue-router";
// lazy-loaded
const Login = () => import("./views/Login.vue");
const Register = () => import("./views/Register.vue");
const Profile = () => import("./views/Profile.vue");
const ExerciseTypeList = () => import("./views/ExerciseTypeList.vue");
const ExerciseType = () => import("./views/ExerciseType.vue");
const WorkoutsList = () => import("./views/WorkoutsList.vue");
const WorkoutForm = () => import("./views/WorkoutForm.vue");
const Workout = () => import("./views/Workout.vue");

const routes = [
  {
    path: "/",
    redirect: "/workouts"
  },
  {
    path: "/login",
    name: "login",
    component: Login
  },
  {
    path: "/register",
    component: Register
  },
  {
    path: "/profile",
    name: "profile",
    component: Profile
  },
  {
    path: "/exercise_type_list",
    name: "exercise_type_list",
    component: ExerciseTypeList
  },
  {
    path: "/exercise_type",
    name: "exercise_type",
    component: ExerciseType
  },
  {
    path: "/workouts",
    name: "workout_list",
    component: WorkoutsList
  },
  {
    path: "/workout_form",
    name: "workout_form",
    component: WorkoutForm
  },
  {
    path: "/workout",
    name: "workout",
    component: Workout
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

router.beforeEach((to, from, next) => {
  const publicPages = ["/login", "/register", "/"];
  const authRequired = !publicPages.includes(to.path);
  const loggedIn = localStorage.getItem("user");

  // trying to access a restricted page + not logged in
  // redirect to login page
  if (authRequired && !loggedIn) {
    next("/login");
  } else {
    next();
  }
});

export default router;
