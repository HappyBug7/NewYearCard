<template>
  <div class="videoWords">
    <div id="part2">
      <span id="textBox" style="user-select: none;"></span>
    </div>
  </div>
</template>

<script>
export default {
  name: "TypingArea",
  comments:{

  },
  data(){
    return{
    }
  },
  methods:{
    typer(){
      // let words=['There is no elevator to success,','you have to take the stairs.','-- Zig Ziglar','Let us cultivate our garden.','-- Voltaire'];
      let words = ['你好2024!','你好2023?']
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
        if(tTurn%words.length == 0 && charater==target.length){
          window.clearInterval(this.timer);
          var sty=document.createElement('style');
          sty.innerText='#Start{display:block}';
          document.body.appendChild(sty);
        }
      },150);
    }
  },
  mounted(){
    this.typer()
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
    width: 100%;
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
    animation: change 0.7s steps(1);
    animation-iteration-count: 7.6
  }
  #textBox
  {
    font-size: 50px;
    display: flex;
    flex-direction: row;
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
</style>