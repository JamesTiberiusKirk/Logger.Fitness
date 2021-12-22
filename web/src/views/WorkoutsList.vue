<template>
  <div class="container">
    <header class="jumbotron">
      <h3>Workouts List</h3>
    </header>
  </div>
  <div v-for="(e, i) in workoutsData" :key="i" class="card mx-auto exercise">
    <div class="card-body">
      <h4 class="card-title">{{ e.workout.title }}</h4>
      <p class="card-text">Notes: {{ e.workout.notes }}</p>
      <p class="card-text" v-if="e.workout.end_time == -1">Still Active</p>
      <button
        v-on:click="showDeleteModalToggle(e.workout.workout_id)"
        href="#"
        class="float-right card-btn btn btn-danger"
      >
        Delete
      </button>
      <router-link
        :to="{ name: 'workout_form', query: { id: e.workout.workout_id } }"
      >
        <a href="#" class="float-right card-btn btn btn-primary">Edit</a>
      </router-link>
      <router-link
        :to="{ name: 'workout', query: { id: e.workout.workout_id } }"
      >
        <a href="#" class="float-right card-btn btn btn-primary"
          >Go To Workout</a
        >
      </router-link>
    </div>
  </div>

  <deleteModal
    class="custom-modal"
    v-if="showDeleteModal"
    :deleteMessage="deleteMessage"
    :continueButtonMessage="continueButtonMessage"
    @closeDeleteModalEvent="closeDeleteModal"
    @deleteRecordEvent="deleteWorkout"
  />

  <fab @fabClickEvent="fabClick" />
</template>

<script>
import DeleteModal from "../components/DeleteModal.vue";
import Fab from "../components/Fab.vue";
import { mapGetters, mapState } from "vuex";

// TODO: do not display the add fab if there is an active workout

export default {
  name: "WorkoutsList",
  components: {
    deleteModal: DeleteModal,
    fab: Fab,
  },
  data() {
    return {
      workoutsData: [],
      showDeleteModal: false,
      deleteWorkoutId: "",
      deleteMessage: "Are you sure?",
      continueButtonMessage: "Yes",
    };
  },
  computed: {
    ...mapState({
      workoutState: "workouts",
    }),
    ...mapGetters({
      getWorkouts: "workouts/getAllInReverse",
    }),
  },
  watch: {
    "$state.state.workouts":{
      deep: true,
      handler: function (oldValue, newValue) {
        console.log("old: ", oldValue);
        console.log("new: ", newValue);
        this.workoutsData = this.getWorkouts;
        console.log("getter: ", this.getWorkouts);
      },
    },
  },

  async created() {
    try {
      await this.$store.dispatch("workouts/fetchAll");
    } catch (err) {
      console.log(err);
    }

    this.workoutsData = this.getWorkouts;
    // this.workoutsData = this.$store.getters["workouts/getAll"];
  },
  methods: {
    showDeleteModalToggle(id) {
      this.deleteWorkoutId = id;
      this.showDeleteModal = true;
    },
    closeDeleteModal() {
      this.showDeleteModal = false;
    },
    async deleteWorkout() {
      this.showDeleteModal = false;
      await this.$store.dispatch("workouts/deleteOne", this.deleteWorkoutId);
      // TODO: BUG: so...this is just here to fix the fact that the page does not update when the data inside the
      this.$router.go()
    },
    fabClick() {
      this.$router.push("workout_form");
    },
  },
};
</script>
<style>
.exercise {
  margin-bottom: 5%;
  width: 90%;
}
.card-btn {
  margin-left: 1%;
}
</style>