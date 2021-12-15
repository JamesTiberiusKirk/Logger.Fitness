import axios from "axios";
import store from "../store";
import router from "../router";

function handle401(err) {
  if (err.response.status !== 401) return;
  store.dispatch("auth/logout");
  router.push("/login");
}

export default function setupHttpInterceptors() {
  axios.interceptors.response.use(
    res => {
      if (res.headers["x-access-token"]) {
        let jwt = res.headers["x-access-token"];
        store.commit("auth/storeJwt", jwt);
      }
      return Promise.resolve(res);
    },
    err => {
      handle401(err);
      return Promise.reject(err);
    }
  );
}
