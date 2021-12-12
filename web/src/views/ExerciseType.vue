<template>
  <div class="container">
    <header class="jumbotron">
      <h3>
        <span v-if="this.$route.query.id">Edit </span>
        <span v-if="!this.$route.query.id">New </span>
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
                v-model="this.data.name"
                name="name"
                type="text"
                class="form-control"
              />
              <ErrorMessage name="name" class="error-feedback" />
            </div>
            <div class="form-group">
              <label for="description">Description</label>
              <Field
                v-model="this.data.description"
                name="description"
                class="form-control"
              >
                <textarea
                  class="form-control"
                  v-model="this.data.description"
                ></textarea>
              </Field>
              <ErrorMessage name="description" class="error-feedback" />
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
              <div v-if="message" class="alert alert-danger" role="alert">
                {{ message }}
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
      name: "",
      description: "",
      dataType: "",
    });

    return {
      loading: false,
      message: "",
      schema,
      data: {},
    };
  },
  mounted() {
    let id = this.$route.query.id;
    if (id) {
      let exerciseType = this.$store.getters["exerciseType/getOneById"](id);
      this.data = exerciseType;
    }
  },
  methods: {
    handleSubmit(exerciseType) {
      this.loading = true;
      if (this.$route.query.id) {
        this.data["name"] = exerciseType.name;
        this.data["description"] = exerciseType.description;

        this.$store.dispatch("exerciseType/updateOne", this.data).then(
          (res) => {
            console.log("success ", res);
            this.loading = false;
          },
          (err) => {
            console.log("error ", err);
            this.loading = false;
          }
        );
      } else {
        this.$store.dispatch("exerciseType/sendOne", exerciseType).then(
          (res) => {
            console.log("success ", res);
            this.loading = false;
          },
          (err) => {
            console.log("error ", err);
            this.loading = false;
          }
        );
      }
    },
  },
};
</script>