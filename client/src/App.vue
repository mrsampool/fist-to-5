<template>
  <top-nav v-bind:mode="this.mode"></top-nav>
  <router-view></router-view>
</template>

<script>
// import HelloWorld from './components/HelloWorld.vue'
import StudentView from './views/StudentView';
import TopNav from "./components/TopNav";
import './global.css';

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
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
}
html{
  background-color: lightblue;
}
</style>
