<template>
    <div class="split">
        <h2>{{ text }}</h2>
        <div style="display:flex">
            <slot name="githublink"></slot>
        </div>    
    </div>    
</template>


<script>
export default {
    props:["text"]
}
</script>


<style lang="css" scoped>
.split{
    max-width: 950px;
    height: 45px;
    margin: auto;
    padding: 0px 10px;
    display: flex;
    justify-content: space-between;
    align-items: center;
}
.split > h2 {
    font-size: 24px;
    color: rgb(65, 65, 65);
}
</style>