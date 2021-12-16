import axios from "axios";
import authHeader from "./auth-header";

const API_URL = "/api/workouts";

class WorkoutsService {
  start(workout) {
    let options = { headers: authHeader() };
    return axios.post(API_URL + "/start", workout, options);
  }

  stop(time) {
    let options = { headers: authHeader() };
    return axios.post(`${API_URL}/stop?end_time=${time}`, {}, options);
  }

  getAll() {
    let options = { headers: authHeader() };
    return axios.get(API_URL, options);
  }

  getActive() {
    let options = { headers: authHeader() };
    return axios.get(API_URL, options);
  }

  delete(id) {
    let options = { headers: authHeader() };
    return axios.delete(`${API_URL}?workout_id=${id}`, options);
  }

  edit(update) {
    let options = { headers: authHeader() };
    return axios.put(API_URL, update, options);
  }
}

export default new WorkoutsService();
