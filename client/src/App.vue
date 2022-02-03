<template>
  <!--<img alt="Vue logo" src="./assets/logo.png">-->
  <!--<HelloWorld msg="Welcome to Your Vue.js App"/>-->
  <!--<StudentView />-->
  <router-view></router-view>
</template>

<script>
// import HelloWorld from './components/HelloWorld.vue'
import StudentView from './views/StudentView';

export default {
  name: 'App',
  components: {
    // HelloWorld,
    StudentView
  }
}
</script>

<style>
#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
}
html, button, label, h1, h2{
  font-family: Avenir, Helvetica, Arial, sans-serif;
}
html{
  background-color: lightblue;
  --transparent-light: hsla(0, 0%, 100%, 0.5)
}
button{
  padding: 1rem 2rem;
  margin: 0.5rem;
  text-transform: uppercase;
  background-color: dodgerblue;
  color: white;
  font-weight: bolder;
  font-size: 1.1rem;
  border-radius: 1rem;
  border: 0;
}
form, main{
  width: fit-content;
  background-color: var(--transparent-light);
  border-radius: 1rem;
  padding: 1rem;
  display: flex;
  flex-direction: column;
  align-items: center;
}
</style>
