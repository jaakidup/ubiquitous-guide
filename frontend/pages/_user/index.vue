<template>
  <div>
    <div class="mt-3 mb-5">
      <b-form-input ref="input" placeholder="Add Task Description" v-model="task.description"></b-form-input>
      <b-button-group class="mt-3">
        <b-button type="submit" variant="outline-primary" @click="submit">Submit</b-button>
        <b-button type="reset" variant="outline-danger" @click="reset">Reset</b-button>
      </b-button-group>
    </div>


    <h3>Tasks:</h3>


    <b-list-group v-show="!editing">
      <div class="pt-3 mt-3" v-for="task in tasks" :key="task.id">
        <b-button-group>
          <b-button
            size="sm"
            :disabled="editing"
            variant="outline-danger"
            @click="deleteTask(task)"
          >DELETE</b-button>
          <b-button size="sm" variant="outline-success" @click="editTask(task)">EDIT</b-button>
        </b-button-group>
          <b-form-checkbox
            v-model="task.state"
            size="sm"
            button
            button-variant="info"
            @change="checker(task)"
            name="state"
            value="done"
            unchecked-value="to do"
          >{{ task.state }}</b-form-checkbox>
        <span class="ml-3">{{task.description}}</span>
      </div>
    </b-list-group>
  </div>
</template>

<script>
export default {
  asyncData({ $axios, params }) {
    console.log(params.id);
    return $axios
      .get("http://localhost:8080/tasks/" + params.user)
      .then(response => {
        return {
          tasks: response.data,
          editing: false,
          params: params,
          task: {},
          user: params.user
        };
      })
      .catch(err => {
        console.log(err);
      });
  },
  // data() {

  //   return $axios
  //     .get("http://localhost:8080/tasks/" + this.user)
  //     .then(response => {
  //       return {
  //         tasks: response.data,
  //         editing: false,
  //         task: {},
  //         user: .user
  //       };
  //     })
  //     .catch(err => {
  //       console.log(err);
  //     });
  // },
  methods: {
    reset() {
      this.task = {}
      this.editing = false
    },
    submit() {
      if (this.task.description == null) {
        console.log("description is empty");
        return;
      }
      this.task.state = "to do";
      this.task.user_id = Number(this.user);
      this.$axios
        .post("http://localhost:8080/task/" + this.user, this.task)
        .then(response => {
          if (!this.editing) {
            this.tasks.push(response.data);
          }
          this.task = {};
          this.editing = false;
        })
        .catch(err => {
          console.log(err);
          this.editing = false;
        });
    },
    checker(task) {
      task.user_id = Number(this.user);
      this.$axios
        .post("http://localhost:8080/task/" + this.user, task)
        .then(response => {})
        .catch(err => {
          console.log(err);
          this.editing = false;
        });
    },
    updateTask() {},
    deleteTask(task) {
      let url = "http://localhost:8080/task/" + this.user + "/" + task.id;
      this.$axios
        .delete(url)
        .then(response => {
          let index = this.tasks.findIndex(item => {
            return item.id == task.id;
          });
          this.tasks.splice(index, 1);
        })
        .catch(err => {
          console.log(err);
        });
    },
    editTask(task) {
      this.task = task;
      this.editing = true;
    }
  }
};
</script>

<style>
</style>