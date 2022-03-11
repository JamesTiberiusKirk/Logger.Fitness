
<template>
  <ion-page>
    <ion-header>
      <ion-toolbar>
        <ion-title>Logger.Fitness</ion-title>
      </ion-toolbar>
    </ion-header>
    <ion-content fullscreen>
      <ion-card>
        <ion-card-header>
          <ion-card-title> Log In </ion-card-title>
        </ion-card-header>
        <ion-card-content>
          <form class="login-form" onSubmit="{validateLoginSubmit()}">
            <ion-item>
              <ion-label>Email:</ion-label>
              <ion-input></ion-input>
            </ion-item>

            <ion-item>
              <ion-label>Password:</ion-label>
              <ion-input type="password"></ion-input>
            </ion-item>
          </form>
        </ion-card-content>
      </ion-card>
    </ion-content>
  </ion-page>
</template>

<script lang="ts">
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
} from "@ionic/vue";

export default {
  components: {
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
  },
};
</script>
<style scoped>
</style>