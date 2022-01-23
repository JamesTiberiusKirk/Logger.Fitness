import axios from "axios";
import authHeader from "./auth-header";

const API_URL = "/api/v2/exercise_type";

class ExerciseTypeService {
  getExerciseTypes() {
    let options = { headers: authHeader() };
    return axios.get(API_URL, options);
  }

  newExerciseType(exerciseType) {
    let options = { headers: authHeader() };
    return axios.post(API_URL, exerciseType, options);
  }

  updateExerciseType(exerciseType) {
    let options = { headers: authHeader() };
    return axios.put(API_URL, exerciseType, options);
  }

  deleteExerciseType(id){
    let options = { headers: authHeader() };
    return axios.delete(`${API_URL}?id=${id}`, options);
  }
}

export default new ExerciseTypeService();
