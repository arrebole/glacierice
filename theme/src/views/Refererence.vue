<template>
  <div class="refererence-page">
    <Banner></Banner>
    <Split :text="title">
      <template v-slot:githublink>
        <img class="link-img" src="../assets/github-link.png" />
        <a class="link-code" :href="remoteData.link">link to code</a>
      </template>
    </Split>

    <article class="min-height10">
      <section
        class="conter min-height6 bg-white max-with9 markdown-body"
        v-if="remoteData.html != ''"
        v-html="remoteData.html"
      ></section>

      <section
        class="conter min-height6 bg-white max-with9 markdown-body"
        v-if="modelsOfDay">
        <ul>
        <template v-for="item in modelsOfDay">
            <li :key="item.title"><a :href="matchLink(item)" target="_blank">{{ item.title }}</a></li>
        </template>
        </ul>
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
  name: "Refererence",
  data() {
    return {
      title: "Title",
      modelsOfDay: null,
      remoteData: {
        link: "",
        html: ""
      }
    };
  },
  created() {
    if (this.$route.name == "CrashCourse") {
      this.title = "Crash Course";
      this.$api.crashCourse().then(res => (this.remoteData = res.data));
    }
    if (this.$route.name == "Papers") {
      this.title = "Papers";
      this.$api.papers().then(res => (this.remoteData = res.data));
    }
    if (this.$route.name == "Models") {
      this.title = "Models of Day";
      this.$api.modelsOfDay().then(res => (this.modelsOfDay = res.data.data));
    }
  },
  methods:{
    matchLink(item){
      return `https://github.com/wuhan2020/Covid-19-data-science/${item.link_github}/${item.link_readme}`
    }
  },
  components: {
    Banner,
    Split,
    Footer
  }
};
</script>

<style scoped>
.refererence-page {
  background-color: #ededed;
}
.min-height10 {
  min-height: 800px;
}
.conter {
  margin: auto;
}
.min-height6 {
  min-height: 600px;
}
.bg-white {
  background-color: #fff;
}
.max-with9 {
  max-width: 900px;
}
.link-code {
  display: inline-block;
  text-decoration: none;
  font-size: 16px;
  margin: 25px 10px;
  color: rgb(92, 92, 92);
}
.link-code:hover {
  color: rgb(58, 119, 211);
}
.link-img {
  width: 40px;
  height: 40px;
  display: inline-block;
  margin: 10px 0px;
}

.markdown-body {
  box-sizing: border-box;
  min-width: 200px;
  max-width: 890px;
  margin: 10px auto;
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