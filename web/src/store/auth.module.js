import AuthService from "../services/auth.service";

const user = JSON.parse(localStorage.getItem("user"));
const initialState = {
  status: {
    loggedIn: user ? true : false
  },
  user
};

export const auth = {
  namespaced: true,
  state: initialState,
  actions: {
    async login({ commit }, user) {
      return AuthService.login(user)
        .then(res => {
          commit("loginSuccess", res.data);
          return Promise.resolve(res);
        })
        .catch(err => {
          commit("loginFailure");
          return Promise.reject(err);
        });
    },
    logout({ commit }) {
      commit("logout");
    },

    // TODO: this needs testing and possibly rework
    async register({ commit }, user) {
      return AuthService.register(user).then(
        response => {
          commit("registerSuccess");
          return Promise.resolve(response.data);
        },
        error => {
          commit("registerFailure");
          return Promise.reject(error);
        }
      );
    }
  },
  mutations: {
    storeJwt(state, jwt) {
      state.user.jwt = jwt;
      localStorage.setItem("user", JSON.stringify(state.user));
    },
    loginSuccess(state, user) {
      state.status.loggedIn = true;
      state.user = user;
      localStorage.setItem("user", JSON.stringify(user));
    },
    loginFailure(state) {
      state.status.loggedIn = false;
      state.user = null;
    },
    logout(state) {
      state.status.loggedIn = false;
      state.user = null;
      localStorage.removeItem("user");
    },
    registerSuccess(state) {
      state.status.loggedIn = false;
    },
    registerFailure(state) {
      state.status.loggedIn = false;
    }
  }
};
