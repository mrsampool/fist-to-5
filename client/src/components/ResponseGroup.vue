<template>
  <li class="response-group">
    <span class="response-group-numbers">
      <p class="value">{{ group.value }}</p>
      <p>{{ group.percent }}</p>
    </span>
    <span>
      <ul>
        <li
          class="response-name"
          v-for="response in group.responses"
        >
          {{ response['Student']['first_name'] }} {{ response['Student']['last_name'] }}
        </li>
      </ul>
      </span>
  </li>
</template>

<script>
export default {
  name: 'ResponseGroup',
  props: ['group']
}
</script>

<style scoped>
.response-group{
  border-bottom: 0.1rem solid dodgerblue;
  margin-bottom: 1rem;
  display: flex;
  flex-direction: row;
}
.response-group-numbers{
  font-weight: bold;
  display: block;
}
.response-group-numbers .value{
  background-color: dodgerblue;
  border-radius: 1rem;
  color: white;
  margin-bottom: 0.1rem;
}
.response-group-numbers p{
  margin: 0;
}
.response-name{
  list-style-type: circle;
}
ul{

}
ul li span{
  display: flex;
  width: fit-content;
  padding: 0 0.5rem;
  align-items: center;
  justify-content: center;
}
</style>