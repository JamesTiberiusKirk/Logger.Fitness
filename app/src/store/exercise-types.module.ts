import { ExerciseType, ExerciseTypeMap } from "@/types/exercise-type";
import ExerciseTypesService from "../services/exercise-types.service";

// const EXERCISE_TYPE_STORE = "exercise_types";

export type ExerciseTypeState = {
  data: ExerciseType[],
  empty: boolean,
}

function getDefaultState(): ExerciseTypeState {
  // let data = JSON.parse(localStorage.getItem(EXERCISE_TYPE_STORE)) || []
  const data: ExerciseType[] = [];
  return {
    data,
    empty: data.length == 0 ? false : true
  };
}

export const exerciseTypes = {
  namespaced: true,
  state: getDefaultState(),
  actions: {
    async fetchAll({ commit }: any) {
      return ExerciseTypesService.getExerciseTypes()
        .then(res => {
          commit("storeAll", res.data);
          return Promise.resolve(res.data);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    async sendOne({ commit }: any, exerciseType: ExerciseType) {
      return ExerciseTypesService.newExerciseType(exerciseType)
        .then(res => {
          commit("addOne", exerciseType);
          return Promise.resolve(res.data);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    // TODO: this again is because the possibility of wiping all the other fields, need testing
    async updateOne({ commit }: any, exerciseType: any) {
      return ExerciseTypesService.updateExerciseType(exerciseType)
        .then(res => {
          commit("updateOne", exerciseType);
          return Promise.resolve(res.data);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    async deleteOne({ commit }: any, id: string) {
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
    storeAll(state: ExerciseTypeState, exerciseTypes: ExerciseType[]) {
      state.data = exerciseTypes;
      state.empty = false;
      // localStorage.setItem(EXERCISE_TYPE_STORE, JSON.stringify(state.data));
    },
    addOne(state: ExerciseTypeState, exerciseType: ExerciseType) {
      state.data.push(exerciseType);
      state.empty = false;
      // localStorage.setItem(EXERCISE_TYPE_STORE, JSON.stringify(state.data));
    },
    updateOne(state: ExerciseTypeState, exerciseType: ExerciseType) {
      state.data.forEach((element, index) => {
        if (element.exercise_type_id == exerciseType.exercise_type_id) {
          state.data[index] = exerciseType;
        }
      });
      // localStorage.setItem(EXERCISE_TYPE_STORE, JSON.stringify(state.data));
    },
    deleteOne(state: ExerciseTypeState, id: string) {
      state.data.forEach((element, index) => {
        if (element.exercise_type_id == id) {
          state.data.splice(index, 1);
        }
      });
      // localStorage.setItem(EXERCISE_TYPE_STORE, JSON.stringify(state.data));
    }
  },
  getters: {
    getAll: (state: ExerciseTypeState) => {
      return state;
    },
    getOneById: (state: ExerciseTypeState) => (id: string) => {
      let result;
      state.data.forEach(element => {
        if (element.exercise_type_id == id) {
          result = element;
        }
      });
      return result;
    },
    getAllAsMap: (state: ExerciseTypeState) => {
      const result: ExerciseTypeMap = {};
      state.data.forEach(element => {
        result[element.exercise_type_id] = element;
      });
      return result;
    }
  }
};
