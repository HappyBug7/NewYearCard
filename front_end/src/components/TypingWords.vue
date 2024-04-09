<template>
  <div class="videoWords">
    <div id="part2">
      <span id="textBox" style="user-select: none;"></span>
    </div>
  </div>
</template>

<script>
export default {
  name: "TypingWords",
  comments:{

  },
  data(){
    return{
      words:this.words
    }
  },
  props:['words'],
  methods:{
    typer(words){
      // let words=['There is no elevator to success,','you have to take the stairs.','-- Zig Ziglar','Let us cultivate our garden.','-- Voltaire'];
      let tTurn=0;
      let charater=0;
      let way=0;
      let wait=0;
      this.timer = setInterval(()=>{
        let target=words[tTurn%words.length];
        if(charater==target.length)
        {
          way=0;
        }
        if(wait==0)
        {
          if(charater<0)
          {
            way=1;
            tTurn++;
          }
          if(way==1)
          {
            charater++;
          }
          if(way==0)
          {
            charater--;
          }
          document.getElementById('textBox').innerHTML=target.slice(0,charater+1);
        }
        if(charater==target.length)
        {
          if(wait==0)
          {
            wait=10;
          }
          if(wait>0)
          {
            wait--;
          }
        }
      },100);
    }
  },
  mounted(){
    this.typer(this.$data.words)
  }
}
</script>

<style scoped>
  *{
    background-color: transparent;
    margin: 0;
    padding: 0;
    box-sizing: content-box;
  }
  .videoWords
  {
    width: 100vw;
    display: flex;
    flex-direction: column;
    align-items: center;
    color: white;
  }
  #textBox::after
  {
    font-size: 20px;
    display: block;
    content: " ";
    height: 70px;
    width: 2px;
    background-color: transparent;
    animation: change 0.7s steps(1) infinite;
  }
  #textBox
  {
    font-size: 50px;
    display: flex;
    flex-direction: row;
    font-family: "myFont";
  }
  @keyframes change 
  {
    0%,100%
    {
      background-color: transparent;
    }
    50%
    {
      background-color: white;
    }
  }
  @font-face {
  font-family: "myFont";
  src: url("../assets/Font.ttf");
  }
</style>