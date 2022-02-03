<template>
  <form @submit.prevent="submitQuestion()">
    <h2>Ask New Question</h2>
    <textarea v-model="newQuestion"></textarea>
    <button>Submit</button>
  </form>
  <router-link to="/responses">
    <button>Previous Responses</button>
  </router-link>
</template>

<script>
import {postQuestion} from "../utils/serverUtils";
export default {
  name: 'Ask',
  data: function(){
    return {
      newQuestion: ''
    }
  },
  methods: {
    submitQuestion(){
      postQuestion(this.newQuestion)
      .then(() => this.newQuestion = '');
    }
  }
}
</script>

<style>
textarea{
  display: block;
  width: 20rem;
  height: 5rem;
  border-radius: 1rem;
  border: 0;
}
</style>