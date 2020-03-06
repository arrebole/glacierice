<template>
    <div class="visualizations">
        <iframe class="visualizations-iframe" src="https://coronavirus.app/map?embed=true" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen>
        </iframe>
    </div>    
</template>

<style scoped>
.visualizations {
  height: 400px;
  width: 100%;
  background-color: rgb(253, 253, 253);
  margin: 10px 0px 30px 0px;
  padding: 10px;
  box-sizing: border-box;
}
.visualizations-iframe{
  width: 100%;
  height: 100%;
}
@media (max-width: 767px) {
  .visualizations {
    height: 300px;
  }
}
</style>
