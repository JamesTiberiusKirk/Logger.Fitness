import { createStore } from "vuex";
import { auth } from "./auth.module";
import { exerciseTypes } from "./exercise-types.module";
import { workouts } from "./workouts.module";
import { exercises } from "./exercises.module";

const store = createStore({
  modules: {
    auth,
    exerciseTypes,
    workouts,
    exercises,
  },
});

export default store;
