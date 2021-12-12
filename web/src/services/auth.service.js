import axios from 'axios';

// const API_URL = 'http://localhost:3000/api/auth/';
const API_URL = 'http:////192.168.1.110:3000/api/auth/';

class AuthService {
  async login(user) {
    return axios
      .post(API_URL + 'login', {
        email: user.email,
        password: user.password
      })
      .then(response => {
        if (response.data.jwt) {
          localStorage.setItem('user', JSON.stringify(response.data))
        }

        return response.data
      });
  }

  logout() {
    localStorage.removeItem('user');
  }

  register(user) {
    return axios.post(API_URL + 'register', user)
  }
}

export default new AuthService();
