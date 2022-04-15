import { Workout, WorkoutGroup } from "@/types/workout";
import axios, { AxiosResponse } from "axios";
import authHeader from "./auth-header";
import API_BASE_URL from "./config.service";

const API_URL = API_BASE_URL+"/workouts";

class WorkoutsService {
  start(workout: Workout): Promise<AxiosResponse<Workout, Error>> {
    const options = { headers: authHeader() };
    return axios.post(API_URL + "/start", workout, options);
  }

  stop(time: string): Promise<AxiosResponse<Workout, Error>> {
    const options = { headers: authHeader() };
    return axios.post(`${API_URL}/stop?end_time=${time}`, {}, options);
  }

  getAll(): Promise<AxiosResponse<WorkoutGroup[], Error>> {
    const options = { headers: authHeader() };
    return axios.get(API_URL, options);
  }

  getActive(): Promise<AxiosResponse<WorkoutGroup, Error>> {
    const options = { headers: authHeader() };
    return axios.get(API_URL, options);
  }

  delete(id: string): Promise<AxiosResponse<void, Error>> {
    const options = { headers: authHeader() };
    return axios.delete(`${API_URL}?workout_id=${id}`, options);
  }

  edit(update: object): Promise<AxiosResponse<Workout, Error>> {
    const options = { headers: authHeader() };
    return axios.put(API_URL, update, options);
  }
}

export default new WorkoutsService();
