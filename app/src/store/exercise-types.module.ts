import { ExerciseType, ExerciseTypeMap } from "@/types/exercise-type";
import ExerciseTypesService from "../services/exercise-types.service";

const EXERCISE_TYPE_STORE = "exercise_types";
const EXERCISE_TYPE_STORE_TX = "exercise_types_tx";

// Transactions are for keeping a list of transactions to be ran on the data
//  this helps with decoupling frontend from server
type ExerciseTypeTransaction = {
  type: string,
  data: ExerciseType,
}

export type ExerciseTypeState = {
  data: ExerciseType[],
  empty: boolean,
  transactions: ExerciseTypeTransaction[],
}

function getDefaultState(): ExerciseTypeState {
  const data = JSON.parse(localStorage.getItem(EXERCISE_TYPE_STORE) || "[]")
  const transactions = JSON.parse(localStorage.getItem(EXERCISE_TYPE_STORE_TX) || "[]")
  return {
    data,
    empty: data.length == 0 ? false : true,
    transactions: transactions as ExerciseTypeTransaction[],
  };
}

export const exerciseTypes = {
  namespaced: true,
  state: getDefaultState(),
  actions: {
    async syncTx({ commit, state }: any) {
      console.log("Exercise type transaction sync");
      console.log(state.transactions);

      for (let i = 0; i++; i < state.transaction.length) {

        try {

        switch (state.transaction[i].type) {
          case "add": {
            await ExerciseTypesService.newExerciseType(state.transaction[i].data)
            break;
          }
          case "update": {
            break;
          }
          case "delete": {
            break;
          }
        }

        } catch (e){
          return Promise.reject(e)
        } finally {
          commit("deleteTx", state.transaction[i].data.name)
        }
      }

    },
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
      commit("addOne", exerciseType);
      commit("addTx", "add", exerciseType)
      return ExerciseTypesService.newExerciseType(exerciseType)
        .then(res => {
          commit("deleteTx", exerciseType.name)
          return Promise.resolve(res.data);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    async updateOne({ commit }: any, exerciseType: any) {
      commit("updateOne", exerciseType);
      commit("addTx", "update", exerciseType)
      return ExerciseTypesService.updateExerciseType(exerciseType)
        .then(res => {
          commit("deleteTx", exerciseType.name)
          return Promise.resolve(res.data);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    async deleteOne({ commit }: any, name: string) {
      commit("deleteOne", name);
      commit("addTx", "delete", { name })
      return ExerciseTypesService.deleteExerciseType(name)
        .then(res => {
          commit("deleteTx", { name })
          return Promise.resolve(res);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    }
  },
  mutations: {
    deleteTx(state: ExerciseTypeState, exerciseTypeName: string) {
      state.transactions = state.transactions.filter((tx: ExerciseTypeTransaction) => {
        return tx.data.name === exerciseTypeName
      })
      localStorage.setItem(EXERCISE_TYPE_STORE_TX, JSON.stringify(state.transactions))
    },
    addTx(state: ExerciseTypeState, type: string, exerciseType: ExerciseType) {
      state.transactions.push({
        type,
        data: exerciseType,
      } as ExerciseTypeTransaction)
      localStorage.setItem(EXERCISE_TYPE_STORE_TX, JSON.stringify(state.transactions))
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
