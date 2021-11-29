<template>
  <div id="login">
    <h1>Login</h1>
    <input
      class ="login_field"
      type="text"
      name="email"
      v-model="input.email"
      placeholder="Email"
    />
    <br/>
    <input
      class ="login_field"
      type="password"
      name="password"
      v-model="input.password"
      placeholder="Password"
    />
    <br/>
    <button type="button" v-on:click="login()">Login</button>
  </div>
</template>

<script>
export default {
  name: 'Login',
  data () {
    return {
      input: {
        email: '',
        password: ''
      }
    }
  },
  methods: {
    login () {
      const api = require('../lib/api')
      if (this.input.email === '' && this.input.password === '') {
        console.log('A username and password must be present')
        return
      }

      api.login(this.input.email, this.input.password)
        .catch(err => {
          console.log(err)
        }).then(res => {
          console.log(res)
        })
    }
  }
}
</script>

<style scoped>
#login {
  width: 95%;
  border: 1px solid #cccccc;
  background-color: #ffffff;
  margin: auto;
  margin-top: 40%;
  padding-bottom: 20px;
}
</style>
