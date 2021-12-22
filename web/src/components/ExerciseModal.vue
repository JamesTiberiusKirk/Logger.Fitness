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
            class="btn btn-primary btn-sm"
            @click="addExercise"
          >
            Add exercise
          </button>
          <button
            type="button"
            class="btn btn-secondary btn-sm close-btn"
            @click="closeModal"
          >
            Cancel
          </button>
        </div>
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
    };
  },
  async created() {
    try {
      await this.$store.dispatch("exerciseTypes/fetchAll");
    } catch (err) {
      console.log(err);
    }

    this.exerciseTypes = this.$store.getters["exerciseTypes/getAll"].data;
  },
  methods: {
    closeModal() {
      this.$emit("closeModalEvent");
    },
    addExercise() {
      this.$emit("addExerciseEvent", this.selectedExerciseType);
      this.$emit("closeModalEvent");

      // TODO: BUG: view not refreshing when exercise added
      this.$router.go()
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
.close-btn {
  margin-left: 1em;
}

.btn-wrapper {
  padding-top: 15px;
  width: 100%;
  /* bottom: 65px; */
  /* Below code aligns it to the center */
  left: 0;
  right: 0;
  margin: 0 auto;

  display: flex;
  flex-wrap: nowrap;
}
.btn-wrapper > button {
  padding: 10px;
  width: 100%;
  text-align: center;
}
</style>