<template>
    <div class="bar">
        <div class="img"></div>
        <div class="bar-text">
            <h4>{{ title }}</h4>
            <p>{{ content }}</p>
        </div>
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
    background-color: #fff;
    font-size: 14px;
    border: 1px solid rgb(255, 255, 255);
    margin: 10px 20px;
}
.bar:hover{
    cursor: pointer;
    border: 1px solid rgb(128, 50, 50);
}

.bar-text{
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
    background-color: rgb(177, 177, 177);
}
</style>