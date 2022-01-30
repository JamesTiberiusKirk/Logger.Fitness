import ExerciseService from "../services/exercise.service";

// const EXERCISE_STORE = "exercises";

function getDefaultState() {
  // let data = JSON.parse(localStorage.getItem(EXERCISE_STORE)) || [];
  let data = [];
  return {
    data,
    empty: data.length == 0 ? false : true
  };
}

export const exercises = {
  namespaced: true,
  state: getDefaultState(),
  actions: {
    async new({ commit }, exercise) {
      return ExerciseService.new(exercise)
        .then(res => {
          commit("storeOne", res.data);
          return Promise.resolve(res.data);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    async fetchAll({ commit }) {
      return ExerciseService.getAll()
        .then(res => {
          commit("storeAll", res.data);
          return Promise.resolve(res.data);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    async fetchInWorkout({ commit }, workoutID) {
      return ExerciseService.getAllInWorkout(workoutID)
        .then(res => {
          commit("storeAll", res.data);
          return Promise.resolve(res.data);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    async updateOne({ commit }, exercise) {
      return ExerciseService.edit(exercise)
        .then(res => {
          commit("updateOne", exercise);
          return Promise.resolve(res.data);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    async deleteOne({ commit }, id) {
      return ExerciseService.delete(id)
        .then(res => {
          commit("deleteOne", id);
          return Promise.resolve(res);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    }
  },
  mutations: {
    storeAll(state, exercises) {
      state.data = exercises;
      state.empty = false;
      // localStorage.setItem(EXERCISE_STORE, JSON.stringify(state.data));
    },
    storeByWorkout(state, exercises) {
      state.data = exercises;
      state.empty = false;
      // localStorage.setItem(EXERCISE_STORE, JSON.stringify(state.data));
    },
    storeOne(state, exercise) {
      state.data.push(exercise);
      state.empty = false;
      // localStorage.setItem(EXERCISE_STORE, JSON.stringify(state.data));
    },
    updateOne(state, exercise) {
      state.data.forEach((element, index) => {
        if (element.exercise_id == exercise.exercise_id) {
          state.data[index] = exercise;
        }
      });
      // localStorage.setItem(EXERCISE_STORE, JSON.stringify(state.data));
    },
    deleteOne(state, id) {
      state.data.forEach((element, index) => {
        if (element.exercise_id == id) {
          state.data.splice(index, 1);
        }
      });
      // localStorage.setItem(EXERCISE_STORE, JSON.stringify(state.data));
    }
  },
  getters: {
    getAllByWorkoutId: state => workoutId => {
      let result = [];
      state.data.forEach(element => {
        if (element.workoutId == workoutId) {
          result.push(element);
        }
      });
      return result;
    },
    getOneById: state => id => {
      let result;
      state.data.forEach(element => {
        if (element.exercise_id == id) {
          result = element;
        }
      });
      return result;
    }
  }
};
