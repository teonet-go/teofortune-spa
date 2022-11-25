<template>
  <div id="app">
    <h1>Teonet Fortune single-page application</h1>
    <p class="subj">
        This simple single-page application calls Teonet fortune Google 
        Cloud Function. Click the 'Click' button to execute 
        fortune Google Cloud Function: 
        <a href="https://fortune-2.teonet.app">https://fortune-2.teonet.app</a>
    </p>
    <button type="button" @click="getFortune()">Click</button>
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

<style>
    .container {
        margin-top: 25px;
        margin-bottom: 25px;
        max-width: 835px;
    }
    .subj {
        color: gray;
        font-size: small;
    }
    .subj a {
        color: gray;
    }
</style>