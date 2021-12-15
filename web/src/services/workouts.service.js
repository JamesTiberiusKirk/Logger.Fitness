import axios from "axios";
import authHeader from "./auth-header";

const API_URL = "/api/workouts";

class WorkoutsService {
  // TODO: Need to reroute to login page if 401, 403 (missing key)
  //  on a second thought, need to implement 401 returns on the backend first...might not need to???
  start(workout) {
    let options = { headers: authHeader() };
    return axios.post(API_URL + "/start", workout, options);
  }

  stop(time) {
    let options = { headers: authHeader() };
    return axios.post(`${API_URL}?stop_time=${time}` + "/stop", options);
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
