<template>
  <ion-page>
    <h1>
    {{JSON.stringify(workout)}}
    </h1>
  </ion-page>
</template>

<script setup lang="ts">
import { IonPage } from "@ionic/vue";
import store from "@/store";
import { useRoute } from "vue-router";
import { ref } from "vue";
import { Workout } from "@/types/workout";

const route = useRoute();

const workout = ref({} as Workout);

const metaData = ref({
  loading: false,
  errMessage: "",
});

workout.value = store.getters["workouts/getOneWorkoutById"](route.query["id"]);
console.log(workout.value);


</script>
