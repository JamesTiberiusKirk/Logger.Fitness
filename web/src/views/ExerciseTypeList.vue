<template>
  <div class="container">
    <header class="jumbotron">
      <h3>Exercise Types</h3>
    </header>
  </div>
  <div
    v-for="(e, i) in this.exerciseTypes"
    :key="i"
    class="card mx-auto exercise"
  >
    <div class="card-body">
      <h4 class="card-title">{{ e.name }}</h4>
      <p class="card-text">{{ e.description }}</p>
      <p class="card-text">Data type: {{ e.data_type}}</p>
      <button
        v-on:click="showDeleteModal(e.exercise_type_id)"
        href="#"
        class="float-right card-btn btn btn-danger"
      >
        Delete
      </button>
      <router-link
        :to="{ name: 'exercise_type', query: { id: e.exercise_type_id } }"
      >
        <a href="#" class="float-right card-btn btn btn-primary">Edit</a>
      </router-link>
    </div>
  </div>

  <deleteModal
    class="custom-modal"
    v-if="deleteModal"
    :deleteMessage="deleteMessage"
    :continueButtonMessage="continueButtonMessage"
    @closeDeleteModalEvent="closeDeleteModal"
    @deleteRecordEvent="deleteExercise"
  />

  <fab @fabClickEvent="fabClick" />
</template>

<script>
import DeleteModal from "../components/DeleteModal.vue";
import Fab from "../components/Fab.vue";

export default {
  name: "ExerciseTypeList",
  components: {
    deleteModal: DeleteModal,
    fab: Fab,
  },
  data() {
    return {
      exerciseTypes: [],
      deleteModal: false,
      deleteExerciseId: "",
      deleteMessage: "Are you sure?",
      continueButtonMessage:"Yes",
    };
  },
  async created() {
    let store = this.$store.getters["exerciseType/getAll"];
    if (store.empty) {
      await this.$store.dispatch("exerciseType/fetchAll");
    }

    this.exerciseTypes = store.data;
  },
  methods: {
    showDeleteModal(id) {
      this.deleteExerciseId = id;
      this.deleteModal = true;
    },
    closeDeleteModal() {
      this.deleteModal = false;
    },
    deleteExercise() {
      this.deleteModal = false;
      this.$store
        .dispatch("exerciseType/deleteOne", this.deleteExerciseId)
        .then(
          (res) => {
            console.log(res);
          },
          (err) => {
            console.log(err);
          }
        );
    },
    fabClick() {
      this.$router.push("exercise_type");
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