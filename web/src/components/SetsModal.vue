<template>
  <div id="modal-dialog" class="modal-dialog-container">
    <div class="modal-dialog-content">
      <Form @submit="addSet" :validation-schema="schema">
        <div class="modal-dialog-header">
          <h4>Set nr#{{ newSetNumber }}</h4>
        </div>
        <div class="modal-dialog-body">
          <div class="form-group">
            <label for="reps">Reps</label>
            <Field
              v-model="set.reps"
              name="reps"
              type="number"
              class="form-control"
              id="reps"
            />
            <ErrorMessage name="reps" class="error-feedback" />
          </div>
          <div class="form-group">
            <label for="resistance">{{ exerciseType.measurement_type }}</label>
            <Field
              v-model="set.resistance"
              name="resistance"
              type="number"
              class="form-control"
            />
            <ErrorMessage name="resistance" class="error-feedback" />
          </div>
        </div>
        <div class="modal-dialog-footer">
          <div class="form-group btn-wrapper">
            <button
              type="button"
              class="btn btn-primary btn-sm"
              @click="addSet"
            >
              Add
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
      </Form>
    </div>
  </div>
</template>

<script>
import { Form, Field, ErrorMessage } from "vee-validate";
import * as yup from "yup";

export default {
  name: "SetsModal",
  components: {
    Form,
    Field,
    ErrorMessage,
  },
  props: {
    exercise: {
      type: Object,
      required: true,
    },
    exerciseType: {
      type: Object,
      required: true,
    },
  },
  emits: ["closeModal"],
  computed: {
    newSetNumber() {
      if (!this.exercise.sets) return 1;
      return this.exercise.sets.length + 1;
    },
  },
  data() {
    const schema = yup.object().shape({
      reps: yup.number().required("Repetitions require"),
      resistance: yup.number().required("Resistance require"),
    });
    return {
      schema,
      set: {
        reps: 0,
        resistance: 0,
      },
      errorMessage: "",
    };
  },
  methods: {
    // TODO: Cleanup the following
    // TODO: Implement the dropset checkbox
    addSet() {
      let exercise = this.exercise;
      if (!Array.isArray(exercise.sets)) exercise.sets = [];
      exercise.sets.push({
        rep: parseInt(this.set.reps),
        resistance: parseInt(this.set.resistance),
      });
      console.log(exercise);
      console.log(this.exerciseType);
      this.$store
        .dispatch("exercises/updateOne", exercise)
        .then(() => {
          this.$emit("closeModal");
        })
        .catch((err) => {
          this.errorMessage = err;
          console.log(err);
        });
    },
    closeModal() {
      this.$emit("closeModal");
    },
  },
};
</script>

<style scoped>
.modal-dialog-container {
  position: fixed;
  z-index: 9999;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  overflow: auto;
  background-color: rgb(0, 0, 0);
  background-color: rgba(0, 0, 0, 0.4);
}
.modal-dialog-content {
  background-color: #fefefe;
  margin: 10% auto;
  margin-top: 50%;
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