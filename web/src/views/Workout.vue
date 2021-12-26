
<template>
  <div class="container">
    <header class="jumbotron">
      <h3>{{ workout.workout.title }}</h3>
      <p class="lead">
        {{ getDate }} | {{ getTime }}
        <br />
        {{ workout.workout.notes }}
      </p>
    </header>

    <span class="exercise-list" v-for="(e, i) in workout.exercises" :key="i">
      <ExerciseCard
        :exercise="e"
        :exerciseType="exerciseTypes[e.exercise_type_id]"
      />
    </span>
    <div class="exercise-list-padding"/>

    <div class="bottom-pane">
      <button
        v-if="workout.workout.end_time == -1"
        v-on:click="stopModalToggle()"
        href="#"
        class="btn btn-danger stop-button"
      >
        Stop Workout
      </button>
      <div/>
      <button
        v-if="workout.workout.end_time == -1"
        v-on:click="exerciseModalToggle"
        href="#"
        class="btn btn-primary add-exercise-button"
      >
       Add Exercise 
      </button>
    </div>


    <ExerciseModal
      v-if="exerciseModal.show"
      @closeModalEvent="exerciseModalToggle"
      @addExerciseEvent="addExercise"
    />

    <DeleteModal
      class="custom-modal"
      v-if="stopModal.show"
      :deleteMessage="stopModal.message"
      :continueButtonMessage="stopModal.continueButtonMessage"
      @closeDeleteModalEvent="stopModalToggle"
      @deleteRecordEvent="stopExercise"
    />
  </div>
</template>

<script>
import DeleteModal from "../components/DeleteModal.vue";
import ExerciseModal from "../components/ExerciseModal.vue";
import ExerciseCard from "../components/ExerciseCard.vue";

// TODO: Need do implement the exercise delete button

export default {
  name: "Workout",
  components: {
    DeleteModal,
    ExerciseModal,
    ExerciseCard,
  },
  data() {
    return {
      stopModal: {
        show: false,
        message: "Are you sure you want to stop this workout?",
        continueButtonMessage: "Yes",
      },
      exerciseModal: {
        show: false,
        message: "Chose exercise",
      },
      loading: false,
      successMessage: "",
      errorMessage: "",
      workout: {
        exercises: [],
        workout: {
          workout_id: "",
          title: "",
          notes: "",
          start_time: -1,
        },
      },
      exerciseTypes: new Map(),
    };
  },
  async mounted() {
    let workoutId = this.$route.query.id;
    if (!workoutId) this.$router.push("workouts_list");

    try {
      await this.$store.dispatch("workouts/fetchAll");
      await this.$store.dispatch("exerciseTypes/fetchAll");
    } catch (err) {
      console.error(err);
    }
    this.workout = this.$store.getters["workouts/getOneById"](workoutId);
    this.exerciseTypes = this.$store.getters["exerciseTypes/getAllAsMap"];
  },
  computed: {
    getDate() {
      if (this.workout.workout.start_time == -1) return;
      return new Date(this.workout.workout.start_time).toLocaleDateString();
    },
    getTime() {
      if (this.workout.workout.start_time == -1) return;
      let time = new Date(this.workout.workout.start_time);
      return time.getHours() + ":" + time.getMinutes();
    },
  },
  methods: {
    getModalName(exerciseType) {
      return exerciseType.data_type === "sets"
        ? "setsModal"
        : "singleValueModal";
    },
    async addExercise(exerciseType) {
      let exercise = {
        exercise_type_id: exerciseType.exercise_type_id,
        workout_id: this.workout.workout.workout_id,
      };
      try {
        await this.$store.dispatch("exercises/new", exercise);
      } catch (err) {
        console.error(err);
      }
    },
    exerciseModalToggle() {
      this.exerciseModal.show = !this.exerciseModal.show;
    },
    stopModalToggle() {
      this.stopModal.show = !this.stopModal.show;
    },
    stopExercise() {
      let unixEndTime = Math.floor(Date.now() / 1000);
      this.$store.dispatch("workouts/stop", unixEndTime).then(
        () => {
          this.stopModal.show = false;

          // TODO: BUG: this is to refresh the page because the stop button
          //  will not disappear on it's own.
          this.$router.go();
        },
        (err) => {
          console.error(err);
        }
      );
    },
  },
};
</script>

<style scoped>
.exercise-add-fab {
  margin-bottom: 40px;
}
.bottom-pane {
  width: 95%;
  position: fixed;
  bottom: 65px;
  z-index: 901;
  /* Below code aligns it to the center */
  left: 0;
  right: 0;
  margin: 0 auto;

  display: flex;
  flex-wrap: nowrap;
}
.bottom-pane > button {
  width: 100%;
  /* margin: 5%; */
  text-align: center;
}
.bottom-pane > div {
  width: 2%;
}
.exercise-list-padding {
  margin-bottom: 40%;
}
</style>