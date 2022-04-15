<template>
  <div>
    <div v-if="id == ''">
      <h2>请输入连接信息</h2>
      <el-form
        ref="form"
        :rules="form_role"
        style="width: 75%"
        :model="form"
        label-width="120px"
      >
        <el-form-item label="ip地址" prop="ip_address">
          <el-input v-model="form.ip_address" />
        </el-form-item>
        <el-form-item label="连接端口" prop="port">
          <el-input v-model="form.port" />
        </el-form-item>
        <el-form-item label="用户名" prop="username">
          <el-input v-model="form.username" />
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input type="password" v-model="form.password" show-password />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="p_host()"> 连接 </el-button>
          <el-button @click="f_cancel"> 取消 </el-button>
        </el-form-item>
      </el-form>
    </div>
    <div v-else class="console" id="terminal"></div>
  </div>
</template>
<script>
import axios from "axios";
import { Terminal } from "xterm";
import * as attach from "xterm/lib/addons/attach/attach";
import * as fit from "xterm/lib/addons/fit/fit";
export default {
  name: "webssh",
  data() {
    return {
      term: null,
      terminalSocket: null,
      order: "",
      id: "",
      form_role: {
        ip_address: [
          { required: true, message: "请输入IP地址", trigger: "blur" },
        ],
        port: [{ required: true, message: "请输入连接端口", trigger: "blur" }],
        username: [
          { required: true, message: "请输入用户名", trigger: "blur" },
        ],
        password: [{ required: true, message: "请输入密码", trigger: "blur" }],
      },
      form: {
        ip_address: "",
        username: "",
        port: "",
        password: "",
      },
    };
  },
  methods: {
    f_cancel() {
      this.form = {
        ip_address: "",
        username: "",
        port: "",
        password: "",
      };
    },
    p_host() {
      console.log(this.form);
      let host = window.location.host.split(":")[0]
      axios
        .post(`http://${host}/api/chat/p-host`, this.form)
        .then((res) => {
          console.log(res);
          if (res.data.code === 200) {
            this.id = res.data.data.id;
            const url =
              window.location.protocol +
              "://" +
              window.location.host +
              "?id=" +
              this.id;
            window.open(url);
          }
        });
    },
  },
  mounted() {
    let id = location.href.split("=")[1];
    if (id == undefined) {
      console.log("?id=xxx 必须携带 id 参数");
    } else {
      this.id = id;
      //实例化一个websocket，用于和django江湖
      let host = window.location.host.split(":")[0]
      this.terminalSocket = new WebSocket(`ws://${host}/api/web/${id}/`);
      //获取到后端传回的信息
      this.terminalSocket.onmessage = (res) => {
        console.log(res);
        const reader = new window.FileReader();
        reader.onload = () => this.term.write(reader.result);
        reader.readAsText(res.data, "utf-8");
        //重置要发送的信息
        this.order = "";
      };
      let terminalContainer = document.getElementById("terminal");
      const width = window.innerWidth;
      const height = window.innerHeight;
      console.log("width: ", width);
      console.log("height: ", height);
      //创建xterm实例
      this.term = new Terminal({
        rendererType: "canvas",
        convertEol: true,
        scrollback: 100,
        fontSize: 14,
        disableStdin: false,
        cursorStyle: "block",
        cursorBlink: true, // 显示光标
        tabStopWidth: 4,
        cols: parseInt(width / 8, 10),
        rows: parseInt(height / 17, 10),
      });

      this.term.open(terminalContainer);

      //在xterm上显示命令行提示
      this.term.write("$ ");
      this.term.on("resize", (size) => {
        this.terminalSocket.send("resize", [size.cols, size.rows]);
      });
      //监听xterm的键盘事件
      this.term.on("key", (key, ev) => {
        // key是输入的字符 ev是键盘按键事件
        this.term.write(key);
        if (ev.keyCode == 13) {
          //使用webscoket将数据发送到django
          const data = this.order;
          this.terminalSocket.send(JSON.stringify({ data }));
          // this.order=''
          console.log("里面的order", this.order);
        } else if (ev.keyCode == 8) {
          //删除按钮
          //截取字符串[0,lenth-1]
          this.order = this.order.substr(0, this.order.length - 1);

          //清空当前一条的命令
          this.term.write("\x1b[2K\r");
          //简化当前的新的命令显示上
          this.term.write("$ " + this.order);

          console.log("截取的字符串" + this.order);
          typeof this.order;
        } else {
          // 将每次输入的字符拼凑起来
          this.order += key;
          console.log("外面的order", this.order);
        }
      });
    }
  },
};
</script>