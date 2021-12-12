import axios from "axios";
import authHeader from "./auth-header";

// const API_URL = "http://localhost:3000/api/exercise_type";
const API_URL = "http:////192.168.1.110:3000/api/exercise_type";

class ExerciseTypeService {
  // TODO: Need to reroute to login page if 401
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
