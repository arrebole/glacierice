<template>
    <div class="bar">
        <div class="img">
            <slot name="cover"></slot>
        </div>
        <div class="bar-text">
            <router-link :to="to">
                <h4>{{ title }}</h4>
            </router-link>
            <p>{{ content }}</p>
        </div>
        <img class="bar-log" src="../assets/github-logo.png"/>
    </div>    
</template>


<script>
export default {
    props:["title", "content", "to"]
}
</script>


<style lang="css" scoped>
.bar{
    display: flex;
    align-items: center;
    height: 90px;
    margin: 10px 20px;
    font-size: 16px;
    background-color: #fff; 
    border: 1px solid rgb(255, 255, 255);
}
.bar:hover{
    border: 1px solid rgb(128, 50, 50);
}

.bar-text{
    flex: 6;
    padding: 20px 20px;
    
}
.bar-text > a > h4{
    margin: 5px;
    color: rgb(65, 65, 65);
}
.bar-text > a{
    text-decoration: none;
}
.bar-text > a:hover{
    text-decoration: underline rgb(34, 22, 136);
}

.bar-text > p{
    margin: 5px;
    font-size: 14px;
    color: rgb(105, 105, 105);
}
.img{
    width: 100px;
    /* background-color: rgb(177, 177, 177); */
}
.bar-log{
    width: 32px;
    height: 40px;
    margin: 0 15px;
}
</style>