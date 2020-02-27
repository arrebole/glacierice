<template>
  <div class="content-page">
    <Banner></Banner>
    <Split :text="$route.name"></Split>
    <article class="min-height10">
      <section class="conter bg-white markdown-body" v-if="openDataApi != ''" v-html="openDataApi"></section>

      <section class="conter bg-white markdown-body" v-if="contributor != ''">
        <table class="contributor-table">
          <thead>
            <tr>
              <th>Name</th>
              <th>GithubID</th>
            </tr>
          </thead>

          <tbody>
            <tr :key="item.id" v-for="item in contributor">
              <td>{{ item.name }}y</td>
              <td>{{ item.github_id }}</td>
            </tr>
          </tbody>
        </table>
      </section>

      <section class="conter bg-white markdown-body" v-if="resource3rd!= ''">
        <table class="contributor-table">
          <thead>
            <tr>
              <th>resource_nam</th>
              <th>description</th>
              <th>contributor_name</th>
            </tr>
          </thead>

          <tbody>
            <tr :key="item.id" v-for="item in resource3rd">
              <td><a :href="item.link">{{ item.resource_name }}</a></td>
              <td>{{ item.description || "/" }}</td>
              <td>{{ item.contributor_name }}</td>
            </tr>
          </tbody>
        </table>
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
  name: "Content",
  data() {
    return {
      contributor: "",
      openDataApi: "",
      resource3rd: ""
    };
  },
  created() {
    if (this.$route.name == "OpenDataAPI") {
      this.$api.openDataApi().then(res => {
        this.openDataApi = res.data;
      });
    }
    if (this.$route.name == "Contributor") {
      this.$api.contributor().then(res => {
        this.contributor = res.data.data;
      });
    }
    if (this.$route.name == "DataResource") {
      this.$api.dataResource().then(res => {
        this.resource3rd = res.data.data;
      });
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
.content-page {
  background-color: #ededed;
}
.min-height10 {
  min-height: 800px;
}
.min-height6 {
  min-height: 600px;
}
.bg-white {
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
.contributor-table {
  width: 100%;
  border: 1px solid rgb(167, 167, 167);
}
table {
  border-collapse: collapse;
  margin: 0 auto;
  text-align: center;
}
table td,
table th {
  border: 1px solid #cad9ea;
  color: #666;
  height: 30px;
}
table thead th {
  background-color: #cce8eb;
  width: 100px;
}
table tr:nth-child(odd) {
  background: #fff;
}
table tr:nth-child(even) {
  background: #f5fafa;
}
</style>