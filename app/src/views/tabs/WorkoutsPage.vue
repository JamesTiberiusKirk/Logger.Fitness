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
            <ion-item button @click="workoutClick(w)">
              <ion-label>
                <h3>
                  {{
                    moment(w.workout.start_time).format("MMMM Do YYYY, h:mm a")
                  }}
                </h3>
                <h1>{{ w.workout.title }}</h1>
                <p>{{ w.workout.notes }}</p>
              </ion-label>
            </ion-item>

            <ion-item-options side="end">
              <ion-item-option @click="edit(w.workout.workout_id)"
                >View</ion-item-option
              >
              <ion-item-option
                color="danger"
                @click="openDeleteModal(w.workout.workout_id)"
                >Delete</ion-item-option
              >
            </ion-item-options>
          </ion-item-sliding>
        </span>
      </ion-list>
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
  IonSearchbar,
} from "@ionic/vue";
import { ref, computed } from "vue";
import { Workouts } from "@/types/workout";
import store from "@/store/index";
import moment from "moment";
import notYetImplementedToast from "@/common/notYetImplementedToast";
import Validate from "@/common/validate";

const workoutsList = ref({} as Workouts[]);
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
      .then((data: Workouts[]) => {
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
  return workoutsList.value.filter((workout:any)=>{
    return (
      workout.workout.title.toLowerCase().includes(st) ||
      workout.workout.notes.toLowerCase().includes(st)
    )
  });
});

function workoutClick(w: any) {
  console.log(w);
  notYetImplementedToast();
}

function doRefresh(event: CustomEvent) {
  getData().then(() => {
    event.target.complete();
  });
}
</script>