<template>
  <div class="home">
    <Header></Header>
    <Banner></Banner>

    <Split text="Export Data"></Split>
    <div class="export-data">
      <Piece title="Data Resource" content="Secondary line text Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam massa quam."></Piece>
      <Piece title="Open Data API" content="Secondary line text Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam massa quam."></Piece>  
      <Piece title="Contribute" content="Share your own work with us on Github. Include but not limited to data, model, code and wikis."></Piece>    
    </div>
  
    <Split text="Papers and Wiki"></Split>
    <div class="center">
        <Bar title="Crash Course" content="course for epidemic modeling"></Bar>
        <Bar title="Papers" content="Collection of recent research"></Bar>
    </div>


    <Split text="Write Your Own Models"></Split>
    <div class="center">
      <Bar title="Models of the Day" content="R, Jupyter Notebook, Python"></Bar>
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
import Header from "../components/Header.vue";
import Banner from "../components/Banner.vue";
import Split from "../components/Split.vue";
import Piece from "../components/Piece.vue";
import Footer from "../components/Footer.vue";
import Visualizations from "../components/Visualizations.vue";
import InteractiveModels from "../components/InteractiveModels.vue";
import Bar from "../components/Bar.vue";

export default {
  name: "Home",
  components: {
    Banner,
    Split,
    Piece,
    Footer,
    Visualizations,
    InteractiveModels,
    Header,
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
.export-data{
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
@media screen and (max-width: 720px) {
   .export-data{
     flex-direction: column;
   } 
}
</style>