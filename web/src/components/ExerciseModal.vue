<template>
  <div id="modal-dialog" class="modal-dialog-container">
    <div class="modal-dialog-content">
      <div class="modal-dialog-header">
        <h4>Select Exercise</h4>
      </div>
      <div class="modal-dialog-body">
        <ul class="list-group">
          <span v-for="(et, i) in exerciseTypes" :key="i">
            <a
              href="#"
              class="
                list-group-item list-group-item-action
                flex-column
                align-items-start
              "
              v-bind:class="{ active: selectedExerciseType == et }"
              @click="selectHandler(et)"
            >
              <div class="d-flex w-100 justify-content-between">
                <h5 class="mb-1">{{ et.name }}</h5>
                <small>{{ et.data_type }}</small>
              </div>
              <p class="mb-1">
                {{ et.description }}
              </p>
            </a>
          </span>
        </ul>
      </div>
      <div class="modal-dialog-footer">
        <div class="btn-wrapper">
          <button
            type="button"
            class="button btn btn-primary btn-sm"
            @click="addExercise"
          >
            Add Exercise
          </button>
          <div />

          <router-link class="button" :to="{ name: 'exercise_type' }">
            <button type="button" class="button btn btn-primary btn-sm">
              New Type
            </button>
          </router-link>
          <div />
          <button
            type="button"
            class="button btn btn-secondary btn-sm close-btn"
            @click="closeModal"
          >
            Cancel
          </button>
        </div>
      </div>

      <div v-if="errorMessage" class="error-message alert alert-danger">
        {{ errorMessage }}
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ExerciseModal",
  emits: ["closeModalEvent", "addExerciseEvent"],
  data() {
    return {
      selectedExerciseType: {},
      exerciseTypes: [],
      errorMessage: "",
    };
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
    closeModal() {
      this.$emit("closeModalEvent");
    },
    addExercise() {
      console.log(this.selectedExerciseType);
      if (JSON.stringify(this.selectedExerciseType) == "{}") {
        this.errorMessage = "Please select an exercise to start";
        return;
      }
      this.$emit("addExerciseEvent", this.selectedExerciseType);
      this.$emit("closeModalEvent");

      // TODO: BUG: view not refreshing when exercise added
      this.$router.go();
    },
    selectHandler(et) {
      this.selectedExerciseType = et;
    },
  },
};
</script>

<style scoped>
.modal-dialog-container {
  position: fixed; /* Stay in place */
  z-index: 9999; /* Sit on top */
  left: 0;
  top: 0;
  width: 100%; /* Full width */
  height: 100%; /* Full height */
  overflow: auto; /* Enable scroll if needed */
  background-color: rgb(0, 0, 0); /* Fallback color */
  background-color: rgba(0, 0, 0, 0.4); /* Black w/ opacity */
}
.modal-dialog-content {
  background-color: #fefefe;
  margin: 10% auto;
  padding: 20px;
  border: 1px solid #888;
  border-radius: 0.3rem;
  width: 80%;
}
.modal-dialog-footer {
  text-align: center;
}

.btn-wrapper {
  padding-top: 15px;
  width: 100%;
  left: 0;
  right: 0;
  display: flex;
  flex-wrap: nowrap;
}
.button {
  width: 100%;
  text-align: center;
}
.btn-wrapper > div {
  width: 3%;
}
.error-message {
  margin-top: 15px;
}
</style>