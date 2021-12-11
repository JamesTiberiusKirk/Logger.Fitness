<template>
  <div id="app">
    <nav class="navbar navbar-expand navbar-dark bg-dark">
      <div class="navbar-nav mr-auto">
        <li class="nav-item">
          <router-link to="/home" class="nav-link">
            <font-awesome-icon icon="home" /> Logger.Fitness
          </router-link>
        </li>
      </div>

      <div v-if="!currentUser" class="navbar-nav ml-auto">
        <li class="nav-item">
          <router-link to="/register" class="nav-link">
            <font-awesome-icon icon="user-plus" />
          </router-link>
        </li>
        <li class="nav-item">
          <router-link to="/login" class="nav-link">
            <font-awesome-icon icon="sign-in-alt" />
          </router-link>
        </li>
      </div>

      <div v-if="currentUser" class="navbar-nav ml-auto">
        <li class="nav-item">
          <router-link to="/profile" class="nav-link">
            <font-awesome-icon icon="user" />
            {{ currentUser.username }}
          </router-link>
        </li>
        <li class="nav-item">
          <a class="nav-link" @click.prevent="logOut">
            <font-awesome-icon icon="sign-out-alt" />
          </a>
        </li>
      </div>
    </nav>

    <link
      href="https://fonts.googleapis.com/icon?family=Material+Icons"
      rel="stylesheet"
    />
    <div class="container">
      <router-view />
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    currentUser() {
      return this.$store.state.auth.user;
    },
  },
  methods: {
    logOut() {
      this.$store.dispatch("auth/logout");
      this.$router.push("/login");
    },
  },
};
</script>


<style>
.fab-container {
  position: fixed;
  bottom: 20px;
  right: 20px;
  z-index: 999;
  cursor: pointer;
}

.fab-icon-holder {
  width: 45px;
  height: 45px;
  border-radius: 100%;
  background: black;
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.9);
}

.fab-icon-holder:hover {
  opacity: 0.8;
}

.fab-icon-holder i {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  font-size: 25px;
  color: #ffffff;
}

.fab {
  width: 60px;
  height: 60px;
  background: var(--primary);
}

.fab-options {
  list-style-type: none;
  margin: 0;
  position: absolute;
  bottom: 70px;
  right: 0px;
  opacity: 0;
  transition: all 0.2s ease;
  transform: scale(0);
  transform-origin: 70% bottom;
}

.fab:hover + .fab-options,
.fab-options:hover {
  opacity: 1;
  transform: scale(1);
}

.fab-options li {
  display: flex;
  justify-content: flex-end;
  padding: 5px;
}
</style>