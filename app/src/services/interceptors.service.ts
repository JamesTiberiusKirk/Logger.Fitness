import axios, { AxiosError, AxiosResponse } from "axios";
import store from "../store";
import router from "../router";

function handle401(err: any) {
  if (err.response.status !== 401) return;
  store.dispatch("auth/logout");
  router.push("/login");
}

export default function setupHttpInterceptors() {
  axios.interceptors.response.use(
    (res: AxiosResponse) => {
      if (res.headers["x-access-token"]) {
        const jwt: string = res.headers["x-access-token"];
        store.commit("auth/storeJwt", jwt);
      }
      return Promise.resolve(res);
    },
    (err: AxiosError) => {
      handle401(err);
      return Promise.reject(err);
    }
  );
}
