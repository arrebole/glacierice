<template>
    <div class="content-page">
        <Banner></Banner>
        <Split :text="$route.name"></Split>
        <article class="min-height10">
            <section class="conter bg-white markdown-body" v-html="source">

            </section>
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
            source: "Loading"
        }
    },
    created(){
        if (this.$route.name == "OpenDataAPI"){
            this.$api.openDataApi().then(res=>{this.source = res.data})
        }
        if (this.$route.name == "Contributor"){
            this.$api.contributor().then(res=>{this.source = res.data})
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
.min-height6{
    min-height: 600px;
}
.bg-white{
    background-color: #fff;
}
.markdown-body {
  box-sizing: border-box;
  min-width: 200px;
  max-width: 980px;
  margin: 0 auto;
  padding: 30px;
  font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Helvetica, Arial,
    sans-serif, Apple Color Emoji, Segoe UI Emoji;
  font-size: 16px;
  line-height: 1.5;
}

@media (max-width: 767px) {
  .markdown-body {
    padding: 15px;
  }
}
</style>