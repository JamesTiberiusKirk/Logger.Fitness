<template>
  <div id="modal-dialog" class="modal-dialog-container">
    <div class="modal-dialog-content">
      <Form @submit="addSingleValue" :validation-schema="schema">
        <div class="modal-dialog-header">
          <h4>Single value</h4>
        </div>
        <div class="modal-dialog-body">
          <div class="form-group input-group input-group mb-3">
            <div class="input-group-prepend">
              <span id="single_value-label" class="input-group-text"
                >{{exerciseType.measurement_type}}</span
              >
            </div>
            <Field
              aria-label="Small"
              aria-describedby="single_value-label"
              v-model="single_value"
              name="single_value"
              type="number"
              class="form-control"
              id="single_value"
            />
            <ErrorMessage name="single_value" class="error-feedback" />
          </div>
        </div>
        <div class="modal-dialog-footer">
          <div class="form-group btn-wrapper">
            <button
              type="button"
              class="btn btn-primary btn-sm"
              @click="addSingleValue"
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
  name: "SingleValueModal",
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
  data() {
    const schema = yup.object().shape({
      single_value: yup.number().required("Required"),
    });
    return {
      schema,
      single_value: 0,
      errorMessage: "",
    };
  },
  methods: {
    addSingleValue() {
      let exercise = this.exercise;
      exercise.single_value = {
        value: parseInt(this.single_value),
      };
      console.log(exercise.single_value);
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
</style>