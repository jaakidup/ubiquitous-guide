<template>
  <div>
    <div class="mt-3 mb-5">
      <b-form-input ref="input" placeholder="Enter a name" v-model="user.name"></b-form-input>
      <b-button-group class="mt-3">
        <b-button type="submit" variant="outline-primary" @click="submit">Submit</b-button>
        <b-button type="reset" variant="outline-danger" @click="reset">Reset</b-button>
      </b-button-group>
    </div>

    <b-list-group v-show="!editing">
      <div class="pt-3 mt-3" v-for="user in users" :key="user.id">
        <b-button-group>
          <b-button
            size="sm"
            :disabled="editing"
            variant="outline-danger"
            @click="deleteUser(user)"
          >DELETE</b-button>
          <b-button size="sm" variant="outline-success" @click="editUser(user)">EDIT</b-button>
          <b-button size="sm" variant="outline-primary" :to="{path:'/'+user.id}">VIEW</b-button>
        </b-button-group>
        <span class="ml-3">{{user.name}}</span>
      </div>
    </b-list-group>
  </div>
</template>

<script>
export default {
  asyncData({ $axios }) {
    return $axios
      .get("http://localhost:8080/users")
      .then(response => {
        return {
          users: response.data,
          user: {},
          editing: false
        };
      })
      .catch(err => {
        console.log(err);
      });
  },
  data() {
    this.resfreshData();
  },
  methods: {
    resfreshData() {
      this.$axios
        .get("http://localhost:8080/users")
        .then(response => {
          return {
            users: response.data,
            name: {},
            editing: false
          };
        })
        .catch(err => {
          console.log(err);
        });
    },
    reset() {
      this.name = "";
    },
    submit() {
      if (this.user.name == null) {
        console.log("Username is empty");
        return;
      }

      this.$axios
        .post("http://localhost:8080/users", this.user)
        .then(response => {
          if (!this.editing) {
            this.users.push(response.data);
          }
          this.user = {};
          this.editing = false;
        })
        .catch(err => {
          console.log(err);
          this.editing = false;
        });
    },
    editUser(user) {
      this.user = user;
      this.editing = true;
    },
    deleteUser(user) {
      let url = "http://localhost:8080/users/" + user.id;
      this.$axios
        .delete(url)
        .then(response => {
          let index = this.users.findIndex(item => {
            return item.id == user.id;
          });
          this.users.splice(index, 1);
        })
        .catch(err => {
          console.log(err);
        });
    }
  }
};
</script>

<style>
</style>
