<template>
  <div class="container">
    <header class="jumbotron">
      <h3>
        <strong>{{currentUser.claim.username}}</strong> Profile
      </h3>
    </header>
    <p>
      <strong>Token:</strong>
      {{currentUser.jwt.substring(0, 20)}} ... {{currentUser.jwt.substr(currentUser.jwt.length - 20)}}
    </p>
    <p>
      <strong>Id:</strong>
      {{currentUser.claim.id}}
    </p>
    <p>
      <strong>Email:</strong>
      {{currentUser.claim.email}}
    </p>
    <strong>Roles:</strong>
    <ul>
      <li v-for="role in currentUser.claim.roles" :key="role">{{role}}</li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'Profile',
  computed: {
    currentUser() {
      let user = this.$store.state.auth.user
      // let roleArray = Array.from(user.claim.roles.values())
      // user.claim.roles = roleArray
      return user
    }
  },
  mounted() {
    if (!this.currentUser) {
      this.$router.push('/login')
    }
  }
};
</script>