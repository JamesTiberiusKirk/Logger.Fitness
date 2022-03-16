import { Exercise } from "@/types/exercise";
import axios, { AxiosResponse } from "axios";
import authHeader from "./auth-header";
import API_BASE_URL from "./config.service";

const API_URL = API_BASE_URL+"/exercises";

class ExerciseService {
  getAll(): Promise<AxiosResponse<Exercise[], any>> {
    const options = { headers: authHeader() };
    return axios.get(API_URL, options);
  }

  getAllInWorkout(workoutID: string): Promise<AxiosResponse<Exercise[], any>> {
    const options = { headers: authHeader() };
    return axios.get(`${API_URL}?workout_id=${workoutID}`, options);
  }

  new(exercise: Exercise): Promise<AxiosResponse<Exercise[], any>> {
    const options = { headers: authHeader() };
    return axios.post(API_URL, exercise, options);
  }

  // If I set the param type as Exercise it might sent the rest of the fields as empty which would wipe them on the server side.
  // TODO: need to experiment with the above
  edit(exercise: object): Promise<AxiosResponse<Exercise, any>> {
    const options = { headers: authHeader() };
    return axios.put(API_URL, exercise, options);
  }

  delete(id: string): Promise<AxiosResponse<void>> {
    const options = { headers: authHeader() };
    return axios.delete(`${API_URL}?exercise_id=${id}`, options);
  }
}

export default new ExerciseService();
