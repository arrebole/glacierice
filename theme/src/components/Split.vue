<template>
    <div class="split">
        <h2>{{ text }}</h2>    
    </div>    
</template>


<script>
export default {
    props:["text"]
}
</script>


<style lang="css" scoped>
.split > h2 {
    margin: 15px auto;
    max-width: 900px;
    text-indent: 10px;
    color: rgb(65, 65, 65);
}
</style>