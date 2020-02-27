<template>
  <div class="thanks-page">
    <Banner></Banner>

    <article>
      <section class="markdown-body" v-html="source">
      </section>
    </article>

    <Footer></Footer>
  </div>
</template>


<script>
import Banner from "../components/Banner.vue";
import Footer from "../components/Footer.vue";

export default {
  name: "Thanks",
  data(){
    return {
      source: "loading"
    }
  },
  created(){
    this.$api.thanks().then(res =>{this.source = res.data})
  },
  components: {
    Banner,
    Footer
  }
};
</script>

<style scoped>
.markdown-body {
  box-sizing: border-box;
  min-width: 200px;
  max-width: 980px;
  margin: 0 auto;
  padding: 45px;
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