import { UserLoginDTO, UserRegisterDTO } from "@/types/user";
import axios from "axios";

const API_URL = "/api/auth";

class AuthService {
  login(user: UserLoginDTO) {
    return axios.post(API_URL + "/login", user);
  }

  register(user: UserRegisterDTO) {
    return axios.post(API_URL + "/register", user);
  }
}

export default new AuthService();
