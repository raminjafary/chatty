<template>
  <div>
    <h1>Rooms</h1>

    <div>
      <input placeholder="Insert your room name..." v-model="newRoomName">
      <button @click="createRoom"> Create new room</button>
    </div>  
    <div class="rooms-wrapper">
      <div class="room-card" :key="room.id" v-for="room in rooms">
        <router-link :to="`/rooms/${room.id}`"> {{ room.name }}</router-link>
      </div>
    </div>
  </div>
</template>
  
<script>
import axios from 'axios';
export default {
  data() {
    return {
      newRoomName: '',
      rooms: []
    };
  },
  mounted() {
    this.fetchRooms();
  },
  methods: {

    fetchRooms() {
      axios.get('http://localhost:8000/v1/rooms', {
        withCredentials: true
      })
        .then(response => {
          // Handle successful login
          this.rooms = response.data;
          //to do redirect user to new room
        })
        .catch(error => {
          // Handle login error
          console.error(error);
        });
    },
    createRoom() {

      const payload = {
        id: String(Math.floor(Math.random() * 100000) + 1),
        userId: String(Math.floor(Math.random() * 100000) + 1),
        roomName: this.newRoomName,
        // Members: []
      };

      axios.post('http://localhost:8000/v1/rooms/create', payload, {
        withCredentials: true
      })
        .then(response => {
          // Handle successful login
          console.log(response.data);
          this.fetchRooms()
          //to do redirect user to new room
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
.room-card {
  background-color: #dfdfdf;
  margin: 5px;
  width: 200px;
  text-align: center;
  height: 50px;
  line-height: 50px;
  border-radius: 5px;
}

.room-card a {
  color: #494949;
}

.rooms-wrapper {
  display: flex;
  margin-top: 10px;
}
</style>