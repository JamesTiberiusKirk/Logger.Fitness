import { createStore } from "vuex";
import { auth } from "./auth.module";
import { exerciseType } from "./exercise-types.module";

const store = createStore({
  modules: {
    auth,
    exerciseType,
  },
});

export default store;
