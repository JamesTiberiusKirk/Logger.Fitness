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
      <!-- <span v-for="(e, i) in exerciseTypeListFilter" :key="i">
        <ion-card>
          <ion-card-header>
            <ion-card-subtitle>Type: {{ e.data_type }}</ion-card-subtitle>
            <ion-card-title>{{ e.name }}</ion-card-title>
          </ion-card-header>

          <ion-card-content>
            <ion-item>
              <ion-label>{{ e.description }}</ion-label>
              <ion-icon
                :icon="create"
                slot="end"
                @click="edit(e.exercise_type_id)"
              ></ion-icon>

              <ion-icon
                :icon="trash"
                slot="end"
                @click="openDeleteModal(e.exercise_type_id)"
              ></ion-icon>
            </ion-item>
          </ion-card-content>
        </ion-card>
      </span> -->

      <ion-list>
        <span v-for="(e, i) in exerciseTypeListFilter" :key="i">
          <ion-item-sliding>
            <!-- <ion-item-options side="start">
              <ion-item-option @click="favorite(item)"
                >Favorite</ion-item-option
              >
              <ion-item-option color="danger" @click="share(item)"
                >Share</ion-item-option
              >
            </ion-item-options> -->
            <ion-item>
              <ion-label>
                <h3>Type: {{ e.data_type }}</h3>
                <h1>{{ e.name }}</h1>
                <p>{{ e.description }}</p>
              </ion-label>

              <!-- <ion-icon
                :icon="create"
                slot="end"
                @click="edit(e.exercise_type_id)"
              ></ion-icon>

              <ion-icon
                :icon="trash"
                slot="end"
                @click="openDeleteModal(e.exercise_type_id)"
              ></ion-icon> -->
            </ion-item>

            <ion-item-options side="end">
              <ion-item-option @click="edit(e.exercise_type_id)"
                >Edit</ion-item-option
              >
              <ion-item-option color="danger" @click="openDeleteModal(e.exercise_type_id)"
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

    <!-- Card Modal -->
    <ion-modal
      :is-open="isDeleteModalOpen"
      :swipe-to-close="true"
      :presenting-element="$parent.$refs.ionRouterOutlet"
    >
      <ion-content>
        <ion-page>
          <ion-content fullscreen>
            <ion-card>
              <ion-card-header>
                <ion-card-title> Delete exercise type IDxxxxx </ion-card-title>
              </ion-card-header>
              <ion-card-content>
                <ion-button @click="deleteExercise()" color="danger">
                  Delete</ion-button
                >
                <ion-button @click="closeDeleteModal()"> Cancel </ion-button>
              </ion-card-content>
            </ion-card>
          </ion-content>
        </ion-page>
      </ion-content>
    </ion-modal>
    <ion-loading :is-open="metaData.loading" message="Loading please wait..." />
  </ion-page>
</template>

<script setup lang="ts">
import {
  IonPage,
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
  IonItem,
  IonRefresherContent,
  IonModal,
  IonButton,
  IonLoading,
  IonRefresher,
  IonList,
  IonItemSliding,
  IonItemOptions,
  IonItemOption,
} from "@ionic/vue";
import { add, create, trash } from "ionicons/icons";
import { computed, ref } from "vue";

import { ExerciseType } from "@/types/exercise-type";
import store from "@/store";
import Validate from "@/common/validate";
import router from "@/router";

const exerciseTypeList = ref({} as ExerciseType[]);
const searchTerm = ref("");
const filter = ref("all");
const isDeleteModalOpen = ref(false);
const deleteExerciseTypeID = ref(-1);
const metaData = ref({
  loading: false,
  errMessage: "",
});

function getData() {
  return new Promise((resolve, reject) => {
    store.dispatch("exerciseTypes/fetchAll").then(
      (data: ExerciseType[]) => {
        console.log(data);
        exerciseTypeList.value = data;
        resolve(data);
      },
      (err) => {
        console.log(err);
        reject(err);
      }
    );
  });
}
getData();

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

function openDeleteModal(id: string) {
  console.log(id);
  deleteExerciseTypeID.value = id;
  isDeleteModalOpen.value = true;
}

function closeDeleteModal() {
  isDeleteModalOpen.value = false;
}

function deleteExercise() {
  if (!deleteExerciseTypeID.value) return;

  metaData.value.loading = true;
  store
    .dispatch("exerciseTypes/deleteOne", deleteExerciseTypeID.value)
    .then(() => {
      metaData.value.loading = false;
      isDeleteModalOpen.value = false;
    })
    .catch((err) => {
      console.log(err);

      isDeleteModalOpen.value = false;
      metaData.value.loading = false;
    });
}

function fabClick() {
  router.push(`/tabs/exercise`);
}
</script>