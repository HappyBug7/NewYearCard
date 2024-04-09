<template>
  <div id="Home">
    <typing-area></typing-area>
    <firework style="z-index: -1;"></firework>
    <button @click="ToCards" id="Start">开启祝福</button>
    <router-view>

    </router-view>
  </div>
</template>

<script>

import TypingArea from '../components/TypingArea.vue';
import Firework from '../components/Firework.vue';

export default {
  name: 'Home',
  components: {
    TypingArea,
    Firework
  },
  methods:{
    ToCards(){
      this.$router.push('/cards')
    }
  }
}
</script>

<style scoped>
  #Start{
    width: 150px;
    height: 50px;
    background-color: #d14851;
    border-style: none;
    border-radius: 10px;
  }
  #Home {
    display: flex;
    align-items: center;
    flex-direction: column;
    font-family: Avenir, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
  }
</style>
