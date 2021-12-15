
<template>
  <div class="container">
    <header class="jumbotron">
      <h3>
        <span v-if="this.$route.query.id">Edit </span>
        <span v-if="!this.$route.query.id">New </span>
        Workout
      </h3>
    </header>
    <div class="col-md-12">
      <div class="card card-container">
        <div class="card-body">
          <Form @submit="handleSubmit" :validation-schema="schema">
            <div class="form-group">
              <label for="title">Title</label>
              <Field
                v-model="this.data.title"
                name="title"
                type="text"
                class="form-control"
              />
              <ErrorMessage name="title" class="error-feedback" />
            </div>
            <div class="form-group">
              <label for="notes">Notes</label>
              <Field
                v-model="this.data.notes"
                name="notes"
                class="form-control"
              >
                <textarea
                  class="form-control"
                  v-model="this.data.notes"
                ></textarea>
              </Field>
              <ErrorMessage name="notes" class="error-feedback" />
            </div>

            <div class="form-group">
              <button class="btn btn-primary btn-block" :disabled="loading">
                <span
                  v-show="loading"
                  class="spinner-border spinner-border-sm"
                ></span>
                <span v-if="this.$route.query.id">Update</span>
                <span v-if="!this.$route.query.id">Add</span>
              </button>
            </div>

            <div class="form-group">
              <div v-if="errorMessage" class="alert alert-danger" role="alert">
                {{ errorMessage }}
              </div>
            </div>
            <div class="form-group">
              <div
                v-if="successMessage"
                class="alert alert-success"
                role="alert"
              >
                {{ successMessage }}
              </div>
            </div>
          </Form>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { Form, Field, ErrorMessage } from "vee-validate";
import * as yup from "yup";
export default {
  name: "WorkoutForm",
  components: {
    Form,
    Field,
    ErrorMessage,
  },
  data() {
    const schema = yup.object().shape({
      title: yup.string().required("Please provide a title"),
      notes: yup.string(),
    });

    return {
      loading: false,
      successMessage: "",
      errorMessage: "",
      schema,
      data: {
        title: "",
        notes: "",
        start_time: "",
      },
    };
  },
  mounted() {
    let id = this.$route.query.id;
    if (id) {
      let workoutId = this.$store.getters["workouts/getOneById"](id);
      this.data = workoutId;
    }
  },
  methods: {
    handleSubmit(workout) {
      this.loading = true;
      if (this.$route.query.id) {
        this.data["title"] = workout.title;
        this.data["notes"] = workout.notes;

        this.$store.dispatch("workouts/updateOne", this.data).then(
          (res) => {
            console.log("success ", res);
            this.successMessage = "Success";
            this.loading = false;
          },
          (err) => {
            console.log("error ", err);
            this.errorMessage = err;
            this.loading = false;
          }
        );
      } else {
        this.data.start_time = Date.now();
        this.$store.dispatch("workouts/start", this.data).then(
          (res) => {
            console.log("success: ", res);
            this.errorMessage = undefined;
            this.successMessage = "Workout started";
            this.loading = false;
          },
          (err) => {
            this.successMessage = undefined;
            this.errorMessage = err.response.data;
            this.loading = false;
          }
        );
      }
    },
  },
};
</script>