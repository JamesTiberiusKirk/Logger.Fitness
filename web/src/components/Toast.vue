
<template>
  <!-- <div id="snackbar">{{ title }}</div> -->
  <div id="snackbar" class="card card-container custom-toast">
    <div class="card-body">
      <div class="d-flex w-100 justify-content-between">
        <strong class="mb-1">
          {{ title }}
        </strong>
        <small>
          <i @click="closeClick" class="fas fa-question material-icons">
            clear
          </i>
        </small>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Toast",
  props: {
    title: {
      type: String,
      required: true,
    },
  },
  emits: ["closeToast"],
  mounted() {
    // Get the snackbar DIV
    var toast = document.getElementById("snackbar");

    // Add the "show" class to DIV
    toast.className = "show";

    // After 3 seconds, remove the show class from DIV
    setTimeout(() => {
      toast.className = toast.className.replace("show", "");
      this.$emit("closeToast");
    }, 3000);
  },
  created() {
    console.log("title:", this.title);
    console.log("smallTitle:", this.smallTitle);
    console.log("message:", this.message);
  },
  methods: {
    closeClick() {
      this.$emit("closeToast");
    },
  },
};
</script>

<style scoped>
.card-body {
  padding: 10px;
}

/* From w3 schools */
/* The snackbar - position it at the bottom and in the middle of the screen */
#snackbar {
  visibility: hidden; /* Hidden by default. Visible on click */
  position: fixed;
  bottom: 45px;
  z-index: 9999;
  width: 90%;
  left: 0px;
  margin: 5%;
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.9);
  background: white;
  border-radius: 5px;
  border: 1px white solid ;
}

/* Show the snackbar when clicking on a button (class added with JavaScript) */
#snackbar.show {
  visibility: visible; /* Show the snackbar */
  /* Add animation: Take 0.5 seconds to fade in and out the snackbar.
  However, delay the fade out process for 2.5 seconds */
  -webkit-animation: fadein 0.5s, fadeout 0.5s 2.5s;
  animation: fadein 0.5s, fadeout 0.5s 2.5s;
}

/* Animations to fade the snackbar in and out */
@-webkit-keyframes fadein {
  from {
    bottom: 0;
    opacity: 0;
  }
  to {
    bottom: 45px;
    opacity: 1;
  }
}

@keyframes fadein {
  from {
    bottom: 0;
    opacity: 0;
  }
  to {
    bottom: 45px;
    opacity: 1;
  }
}

@-webkit-keyframes fadeout {
  from {
    bottom: 45px;
    opacity: 1;
  }
  to {
    bottom: 0;
    opacity: 0;
  }
}

@keyframes fadeout {
  from {
    bottom: 45;
    opacity: 1;
  }
  to {
    bottom: 0;
    opacity: 0;
  }
}
</style>