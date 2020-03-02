<template>
  <div class="home">
    <!-- <Header></Header> -->
    <Banner></Banner>

    <Split text="Explore Data"></Split>
    <div class="explore-data">
      <Piece to="data-resource" title="Data Resource" content="Secondary line text Lorem ipsum dolor sit amet, consectetur adipiscing elit."></Piece>
      <Piece to="open-data-api" title="Open Data API" content="This API primarily aims at feeding standardized data to data-science use cases."></Piece>  
      <Piece to="contributor" title="Contribute" content="Share your own work with us. Include but not limited to data, model, code and wikis."></Piece>    
    </div>
  
    <Split text="Papers and Wiki"></Split>
    <div class="center">
        <Bar to="crash-course" title="Crash Course" content="course for epidemic modeling">
          <template v-slot:cover><img class="bar-cover" src="../assets/crash-course.png"/></template>
        </Bar>
        
        <Bar to="papers" title="Papers" content="Collection of recent research">
          <template v-slot:cover><img class="bar-cover" src="../assets/paper.png" /></template>
        </Bar>
    </div>


    <Split text="Write Your Own Models"></Split>
    <div class="center">
      <Bar to="models" title="Models of the Day" content="R, Jupyter Notebook, Python">
        <template v-slot:cover><img class="bar-cover" src="../assets/math.png"/></template>
      </Bar>
    </div>

    <Split text="Interactive Models"></Split>
    <div class="center">
      <InteractiveModels></InteractiveModels>
    </div>

    <Split text="Visualizations"></Split>
    <div class="center">
      <Visualizations></Visualizations>
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

export default {
  name: "Home",
  methods:{
  },
  components: {
    Banner,
    Split,
    Piece,
    Footer,
    Visualizations,
    InteractiveModels,
    Bar,
  }
};
</script>

<style scoped>
.home {
  min-height: 900px;
  background-color: #ededed;
  font-family: 'Helvetica';
}
.explore-data{
  padding: 0 20px;
  margin: auto;
  display: flex;
  max-width: 900px;
  justify-content:space-between;
}
.center{
  margin: auto;
  width: 100%;
  max-width: 930px;
}
.bar-cover{
  width: 90px;
  height: 80px;
  padding: 5px;
}
@media screen and (max-width: 720px) {
   .export-data{
     flex-direction: column;
   } 
}
</style>