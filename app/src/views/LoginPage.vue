
<template>
  <ion-page>
    <ion-header>
      <ion-toolbar>
        <ion-title>Logger.Fitness</ion-title>
      </ion-toolbar>
    </ion-header>
    <ion-content fullscreen>
      <ion-card class="ion-align-items-center">
        <ion-card-header>
          <ion-card-title> Log In </ion-card-title>
        </ion-card-header>
        <ion-card-content>
          <form class="login-form">
            <ion-item>
              <ion-label>Email:</ion-label>
              <ion-input v-model="user.email"></ion-input>
              <ion-label v-if="!isValid.email" color="warning"
                >Input is not valid</ion-label
              >
            </ion-item>

            <ion-item>
              <ion-label>Password:</ion-label>
              <ion-input v-model="user.password" type="password"></ion-input>
              <ion-label v-if="!isValid.password" color="warning"
                >Input is not valid</ion-label
              >
            </ion-item>

            <ion-label v-if="metaData.errMessage" color="danger">{{
              metaData.errMessage
            }}</ion-label>
          </form>
          <ion-button @click="validateLoginSubmit()" expand="block"
            >Login</ion-button
          >
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
  IonCard,
  IonCardContent,
  IonCardHeader,
  IonCardTitle,
  IonItem,
  IonLabel,
  IonInput,
  IonHeader,
  IonPage,
  IonContent,
  IonToolbar,
  IonTitle,
  IonButton,
  IonLoading,
} from "@ionic/vue";
import { UserLoginDTO } from "@/types/user";
import store from "@/store";
import { Ref, ref } from "vue";
import router from "@/router";
import Validate from "@/common/validate";

const user: Ref<UserLoginDTO> = ref<UserLoginDTO>({
  email: "",
  password: "",
} as UserLoginDTO);

const isValid = ref({
  email: true,
  password: true,
});

const metaData = ref({
  errMessage: "",
  successful: false,
  loading: false,
});

console.log(store.state.auth);

if (store.state.auth.status.loggedIn) {
  router.push("/tabs");
}

function validateLoginSubmit() {
  isValid.value.email = Validate.email(user.value.email);
  isValid.value.password = !Validate.isEmpty(user.value.password);
  if (!isValid.value.email || !isValid.value.password) return;

  metaData.value.loading = true;
  store.dispatch("auth/login", user.value).then(
    () => {
      metaData.value.loading = false;
      metaData.value.successful = true;
      router.push("/tabs");
    },
    (err: any) => {
      metaData.value.loading = false;
      metaData.value.successful = false;
      metaData.value.errMessage = err.request.responseText || err.data;
      if (err.request.status === 401)
        metaData.value.errMessage = "Wrong password or email";
    }
  );
}
</script>

<style scoped>
</style>
