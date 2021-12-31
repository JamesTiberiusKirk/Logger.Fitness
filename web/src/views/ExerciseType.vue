<template>
  <div class="container">
    <header class="jumbotron">
      <h3>
        <span v-if="$route.query.id">Edit </span>
        <span v-if="!$route.query.id">New </span>
        Exercise Type
      </h3>
    </header>
    <div class="col-md-12">
      <div class="card card-container">
        <div class="card-body">
          <Form @submit="handleSubmit" :validation-schema="schema">
            <div class="form-group">
              <label for="name">Name</label>
              <Field
                v-model="exercise_type.name"
                name="name"
                type="text"
                class="form-control"
              />
              <ErrorMessage name="name" class="error-feedback" />
            </div>
            <div class="form-group">
              <label for="description">Description</label>
              <Field
                v-model="exercise_type.description"
                name="description"
                class="form-control"
              >
                <textarea
                  class="form-control"
                  v-model="exercise_type.description"
                ></textarea>
              </Field>
              <ErrorMessage name="description" class="error-feedback" />
            </div>

            <div class="form-group">
              <Field name="data_type" v-model="exercise_type.data_type">
                <label for="dataTypeSelect"> Chose data type </label>
                <select
                  id="dataTypeSelect"
                  class="form-control"
                  v-model="exercise_type.data_type"
                >
                  <option value="">Choose</option>
                  <option value="sets">Sets</option>
                  <option value="single-value">Single Value</option>
                </select>
              </Field>
              <ErrorMessage name="data_type" class="error-feedback" />
            </div>

            <div class="form-group">
              <label for="name">Unit of measurement</label>
              <Field
                v-model="exercise_type.measurement_type"
                name="measurement_type"
                type="text"
                class="form-control"
              />
              <ErrorMessage name="measurement_type" class="error-feedback" />
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
  name: "ExerciseType",
  components: {
    Form,
    Field,
    ErrorMessage,
  },
  data() {
    const schema = yup.object().shape({
      name: yup.string().required("Please provide a name"),
      description: yup.string(),
      data_type: yup.string().required("Please chose one"),
      measurement_type: yup.string().required("Please add a measurement unit"),
    });

    return {
      loading: false,
      successMessage: "",
      errorMessage: "",
      schema,
      exercise_type: {
        exercise_type_id: "",
        user_id: "",
        name: "",
        description: "",
        data_type: "",
        measurement_type: "",
      },
    };
  },
  mounted() {
    let id = this.$route.query.id;
    if (id) {
      this.exercise_type = this.$store.getters["exerciseTypes/getOneById"](id);
    }
  },
  methods: {
    handleSubmit(exercise_type) {
      this.loading = true;
      if (this.$route.query.id) {
        this.exercise_type["name"] = exercise_type.name;
        this.exercise_type["description"] = exercise_type.description;
        this.exercise_type["measurement_type"] = exercise_type.measurement_type;

        this.$store.dispatch("exerciseTypes/updateOne", this.exercise_type).then(
          () => {
            this.successMessage = "Success";
            this.loading = false;
            this.$router.back()
          },
          (err) => {
            console.error(err);
            this.errorMessage = err;
            this.loading = false;
          }
        );
      } else {
        this.$store.dispatch("exerciseTypes/sendOne", exercise_type).then(
          () => {
            this.successMessage = "Success";
            this.loading = false;
            this.$router.back()
          },
          (err) => {
            console.error(err);
            this.errorMessage = err.request.requestText
            this.loading = false;
          }
        );
      }
    },
  },
};
</script>