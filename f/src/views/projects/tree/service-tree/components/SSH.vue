<template>
  <div>
    <div
      id="terminal"
      v-loading="loading"
      class="console"
      element-loading-text="拼命连接中"
    />
  </div>
</template>
<script>
import "xterm/css/xterm.css";
import { Terminal } from "xterm";
import { FitAddon } from "xterm-addon-fit";
import { AttachAddon } from "xterm-addon-attach";

export default {
  name: "SSH",
  props: ["machine", "instance_name"],
  data() {
    return {
      loading: true,
      term: null,
      socket: null,
      rows: 150,
      cmd: "",
    };
  },
  mounted() {
    this.initSocket();
  },
  beforeDestroy() {
    this.socket.close();
  },
  methods: {
    // Xterm主题
    initTerm() {
      const term = new Terminal({
        rendererType: "canvas", // 渲染类型
        rows: this.rows, // 行数
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
      this.loading = false;
      const attachAddon = new AttachAddon(this.socket);
      const fitAddon = new FitAddon();
      term.loadAddon(attachAddon);
      term.loadAddon(fitAddon);
      term.open(document.getElementById("terminal"));
      fitAddon.fit();
      term.focus();
      const _this = this;
      // 限制和后端交互，只有输入回车键才显示结果
      term.prompt = () => {
        term.write("\r\n");
      };
      term.prompt();
      function runFakeTerminal(_this) {
        if (term._initialized) {
          return;
        }
        // 初始化
        term._initialized = true;
        term.writeln(
          "$ Welcome to \x1B[1;3;31m our " +
            (_this.instance_name
              ? _this.instance_name + " instance"
              : "instance") +
            " \x1B[0m \r"
        );
        term.write("\r$ ");
        term.onData(function(key) {
          console.log(key.charCodeAt());
          // 上下左右方向键: 27
          if (key.charCodeAt() != 27) {
            if (key.charCodeAt() == 3) {
              term.write("\x1b[2K\r");
              term.write("$ ");
            } else {
              term.write(key);
              _this.cmd += key;
              if (key == "\r\n" || key == "\r") {
                const order = {
                  Data: key,
                  Op: "stdin",
                };

                if (_this.cmd.indexOf("clear") !== -1) {
                  term.clear();
                  _this.cmd = "";
                } else {
                  _this.onSend(order);
                  term.write("\r\n");
                }
              }
            }
          } else {
            term.write("\x1b[2K\r");
            term.write("$ ");
          }
        });
        _this.term = term;
      }
      runFakeTerminal(_this);
    },
    // webShell主题
    initSocket() {
      const WebSocketUrl =
        "ws://localhost:9092/assets/v1/ws/" + this.machine.id;
      this.socket = new WebSocket(WebSocketUrl);
      this.socketOnClose();
      this.socketOnOpen();
      this.socketOnError();
    },
    // webshell链接成功之后操作
    socketOnOpen() {
      this.socket.onopen = () => {
        // 链接成功后
        console.log("socket 连接成功");
        this.initTerm();
      };
    },
    // webshell关闭之后操作
    socketOnClose() {
      this.socket.onclose = () => {
        console.log("close socket");
      };
    },
    // webshell错误信息
    socketOnError() {
      this.socket.onerror = () => {
        console.log("socket 链接失败");
      };
    },
    // 特殊处理
    onSend(data) {
      data = JSON.stringify(data);
    },
    // 删除左右两端的空格
    trim(str) {
      return str.replace(/(^\s*)|(\s*$)/g, "");
    },
  },
};
</script>

<style lang="scss" scoped>
#terminal {
  width: 100%;
  height: 650px;
}
</style>
