<template>
  <div class="container">
    <header class="jumbotron">
      <h3>Exercise Types</h3>
    </header>
  </div>
  <div v-for="(e, i) in this.exerciseTypes" :key="i" class="card mx-auto exercise">
    <div class="card-body">
      <h4 class="card-title">{{ e.name }}</h4>
      <p class="card-text">{{ e.description }}</p>
      <a href="#" class="float-right card-btn btn btn-danger">Delete</a>
      <a href="#" class="float-right card-btn btn btn-primary">Edit</a>
    </div>
  </div>
</template>

<script>
import ExerciseTypesService from "../services/exercise-types.service";

export default {
  name: "ExerciseType",
  data() {
    return {
      exerciseTypes: null,
    };
  },
  async created() {
    ExerciseTypesService.getExerciseTypes().then((res) => {
      this.exerciseTypes = res.data
      console.log(this.exerciseTypes)
    });
  },
};
</script>
<style>
.exercise {
  margin-bottom: 5%;
  width: 90%;
}
.card-btn{
  margin-left:1%;
}
</style>