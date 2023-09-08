<template>
  <div>
    <h1>welcome to Room #{{ $route.params.id }} - {{ connectionStatus }}</h1>
    <div class="members-wrapper">
      <div class="member-card" :key="memebr.id" v-for="memebr in memebrs">
        <div> {{ memebr.name }}</div>
      </div>
    </div>
    <hr>
    <div class="form">
      <textarea v-model="newMessage"></textarea>
      <button @click="sendMessage">send</button>
    </div>
    <div class="messages-wrapper">
      <div class="message-card" :key="i" v-for="message, i in messages">
        <div> 
          <div>{{ message.username }}</div>
          <div>{{ message.content }}</div>
        </div>
      </div>
    </div>
  </div>
</template>
  
<script>
import axios from 'axios';
export default {
  data() {
    return {
      clientId: null,
      roomId: this.$route.params.id,
      memebrs: [],
      messages: [],
      connectionStatus: 'Disconnected',
      ws: null,
      newMessage: null,
    };
  },
  mounted() {
    this.fetchMembers();
    this.joinToRoom();
  },
  methods: {

    sendMessage() {
      if (this.connectionStatus == 'Disconnected')
        return
      
      this.ws.send(this.newMessage);
    },
    fetchMembers() {
      axios.get(`http://localhost:8080/ws/getClients/${this.roomId}`, {
        withCredentials: true
      })
        .then(response => {
          // Handle successful login
          this.memebrs = response.data;
          //to do redirect user to new room
        })
        .catch(error => {
          // Handle login error
          console.error(error);
        });
    },
    joinToRoom() {

      const clientId = String(Math.floor(Math.random() * 100000) + 1);
      const Username = String(Math.floor(Math.random() * 100000) + 1);

      this.ws = new WebSocket(`ws://localhost:8080/ws/joinRoom/${this.roomId}?userId=${clientId}&username=${Username}`);

      this.ws.addEventListener('open', () => {
        this.connectionStatus = 'Connected';

        // Handle any other events or logic on connection
      });

      this.ws.addEventListener('close', () => {
        this.connectionStatus = 'Disconnected';
        // Handle any other events or logic on disconnection
      });

      this.ws.addEventListener('message', (event) => {
        // Handle the data received from the server
        this.messages.push(JSON.parse(event.data))
      });



    },
    beforeDestroy() {
      // Close the WebSocket connection when the component is destroyed
      if (this.ws && this.ws.readyState === WebSocket.OPEN) {
        this.ws.close();
      }
    },
  }
};
</script>
  