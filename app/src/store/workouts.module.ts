import { Workout, WorkoutGroup } from "@/types/workout";
import WorkoutsService from "../services/workouts.service";

const WORKOUTS_STORE = "workouts";

export type WorkoutsState = {
  data: WorkoutGroup[],
  empty: boolean,
}

function getDefaultState(): WorkoutsState {
  const data = JSON.parse(localStorage.getItem(WORKOUTS_STORE) || "[]");
  // const data: WorkoutGroup[] = [];
  return {
    data,
    empty: data.length == 0 ? false : true
  };
}

export const workouts = {
  namespaced: true,
  state: getDefaultState(),
  actions: {
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
    async updateOne({ commit }: any, workout: any) {
      return WorkoutsService.edit(workout)
        .then(res => {
          commit("updateOne", workout);
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
    updateOne(state: WorkoutsState, workout_group: WorkoutGroup) {
      state.data.forEach((element, index) => {
        if (element.workout.workout_id == workout_group.workout.workout_id) {
          state.data[index] = workout_group;
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
    getOneById: (state: WorkoutsState) => (id: string) => {
      let result;
      state.data.forEach(workout_group=> {
        if (workout_group.workout.workout_id == id) {
          result = workout_group;
        }
      });
      return result;
    }
  }
};
