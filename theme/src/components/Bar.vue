<template>
    <div class="bar">
        <div class="img">
            <slot name="cover"></slot>
        </div>
        <div class="bar-text">
            <h4>{{ title }}</h4>
            <p>{{ content }}</p>
        </div>
        <img class="bar-log" src="../assets/github-logo.png"/>
    </div>    
</template>


<script>
export default {
    props:["title", "content"]
}
</script>


<style lang="css" scoped>
.bar{
    display: flex;
    height: 90px;
    margin: 10px 20px;
    font-size: 14px;
    background-color: #fff; 
    border: 1px solid rgb(255, 255, 255);
}
.bar:hover{
    cursor: pointer;
    border: 1px solid rgb(128, 50, 50);
}

.bar-text{
    flex: 6;
    padding: 20px 20px;
    
}
.bar-text > h4{
    margin: 5px;
    color: rgb(65, 65, 65);
}
.bar-text > p{
    margin: 5px;
    color: rgb(105, 105, 105);
}
.img{
    width: 100px;
    /* background-color: rgb(177, 177, 177); */
}
.bar-log{
    max-width: 40px;
    height: 50px;
    padding: 20px 30px;
}
</style>