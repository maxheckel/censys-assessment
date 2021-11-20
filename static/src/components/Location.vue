<template>

  <div v-if="location !== undefined">
    <h1>Location Found!</h1>
    Latitude: {{location.lat}} <br> Longitude: {{location.lng}}
  </div>
</template>

<script>
export default {
  name: "Location",

  props: ['location'],
  mounted() {

  }
}
</script>

<style scoped>

</style>