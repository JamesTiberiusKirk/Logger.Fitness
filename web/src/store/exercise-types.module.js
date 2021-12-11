import ExerciseTypesService from "../services/exercise-types.service";

const EXERCISE_TYPE_STORE = "exercise_types";
// const all = JSON.parse(localStorage.getItem(EXERCISE_TYPE_STORE))
// const all = null

function getDefaultState() {
  let data = JSON.parse(localStorage.getItem(EXERCISE_TYPE_STORE));
  return {
    data,
    empty: data ? false : true
  };
}

export const exerciseType = {
  namespaced: true,
  state: getDefaultState(),
  actions: {
    async fetchAll({ commit }) {
      return ExerciseTypesService.getExerciseTypes()
        .then(res => {
          commit("storeAll", res.data);
          return Promise.resolve(res.data);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    async sendOne({ commit }, exerciseType) {
      return ExerciseTypesService.newExerciseType(exerciseType)
        .then(res => {
          commit("addOne", exerciseType);
          return Promise.resolve(res.data);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    async updateOne({ commit }, exerciseType) {
      return ExerciseTypesService.updateExerciseType(exerciseType)
        .then(res => {
          commit("updateOne", exerciseType);
          return Promise.resolve(res.data);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    async deleteOne({ commit }, id) {
      return ExerciseTypesService.deleteExerciseType(id)
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
    storeAll(state, exerciseTypes) {
      state.data = exerciseTypes;
      state.empty = false;
      localStorage.setItem(EXERCISE_TYPE_STORE, JSON.stringify(state.data));
    },
    addOne(state, exerciseType) {
      state.data.push(exerciseType);
      state.empty = false;
      localStorage.setItem(EXERCISE_TYPE_STORE, JSON.stringify(state.data));
    },
    updateOne(state, exerciseType) {
      state.data.forEach((element, index) => {
        if (element.exercise_type_id == exerciseType.exercise_type_id) {
          state.data[index] = exerciseType;
        }
      });
      localStorage.setItem(EXERCISE_TYPE_STORE, JSON.stringify(state.data));
    },
    deleteOne(state, id) {
      state.data.forEach((element, index) => {
        if (element.exercise_type_id == id) {
          state.data.splice(index, 1);
        }
      });
      localStorage.setItem(EXERCISE_TYPE_STORE, JSON.stringify(state.data));
    }
  },
  getters: {
    getAll: state => {
      return state;
    },
    getOneById: state => id => {
      let result;
      state.data.forEach(element => {
        if (element.exercise_type_id == id) {
          result = element;
        }
      });
      return result;
    }
  }
};
