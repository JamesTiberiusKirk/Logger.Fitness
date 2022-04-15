<template>
  <ion-page>
    <ion-content fullscreen>
      <ion-card class="ion-align-items-center">
        <ion-card-header>
          <ion-card-title v-if="isEdit">Workout Edit</ion-card-title>
          <ion-card-title v-if="!isEdit">Start Workout</ion-card-title>
        </ion-card-header>
        <ion-card-content>
          <form class="login-form">
            <ion-item>
              <ion-label position="floating">Title:</ion-label>
              <ion-input v-model="workout.title"></ion-input>
            </ion-item>

            <ion-item>
              <ion-label position="floating">Notes:</ion-label>
              <ion-textarea v-model="workout.notes"></ion-textarea>
            </ion-item>

            <ion-label color="danger">{{ metaData.errMessage }}</ion-label>
          </form>

          <span v-if="isEdit">
            <ion-item>
              <ion-label position="floating">Start time:</ion-label>
              <ion-button id="open-modal">Open Datetime Modal</ion-button>
              <ion-modal trigger="open-modal">
                <ion-content>
                  <ion-datetime></ion-datetime>
                </ion-content>
              </ion-modal>
            </ion-item>

            <!-- Datetime in popover with input -->
            <ion-item>
              <ion-input :value="date" />
              <ion-button fill="clear" id="open-date-input-2">
                <ion-icon icon="calendar" />
              </ion-button>
              <ion-popover trigger="open-date-input-2" :show-backdrop="false">
                <ion-datetime
                  presentation="date"
                  @ionChange="(ev: DatetimeCustomEvent) => date = formatDate(ev.detail.value)"
                />
              </ion-popover>
            </ion-item>
          </span>

          <ion-button @click="submit()" expand="block">
            <span v-if="!isEdit">Add</span>
            <span v-if="isEdit">Update</span>
          </ion-button>
        </ion-card-content>
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
  IonDatetime,
  IonModal,
  IonPopover,
  IonIcon,
} from "@ionic/vue";
import { calendar } from "ionicons/icons";
import { computed, ref } from "vue";
import { useRoute } from "vue-router";
import store from "@/store";
import router from "@/router";
import { Workout } from "@/types/workout";
import { format, parseISO } from "date-fns";

const metaData = ref({
  loading: false,
  errMessage: "",
});

const workout = ref({} as Workout);
const date = ref("");

const route = useRoute();

const isEdit = computed(() => route.query["id"]);

if (isEdit.value) {
  workout.value = store.getters["workouts/getOneWorkoutById"](
    route.query["id"]
  );
}

function formatDate (value: string) {
  return format(parseISO(value), "MMM dd yyyy");
}

function submit() {
  metaData.value.loading = true;
  console.log(workout.value);

  if (isEdit.value) {
    store
      .dispatch("workouts/updateOneWorkout", workout.value)
      .then(() => {
        metaData.value.loading = false;
      })
      .catch((err) => {
        metaData.value.loading = false;
        metaData.value.errMessage = err.message;
        console.log(err);
      });
  } else {
    store
      .dispatch("workouts/start", workout.value)
      .then(() => {
        metaData.value.loading = false;
        router.back();
      })
      .catch((err) => {
        metaData.value.loading = false;
        metaData.value.errMessage = err.message;
      });
  }
}
</script>