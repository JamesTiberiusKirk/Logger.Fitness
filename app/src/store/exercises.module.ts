import { Exercise } from "@/types/exercise";
import ExerciseService from "../services/exercise.service";

// const EXERCISE_STORE = "exercises";

export type ExerciseState = {
  data: Exercise[],
  empty: boolean,
}

function getDefaultState() {
  // let data = JSON.parse(localStorage.getItem(EXERCISE_STORE)) || [];
  const data: Exercise[] = [];
  return {
    data,
    empty: data.length == 0 ? false : true
  };
}

export const exercises = {
  namespaced: true,
  state: getDefaultState(),
  actions: {
    async new({ commit }: any, exercise: Exercise) {
      return ExerciseService.new(exercise)
        .then(res => {
          commit("storeOne", res.data);
          return Promise.resolve(res.data);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    async fetchAll({ commit }: any) {
      return ExerciseService.getAll()
        .then(res => {
          commit("storeAll", res.data);
          return Promise.resolve(res.data);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    async fetchInWorkout({ commit }: any, workoutID: string) {
      return ExerciseService.getAllInWorkout(workoutID)
        .then(res => {
          commit("storeAll", res.data);
          return Promise.resolve(res.data);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    // TODO: again any bc dont want to wipe fields, need testing
    async updateOne({ commit }: any, exercise: any) {
      return ExerciseService.edit(exercise)
        .then(res => {
          commit("updateOne", exercise);
          return Promise.resolve(res.data);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    async deleteOne({ commit }: any, id: string) {
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
    storeAll(state: ExerciseState, exercises: Exercise[]) {
      state.data = exercises;
      state.empty = false;
      // localStorage.setItem(EXERCISE_STORE, JSON.stringify(state.data));
    },
    // EMMM idk whats happening here, looks the same as storeAlll
    // TODO: find where it is used and figure out what its meant todo
    storeByWorkout(state: ExerciseState, exercises: Exercise[]) {
      state.data = exercises;
      state.empty = false;
      // localStorage.setItem(EXERCISE_STORE, JSON.stringify(state.data));
    },
    storeOne(state: ExerciseState, exercise: Exercise) {
      state.data.push(exercise);
      state.empty = false;
      // localStorage.setItem(EXERCISE_STORE, JSON.stringify(state.data));
    },
    updateOne(state: ExerciseState, exercise: Exercise) {
      state.data.forEach((element, index) => {
        if (element.exercise_id == exercise.exercise_id) {
          state.data[index] = exercise;
        }
      });
      // localStorage.setItem(EXERCISE_STORE, JSON.stringify(state.data));
    },
    deleteOne(state: ExerciseState, id: string) {
      state.data.forEach((element, index) => {
        if (element.exercise_id == id) {
          state.data.splice(index, 1);
        }
      });
      // localStorage.setItem(EXERCISE_STORE, JSON.stringify(state.data));
    }
  },
  getters: {
    // TODO: maybe go back and use something like the filter method?
    getAllByWorkoutId: (state: ExerciseState) => (workoutId: string): Exercise[] => {
      const result: Exercise[] = [];
      state.data.forEach(element => {
        if (element.workout_id == workoutId) {
          result.push(element);
        }
      });
      return result;
    },
    // TODO: maybe go back and use something like the filter method?
    getOneById: (state: ExerciseState) => (id: string): Exercise => {
      let result: Exercise = {} as Exercise;
      state.data.forEach(element => {
        if (element.exercise_id == id) {
          result = element;
        }
      });
      return result;
    }
  }
};
