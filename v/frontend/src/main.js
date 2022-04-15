import { createApp } from "vue";
import App from "./App.vue";
import "xterm/dist/xterm.css";
import ElementPlus from 'element-plus';
import 'element-plus/theme-chalk/index.css';

const app = createApp(App)
app.use(ElementPlus)
app.mount('#app')

// Vue.config.productionTip = false;
// /* eslint-disable no-new */
// new Vue({
//   el: "#app",
//   components: { App },
//   template: "<App/>",
// });
