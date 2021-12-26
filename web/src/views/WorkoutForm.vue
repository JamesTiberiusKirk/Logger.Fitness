
<template>
  <div class="container">
    <header class="jumbotron">
      <h3>
        <span v-if="$route.query.id">Edit </span>
        <span v-if="!$route.query.id">New </span>
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
                v-model="workout.title"
                name="title"
                type="text"
                class="form-control"
              />
              <ErrorMessage name="title" class="error-feedback" />
            </div>
            <div class="form-group">
              <label for="notes">Notes</label>
              <Field v-model="workout.notes" name="notes" class="form-control">
                <textarea
                  class="form-control"
                  v-model="workout.notes"
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
                <span v-if="$route.query.id">Update</span>
                <span v-if="!$route.query.id">Add</span>
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
      workout: {
        workout_id: "",
        title: "",
        notes: "",
        start_time: "",
      },
    };
  },
  created() {
    let id = this.$route.query.id;
    if (id) {
      this.workout = this.$store.getters["workouts/getOneById"](id).workout;
      console.log(this.workout);
    }
  },
  methods: {
    handleSubmit(workout) {
      this.loading = true;
      if (this.$route.query.id) { // For updating a current workout
        this.workout.title = workout.title;
        this.workout.notes = workout.notes;

        this.$store.dispatch("workouts/updateOne", this.workout).then(
          () => {
            this.successMessage = "Success";
            this.loading = false;
          },
          (err) => {
            console.error(err);
            this.errorMessage = err;
            this.loading = false;
          }
        );
      } else { // For new workout
        this.workout.start_time = Date.now();
        this.$store
          .dispatch("workouts/start", this.workout)
          // TODO: BUG: so when there is no workout in the list, this whole
          //  .then does not run for whatever fucking bizarre reason...
          //  After there is at least one item in the list, the .then runs and
          //    the app navigates back to an updated workouts list.
          .then(() => {
            this.errorMessage = undefined;
            this.successMessage = "Workout started";
            this.loading = false;
            // TODO: router.push to the new workout
            // this.$router.push("workout",{id:this.data.workout_id})
            // pushing back to workout screen for now

            this.$router.back();
          })
          .catch((err) => {
            this.successMessage = undefined;

            err.response && err.response.data
              ? (this.errorMessage = err.response.data)
              : (this.errorMessage = err.response);

            this.loading = false;
          });
      }
    },
  },
};
</script>