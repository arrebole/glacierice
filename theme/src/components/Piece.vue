<template>
  <div class="piece">
    <router-link :to="to">
      <h2 style="font-size:20px">{{ title }}</h2>
    </router-link>
    <p style="font-size:14px">{{ content }}</p>
  </div>
</template>


<script>
export default {
  props: ["title", "content", "to"]
};
</script>


<style lang="css" scoped>
.piece {
  background-color: #fff;
  box-sizing: border-box;
  height: 160px;
  border: 1px solid rgb(255, 255, 255);
  padding: 40px 15px 10px 15px;
  margin: 2px 0;
}
.piece > a {
  text-decoration: none;
  color: rgb(34, 33, 33);
}
.piece > a:hover {
  text-decoration: underline rgb(34, 22, 136);
  color: rgb(43, 76, 226);
}
.piece:hover {
  border: 1px solid rgb(13, 176, 197);
}
.piece > h2 {
  color: rgb(66, 66, 66);
}
.piece > p {
  color: rgb(65, 65, 65);
}

@media screen and (min-width: 720px) {
  .piece {
    max-width: 210px;
  }
}
@media screen and (max-width: 720px) {
  .piece {
    max-width: 100%;
  }
}
</style>