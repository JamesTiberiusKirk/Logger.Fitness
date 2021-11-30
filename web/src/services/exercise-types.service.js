import axios from 'axios';
import authHeader from './auth-header';

const API_URL = 'http://localhost:3000/extp/';

class ExerciseTypesService {

  getExerciseTypes() {
    return axios.get(API_URL, { headers: authHeader() });
  }

  postExerciseType() {
    return axios.get(API_URL, { headers: authHeader() });
  }

}

export default new ExerciseTypesService();
