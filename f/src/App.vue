<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
export default {
  name: "App",
};
if (window.console && window.console.log) {
  console.log(
    "%c 欢迎来到 EcoPlants 运维平台",
    'font-family: "Helvetica Neue", Helvetica, Arial,sans-serif;font-size:24px;color:#49b165;-webkit-text-fill-color:#49b165;-webkit-text-stroke: 1px #49b165;'
  );
  console.log(
    "%c 遵纪守法，违规必究",
    'font-family: "Helvetica Neue", Helvetica, Arial,sans-serif;font-size:12px;color:#fc5531;-webkit-text-fill-color:#fc5531;-webkit-text-stroke: 1px #fc5531;'
  );
}
</script>
