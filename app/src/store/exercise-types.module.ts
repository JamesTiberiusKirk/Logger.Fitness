import { ExerciseType, ExerciseTypeMap } from "@/types/exercise-type";
import ExerciseTypesService from "../services/exercise-types.service";

const EXERCISE_TYPE_STORE = "exercise_types";

export type ExerciseTypeState = {
  data: ExerciseType[],
  lastUpdated: number,
  lastSynced: number,
  empty: boolean,
}

function getDefaultState(): ExerciseTypeState {
  const data = JSON.parse(localStorage.getItem(EXERCISE_TYPE_STORE) || "[]")
  // const data: ExerciseType[] = [];
  return {
    data,
    lastUpdated: -1,
    lastSynced: -1,
    empty: data.length == 0 ? false : true
  };
}

export const exerciseTypes = {
  namespaced: true,
  state: getDefaultState(),
  actions: {
    async syncChanges({ state }: any) {
      return ExerciseTypesService.syncAll(state.data)
        .then(() => {
          return Promise.resolve()
        })
        .catch(err => {
          return Promise.reject(err)
        })
    },
    async fetchAll({ commit }: any) {
      return ExerciseTypesService.getExerciseTypes()
        .then(res => {
          commit("storeAll", res.data);
          commit("updateLastUpdated");
          return Promise.resolve(res.data);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    async sendOne({ commit }: any, exerciseType: ExerciseType) {
      commit("addOne", exerciseType);
      return ExerciseTypesService.newExerciseType(exerciseType)
        .then(res => {
          commit("updateLastUpdated");
          return Promise.resolve(res.data);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    async updateOne({ commit }: any, exerciseType: any) {
      commit("updateOne", exerciseType);
      return ExerciseTypesService.updateExerciseType(exerciseType)
        .then(res => {
          commit("updateLastUpdated");
          return Promise.resolve(res.data);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    async deleteOne({ commit }: any, name: string) {
      commit("deleteOne", name);
      return ExerciseTypesService.deleteExerciseType(name)
        .then(res => {
          return Promise.resolve(res);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    }
  },
  mutations: {
    updateLastUpdated(state: ExerciseTypeState) {
      state.lastUpdated = Date.now();
    },
    updateLastSynced(state: ExerciseTypeState) {
      state.lastUpdated = Date.now();
    },
    storeAll(state: ExerciseTypeState, exerciseTypes: ExerciseType[]) {
      state.data = exerciseTypes;
      state.empty = false;
      localStorage.setItem(EXERCISE_TYPE_STORE, JSON.stringify(state.data));
    },
    addOne(state: ExerciseTypeState, exerciseType: ExerciseType) {
      state.data.push(exerciseType);
      state.empty = false;
      localStorage.setItem(EXERCISE_TYPE_STORE, JSON.stringify(state.data));
    },
    updateOne(state: ExerciseTypeState, exerciseType: ExerciseType) {
      state.data.forEach((element, index) => {
        if (element.name == exerciseType.name) {
          state.data[index] = exerciseType;
        }
      });
      localStorage.setItem(EXERCISE_TYPE_STORE, JSON.stringify(state.data));
    },
    deleteOne(state: ExerciseTypeState, id: string) {
      state.data.forEach((element, index) => {
        if (element.exercise_type_id == id) {
          state.data.splice(index, 1);
        }
      });
      localStorage.setItem(EXERCISE_TYPE_STORE, JSON.stringify(state.data));
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
    },
    getAllAsMapByName: (state: ExerciseTypeState) => {
      const result: ExerciseTypeMap = {};
      state.data.forEach(element => {
        result[element.name] = element;
      });
      return result;
    },
    getOneByName: (state: ExerciseTypeState) => (name: string) => {
      let result = {} as ExerciseType;
      state.data.forEach(element => {
        if (element.name == name) {
          result = element;
        }
      });
      return result;
    },

  }
};
