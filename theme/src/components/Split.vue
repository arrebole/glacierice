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
    max-width: 900px;
    margin: auto;
    display: flex;
    justify-content: space-between;
}
.split > h2 {
    color: rgb(65, 65, 65);
}
</style>