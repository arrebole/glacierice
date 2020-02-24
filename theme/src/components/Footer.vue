<template>
    <footer>
        <div><router-link :to="{ path: 'Thanks' }">Thanks to EVERYONE</router-link></div>
        <div><a>ABOUT</a> <a>LICENCE</a> <a>YEAR</a></div>
    </footer> 
</template>

<style scoped>
footer{
    height: 100px;
    padding: 20px 0;
    font-size: 20px;
    text-align: center;
    background-color: #009687;
    color: #fff;
}
footer > div > a{
    text-decoration: none;
    color: #fff;
}

footer > div > a:hover{
    color: rgb(216, 216, 216);
}

footer > div{
    padding: 10px 0;
}
footer > div > a{
    padding: 0 20px;
}
</style>