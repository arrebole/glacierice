<template>
  <div class="interactive-models">
    <canvas id="models" class="canvas"></canvas>
  </div>
</template>

<script>
import Chart from "chart.js";

export default {
  data() {
    return {
      modelData: {
        type: "line",
        data: {
          labels: [],
          datasets: []
        },
        options: {}
      }
    };
  },
  methods: {
    canvasCtx() {
      return document.getElementById("models").getContext("2d");
    },
    createModelLabels() {
      const result = new Array(22);
      for (let i = 0; i < result.length; i++) {
        result[i] = i;
      }
      return result;
    },
    createModelDatasets() {
      const result = {
        label: "预测",
        backgroundColor: "rgba(255, 255, 255, 0)",
        borderColor: "rgba(100,157,205,1)",
        data: []
      };
      const f = x => 100 - (x - 10) ** 2;
      for (let i = 0; i <= 20; i+=1) {
        result.data.push({ x: i, y: f(i) });
      }
      return result;
    }
  },
  mounted() {
    this.modelData.data.labels = this.createModelLabels();
    this.modelData.data.datasets.push(this.createModelDatasets());
    new Chart(this.canvasCtx(), this.modelData);
  }
};
</script>

<style scoped>
.interactive-models {
  padding: 10px;
  height: 400px;
  background-color: #fff;
  margin: 20px;
}
.canvas {
  width: 100%;
  height: 100%;
}
@media (max-width: 767px) {
  .interactive-models {
    height: 300px;
  }
}
</style>