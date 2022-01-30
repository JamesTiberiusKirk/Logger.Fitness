<template>
  <div id="modal-dialog" class="modal-dialog-container">
    <div class="modal-dialog-content">
      <Form @submit="addSet" :validation-schema="schema">
        <div class="modal-dialog-header">
          <h4>Set #{{ newSetNumber }}</h4>
        </div>
        <div class="modal-dialog-body">
          <div class="form-group input-group input-group mb-3">
            <div class="input-group-prepend">
              <span id="reps-label" class="label input-group-text">Reps</span>
              <i
                class="input-group-text fas material-icons"
                @click="removeFromReps"
              >
                remove
              </i>
            </div>
            <Field
              aria-label="Small"
              aria-describedby="reps-label"
              v-model="set.reps"
              name="reps"
              type="number"
              class="form-control"
              id="reps"
            />
            <div class="input-group-append" @click="addToReps">
              <i class="input-group-text fas material-icons"> add </i>
            </div>
            <ErrorMessage name="reps" class="error-feedback" />
          </div>

          <div class="form-group input-group input-group mb-3">
            <div class="input-group-prepend">
              <span id="resistance-label" class="label input-group-text">{{
                exerciseType.measurement_type
              }}</span>
              <i
                class="input-group-text fas material-icons"
                @click="removeFromResistance"
              >
                remove
              </i>
            </div>
            <Field
              aria-label="Small"
              aria-describedby="resistance-label"
              v-model="set.resistance"
              name="resistance"
              type="number"
              class="form-control"
            />
            <div class="input-group-append" @click="addToResistance">
              <i class="input-group-text fas material-icons"> add </i>
            </div>
            <ErrorMessage name="resistance" class="error-feedback" />
          </div>

          <div class="form-group input-group input-group mb-3">
            <div class="input-group-prepend">
              <span id="is_drop_set-label" class="form-check-label"
                >Drop-set?</span
              >
            </div>
            <Field
              aria-label="Small"
              aria-describedby="is_drop_set-label"
              id="is_drop_set"
              v-model="set.is_drop_set"
              name="is_drop_set"
              type="checkbox"
              class="form-check-input"
            />
            <ErrorMessage name="is_drop_set" class="error-feedback" />
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
    let reps = 0;
    let resistance = 0;
    const lastSet = this.exercise.sets[this.exercise.sets.length - 1];
    const schema = yup.object().shape({
      reps: yup.number().required("Repetitions require"),
      resistance: yup.number().required("Resistance require"),
    });

    if (this.exercise.sets.length > 0) {
      reps = lastSet.reps;
      resistance = lastSet.resistance;
    }
    return {
      schema,
      set: {
        reps: reps,
        resistance: resistance,
        is_drop_set: false,
      },
      errorMessage: "",
    };
  },
  methods: {
    removeFromReps() {
      if (this.set.reps<=0){
        return
      }
      this.set.reps--
    },
    addToReps() {
      if (this.set.reps>=1000){
        return
      }
      this.set.reps++
    },
    removeFromResistance() {
      if (this.set.resistance<=0){
        return
      }
      this.set.resistance--
    },
    addToResistance() {
      if (this.set.resistance>=1000){
        return
      }
      this.set.resistance++
    },
    // TODO: Cleanup the following
    addSet() {
      let exercise = this.exercise;
      if (!Array.isArray(exercise.sets)) exercise.sets = [];
      exercise.sets.push({
        reps: parseInt(this.set.reps),
        resistance: parseInt(this.set.resistance),
        is_drop_set: this.set.is_drop_set === "on",
      });
      console.log(exercise.sets);
      this.$store
        .dispatch("exercises/updateOne", exercise)
        .then(() => {
          this.$emit("closeModal");
        })
        .catch((err) => {
          this.errorMessage = err;
          console.error(err);
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

.checkbox {
  text-align: center;
  vertical-align: center;
}

.label {
  width: 4rem;
}
</style>