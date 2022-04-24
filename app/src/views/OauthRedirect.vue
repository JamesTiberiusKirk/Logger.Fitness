<template>
  <ion-page>
    <p>Redirect page</p>
    <!-- <p v-if="metaData.success">{{ metaData.message }}</p> -->
  </ion-page>
</template>

<script setup lang="ts">
import { IonPage } from "@ionic/vue";
import store from "@/store";
import { useRoute } from "vue-router";
import router from "@/router";
import { onMounted, ref } from "vue";

const metaData = ref({
  success: false,
  message: "",
});

onMounted(() => {
  const route = useRoute();

  const params = {
    state: route.query["state"],
    code: route.query["code"],
    scope: route.query["scope"],
    authUser: route.query["authuser"],
    prompt: route.query["prompt"],
  }

  store.dispatch("auth/oauthCallback", params).then(
    (data) => {
      console.log(data);
      metaData.value.success = true;
      metaData.value.message = data;
      router.push("/tabs")
    },
    (error) => {
      console.log(error);
      metaData.value.success = false;
      metaData.value.message = error.message;
    }
  );
});
</script>