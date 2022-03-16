import { ExerciseType } from "@/types/exercise-type";
import axios, { AxiosResponse } from "axios";
import authHeader from "./auth-header";
import API_BASE_URL from "./config.service";

const API_URL = API_BASE_URL+"/exercise_types";

class ExerciseTypeService {
  getExerciseTypes(): Promise<AxiosResponse<ExerciseType[], Error>> {
    const options = { headers: authHeader() };
    return axios.get(API_URL, options);
  }

  newExerciseType(exerciseType: ExerciseType): Promise<AxiosResponse<ExerciseType>> {
    const options = { headers: authHeader() };
    return axios.post(API_URL, exerciseType, options);
  }

  // If I set the param type as ExerciseType it might sent the rest of the fields as empty which would wipe them on the server side.
  // TODO: need to experiment with the above
  updateExerciseType(exerciseType: object): Promise<AxiosResponse<ExerciseType>> {
    const options = { headers: authHeader() };
    return axios.put(API_URL, exerciseType, options);
  }

  deleteExerciseType(id: string): Promise<AxiosResponse<void>> {
    const options = { headers: authHeader() };
    return axios.delete(`${API_URL}?id=${id}`, options);
  }
}

export default new ExerciseTypeService();
