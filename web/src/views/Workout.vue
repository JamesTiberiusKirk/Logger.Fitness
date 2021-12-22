
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
    <span v-for="(e, i) in workout.exercises" :key="i">
      <div class="col-md-12">
        <div class="card card-container">
          <div class="card-body">
            <p class="card-text">
              {{ exerciseTypes[e.exercise_type_id].name }}
            </p>
            <p>
              {{ exerciseTypes[e.exercise_type_id].data_type }}
              <span
                v-if="exerciseTypes[e.exercise_type_id].data_type === 'sets'"
              >
                <span v-if="e.sets">
                  <span v-for="(s, i) in e.sets" :key="i">
                    <!-- TODO: properly display the sets -->
                    {{ JSON.stringify(e.sets[i]) }}
                  </span>
                </span>
                <button @click="setModalToggle(e)" class="btn btn-primary">
                  Add sets
                </button>
              </span>
              <span
                v-if="
                  exerciseTypes[e.exercise_type_id].data_type === 'single_value'
                "
              >
                {{ JSON.stringify(e.single_value) }}
              </span>
            </p>
          </div>
        </div>
      </div>
    </span>
    <div class="bottom-pane">
      <button
        v-if="workout.workout.end_time == -1"
        v-on:click="showStopModal(workout.workout.workout_id)"
        href="#"
        class="btn btn-danger stop-button"
      >
        Stop Workout
      </button>
    </div>

    <setsModal
      v-if="setsModal.show"
      :exercise="setsModal.currentExercise"
      :exerciseType="setsModal.currentExerciseType"
      @closeModal="setModalToggle"
    />

    <fab class="exercise-add-fab" @fabClickEvent="exerciseModalToggle" />

    <exerciseModal
      v-if="exerciseModal.show"
      @closeModalEvent="exerciseModalToggle"
      @addExerciseEvent="addExercise"
    />

    <stopModal
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
import SetsModal from "../components/SetsModal.vue";
import Fab from "../components/Fab.vue";

export default {
  name: "Workout",
  components: {
    stopModal: DeleteModal,
    fab: Fab,
    exerciseModal: ExerciseModal,
    setsModal: SetsModal,
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
      setsModal: {
        show: false,
        currentExercise: {},
        currentExerciseType: {},
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
      console.log(err);
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
    async addExercise(exerciseType) {
      let exercise = {
        exercise_type_id: exerciseType.exercise_type_id,
        workout_id: this.workout.workout.workout_id,
      };
      try {
        await this.$store.dispatch("exercises/new", exercise);
      } catch (err) {
        console.log(err);
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
          this.stopModal = false;

          // TODO: BUG: this is to refresh the page because the stop button
          //  will not disappear on it's own.
          this.$router.go();
        },
        (err) => {
          console.log(err);
        }
      );
    },
    setModalToggle(e) {
      if (e) {
        this.setsModal.currentExercise = e;
        this.setsModal.currentExerciseType =
          this.exerciseTypes[e.exercise_type_id];
      } else {
        this.setsModal.currentExercise = {};
        this.setsModal.currentExerciseType = {};
      }
      this.setsModal.show = !this.setsModal.show;
      console.log(this.setsModal);
    },
  },
};
</script>

<style>
.exercise-add-fab {
  margin-bottom: 40px;
}
.bottom-pane {
  width: 90%;
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
  margin-right: 5px;
  text-align: center;
}
</style>