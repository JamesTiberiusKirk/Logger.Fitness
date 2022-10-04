import { Workout, WorkoutGroup } from "@/types/workout";
import WorkoutsService from "../services/workouts.service";

const WORKOUTS_STORE = "workouts";
const EXERCISE_TYPE_STORE_TX = "workouts_types_tx";

type WorkoutGroupTransaction = {
  type: string,
  data: WorkoutGroup,
}

export type WorkoutsState = {
  data: WorkoutGroup[],
  empty: boolean,
  transactions: WorkoutGroupTransaction[],
}

function getDefaultState(): WorkoutsState {
  const data = JSON.parse(localStorage.getItem(WORKOUTS_STORE) || "[]");
  const transactions = JSON.parse(localStorage.getItem(EXERCISE_TYPE_STORE_TX) || "[]")
  return {
    data,
    empty: data.length == 0 ? false : true,
    transactions: transactions as WorkoutGroupTransaction[],
  };
}

export const workouts = {
  namespaced: true,
  state: getDefaultState(),
  actions: {
    async syncTx({ commit, state }: any) {
      console.log("Workout state transaction sync");
      console.log(state.transactions);

      for (let i = 0; i++; i < state.transaction.length) {
        try {
          switch (state.transaction[i].type) {
            case "start": {
              await WorkoutsService.start(state.transaction[i].data)
              break;
            }
            case "stop": {
              await WorkoutsService.stop(state.transaction[i].data.end_time)
              break;
            }
            // TODO: Need to finish the rest of the operations of the workout service
            //case "delete": {
              //await WorkoutsService.update(state.transaction[i].data.name)
              //break;
            //}
          }
        } catch (e) {
          return Promise.reject(e)
        } finally {
          commit("deleteTx", state.transaction[i].data.name)
        }
      }
      return Promise.resolve()
    },
    async start({ commit }: any, workout: Workout) {
      return WorkoutsService.start(workout)
        .then(res => {
          commit("storeOne", res.data);
          return Promise.resolve(res.data);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    async stop({ commit }: any, time: string) {
      return WorkoutsService.stop(time)
        .then(res => {
          commit("updateOne", res.data);
          return Promise.resolve(res.data);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    async fetchAll({ commit }: any) {
      return WorkoutsService.getAll()
        .then(res => {
          commit("storeAll", res.data);
          return Promise.resolve(res.data);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    async updateOneWorkout({ commit }: any, workout: Workout) {
      return WorkoutsService.edit(workout)
        .then(res => {
          commit("updateOneWorkout", workout);
          return Promise.resolve(res.data);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    async deleteOne({ commit }: any, id: string) {
      return WorkoutsService.delete(id)
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
    deleteTx(state: WorkoutsState, start_time: string) {
      console.log("deleting tx with name", start_time);

      let toDelete = -1
      for (let i = 0; i < state.transactions.length; i++) {
        if (state.transactions[i].data.workout.start_time === start_time) {
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
    addTx(state: WorkoutsState, tx: WorkoutGroupTransaction) {
      console.log("adding tx", tx);

      state.transactions.push(tx)
      localStorage.setItem(EXERCISE_TYPE_STORE_TX, JSON.stringify(state.transactions))
    },
    storeAll(state: WorkoutsState, workouts: WorkoutGroup[]) {
      state.data = workouts;
      state.empty = false;
      localStorage.setItem(WORKOUTS_STORE, JSON.stringify(state.data));
    },
    storeOne(state: WorkoutsState, workout_group: WorkoutGroup) {
      state.data.push(workout_group);
      state.empty = false;
      localStorage.setItem(WORKOUTS_STORE, JSON.stringify(state.data));
    },
    updateOneGroup(state: WorkoutsState, workout_group: WorkoutGroup) {
      state.data.forEach((element, index) => {
        if (element.workout.workout_id == workout_group.workout.workout_id) {
          state.data[index] = workout_group;
        }
      });
      localStorage.setItem(WORKOUTS_STORE, JSON.stringify(state.data));
    },
    updateOneWorkout(state: WorkoutsState, workout: Workout) {
      state.data.forEach((element, index) => {
        if (element.workout.workout_id == workout.workout_id) {
          state.data[index].workout = workout;
        }
      });
      localStorage.setItem(WORKOUTS_STORE, JSON.stringify(state.data));
    },
    deleteOne(state: WorkoutsState, id: string) {
      state.data.forEach((workout_group, index) => {
        if (workout_group.workout.workout_id == id) {
          state.data.splice(index, 1);
        }
      });
      localStorage.setItem(WORKOUTS_STORE, JSON.stringify(state.data));
    }
  },
  getters: {
    getAll: (state: WorkoutsState) => {
      return state;
    },
    getAllInReverse: (state: WorkoutsState) => {
      if (!state.data) return [];
      return state.data.slice().reverse();
    },
    getOneWorkoutGroupById: (state: WorkoutsState) => (id: string): WorkoutGroup => {
      let result: WorkoutGroup = {} as WorkoutGroup;
      state.data.forEach((workout_group: WorkoutGroup) => {
        if (workout_group.workout.workout_id == id) {
          result = workout_group;
        }
      });
      return result;
    },
    getOneWorkoutById: (state: WorkoutsState) => (id: string): Workout => {
      let result: Workout = {} as Workout;
      state.data.forEach((workout_group: WorkoutGroup) => {
        if (workout_group.workout.workout_id == id) {
          result = workout_group.workout;
        }
      });
      return result;
    }
  }
};
