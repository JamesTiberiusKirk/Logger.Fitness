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
            <ion-label>Inbox</ion-label>
          </ion-item>
          <ion-item>
            <ion-label>Outbox</ion-label>
          </ion-item>
          <ion-item>
            <ion-label>Favorites</ion-label>
          </ion-item>
          <ion-item>
            <ion-label>Archived</ion-label>
          </ion-item>
          <ion-item>
            <ion-label>Trash</ion-label>
          </ion-item>
        </ion-list>
        <ion-button class="logout-btn" expand="block" @click="logout()">
          Logout
        </ion-button>
        <ion-avatar v-if="user">
          <img :src="user.claim.profile_picture" />
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
} from "@ionic/vue";
import { create, fileTrayFull, analyticsOutline } from "ionicons/icons";
import store from "@/store";
import { ref, onMounted } from "@vue/runtime-core";

function logout() {
  // TODO: need to close menu pane on click
  store.dispatch("auth/logout");
  menuController.close();
}

const user = ref(JSON.parse(localStorage.getItem("user") as string));
console.log(user.value);
</script>

<style scoped>
.logout-btn {
  padding-left: 1%;
  padding-right: 1%;
}
</style>