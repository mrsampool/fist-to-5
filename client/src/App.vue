<template>
  <top-nav v-bind:mode="this.mode"></top-nav>
  <router-view></router-view>
</template>

<script>
// import HelloWorld from './components/HelloWorld.vue'
import StudentView from './views/StudentView';
import TopNav from "./components/TopNav";

export default {
  name: 'App',
  // components: [TopNav],
  computed: {
    mode: function(){
      if (this.$route.fullPath.indexOf('teacher') !== -1){
        return 'teacher';
      } else if (this.$route.fullPath.indexOf('student') !== -1) {
        return 'student';
      } else {
        return null;
      }
    }
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
body{
  margin: 0;
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
