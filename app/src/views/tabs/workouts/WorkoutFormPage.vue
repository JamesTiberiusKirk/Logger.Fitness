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
              <ion-textarea
                auto-grow="true"
                v-model="workout.notes"
              ></ion-textarea>
            </ion-item>

            <ion-label color="danger">{{ metaData.errMessage }}</ion-label>
          </form>

          <span v-if="isEdit">
            <ion-item>
              <ion-label
                >Start: {{ parseUnixDateTime(workout.start_time) }}</ion-label
              >
              <ion-button
                class="input-btn"
                fill="clear"
                slot="end"
                id="open-start-date-input"
              >
                Select
              </ion-button>

              <ion-popover
                trigger="open-start-date-input"
                alignment="end"
                :dismiss-on-select="false"
              >
                <ion-datetime
                  presentation="date-time"
                  @ionChange="(ev: DatetimeCustomEvent) => saveDateTimeCustomEvent(ev, 'start')"
                />
              </ion-popover>
            </ion-item>
            <ion-item>
              <ion-label
                >End:
                {{ parseUnixDateTime(workout.end_time * 1000) }}</ion-label
              >
              <ion-button
                class="input-btn"
                fill="clear"
                slot="end"
                id="open-end-date-input"
              >
                Select
              </ion-button>

              <ion-popover
                trigger="open-end-date-input"
                alignment="end"
                :dismiss-on-select="false"
              >
                <ion-datetime
                  presentation="date-time"
                  @ionChange="(ev: DatetimeCustomEvent) => saveDateTimeCustomEvent(ev, 'end')"
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
  IonPopover,
  DatetimeCustomEvent,
} from "@ionic/vue";
import { computed, ref } from "vue";
import { useRoute } from "vue-router";
import store from "@/store";
import router from "@/router";
import { Workout } from "@/types/workout";
import {
  parseHumanReadableDateTimeToUnixEpoch,
  parseUnixDateTime,
} from "@/common/date";

const metaData = ref({
  loading: false,
  errMessage: "",
});

const workout = ref({} as Workout);

const route = useRoute();
const isEdit = computed(() => route.query["id"]);

if (isEdit.value) {
  workout.value = store.getters["workouts/getOneWorkoutById"](
    route.query["id"]
  );
}

function saveDateTimeCustomEvent(date: DatetimeCustomEvent, type: string) {
  if (type === "start")
    workout.value.start_time = parseHumanReadableDateTimeToUnixEpoch(
      date.detail.value
    );

  if (type === "end")
    workout.value.end_time = parseHumanReadableDateTimeToUnixEpoch(
      date.detail.value
    );
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