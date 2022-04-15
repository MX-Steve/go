<template>
  <HelloWorld msg="Welcome to Your Vue.js App" />
</template>

<script>
import HelloWorld from "./components/HelloWorld.vue";

export default {
  name: "App",
  components: {
    HelloWorld,
  },
};
// document.oncontextmenu = function () {
//   //右键禁用
//   return false;
// };
// document.addEventListener("keydown", function (e) {
//   //按键禁用
//   try {
//     e = e || event || window.event;
//     var keycode = e.keyCode || e.which || e.charCode;
//     var ctrlKey = e.ctrlKey || e.metaKey;
//     if (keycode == 123) {
//       //F12
//       e.preventDefault();
//       window.event.returnValue = false;
//     } else if (keycode == 8) {
//       var act = document.activeElement.tagName.toLowerCase();
//       if (act.indexOf("body") != -1 || act.indexOf("html") != -1) {
//         window.history.back();
//         return false;
//       }
//       return true;
//     } else if (ctrlKey && keycode == 83) {
//       //ctrl+S
//       e.preventDefault();
//       window.event.returnValue = false;
//     } else if (ctrlKey && keycode == 87) {
//       //ctrl+W
//       e.preventDefault();
//       window.event.returnValue = false;
//     } else if (ctrlKey && keycode == 107) {
//       //ctrl++
//       e.preventDefault();
//       window.event.returnValue = false;
//     } else if (ctrlKey && keycode == 109) {
//       //ctrl+-
//       e.preventDefault();
//       window.event.returnValue = false;
//     } else if (keycode == 122) {
//       //F11
//       $("#switchFullScreenCloseBtn").trigger("click");
//     }
//   } catch (e) {
//     console.error(e);
//   }
// });

</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  position: fixed;
  top: 5px;
  left: 5px;
  right: 5px;
  bottom: 5px;
}
</style>
