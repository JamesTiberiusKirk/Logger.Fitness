<template>
  <ion-page>
    <ion-content>
      <ion-refresher slot="fixed" @ionRefresh="doRefresh($event)">
        <ion-refresher-content></ion-refresher-content>
      </ion-refresher>
      <ion-searchbar v-model="searchTerm"></ion-searchbar>
      <ion-segment v-model="filter">
        <ion-segment-button value="all">
          <ion-label>All</ion-label>
        </ion-segment-button>
        <ion-segment-button value="sets">
          <ion-label>Sets</ion-label>
        </ion-segment-button>
        <ion-segment-button value="single-value">
          <ion-label>Single Values</ion-label>
        </ion-segment-button>
        <ion-segment-button value="custom">
          <ion-label>Custom</ion-label>
        </ion-segment-button>
      </ion-segment>
      <ion-list>
        <span v-for="(e, i) in exerciseTypeListFilter" :key="i">
          <ion-item-sliding>
            <ion-item>
              <ion-label>
                <h3>Type: {{ e.data_type }}</h3>
                <h1>{{ e.name }}</h1>
                <p>{{ e.description }}</p>
              </ion-label>
            </ion-item>

            <ion-item-options side="end">
              <ion-item-option @click="edit(e.exercise_type_id)"
                >Edit</ion-item-option
              >
              <ion-item-option color="danger" @click="setModalState(true, e)"
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

    <ion-action-sheet
      :is-open="modal.isOpen"
      header="Are you sure?"
      :buttons="modal.buttons"
      @didDismiss="setModalState(false, -1)"
    >
    </ion-action-sheet>
  </ion-page>
</template>

<script setup lang="ts">
import {
  IonPage,
  IonContent,
  IonFab,
  IonFabButton,
  IonIcon,
  IonSearchbar,
  IonSegment,
  IonSegmentButton,
  IonLabel,
  IonItem,
  IonRefresherContent,
  IonLoading,
  IonRefresher,
  IonList,
  IonItemSliding,
  IonItemOptions,
  IonItemOption,
  IonActionSheet,
} from "@ionic/vue";
import { add, trash } from "ionicons/icons";
import { computed, onMounted, ref } from "vue";

import { ExerciseType } from "@/types/exercise-type";
import store from "@/store";
import Validate from "@/common/validate";
import router from "@/router";

const exerciseTypeList = ref({} as ExerciseType[]);
const searchTerm = ref("");
const filter = ref("all");
const metaData = ref({
  loading: false,
  errMessage: "",
});

const modal = ref({
  isOpen: false,
  toDelete: -1,
  buttons: [
    {
      text: "Delete",
      role: "destructive",
      icon: trash,
      data: {
        type: "delete",
      },
      handler: deleteExercise,
    },
    {
      text: "Cancel",
      icon: close,
      role: "cancel",
    },
  ],
});

// BUG: So theres a bug here where the toDelete id does not get set if you're trying to delete the exercise type right after creating it
// let exerciseTypeID = -1;
function setModalState(state: boolean, exerciseId: ExerciseType) {
  // exerciseTypeID = exerciseId.exercise_type_id;
  modal.value.toDelete = exerciseId.exercise_type_id;
  modal.value.isOpen = state;
  // console.log(exerciseTypeID);
  // console.log(modal.value.toDelete);
  // console.log(exerciseId);
}

function getData() {
  metaData.value.loading = true;
  return new Promise((resolve, reject) => {
    store.dispatch("exerciseTypes/fetchAll").then(
      (data: ExerciseType[]) => {
        metaData.value.loading = false;
        console.log(data);
        exerciseTypeList.value = data;
        resolve(data);
      },
      (err) => {
        metaData.value.loading = false;
        console.log(err);
        reject(err);
      }
    );
  });
}

onMounted(async () => {
  await getData();
});

const exerciseTypeListFilter = computed(() => {
  if (!Validate.isEmpty(searchTerm.value)) {
    const st = searchTerm.value.toLowerCase();
    console.log("search");

    return exerciseTypeList.value.filter((exerciseType: ExerciseType) => {
      if (
        exerciseType.name.toLowerCase().includes(st) ||
        exerciseType.description.toLowerCase().includes(st)
      ) {
        if (filter.value !== "all") {
          return exerciseType.data_type === filter.value;
        }
        return true;
      }
    });
  } else if (filter.value !== "all") {
    return exerciseTypeList.value.filter((exerciseType: ExerciseType) => {
      return exerciseType.data_type === filter.value;
    });
  }
  return exerciseTypeList.value;
});

function doRefresh(event: CustomEvent) {
  getData().then(() => {
    event.target.complete();
  });
}

function edit(id: string) {
  router.push(`/tabs/exercise?id=${id}`);
}

function deleteExercise() {
  if (modal.value.toDelete === -1 || !modal.value.toDelete)
    return console.log("nope");

  metaData.value.loading = true;
  store
    .dispatch("exerciseTypes/deleteOne", modal.value.toDelete)
    .then(() => {
      metaData.value.loading = false;
      modal.value.isOpen = false;
    })
    .catch((err) => {
      console.log(err);
      modal.value.isOpen = false;
      metaData.value.loading = false;
    });
}

function fabClick() {
  router.push(`/tabs/exercise`);
}
</script>