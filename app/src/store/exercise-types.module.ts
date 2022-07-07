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
    data: data as ExerciseType[],
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

      for (let i = 0; i < state.transactions.length; i++) {
        const tx = state.transactions[i]
        try {
          switch (tx.type) {
            case "add": {
              await ExerciseTypesService.newExerciseType(tx.data)
              break;
            }
            case "update": {
              await ExerciseTypesService.updateExerciseType(tx.data)
              break;
            }
            case "delete": {
              await ExerciseTypesService.deleteExerciseType(tx.data.name)
              break;
            }
          }
        } catch (e) {
          return Promise.reject(e)
        } finally {
          commit("deleteTx", tx.data.name)
        }
      }
      return Promise.resolve()
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
      const tx = { type: "add", data: exerciseType } as ExerciseTypeTransaction;
      commit("addTx", tx)
      return ExerciseTypesService.newExerciseType(exerciseType)
        .then(res => {
          commit("deleteTx", exerciseType.name)
          return Promise.resolve(res.data);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    async updateOne({ commit }: any, exerciseType: ExerciseType) {
      commit("updateOne", exerciseType);

      const tx = { type: "update", data: exerciseType } as ExerciseTypeTransaction;
      commit("addTx", tx)

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

      const tx = { type: "delete", data: { name } } as ExerciseTypeTransaction
      commit("addTx", tx)

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
    deleteTx(state: ExerciseTypeState, name: string) {
      console.log("deleting tx with name", name);

      let toDelete = -1
      for (let i = 0; i < state.transactions.length; i++) {
        if (state.transactions[i].data.name === name) {
          toDelete = i
        }
      }

      if (toDelete != -1) {
        state.transactions.splice(toDelete, 1);
      } else {
        console.log("nothing to delete");
      }

      localStorage.setItem(EXERCISE_TYPE_STORE_TX, JSON.stringify(state.transactions))
    },
    addTx(state: ExerciseTypeState, tx: ExerciseTypeTransaction) {
      console.log("adding tx", tx);

      state.transactions.push(tx)
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
    deleteOne(state: ExerciseTypeState, name: string) {
      let toDelete = -1 
      state.data.forEach((element, index) => {
        console.log(element, name)
        if (element.name == name) {
          toDelete = index;
        }
      });
      state.data.splice(toDelete, 1);
      localStorage.setItem(EXERCISE_TYPE_STORE, JSON.stringify(state.data));
    }
  },
  getters: {
    needToSync: (state: ExerciseTypeState): boolean =>{
      return (state.transactions.length > 0)
    }, 
    txAmount: (state: ExerciseTypeState): number => {
      return state.transactions.length
    },
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
