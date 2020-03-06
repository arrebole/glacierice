<template>
    <footer>
        <div><router-link :to="{ path: 'thanks' }" style="font-size:14px">Thanks to EVERYONE</router-link></div>
        <div style="font-size:14px"><a>ABOUT</a> <a>LICENCE</a> <a>YEAR</a></div>
    </footer> 
</template>

<style scoped>
footer{
    height: 55px;
    overflow: hidden;
    text-align: center;
    background-color: #009687;
    color: #fff;
}
footer > div > a{
    text-decoration: none;
    color: #fff;
}
footer > div{
    margin: 6px 6px;
}
footer > div > a:hover{
    color: rgb(216, 216, 216);
}
footer > div > a{
    margin: 0px 8%;
}
</style>