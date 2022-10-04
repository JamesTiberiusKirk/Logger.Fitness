<template>
  <ion-page>
    <ion-menu side="start" content-id="main">
      <ion-header>
        <ion-toolbar translucent>
          <ion-title>Menu</ion-title>
        </ion-toolbar>
      </ion-header>
      <ion-content class="menu-content">
        <ion-list>
          <ion-item>
            <ion-label @click="syncData">Sync Data</ion-label>
          </ion-item>
          <ion-item>
            <ion-label>Transactions: {{txAmount}}</ion-label>
          </ion-item>
        </ion-list>
        <ion-button class="logout-btn" expand="block" @click="logout()">
          Logout
        </ion-button>
        <ion-avatar v-if="user">
          <img v-if="user.claim.profile_picture" :src="user.claim.profile_picture" />
        </ion-avatar>
      </ion-content>
    </ion-menu>

    <ion-page>
      <ion-tabs>
        <ion-header>
          <ion-toolbar>
            <ion-buttons slot="start">
              <ion-menu-button></ion-menu-button>
            </ion-buttons>
            <ion-title>Logger.Fitness</ion-title>
          </ion-toolbar>
        </ion-header>
        <ion-router-outlet id="main"></ion-router-outlet>
        <ion-tab-bar>
          <ion-tab-button tab="exercise_list" href="/tabs/exercise/list">
            <ion-icon :icon="create" />
            <ion-label>Exercise List</ion-label>
          </ion-tab-button>

          <ion-tab-button tab="workouts" href="/tabs/workouts/list">
            <ion-icon :icon="fileTrayFull" />
            <ion-label>Workouts</ion-label>
          </ion-tab-button>

          <ion-tab-button tab="analytics" href="/tabs/analytics">
            <ion-icon :icon="analyticsOutline" />
            <ion-label>Analytics</ion-label>
          </ion-tab-button>
        </ion-tab-bar>
      </ion-tabs>
    </ion-page>
  </ion-page>
</template>

<script setup lang="ts">
import {
  IonTabBar,
  IonTabButton,
  IonTabs,
  IonLabel,
  IonIcon,
  IonPage,
  IonRouterOutlet,
  IonContent,
  IonMenu,
  IonTitle,
  IonToolbar,
  IonHeader,
  IonItem,
  IonList,
  IonButtons,
  IonButton,
  IonMenuButton,
  menuController,
  IonAvatar,
  toastController,
} from "@ionic/vue";
import { create, fileTrayFull, analyticsOutline } from "ionicons/icons";
import store from "@/store";
import { ref, computed } from "@vue/runtime-core";

const txAmount = computed(()=>  store.getters["exerciseTypes/txAmount"])

function logout() {
  // TODO: need to close menu pane on click
  store.dispatch("auth/logout");
  menuController.close();
}

const user = ref(JSON.parse(localStorage.getItem("user") as string));

async function syncData(manual: boolean){
  const errorToast = await toastController.create({
    message: "Error syncig",
    duration: 2000
  });

  const nothingToSyncToast = await toastController.create({
    message: "Nothing to Sync",
    duration: 2000
  });
  const manualToast = await toastController.create({
    message: "Manual Sync",
    duration: 1000
  });

  if (manual) manualToast.present()

  if (manual && !store.getters["exerciseTypes/needToSync"]){
    nothingToSyncToast.present()
  }

  if (store.getters["exerciseTypes/needToSync"]){
    store.dispatch("exerciseTypes/syncTx")
    .then(()=>{
      if (manual) manualToast.dismiss()
    })
    .catch(()=>{
      if (manual) errorToast.present()
    });
  }
}

/*setInterval(syncData, 3000)*/

</script>

<style scoped>
.logout-btn {
  padding-left: 1%;
  padding-right: 1%;
}
</style>
