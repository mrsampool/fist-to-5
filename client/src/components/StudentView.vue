<template>
  <form>
    <h1>Student View</h1>
    <h2>{{ currQuestion }}</h2>
    <div id="values-list">
      <label v-for="n in 5" :key="'option-'+n">
        {{  !(n - 1) ? 'fist' : n  }}
        <input type="radio" v-bind:value="n - 1">
      </label>
    </div>
  </form>
</template>

<script>
import axios from 'axios';
export default {
  name: 'StudentView',
  data: function(){
    return {
      currQuestion: 'Loading question...'
    }
  },
  methods: {
    fetchQuestion() {
      axios.get(`/api/question/current`).then(({data}) => {
        this.currQuestion = data.text;
      } );
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
}
label{
  display: flex;
  flex-direction: column;
  width: fit-content;
  margin: 1rem;
}
#values-list{
  display: flex;
}
</style>