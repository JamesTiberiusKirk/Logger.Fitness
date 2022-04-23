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
import { onMounted, ref } from "vue";

const metaData = ref({
  success: false,
  message: "",
});

onMounted(() => {
  const route = useRoute();
  store.dispatch("auth/oauthCallback", route.query).then(
    (data) => {
      console.log(data);
      metaData.value.success = true;
      metaData.value.message = data;
    },
    (error) => {
      console.log(error);
      metaData.value.success = false;
      metaData.value.message = error.message;
    }
  );
});
</script>