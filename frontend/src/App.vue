<template>
  <div id="app">
    <button type="button" @click="getFortune()">Get a new fortune message</button>
    <div v-if="fortune != null">
      <p>{{ fortune }}</p>
    </div>
  </div>
</template>

<script>

export default {
  name: "App",
  components: {},
  data() {
    return {
      fortune: null,
    };
  },
  methods: {
    getFortune() {
       this.axios.get("https://fortune-2.teonet.app").then((response) => (this.fortune = response.data));
    },
  },
};
</script>