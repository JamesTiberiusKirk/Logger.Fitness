<template>
  <ion-page>
    <ion-content>
      <ion-searchbar v-model="searchTerm"></ion-searchbar>
      <ion-segment v-model="filter">
        <ion-segment-button value="all">
          <ion-label>All</ion-label>
        </ion-segment-button>
        <ion-segment-button value="sets">
          <ion-label>Sets</ion-label>
        </ion-segment-button>
        <ion-segment-button value="single-values">
          <ion-label>Single Values</ion-label>
        </ion-segment-button>
        <ion-segment-button value="custom">
          <ion-label>Custom</ion-label>
        </ion-segment-button>
      </ion-segment>

      <span v-for="(e, i) in exerciseTypeListFilter" :key="i">
        <ion-card>
          <ion-card-header>
            <ion-card-subtitle>Type: {{ e.data_type }}</ion-card-subtitle>
            <ion-card-title>{{ e.name }}</ion-card-title>
          </ion-card-header>

          <ion-card-content>{{ e.description }}</ion-card-content>
        </ion-card>
      </span>
      <ion-fab vertical="bottom" horizontal="end" slot="fixed">
        <ion-fab-button @click="fabClick()">
          <ion-icon :icon="add"></ion-icon>
        </ion-fab-button>
      </ion-fab>
    </ion-content>
  </ion-page>
</template>

<script setup lang="ts">
// TODO: Actually come in and finish this
// NOTE: This is an old version of the Exercise List Page...
import { ExerciseType } from "@/types/exercise-type";
import {
  IonPage,
  IonTitle,
  IonContent,
  IonCard,
  IonCardContent,
  IonCardHeader,
  IonCardTitle,
  IonCardSubtitle,
  IonFab,
  IonFabButton,
  IonIcon,
  IonSearchbar,
  IonSegment,
  IonSegmentButton,
  IonLabel,
} from "@ionic/vue";
import { add } from "ionicons/icons";
import { computed, ref } from "vue";
import store from "@/store";
import Validate from "@/common/validate";

const exerciseTypeList = ref({} as ExerciseType[]);
const searchTerm = ref("");
const filter = ref("all");

store.dispatch("exerciseTypes/fetchAll").then(
  (data: ExerciseType[]) => {
    console.log(data);
    exerciseTypeList.value = data;
  },
  (err) => {
    console.log(err);
  }
);

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

function fabClick() {
  console.log("fab click");
  console.log(searchTerm.value);
}

</script>