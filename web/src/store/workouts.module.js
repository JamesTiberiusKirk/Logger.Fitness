import WorkoutsService from "../services/workouts.service";

const WORKOUTS_STORE = "workouts";

function getDefaultState() {
  let data = JSON.parse(localStorage.getItem(WORKOUTS_STORE));
  return {
    data,
    empty: data ? false : true
  };
}

export const workouts = {
  namespaced: true,
  state: getDefaultState(),
  actions: {
    async start({ commit }, workout) {
      return WorkoutsService.start(workout)
        .then(res => {
          commit("storeOne", res.data);
          return Promise.resolve(res.data);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    async stop({ commit }, time) {
      return WorkoutsService.stop(time)
        .then(res => {
          commit("updateOne", res.data);
          return Promise.resolve(res.data);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    async fetchAll({ commit }) {
      return WorkoutsService.getAll()
        .then(res => {
          commit("storeAll", res.data);
          return Promise.resolve(res.data);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    async updateOne({ commit }, workout) {
      return WorkoutsService.edit(workout)
        .then(res => {
          commit("updateOne", workout);
          return Promise.resolve(res.data);
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    async deleteOne({ commit }, id) {
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
    storeAll(state, workout) {
      state.data = workout;
      state.empty = false;
      localStorage.setItem(WORKOUTS_STORE, JSON.stringify(state.data));
    },
    storeOne(state, workout) {
      state.data.push(workout);
      state.empty = false;
      localStorage.setItem(WORKOUTS_STORE, JSON.stringify(state.data));
    },
    updateOne(state, workout) {
      state.data.forEach((element, index) => {
        if (element.workout_id == workout.workout_id) {
          state.data[index] = workout;
        }
      });
      localStorage.setItem(WORKOUTS_STORE, JSON.stringify(state.data));
    },
    deleteOne(state, id) {
      state.data.forEach((element, index) => {
        if (element.workout_id == id) {
          state.data.splice(index, 1);
        }
      });
      localStorage.setItem(WORKOUTS_STORE, JSON.stringify(state.data));
    }
  },
  getters: {
    getAll: state => {
      return state;
    },
    getOneById: state => id => {
      let result;
      state.data.forEach(element => {
        if (element.workout_id == id) {
          result = element;
        }
      });
      return result;
    }
  }
};
