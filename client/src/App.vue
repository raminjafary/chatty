<template>
  <div>
    <div>
      <router-link class="p-10" to="/">Home</router-link>
      <router-link class="p-10" to="/login">Login</router-link>
      <router-link class="p-10" to="/register">Register</router-link>
    </div>
    <router-view></router-view>
  </div>
</template>

<script>

export default {
  name: 'App',
  components: {

  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

.p-10{
  padding: 10px;
}
</style>
