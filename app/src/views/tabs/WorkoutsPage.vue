<template>
  <ion-page>
    <ion-content>
      <ion-refresher slot="fixed" @ionRefresh="doRefresh($event)">
        <ion-refresher-content></ion-refresher-content>
      </ion-refresher>

      <ion-searchbar v-model="searchTerm"></ion-searchbar>
      <ion-list>
        <span v-for="(w, i) in workoutsListFilter" :key="i">
          <ion-item-sliding>
            <ion-item button @click="workoutClick(w.workout.workout_id)">
              <ion-label>
                <h3 v-if="hasEnded(w.workout)" class="active-workout">
                  Active
                </h3>
                <h3>
                  {{
                    moment(w.workout.start_time).format(
                      "dddd, MMMM Do YYYY, h:mm a"
                    )
                  }}
                </h3>
                <h1>{{ w.workout.title }}</h1>
                <p>{{ w.workout.notes }}</p>
              </ion-label>
            </ion-item>

            <ion-item-options side="end">
              <ion-item-option @click="workoutEditClick(w.workout.workout_id)"
                >View</ion-item-option
              >
              <ion-item-option
                color="danger"
                @click="workoutDeleteClick(w.workout.workout_id)"
                >Delete</ion-item-option
              >
            </ion-item-options>
          </ion-item-sliding>
        </span>
      </ion-list>

      <ion-fab vertical="bottom" horizontal="end" slot="fixed">
        <ion-fab-button @click="fabClick()">
          <ion-icon :icon="add"></ion-icon>
        </ion-fab-button>
      </ion-fab>
    </ion-content>

    <ion-loading :is-open="metaData.loading" message="Loading please wait..." />
  </ion-page>
</template>

<script setup lang="ts">
import {
  IonPage,
  IonContent,
  IonLoading,
  IonRefresher,
  IonRefresherContent,
  IonList,
  IonItemSliding,
  IonItemOption,
  IonItemOptions,
  IonItem,
  IonLabel,
  IonIcon,
  IonFab,
  IonFabButton,
  IonSearchbar,
} from "@ionic/vue";
import { add } from "ionicons/icons";
import { ref, computed } from "vue";
import { WorkoutGroup, Workout } from "@/types/workout";
import store from "@/store/index";
import moment from "moment";
import notYetImplementedToast from "@/common/notYetImplementedToast";
import Validate from "@/common/validate";

const workoutsList = ref({} as Workout[]);
const searchTerm = ref("");
const metaData = ref({
  loading: false,
  errMessage: "",
});

function getData() {
  console.log("workouts");
  return new Promise((resolve, reject) => {
    store
      .dispatch("workouts/fetchAll")
      .then((data: Workout[]) => {
        console.log("data: ", data);
        workoutsList.value = data.reverse();
        resolve(data);
      })
      .catch((err) => {
        console.log("err: ", err);
        reject(err);
      });
  });
}
getData();

const workoutsListFilter = computed(() => {
  if (Validate.isEmpty(searchTerm.value)) return workoutsList.value;
  const st = searchTerm.value.toLowerCase();
  return workoutsList.value.filter((workout_group: WorkoutGroup) => {
    return (
      workout_group.workout.title.toLowerCase().includes(st) ||
      workout_group.workout.notes.toLowerCase().includes(st)
    );
  });
});

function hasEnded(workout: Workout) {
  return !workout.end_time;
}

function workoutClick(id: string) {
  console.log(id);
  notYetImplementedToast();
}

function doRefresh(event: CustomEvent) {
  getData().then(() => {
    event.target.complete();
  });
}

function workoutEditClick(id: string) {
  console.log(id);
  notYetImplementedToast();
}

function workoutDeleteClick(id: string) {
  console.log(id);
  notYetImplementedToast();
}

function fabClick() {
  notYetImplementedToast();
}
</script>

<style scoped>
.active-workout {
  text-decoration-color: red;
}
</style>