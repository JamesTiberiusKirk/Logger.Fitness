import { UserLoginDTO, UserRegisterDTO, UserStore } from "@/types/user";
import axios, { AxiosResponse } from "axios";
import API_BASE_URL from "@/services/config.service";

const API_URL = API_BASE_URL+"/auth";

class AuthService {
  login(user: UserLoginDTO): Promise<AxiosResponse<any, any>> {
    return axios.post(API_URL + "/login", user);
  }

  register(user: UserRegisterDTO): Promise<AxiosResponse<any, any>> {
    return axios.post(API_URL + "/register", user);
  }

  oauthCallback({state, code, scope, authUser, prompt}:any): Promise<AxiosResponse<UserStore, any>> {
    const uri = API_URL + `/google/callback?state=${state}&code=${code}&scope=${scope}&authuser=${authUser}&prompt=${prompt}`
    return axios.get(encodeURI(uri));
  }
}

export default new AuthService();
