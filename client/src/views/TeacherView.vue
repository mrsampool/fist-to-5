<template>
  <div id="student-view-wrap">
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: 'TeacherView'
}
</script>

<style scoped>

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