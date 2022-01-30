<template>
  <div class="container">
    <header class="jumbotron">
      <h3>Exercise Types</h3>
    </header>
  </div>

  <div class="input-group mb-3 search-bar">
    <input
      type="text"
      class="form-control"
      placeholder="Search"
      v-model="searchTerm"
    />
  </div>

  <dev class="filter">
    <ul class="pagination justify-content-center">
      <li class="page-item" @click="selectFilter('All')">
        <a class="page-link">All</a>
      </li>
      <li class="page-item" @click="selectFilter('sets')">
        <a class="page-link">Sets</a>
      </li>
      <li class="page-item" @click="selectFilter('single-value')">
        <a class="page-link">Single values</a>
      </li>
    </ul>
  </dev>

  <div
    v-for="(e, i) in exerciseTypesFilter"
    :key="i"
    class="card mx-auto exercise"
  >
    <div class="card-body">
      <h4 class="card-title">{{ e.name }}</h4>
      <p class="card-text">{{ e.description }}</p>
      <p class="card-text">{{ e.data_type }} : {{ e.measurement_type }}</p>

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
      searchTerm: "",
      filter: "All",
      exerciseTypes: [],
      deleteModal: false,
      deleteExerciseId: "",
      deleteMessage: "Are you sure?",
      continueButtonMessage: "Yes",
    };
  },
  computed: {
    exerciseTypesFilter() {
      let searchTerm = this.searchTerm.toLowerCase();
      return this.exerciseTypes.filter((exerciseType) => {
        if (
          exerciseType.name.toLowerCase().includes(searchTerm) ||
          exerciseType.description.toLowerCase().includes(searchTerm)
        ) {
          if (this.filter !== "All") {
            return exerciseType.data_type === this.filter;
          }
          return true;
        }
      });
    },
  },
  async created() {
    try {
      await this.$store.dispatch("exerciseTypes/fetchAll");
    } catch (err) {
      console.error(err);
    }
    this.exerciseTypes = this.$store.getters["exerciseTypes/getAll"].data;
  },
  methods: {
    selectFilter(filter) {
      this.filter = filter;
    },
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
        .dispatch("exerciseTypes/deleteOne", this.deleteExerciseId)
        .then(
          () => {},
          (err) => {
            console.error(err);
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
.search-bar{
  padding: 1.25rem;
}
</style>