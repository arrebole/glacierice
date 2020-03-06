<template>
  <div class="sidebar">
    <div>Project - Nav</div>
    <div>
      <div style="font-size:14px; padding-bottom: 10px; color: #05ad9d;">Data</div>
      <div style="padding:10px;font-size:16px">
        <span class="shape"></span>
        <router-link :to="{name: 'DataResource'}">Data Resources</router-link>
      </div>
      <div style="padding:10px;font-size:16px">
        <span class="shape"></span>
        <router-link :to="{name: 'OpenDataAPI'}">Open Data API</router-link>
      </div>
    </div>
    <div style="font-size:16px">
      <div style="font-size:14px; padding-bottom: 10px;color: #05ad9d;">Developers</div>
      <div style="padding:10px">
        <span class="oval"></span>
        <a href="https://github.com/wuhan2020/Covid-19-data-science" target="_bank">Github Repo</a>
      </div>
      <div style="padding:10px">
        <span class="oval"></span>
        <router-link :to="{name: 'Contributor'}">Contribution List</router-link>
      </div>
    </div>
    <div style="font-size:14px; padding-bottom: 10px;color: #05ad9d;">Notes</div>
    <div style="font-size:14px">Chambray Carles Terry Gibson balls plaid wolf.</div>
  </div>
</template>


<script>
export default {};
</script>

<style scoped>
.sidebar {
  width: 250px;
  height: 575px;
  background-color: rgba(69, 90, 100, 0.79);
  color: rgba(255, 255, 255, 0.87);
  font-size: 20px;
  padding: 10px;
}
.sidebar > div {
  padding: 10px;
}
.sidebar > div > div > a {
  color: rgb(226, 226, 226);
  text-decoration: none;
}
.sidebar > div > div > a:hover {
  color: rgb(255, 255, 255);
}
.shape {
  display: inline-block;
  width: 14px;
  height: 14px;
  background-color: rgb(226, 226, 226);
  margin: 0px 8px 0 0;
}
.oval {
  display: inline-block;
  width: 18px;
  height: 18px;
  background-color: rgb(226, 226, 226);
  margin: 0px 8px 0 0;
  border-radius: 50%;
}
@media screen and (max-width: 1350px) {
  .sidebar {
    display: none;
  }
}
</style>