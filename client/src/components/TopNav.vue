<template>
  <nav>
    <span>Fist to 5</span>
    <span>
      <p>{{ mode }}</p>
      <router-link to="/">
        <button>Logout</button>
      </router-link>
    </span>
  </nav>
</template>

<script>
export default {
  name: 'TopNav',
  props: ['mode']
}
</script>

<style scoped>
nav{
  width: 100%;
  background-color: dodgerblue;
  color: white;
  font-weight: bold;
  font-size: 2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}
nav span{
  padding: 1rem;
  display: flex;
  align-items: center;
}
p{
  margin: 0;
  font-size: 1.25rem;
}
button{
  background-color: salmon;
  color: white;
  padding: 0.25rem 0.5rem;
  font-size: 1rem;
}
</style>