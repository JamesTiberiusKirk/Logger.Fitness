<template>
  <div id="app">
    <router-view/>
  </div>
</template>
<script>
export default {
  name: 'App'
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

input{
  width: 90%;
  height: 35px;
  margin: 5px;
  text-align: center;
}

button{
  width: 90%;
  height: 35px;
  margin: 5px;
  text-align: center;
}
</style>
