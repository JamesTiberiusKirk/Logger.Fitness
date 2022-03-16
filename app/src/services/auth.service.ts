import { UserLoginDTO, UserRegisterDTO } from "@/types/user";
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
}

export default new AuthService();
