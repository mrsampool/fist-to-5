<template>
  <div id="student-view-wrap">
    <h1>Student View</h1>
    <form v-if="currQuestion" @submit.prevent="postQuestion">
      <h2>{{ currQuestion.text }}</h2>
      <div id="values-list">
        <label v-for="n in 6" :key="'option-'+n">
          {{  !(n - 1) ? 'fist' : n - 1  }}
          <input type="radio" v-bind:value="n - 1" v-model="value">
        </label>
      </div>
      <button type="submit">Submit</button>
    </form>
    <button type="submit" v-if="!currQuestion" @click="fetchQuestion">Get Next Question</button>
  </div>
</template>

<script>
import axios from 'axios';
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
      axios.get(`/api/question/current`).then(({data}) => {
        this.currQuestion = data;
      });
    },
    postQuestion() {
      if (this.currQuestion && this.currQuestion.id && this.value !== null ){
        axios.post(`/api/responses/${this.currQuestion.id}`,
            {
              studentId: this.currStudent.id,
              value: this.value,
            })
            .then(({data}) => {
              console.log(data);
              this.currQuestion = '';
            })
      }
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