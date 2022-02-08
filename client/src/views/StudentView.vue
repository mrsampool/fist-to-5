<template>
  <div id="student-view-wrap">
    <form
        v-if="currQuestion"
        @submit.prevent="submitResponse()"
    >
      <h1>{{ currQuestion.text }}</h1>
      <div id="values-list">
        <label v-for="n in 6" :key="'option-'+ n">
          {{  !(n - 1) ? 'fist' : n - 1  }}
          <input type="radio" v-bind:value="n - 1" v-model="value" :id="'radio-' + (n - 1)">
        </label>
      </div>
      <button type="submit">Submit</button>
    </form>
    <button
        type="submit"
        v-if="!currQuestion"
        @click="fetchQuestion"
    >Get Next Question
    </button>
  </div>
</template>

<script>
import {getCurrentQuestion, postResponse} from "../utils/serverUtils";

export default {
  name: 'StudentView',
  data: function(){
    return {
      currStudent: { id: 1, firstName: 'Sam', lastName: 'Pool' },
      currQuestion: null,
      value: null
    }
  },
  methods: {
    fetchQuestion() {
      getCurrentQuestion()
        .then((question) => this.currQuestion = question)
        .catch(err => console.log(err))
    },
    submitResponse() {
      postResponse(this.currQuestion, this.currStudent, this.value)
        .then(() => this.currQuestion = '')
        .catch(err => console.log(err))
    }
  },
  mounted: function() {
    this.fetchQuestion()
  }
}
</script>

<style scoped>
form{
  width: fit-content;
  background-color: var(--transparent-light);
  border-radius: 1rem;
  padding: 1rem;
  display: flex;
  flex-direction: column;
  align-items: center;
}
label{
  display: flex;
  flex-direction: column;
  width: fit-content;
  margin: 1rem;
}
#student-view-wrap{
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
}
#values-list{
  display: flex;
}
</style>