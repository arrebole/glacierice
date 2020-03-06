<template>
  <div class="home">
    <!-- <Header></Header> -->
    <Banner></Banner>
    <div style="display:flex;">
      <div class="side"></div>
      <div class="center">
        <Split text="Explore Data"></Split>
        <div class="explore-data">
          <Piece
            to="data-resource"
            title="Data Resource"
            content="Secondary line text Lorem ipsum dolor sit amet, consectetur adipiscing elit."
          ></Piece>
          <Piece
            to="open-data-api"
            title="Open Data API"
            content="This API primarily aims at feeding standardized data to data-science use cases."
          ></Piece>
          <Piece
            to="contributor"
            title="Contribute"
            content="Share your own work with us. Include but not limited to data, model, code and wikis."
          ></Piece>
        </div>

        <Split text="Papers and Wiki"></Split>

        <Bar to="crash-course" title="Crash Course" content="course for epidemic modeling">
          <template v-slot:cover>
            <img class="bar-cover" src="../assets/crash-course.png" />
          </template>
        </Bar>

        <Bar to="papers" title="Papers" content="Collection of recent research">
          <template v-slot:cover>
            <img class="bar-cover" src="../assets/paper.png" />
          </template>
        </Bar>

        <Split text="Write Your Own Models"></Split>
        <Bar to="models" title="Models of the Day" content="R, Jupyter Notebook, Python">
          <template v-slot:cover>
            <img class="bar-cover" src="../assets/math.png" />
          </template>
        </Bar>

        <Split text="Interactive Models"></Split>
        <InteractiveModels></InteractiveModels>

        <Split text="Visualizations"></Split>
        <Visualizations></Visualizations>
      </div >
      <div class="side" style="display:flex;flex-direction:row-reverse">
        <SideBar/>
      </div>
    </div>
    <Footer></Footer>
  </div>
</template>

<script>
// @ is an alias to /src
import Banner from "../components/Banner.vue";
import Split from "../components/Split.vue";
import Piece from "../components/Piece.vue";
import Footer from "../components/Footer.vue";
import Visualizations from "../components/Visualizations.vue";
import InteractiveModels from "../components/InteractiveModels.vue";
import Bar from "../components/Bar.vue";
import SideBar from "../components/Sidebar.vue";

export default {
  name: "Home",
  methods: {},
  components: {
    Banner,
    Split,
    Piece,
    Footer,
    Visualizations,
    InteractiveModels,
    Bar,
    SideBar
  }
};
</script>

<style scoped>
.home {
  min-height: 900px;
  background-color: #ededed;
  font-family: "Helvetica";
}
.explore-data {
  display: flex;
  max-width: 950px;
  justify-content: space-between;
}
.center {
  padding: 0px 10px;
  width: 100%;
  max-width: 950px;
  box-sizing: border-box;
}
.side{
  flex: 1;
  width: 250px;
}
.bar-cover {
  width: 90px;
  height: 80px;
  padding: 5px;
}
@media screen and (max-width: 720px) {
  .explore-data {
    flex-direction: column;
  }
}
</style>