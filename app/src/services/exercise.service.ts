import axios from "axios";
import authHeader from "./auth-header";

const API_URL = "/api/exercises";

class ExerciseService {
  getAll() {
    const options = { headers: authHeader() };
    return axios.get(API_URL, options);
  }

  getAllInWorkout(workoutID: string) {
    const options = { headers: authHeader() };
    return axios.get(`${API_URL}?workout_id=${workoutID}`, options);
  }

  new(exercise) {
    const options = { headers: authHeader() };
    return axios.post(API_URL, exercise, options);
  }

  edit(exercise) {
    const options = { headers: authHeader() };
    return axios.put(API_URL, exercise, options);
  }

  delete(id: string){
    const options = { headers: authHeader() };
    return axios.delete(`${API_URL}?exercise_id=${id}`, options);
  }
}

export default new ExerciseService();
