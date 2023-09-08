<template>
  <div>
    <h1>Login</h1>
    <form class="form" @submit.prevent="login">
      <label for="email">Email:</label>
      <input type="email" id="email" v-model="email" required>
      <br>
      <label for="password">Password:</label>
      <input type="password" id="password" v-model="password" required>
      <br>
      <button type="submit">Login</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      email: '',
      password: ''
    };
  },
  methods: {
    login() {
      const credentials = {
        email: this.email,
        password: this.password
      };

      axios.post('http://localhost:8080/login', credentials)
        .then(response => {

          const jwtToken = response.data.access_token
          document.cookie = `jwt=${jwtToken}; path=/`;
        
          // Handle successful login
          console.log(response.data);
        })
        .catch(error => {
          // Handle login error
          console.error(error);
        });
    }
  }
};
</script>

<style scoped>

.form{
  display: flex;
    flex-direction: column;
    max-width: 200px;
    margin: auto;
    text-align: left;
}


</style>