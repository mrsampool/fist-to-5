<template>
  <!--<img alt="Vue logo" src="./assets/logo.png">-->
  <!--<HelloWorld msg="Welcome to Your Vue.js App"/>-->
  <StudentView />
</template>

<script>
// import HelloWorld from './components/HelloWorld.vue'
import StudentView from './components/StudentView';

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
  text-transform: uppercase;
  background-color: white;
  color: #2c3e50;
  font-weight: bolder;
  font-size: 1.1rem;
  border-radius: 1rem;
  border: 0;
}
</style>
