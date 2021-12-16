
<template>
  <div class="container">
    <header class="jumbotron">
      <h3>{{ this.workout.title }}</h3>
      <p class="lead">
        {{ getDate }} | {{ getTime }}
        <br />
        {{ this.workout.notes }}
      </p>
    </header>
    <!-- <div class="col-md-12">
      <div class="card card-container">
        <div class="card-body">
          <p class="card-text"></p>
        </div>
      </div>
    </div> -->
    <div class="bottom-pane">
      <button
        v-if="workout.end_time == -1"
        v-on:click="showStopModal(workout.workout_id)"
        href="#"
        class="btn btn-danger stop-button"
      >
        Stop Workout
      </button>
    </div>

    <deleteModal
      class="custom-modal"
      v-if="stopModal"
      :deleteMessage="stopModalMessage"
      :continueButtonMessage="continueButtonMessage"
      @closeDeleteModalEvent="closeStopModal"
      @deleteRecordEvent="stopExercise"
    />
  </div>
</template>

<script>
import DeleteModal from "../components/DeleteModal.vue";

// TODO: hide stop button if not active

export default {
  name: "Workout",
  components: {
    deleteModal: DeleteModal,
  },
  data() {
    return {
      stopModal: false,
      loading: false,
      stopModalMessage: "Are you sure you want to stop this workout?",
      continueButtonMessage: "Yes",
      successMessage: "",
      errorMessage: "",
      workout: {
        workout_id: "",
        title: "",
        notes: "",
        start_time: -1,
      },
      exercises: [],
    };
  },
  mounted() {
    let id = this.$route.query.id;
    if (id) {
      this.workout = this.$store.getters["workouts/getOneById"](id);
    } else {
      // Error
      console.log("no id provided");
    }
  },
  computed: {
    getDate() {
      if (this.workout.start_time == -1) return;
      return new Date(this.workout.start_time).toLocaleDateString();
    },
    getTime() {
      if (this.workout.start_time == -1) return;
      let time = new Date(this.workout.start_time);
      return time.getHours() + ":" + time.getMinutes();
    },
  },
  methods: {
    closeStopModal() {
      this.stopModal = false;
    },
    showStopModal() {
      this.stopModal = true;
    },
    stopExercise() {
      let unixEndTime = Math.floor(Date.now() / 1000);
      this.$store.dispatch("workouts/stop", unixEndTime).then(
        (res) => {
          console.log(res);
          this.stopModal = false;
          
        },
        (err) => {
          console.log(err);
        }
      );
    },
  },
};
</script>

<style>
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