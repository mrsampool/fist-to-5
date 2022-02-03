<template>
  <h2>Responses</h2>
  <main>
    <h3>{{ question.text }}</h3>
    <ul>
      <response-group
          v-for="group in responseGroups"
          v-bind:group="group">
      </response-group>
    </ul>
  </main>
  <router-link to="/ask">
    <button>New Question</button>
  </router-link>
</template>

<script>
import { getCurrentResponses } from "../utils/serverUtils";
import ResponseGroup from "./ResponseGroup";
export default {
  name: 'Responses',
  components: {ResponseGroup},
  data: function(){
    return {
      question: '',
      responses: []
    }
  },
  mounted() {
    this.fetchData();
  },
  computed: {
    responseGroups: function(){
      const groups = {};
      this.responses.forEach(response => {
        if (!groups[response.value]) {
          groups[response.value] = { value: response.value, responses: [response]};
        } else {
          groups[response.value].responses.push(response);
        }
      });
      Object.values(groups).forEach(group => {
        group.percent = `${
            ((100 / this.responses.length) * group.responses.length).toFixed(2)
        }%`;
      })
      return groups;
    }
  },
  methods: {
    fetchData(){
      getCurrentResponses()
          .then(resData => {
            this.question = resData.question;
            this.responses = resData.responses;
          })
          .catch(err => console.error(err));
    }
  },
}
</script>

<style scoped>
textarea{
  display: block;
  width: 20rem;
  height: 5rem;
  border-radius: 1rem;
  border: 0;
}
</style>