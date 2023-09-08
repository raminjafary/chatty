<template>
  <div>
    <h1>Register</h1>
    <form class="form" @submit.prevent="register">
      <label for="username">Username:</label>
      <input type="text" id="username" v-model="username" required>
      <br>
      <label for="email">Email:</label>
      <input type="email" id="email" v-model="email" required>
      <br>
      <label for="password">Password:</label>
      <input type="password" id="password" v-model="password" required>
      <br>
      <button type="submit">Register</button>
    </form>
  </div>
</template>
  
<script>
import axios from 'axios';

export default {
  data() {
    return {
      username: '',
      email: '',
      password: ''
    };
  },
  methods: {
    register() {
      const data = {
        username: this.username,
        email: this.email,
        password: this.password,
      };

      axios.post('http://localhost:8080/register', data)
        .then(response => {

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
.form {
  display: flex;
  flex-direction: column;
  width: 300px;
  max-width: 100%;
  margin: auto;
  text-align: left;
}
</style>