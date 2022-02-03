<template>
  <div id="student-view-wrap">
    <h1>I am a...</h1>
    <router-link to="/student"><button>Student</button></router-link>
    <router-link to="/teacher"><button>Teacher</button></router-link>
  </div>
</template>

<script>
export default {
  name: 'StudentView'
}
</script>

<style scoped>
button{
  margin: 0.5rem;
}
</style>