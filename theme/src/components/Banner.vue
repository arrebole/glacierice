<template>
    <div id="banner">
        <div class="banner-text">
            <div class="banner-text-line">
                <div>nCov</div><div>Open</div><div>Data</div><div>Project</div>
            </div>
            <p>Under Epedemic outbreak, We are as strong as we are united</p>
        </div>
    </div>    
</template>


<script>
export default {}
</script>


<style lang="css" scoped>
#banner{
    text-align: center;
    color: #fff;
    overflow: hidden;
    padding: 20px 0;
    border: rgba(0, 0, 0, 0.329) solid 1px;
    border-radius: 0.6%;
    box-shadow: 5px 6px 4px 1px rgba(0, 0, 0, 0.308);
    background-color: #009687;
}

.banner-text:hover{
    color: rgb(233, 233, 233);
}
.banner-text-line > div{
    display: inline;
    padding: 5px;
    font-size: 28px;
}
.banner-text >p{
    font-size: 13px;
}
</style>