import axios from "axios";

const API_URL = "/api/v2/auth";

class AuthService {
  login(user) {
    return axios.post(API_URL + "/login", user);
  }

  register(user) {
    return axios.post(API_URL + "/register", user);
  }
}

export default new AuthService();
