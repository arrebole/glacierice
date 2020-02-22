<template>
    <div class="content-page">
        <Banner></Banner>
        <Split :text="title"></Split>
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
    name:"Content",
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
.content-page{
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
</style>