
<template>
  <div class="container">
    <header class="jumbotron">
      <h3>Workouts List</h3>
    </header>
  </div>
  <div v-for="(e, i) in this.workouts" :key="i" class="card mx-auto exercise">
    <div class="card-body">
      <h4 class="card-title">Title: {{ e.title }}</h4>
      <p class="card-text">ID: {{ e.workout_id }}</p>
      <p class="card-text">Notes: {{ e.notes }}</p>
      <button
        v-on:click="showDeleteModal(e.workout_id)"
        href="#"
        class="float-right card-btn btn btn-danger"
      >
        Delete
      </button>
      <router-link :to="{ name: 'workout_form', query: { id: e.workout_id } }">
        <a href="#" class="float-right card-btn btn btn-primary">Edit</a>
      </router-link>
      <router-link :to="{ name: 'workout', query: { id: e.workout_id } }">
        <a href="#" class="float-right card-btn btn btn-primary">Go To Workout</a>
      </router-link>
    </div>
  </div>

  <deleteModal
    class="custom-modal"
    v-if="deleteModal"
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

export default {
  name: "WorkoutsList",
  components: {
    deleteModal: DeleteModal,
    fab: Fab,
  },
  data() {
    return {
      workouts: [],
      deleteModal: false,
      deleteWorkoutId: "",
      deleteMessage: "Are you sure?",
      continueButtonMessage: "Yes",
    };
  },
  async created() {
    let store = this.$store.getters["workouts/getAll"];
    if (store.empty) {
      try {
        await this.$store.dispatch("workouts/fetchAll");
      } catch (err) {
        console.log(err);
      }
    }
    this.workouts = store.data;
  },
  methods: {
    showDeleteModal(id) {
      this.deleteWorkoutId = id;
      this.deleteModal = true;
      console.log(this.deleteWorkoutId);
    },
    closeDeleteModal() {
      this.deleteModal = false;
    },
    deleteWorkout() {
      this.deleteModal = false;
      this.$store.dispatch("workouts/deleteOne", this.deleteWorkoutId).then(
        (res) => {
          console.log(res);
        },
        (err) => {
          console.log(err);
        }
      );
    },
    fabClick() {
      this.$router.push("workout");
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