import { Workout } from "@/types/workout";
import axios, { AxiosResponse } from "axios";
import authHeader from "./auth-header";

const API_URL = "/api/workouts";

class WorkoutsService {
  start(workout: Workout): Promise<AxiosResponse<Workout, any>> {
    const options = { headers: authHeader() };
    return axios.post(API_URL + "/start", workout, options);
  }

  stop(time: string): Promise<AxiosResponse<Workout, any>> {
    const options = { headers: authHeader() };
    return axios.post(`${API_URL}/stop?end_time=${time}`, {}, options);
  }

  getAll(): Promise<AxiosResponse<Workout[], any>> {
    const options = { headers: authHeader() };
    return axios.get(API_URL, options);
  }

  getActive(): Promise<AxiosResponse<Workout, any>> {
    const options = { headers: authHeader() };
    return axios.get(API_URL, options);
  }

  delete(id: string): Promise<AxiosResponse<void, any>> {
    const options = { headers: authHeader() };
    return axios.delete(`${API_URL}?workout_id=${id}`, options);
  }

  // If I set the param type as Exercise it might sent the rest of the fields as empty which would wipe them on the server side.
  // TODO: need to experiment with the above
  edit(update: object): Promise<AxiosResponse<Workout, any>> {
    const options = { headers: authHeader() };
    return axios.put(API_URL, update, options);
  }
}

export default new WorkoutsService();
