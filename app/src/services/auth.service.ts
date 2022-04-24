import { UserLoginDTO, UserRegisterDTO, UserStore } from "@/types/user";
import axios, { AxiosError, AxiosResponse } from "axios";
import API_BASE_URL from "@/services/config.service";

const API_URL = API_BASE_URL+"/auth";

class AuthService {
  login(user: UserLoginDTO): Promise<AxiosResponse<any, any>> {
    return axios.post(API_URL + "/login", user);
  }

  register(user: UserRegisterDTO): Promise<AxiosResponse<any, any>> {
    return axios.post(API_URL + "/register", user);
  }

  oauthCallback({state, code, scope, authUser, prompt}:any): Promise<AxiosResponse<UserStore, AxiosError>> {
    const uri = API_URL + `/google/callback?state=${encodeURI(state)}&code=${encodeURI(code)}&scope=${encodeURI(scope)}&authuser=${encodeURI(authUser)}&prompt=${encodeURI(prompt)}`
    return axios.get(uri);
  }
}

export default new AuthService();
