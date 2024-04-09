<template>
  <div id="main">
    <img src="../assets/background.jpg" id="back">
    <div class="container">
      <typing-words style="margin: auto;" :words=word></typing-words>
    </div>
    <!-- <button id="Write" @click="ToWrite">我也要写</button> -->
  </div>
</template>

<script>
import TypingWords from '@/components/TypingWords.vue';
import axios from 'axios'

export default {
  components: { TypingWords },
  name:"Celebration",
  methods:{
    ToWrite(){
      axios.get('/api/hello')
      .then(res => {console.log(res.data["data"]);this.word.push(res.data["data"]);})
      .catch(err => console.log(err))
    }
  },
  data(){
    return{
      word:['愿新春以后，吉吉利利，百事都如意',
      '一愿识尽天下好人，二愿读尽世间好书，三愿看尽世间好山水',
      '从教一岁大家添，但只要、明年强健',
      '一岁一礼，一寸欢喜',
      '如花似叶，岁岁年年，共占春风',
      '故岁今宵尽，新年明旦来',
      '一愿郎君千岁，二愿妾身常健，三愿如同梁上燕，岁岁长相见',
      '愿天上人间，占得欢娱，年年今夜',
      '天地风霜尽，乾坤气象和，历添新岁月，春满旧山河',
      '岁岁年年，共欢同乐，嘉庆与时新',
      '把酒君前欲问年，笑指松椿，当是同年，愿从今后八千年，长似今年，长似今年',
      '听烧爆竹童心在，看换桃符老兴偏，鼓角梅花添一部，五更欢笑拜新年',
      '邻墙旋打娱宾酒，稚子齐歌碎乐岁诗，老去又逢新岁月，春来更有好花枝'
    ]
    }
  }
}
</script>

<style scoped>
  #main{
    display: flex;
    flex-direction: column;
    width: 100vw;
    height: 100vh;
  }
  #back{
    width: 100%;
    height: 100%;
    z-index: -1;
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
  }
  .container{
    display: flex;
    flex-direction: row;
    width: 100vw;
    height: 100vh;
  }

  #Write{
    width: 300px;
    height: 100px;
    margin: auto;
    background-color: rgba(237, 110, 110, 0.454);
    font-family: "myFont";
    border-style: none;
    font-size: 50px;
  }
</style>