import { createStore } from "vuex";
import { auth } from "./auth.module";
import { exerciseTypes } from "./exercise-types.module";
import { workouts } from "./workouts.module";
import { exercises } from "./exercises.module";

// YET another example of js/ts being broken af, have to put <any> in order
//  to get the editor to stop complaining about code that works, yes the
//  Store type does not work if you want to access state.store...
const store = createStore<any>({
  modules: {
    auth,
    exerciseTypes,
    workouts,
    exercises,
  },
});

export default store;
