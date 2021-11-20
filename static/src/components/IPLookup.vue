<template>
  <div class="hello">
     <input type="text" placeholder="address" v-model="currentIP">
    <button @click="lookupIP"  type="submit">Lookup</button>
  </div>
  <Location v-bind:location="location ? location.location : undefined"></Location>
</template>

<script>
import Location from "./Location";
export default {
  name: 'IPLookup',
  components: {Location},
  props: {
    msg: String
  },
  data() {
    return {
      currentIP: '',
      location: {},
      error: undefined
    }
  },
  methods: {
    lookupIP(){
      const url = process.env.VUE_APP_API_BASE + "/ip/" + this.currentIP

      fetch(url)
          .then(response => response.json())
          .then(data => {
            this.location = data
          })
          .catch(err => {console.log(err)})

    }
  },
  mounted(){
    fetch('http://jsonip.com/?callback=')
        .then(response => response.json())
        .then(data => {
          this.currentIP = data.ip
        })
        .catch(err => {console.log(err)})
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}

input, button{
  padding: 10px;
  font-size: 18px;
}

button{
  margin-left: 10px
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
