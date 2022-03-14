
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

            <ion-label v-if="metaData.message" color="danger">{{
              metaData.message
            }}</ion-label>
            <ion-button @click="validateLoginSubmit()" expand="block"
              >Login</ion-button
            >
          </form>
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
  IonLoading
} from "@ionic/vue";
import { UserLoginDTO } from "@/types/user";
import store from "@/store";
import { computed, Ref, ref } from "vue";
import router from "@/router";

const user: Ref<UserLoginDTO> = ref<UserLoginDTO>({
  email: "",
  password: "",
} as UserLoginDTO);

const isValid = ref({
  email: true,
  password: true,
});

const metaData = ref({
  message: "",
  successful: false,
  loading: false,
});

console.log(store.state.auth)

if (store.state.auth.status.loggedIn){
  router.push('/tabs')
}

function validateLoginSubmit() {
  
  // TODO: Need to validate fields first

  metaData.value.loading = true;
  store.dispatch("auth/login", user.value).then(
    () => {
      metaData.value.loading = false;
      metaData.value.successful = true;
      metaData.value.message = "Successful";
      router.push("/tabs");
    },
    (err: any) => {
      metaData.value.loading = false;
      metaData.value.successful = false;
      metaData.value.message = err.request.responseText || err.data;
      if (err.request.status === 401)
        metaData.value.message = "Wrong password or email";
    }
  );
}

// export default {
//   components: {
//     IonCard,
//     IonCardContent,
//     IonCardHeader,
//     IonCardTitle,
//     IonItem,
//     IonLabel,
//     IonInput,
//     IonHeader,
//     IonPage,
//     IonContent,
//     IonToolbar,
//     IonTitle,
//     IonButton,
//   },
//   data() {
//     const user: UserLoginDTO = {} as UserLoginDTO;
//     const isValid = {
//       email: true,
//       password: true,
//     };

//     return {
//       message: "",
//       successful: false,
//       loading: false,
//       user,
//       isValid,
//     };
//   },
//   mounted() {
//     console.log("mounted");
//   },
//   computed: {
//     loggedIn(): boolean {
//       // return store.state.auth.status.loggedIn;
//       return true
//     },
//   },
//   created() {
//     console.log(store);

//     if (this.loggedIn) {
//       // this.$router.push("/profile");
//       // Router push to a logged in page
//       console.log("router push");
//     }
//   },
//   methods: {
//     handleLogin(user: UserLoginDTO) {
//       this.loading = true;
//       this.$store.dispatch("auth/login", user).then(
//         () => {
//           this.successful = true;
//           this.message = "Successful";
//           // this.$router.push("/workouts");
//         },
//         (err: any) => {
//           this.loading = false;
//           this.successful = false;
//           this.message = err.request.responseText || err.data;
//           if (err.request.status === 401)
//             this.message = "Wrong password or email";
//         }
//       );
//     },
//   },
// };
</script>

<style scoped>
</style>
