import { UserLoginDTO, UserRegisterDTO, UserStore } from "@/types/user";
import AuthService from "@/services/auth.service";
import router from '@/router';

// TODO: figure out how typescript works with vuex
// TODO: blow my brains out bc of vuex + composition api + ionic + typescript....

// NOTE: separating this out and using rawUser to clarify weather the user object
//  existed was the only way I could get it to work.... user === {} did not work 
//  properly...
const rawUser = localStorage.getItem("user")
const user = JSON.parse(rawUser || "{}");
console.log(user);


export type AuthState = {
  status: {
    loggedIn: boolean,
  },
  user?: UserStore,
}

const loggedIn = rawUser ? true : false
const initialState: AuthState = {
  status: {
    loggedIn
  },
  user
};


export const auth = {
  namespaced: true,
  state: initialState,
  actions: {
    async login({ commit }: any, user: UserLoginDTO) {
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
    logout({ commit }: any) {
      commit("logout");
      router.push('/login')
    },

    // TODO: this needs testing and possibly rework
    async register({ commit }: any, user: UserRegisterDTO) {
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
    storeJwt(state: AuthState, jwt: string) {
      if (state.user) state.user.jwt = jwt;
      localStorage.setItem("user", JSON.stringify(state.user));
    },
    loginSuccess(state: AuthState, user: UserStore) {
      state.status.loggedIn = true;
      state.user = user;
      localStorage.setItem("user", JSON.stringify(user));
    },
    loginFailure(state: AuthState) {
      state.status.loggedIn = false;
      state.user = undefined;
    },
    logout(state: AuthState) {
      state.status.loggedIn = false;
      state.user = undefined;
      localStorage.removeItem("user");
    },
    registerSuccess(state: any) {
      state.status.loggedIn = false;
    },
    registerFailure(state: any) {
      state.status.loggedIn = false;
    }
  },
  getters: {
    statusLoggedIn(state: AuthState) { return state.status.loggedIn },
  }
};
