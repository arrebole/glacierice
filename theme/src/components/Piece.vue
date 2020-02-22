<template>
    <div class="piece">
        <h2>{{ title }}</h2>
        <p>{{ content }}</p>    
    </div>    
</template>


<script>
export default {
    props:["title","content"]
}
</script>


<style lang="css" scoped>
.piece{
    background-color: #fff;
    border: 1px solid rgb(255, 255, 255);
    padding: 60px 20px 10px 20px;
    margin: 5px;
}
.piece:hover{
    cursor: pointer;
    border: 1px solid rgb(13, 176, 197);
}
.piece > h2{
    color: rgb(66, 66, 66);
}
.piece > p{
    color: rgb(65, 65, 65);
}


@media screen and (max-width: 720px) {
   .piece{
     max-width: 100%;
   } 
}
</style>