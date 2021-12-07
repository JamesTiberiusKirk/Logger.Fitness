import axios from 'axios';
import authHeader from './auth-header';

const API_URL = 'http://localhost:3000/api/exercise_types/';

class ExerciseTypesService {

  getExerciseTypes() {
    return axios.get(API_URL, { headers: authHeader() });
  }

  newExerciseType(exerciseType) {
    return axios.get(API_URL, { body: exerciseType,headers: authHeader() });
  }

}

export default new ExerciseTypesService();
