<template>
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
            <ion-textarea
              auto-grow="true"
              v-model="exerciseType.description"
            ></ion-textarea>
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
        <ion-button @click="submit()" expand="block">
          {{ metaData.type.charAt(0).toUpperCase() }}
        </ion-button>
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
  </ion-content>
</template>

<script lang="ts">
import {
  IonContent,
  IonLabel,
  IonInput,
  IonTextarea,
  IonCard,
  IonCardTitle,
  IonCardHeader,
  IonCardContent,
  IonItem,
  IonButton,
  IonSelect,
  IonSelectOption,
} from "@ionic/vue";
import { defineProps, defineEmits, ref, onMounted } from "vue";
import { ExerciseType } from "@/types/exercise-type";

export default {
  components: {
    IonContent,
    IonLabel,
    IonInput,
    IonTextarea,
    IonCard,
    IonCardTitle,
    IonCardHeader,
    IonCardContent,
    IonItem,
    IonButton,
    IonSelect,
    IonSelectOption,
  },
  name: "ExerciseTypeForm",
  setup() {
    const metaData = ref({
      type: "",
      errMessage: "",
    });

    const props = defineProps<{
      exerciseType?: ExerciseType;
    }>();
    const exerciseType = ref({} as ExerciseType);

    const emit = defineEmits<{
      (e: "new", exerciseType: ExerciseType): void;
      (e: "update", exerciseType: ExerciseType): void;
    }>();

    onMounted(() => {
      console.log(props);

      if (props.exerciseType) {
        metaData.value.type = "update";
        exerciseType.value = props.exerciseType;
      } else {
        metaData.value.type = "new";
      }
    });

    function submit() {
      if (props.exerciseType) {
        emit("update", exerciseType.value);
      } else {
        emit("new", exerciseType.value);
      }
    }

    return {
      metaData,
      exerciseType,
      submit,
    };
  },
};
</script>

