import { UserLoginDTO, UserRegisterDTO, UserStore } from "@/types/user";
import axios, { AxiosError, AxiosResponse } from "axios";
import API_BASE_URL from "@/services/config.service";

const API_URL = API_BASE_URL + "/auth";

interface ServicesResponse<D> {
  data?: D,
  err?: AxiosError,
}
class AuthService {
  login(user: UserLoginDTO): Promise<AxiosResponse<any, any>> {
    return axios.post(API_URL + "/login", user);
  }

  register(user: UserRegisterDTO): Promise<AxiosResponse<any, any>> {
    return axios.post(API_URL + "/register", user);
  }

  googleOauth2Login(): Promise<ServicesResponse<string>> {
    return new Promise(resolve => {
      const resp = {} as ServicesResponse<string>
      axios.get(`${API_URL}/google/login`).then(data => {
        resp.data = data.data;
        return resolve(resp);
      }).catch(err => {
        resp.err = err;
        return resolve(resp);
      })
    });
  }

  oauthCallback({ state, code, scope, authUser, prompt }: any): Promise<AxiosResponse<UserStore, AxiosError>> {
    const uri = API_URL + `/google/callback?state=${encodeURI(state)}` +
      `&code=${encodeURI(code)}&scope=${encodeURI(scope)}` +
      `&authuser=${encodeURI(authUser)}&prompt=${encodeURI(prompt)}`;

    return axios.get(uri);
  }
}

export default new AuthService();
