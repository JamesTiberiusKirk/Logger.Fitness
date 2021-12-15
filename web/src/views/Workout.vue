
<template>
  <div class="container">
    <header class="jumbotron">
      <h3>{{ this.workout.title }}</h3>
      <p class="lead">
        {{ getDate() }} | {{ getTime() }}
        <br/>
        {{this.workout.notes}}
      </p>
    </header>
    <div class="col-md-12">
      <div class="card card-container">
        <div class="card-body">
          <p class="card-text"></p>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  name: "Workout",
  data() {
    return {
      loading: false,
      successMessage: "",
      errorMessage: "",
      workout: {
        title: "",
        notes: "",
        start_time: "",
      },
      exercises: [],
    };
  },
  mounted() {
    let id = this.$route.query.id;
    if (id) {
      let workout = this.$store.getters["workouts/getOneById"](id);
      this.workout = workout;
    } else {
      // Error
      console.log("no id provided");
    }
  },
  methods: {
    getDate() {
      return new Date(this.workout.start_time).toLocaleDateString();
    },
    getTime() {
      let time = new Date(this.workout.start_time);
      return time.getHours() +":"+time.getMinutes();
    },
  },
};
</script>