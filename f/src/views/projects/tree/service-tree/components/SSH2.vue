<template>
  <div class="box">
    <div id="xterm" />
  </div>
</template>

<script>
import { Terminal } from "xterm";
import { FitAddon } from "xterm-addon-fit";
import { AttachAddon } from "xterm-addon-attach";
import "xterm/css/xterm.css";
import "xterm/lib/xterm.js";
export default {
  name: "SSH2",
  data() {
    return {
      term: "",
      socket: "",
    };
  },
  mounted() {
    const url = "ws://localhost:9092/assets/v1/ws?host_id=12";
    this.init(url);
  },
  methods: {
    initXterm() {
      this.term = new Terminal({
        rendererType: "canvas", // 渲染类型
        rows: 50, // 行数
        // cols: this.cols,// 设置之后会输入多行之后覆盖现象
        convertEol: true, // 启用时，光标将设置为下一行的开头
        scrollback: 100, // 终端中的回滚量
        fontSize: 14, // 字体大小
        disableStdin: false, // 是否应禁用输入。
        cursorStyle: "block", // 光标样式
        cursorBlink: true, // 光标闪烁
        scrollback: 50,
        tabStopWidth: 4,
        theme: {
          foreground: "yellow", // 字体
          background: "#060101", // 背景色
          cursor: "help", // 设置光标
        },
      });
      this.term.open(document.getElementById("xterm"));
      const fitAddon = new FitAddon();
      const attachAddon = new AttachAddon(this.socket);
      this.term.loadAddon(attachAddon);
      this.term.loadAddon(fitAddon);
      // 支持输入与粘贴方法
      const _this = this; // 一定要重新定义一个this，不然this指向会出问题
      this.term.onData(function(key) {
        console.log(key);
        const order = {
          Data: key,
          Op: "stdin",
        }; // 这里key值是你输入的值，数据格式一定要找后端要！！！！
        _this.socket.onsend(JSON.stringify(order)); // 转换为字符串
      });
    },
    init(url) {
      // 实例化socket
      this.socket = new WebSocket(url);
      // 监听socket连接
      this.socket.onopen = this.open;
      // 监听socket错误信息
      this.socket.onerror = this.error;
      // 监听socket消息
      this.socket.onmessage = this.getMessage;
      // 发送socket消息
      this.socket.onsend = this.send;
    },
    open: function() {
      console.log("socket连接成功");
      this.initXterm();
    },
    error: function() {
      console.log("连接错误");
    },
    close: function() {
      this.socket.close();
      console.log("socket已经关闭");
    },
    getMessage: function(msg) {
      this.term.write(msg["Data"]);
    },
    send: function(order) {
      this.socket.send(order);
    },
  },
};
</script>

<style lang="scss" scoped>
.box {
  width: 100%;
  height: 100%;
}
</style>
