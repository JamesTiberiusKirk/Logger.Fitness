import { createStore } from "vuex";
import { auth } from "./auth.module";
import { exerciseType } from "./exercise-types.module";
import { workouts } from "./workouts.module";

const store = createStore({
  modules: {
    auth,
    exerciseType,
    workouts,
  },
});

export default store;
