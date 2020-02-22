<template>
    <div class="refererence-page">
        <Banner></Banner>
        <Split :text="title">
            <template v-slot:githublink>
                <img class="link-img" src="../assets/github-link.png"/>
                <a class="link-code" href="#">link to code</a>
            </template>
        </Split>
        <article class="min-height10">
            <section class="conter min-height6 bg-white max-with9"></section>
        </article>
        <Footer></Footer>
    </div>
</template>


<script>
import Banner from "../components/Banner.vue";
import Split from "../components/Split.vue";
import Footer from "../components/Footer.vue";

export default {
    name:"Refererence",
    data(){
        return {
            title:"Title"
        }
    },
    components:{
        Banner,
        Split,
        Footer,
    }
}
</script>

<style scoped>
.refererence-page{
    background-color: #ededed;
}
.min-height10{
    min-height: 800px;
}
.conter{
    margin: auto;
}
.min-height6{
    min-height: 600px;
}
.bg-white{
    background-color: #fff;
}
.max-with9{
    max-width: 900px;
}
.link-code{
    display: inline-block;
    text-decoration: none;
    font-size: 16px;
    margin: 25px 10px;
    color: rgb(92, 92, 92);
}
.link-code:hover{
    color: rgb(58, 119, 211);
}
.link-img{
    width: 40px;
    height: 40px;
    display: inline-block;
    margin: 10px 0px;
}
</style>