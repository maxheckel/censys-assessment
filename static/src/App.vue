<template>
  <h1> IP Lookup</h1>
  <IPLookup/>
</template>

<script>
import IPLookup from './components/IPLookup.vue'

export default {
  name: 'App',
  components: {
    IPLookup
  }
}
</script>


<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
