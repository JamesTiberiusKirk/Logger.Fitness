<template>
  <ion-page>
    <ion-content fullscreen>
      <ion-card class="ion-align-items-center">
        <ion-card-header>
          <ion-card-title>Exercise Type</ion-card-title>
        </ion-card-header>
        <ion-card-content>
          <form class="login-form">
            <ion-item>
              <ion-label position="floating">Name:</ion-label>
              <ion-input v-model="exerciseType.name"></ion-input>
              <!-- <ion-label color="warning">Input is not valid</ion-label> -->
            </ion-item>

            <ion-item>
              <ion-label position="floating">Description:</ion-label>
              <ion-textarea v-model="exerciseType.description"></ion-textarea>
              <!-- <ion-label color="warning">Input is not valid</ion-label> -->
            </ion-item>

            <ion-item>
              <ion-label>Type:</ion-label>
              <ion-select v-model="exerciseType.data_type">
                <ion-select-option value="sets">Sets</ion-select-option>
                <ion-select-option value="single-value"
                  >Single Value</ion-select-option
                >
                <ion-select-option value="custom">Custom</ion-select-option>
              </ion-select>
            </ion-item>

            <ion-item>
              <ion-label position="floating">Measurement Unit:</ion-label>
              <ion-input
                v-model="exerciseType.measurement_type"
                type="text"
              ></ion-input>
            </ion-item>
            <ion-label color="danger">{{ metaData.errMessage }}</ion-label>
          </form>

          <ion-button @click="submit()" expand="block">Add</ion-button>
        </ion-card-content>
      </ion-card>

      <ion-card
        class="ion-align-items-center"
        v-if="exerciseType.data_type === 'custom'"
      >
        <ion-card-header>
          <ion-card-title> Custom types not yet implemented </ion-card-title>
        </ion-card-header>
        <ion-card-content> </ion-card-content>
      </ion-card>

      <ion-loading
        :is-open="metaData.loading"
        message="Loading please wait..."
      />
    </ion-content>
  </ion-page>
</template>

<script setup lang="ts">
import {
  IonPage,
  IonContent,
  IonLabel,
  IonInput,
  IonTextarea,
  IonLoading,
  IonCard,
  IonCardTitle,
  IonCardHeader,
  IonCardContent,
  IonItem,
  IonButton,
  IonSelect,
  IonSelectOption,
} from "@ionic/vue";
import { ref } from "vue";
import { ExerciseType } from "@/types/exercise-type";
import { useRoute } from "vue-router";
import store from "@/store";
import router from '@/router'

const metaData = ref({
  loading: false,
  errMessage: "",
});

const exerciseType = ref({} as ExerciseType);

const route = useRoute();
console.log(route.query);
if (route.query["id"]) {
  exerciseType.value = store.getters["exerciseTypes/getOneById"](
    route.query["id"]
  );
}

function submit() {
  // Should validate the exercise type

  metaData.value.loading = true;
  console.log(exerciseType.value);

  if (route.query["id"]) {
    store
      .dispatch("exerciseTypes/updateOne", exerciseType.value)
      .then(() => {
        metaData.value.loading = false;
        router.back()
      })
      .catch((err) => {
        metaData.value.errMessage = err.message;
      });
  } else {
    store
      .dispatch("exerciseTypes/sendOne", exerciseType.value)
      .then(() => {
        metaData.value.loading = false;
        router.back()
      })
      .catch((err) => {
        metaData.value.errMessage = err.message;
      });
  }
}
</script>