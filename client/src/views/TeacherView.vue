<template>
  <div id="student-view-wrap">
    <h1>Teacher View</h1>
  </div>
</template>

<script>
import axios from 'axios';
import { getCurrentQuestion } from "../utils/serverUtils";

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
    postQuestion() {
      if (this.currQuestion && this.currQuestion.id && this.value !== null ){
        axios.post(`/api/responses/${this.currQuestion.id}`,
            {
              studentId: this.currStudent.id,
              value: this.value,
            })
            .then(({data}) => {
              this.currQuestion = '';
            })
      }
    }
  },
  mounted: function() {
    // getCurrentQuestion(this);
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