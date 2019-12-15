<template>
  <div class="container">


<b-list-group>
  <b-list-group-item button v-for="user in users" :key="user.id">{{user.name}}</b-list-group-item>
  <!-- <b-list-group-item button>I am a button</b-list-group-item>
  <b-list-group-item button>Disabled button</b-list-group-item>
  <b-list-group-item button>This is a button too</b-list-group-item> -->
</b-list-group>







  </div>
</template>

<script>
export default {
  asyncData({ $axios }) {
    return $axios.get("http://localhost:8080/user").then((response) => {
      return {
        users: response.data
      }
    // console.log(result)
    }).catch((err) => {
      console.log(err)      
    });

return {
      users: []
    }
  },
};
</script>

<style>
.container {
  margin: 0 auto;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.title {
  font-family: "Quicksand", "Source Sans Pro", -apple-system, BlinkMacSystemFont,
    "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
  display: block;
  font-weight: 300;
  font-size: 100px;
  color: #35495e;
  letter-spacing: 1px;
}

.subtitle {
  font-weight: 300;
  font-size: 42px;
  color: #526488;
  word-spacing: 5px;
  padding-bottom: 15px;
}

.links {
  padding-top: 15px;
}
</style>
