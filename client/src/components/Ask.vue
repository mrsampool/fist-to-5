<template>
  <div>
    <form @submit.prevent="submitQuestion()">
      <h2>Ask New Question</h2>
      <textarea v-model="newQuestion"></textarea>
      <button type="submit">Submit</button>
    </form>
    <router-link to="/teacher/responses">
      <button>Previous Responses</button>
    </router-link>
  </div>
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

<style scoped>
div{
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
textarea{
  display: block;
  width: 20rem;
  height: 5rem;
  border-radius: 1rem;
  border: 0;
}
</style>