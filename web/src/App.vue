<template>
  <!-- TODO: this needs to be removed -->
  <link
    href="https://fonts.googleapis.com/icon?family=Material+Icons"
    rel="stylesheet"
  />
  <div id="app">
    <!-- <nav class="navbar navbar-expand navbar-dark bg-dark">
      <div class="navbar-nav mr-auto">
        <li class="nav-link nav-item">
          <i @click="hamburgerMenu" class="fas fa-question material-icons">
            menu
          </i>
        </li>
      </div>

      <div class="collapse navbar-collapse" v-if="menuShow">
        <ul class="navbar-nav mr-auto">
          <li class="nav-item active">
            <a class="nav-link" href="#"
              >Home <span class="sr-only">(current)</span></a
            >
          </li>
          <li class="nav-item">
            <a class="nav-link" href="#">Features</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="#">Pricing</a>
          </li>
        </ul>
      </div>

      <div class="navbar-nav mr-auto">
        
      </div>

    </nav> -->
    <nav class="navbar navbar-dark bg-dark">
      <a class="navbar-brand" href="#">Logger.Fitness</a>

      <button
        class="navbar-toggler toggler-example"
        type="button"
        data-toggle="collapse"
        data-target="#navbarSupportedContent1"
        aria-controls="navbarSupportedContent1"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>

      <div class="collapse navbar-collapse" id="navbarSupportedContent1">
        <ul class="navbar-nav mr-auto">
          <li class="nav-item">
            <router-link to="/home" class="nav-link"> Home </router-link>
          </li>
          <li class="nav-item">
            <router-link to="/exercise_type_list" class="nav-link">
              Exercise Types
            </router-link>
          </li>
          <li class="nav-item">
            <router-link to="/workouts" class="nav-link">Workouts</router-link>
          </li>

          <div v-if="!currentUser" class="navbar-nav ">
            <li class="nav-item">
              <router-link to="/register" class="nav-link">
                <font-awesome-icon icon="user-plus" />
                Register
              </router-link>
            </li>
            <li class="nav-item">
              <router-link to="/login" class="nav-link">
                <font-awesome-icon icon="sign-in-alt" />
                Login
              </router-link>
            </li>
          </div>

          <div v-if="currentUser" class="navbar-nav">
            <li class="nav-item">
              <router-link to="/profile" class="nav-link">
                <font-awesome-icon icon="user" />
                {{ currentUser.claim.username}}
              </router-link>
            </li>
            <li class="nav-item">
              <a class="nav-link" @click.prevent="logOut">
                <!-- <font-awesome-icon icon="sign-out-alt" /> -->
                Logout
              </a>
            </li>
          </div>
        </ul>
      </div>
    </nav>

    <div class="container">
      <router-view />
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return { menuShow: false };
  },
  computed: {
    currentUser() {
      return this.$store.state.auth.user;
    },
  },
  methods: {
    hamburgerMenu() {
      this.menuShow = !this.menuShow;
    },
    logOut() {
      this.$store.dispatch("auth/logout");
      this.$router.push("/login");
    },
  },
};
</script>

<style scoped>
.fa-1x {
  font-size: 1.5rem;
}
.navbar-toggler.toggler-example {
  cursor: pointer;
}
.dark-blue-text {
  color: #0a38f5;
}
.dark-pink-text {
  color: #ac003a;
}
.dark-amber-text {
  color: #ff6f00;
}
.dark-teal-text {
  color: #004d40;
}
</style>
