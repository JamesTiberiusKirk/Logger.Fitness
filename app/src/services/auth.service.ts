import { UserLoginDTO, UserRegisterDTO } from "@/types/user";
import axios, { AxiosResponse } from "axios";

const API_URL = "/api/auth";

class AuthService {
  login(user: UserLoginDTO): Promise<AxiosResponse<any, any>> {
    return axios.post(API_URL + "/login", user);
  }

  register(user: UserRegisterDTO): Promise<AxiosResponse<any, any>> {
    return axios.post(API_URL + "/register", user);
  }
}

export default new AuthService();
