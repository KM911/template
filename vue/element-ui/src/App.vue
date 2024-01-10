<script setup>
import axios from 'axios'

function Get(){
  // let url = "https://www.baidu.com"
  //  不要写绝对路径
  let url = "/hello"
  axios.get(url).then(res => {
    console.log(res)
  }).catch(err => {
    console.log(err)
  })
  // const xhr = new XMLHttpRequest();
  // const url = "/api";
  // axios.get(url).then(res=>{
  //   console.log(res)
  // })
    // fetch('localhost:3000/hello').then(res => {
    //   console.log(res)

    // }).catch(err => {
    //   console.log(err)
    // })
}


function HelloWorld() {
  return console.log('Hello World')
}
</script>

<template>
    <el-button @click="Get()">Default</el-button>

</template>

<style scoped>
body{
    display: flex;
}
</style>
