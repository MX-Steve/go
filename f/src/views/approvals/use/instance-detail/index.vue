<template>
  <div class="instance-detail">
    <div v-if="!is_mobile">
      <div class="sub-title">审批实例</div>
      <el-card style="padding: 15px 30px">
        <el-table :data="instance_list" border style="width: 100%">
          <el-table-column prop="id" label="ID" width="55" />
          <el-table-column prop="instance_code" label="实例编码" width="300" />
          <el-table-column prop="status" label="当前状态" width="80" />
          <el-table-column prop="username" label="用户名" width="100" />
          <el-table-column prop="descriptions" label="描述信息" />
          <el-table-column prop="appendix" label="附件信息">
            <template v-if="scope.row.appendix !== null" slot-scope="scope">
              <div v-for="item in scope.row.appendix.split(',')" :key="item">
                <div v-if="item !== ''">
                  <span>{{ item }}</span>
                  <el-button
                    style="margin-left: 15px"
                    size="mini"
                    type="text"
                    icon="el-icon-download"
                    @click="download(item)"
                  />
                </div>
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="job_name" label="job名称" width="120" />
          <el-table-column prop="job_number" label="job编号" width="70" />
          <el-table-column prop="approval_name" label="所属工单" />
          <el-table-column label="操作" width="60">
            <template slot-scope="scope">
              <el-button
                v-if="
                  scope.row.status === '驳回' && scope.row.username === username
                "
                type="text"
                @click="restart()"
              >
                启动
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-card>
      <div class="sub-title">实例表单</div>
      <el-card style="padding: 15px 30px">
        <el-table :data="services" style="width: 100%">
          <el-table-column
            v-for="key in tableKeys"
            :key="key"
            :prop="key"
            :label="key"
            :width="key === '服务编号' ? 80 : ''"
          />
          <el-table-column label="操作" width="60">
            <template slot-scope="scope">
              <el-button
                v-if="trigger_btn_show === 1"
                type="text"
                @click="build_one(scope.row)"
                >触发</el-button
              >
            </template>
          </el-table-column>
        </el-table>
      </el-card>
      <div id="operator_btn" class="sub-title">实例流程</div>
      <el-card v-if="!old_version" style="padding: 15px 30px">
        <el-table
          :data="task_list"
          border
          style="width: 100%"
          :row-class-name="tableRowClassName"
        >
          <el-table-column prop="name" label="节点名称" width="150" />
          <el-table-column prop="type" label="类型" width="80" />
          <el-table-column prop="status" label="状态" width="80" />
          <el-table-column prop="start_time" label="开始时间" />
          <el-table-column prop="end_time" label="结束时间" />
          <el-table-column prop="comment" label="评论信息" />
          <el-table-column prop="username" label="操作人员" width="150">
            <template slot-scope="scope">
              {{
                scope.row.name === "开始"
                  ? sp_person
                  : scope.row.name === "结束"
                  ? sp_person
                  : scope.row.name === "经办人"
                  ? "系统"
                  : scope.row.username
              }}
              <el-button
                v-if="
                  scope.row.start_time !== '' &&
                  scope.row.end_time === '' &&
                  scope.row.username === username &&
                  is_refuted === false
                "
                type="primary"
                size="small"
                @click="approval_btn(scope.row.id)"
              >
                审批
              </el-button>
            </template>
          </el-table-column>
        </el-table>
        <el-dialog title="审批操作" :visible.sync="dialogVisible" width="30%">
          <el-form ref="approval" :model="approval" label-width="120px">
            <el-form-item label="审批">
              <el-select
                v-model="approval.status"
                placeholder="请选择"
                style="width: 100%"
              >
                <el-option label="同意" value="DONE" />
                <el-option label="拒绝" value="REJECTED" />
                <el-option label="转办" value="TRANSFERRED" />
                <el-option label="驳回" value="REFUTED" />
              </el-select>
            </el-form-item>
            <el-form-item
              v-if="approval.status === 'TRANSFERRED'"
              label="操作人"
            >
              <el-select
                v-model="approval.user_id"
                filterable
                clearable
                placeholder="请选择"
                style="width: 100%"
              >
                <el-option
                  v-for="user in user_list"
                  :key="user.id"
                  :label="user.username"
                  :value="user.id"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="评论">
              <el-input v-model="approval.descriptions" type="textarea" />
            </el-form-item>
            <el-form-item>
              <el-button @click="dialog_cancel">取 消</el-button>
              <el-button type="primary" @click="dialog_confirm">
                确 定
              </el-button>
            </el-form-item>
          </el-form>
        </el-dialog>
      </el-card>
      <el-card
        v-else
        style="padding: 15px 30px; text-align: center; color: #999"
      >
        老版本不展示审批流程
      </el-card>
      <div class="sub-title">时间线</div>
      <el-card style="padding: 15px 30px">
        <el-row>
          <el-col v-if="!old_version" :span="8" :offset="8">
            <el-timeline>
              <el-timeline-item
                v-for="(activity, index) in activities"
                :key="index"
                style="position: relative"
                type="primary"
                :color="activity.status === '完成' ? '#0bbd87' : 'red'"
                size="large"
                :timestamp="activity.start_time"
              >
                <span>
                  {{ "[" + activity.name + " 节点] " + activity.comment }}
                </span>
                <span style="float: right">
                  {{
                    activity.name === "开始"
                      ? sp_person
                      : activity.name === "结束"
                      ? sp_person
                      : activity.name === "经办人"
                      ? "系统"
                      : activity.username
                  }}
                </span>
              </el-timeline-item>
            </el-timeline>
          </el-col>
          <el-col
            v-else
            :span="8"
            :offset="8"
            style="text-align: center; color: #999"
          >
            老版本不展示审批流程
          </el-col>
        </el-row>
      </el-card>
    </div>
    <div v-if="is_mobile">
      <div class="sub-title">审批实例</div>
      <el-card>
        <el-table :data="instance_list" border style="width: 100%">
          <el-table-column prop="id" label="ID" width="45" fixed />
          <el-table-column prop="instance_code" label="实例编码" width="300" />
          <el-table-column prop="status" label="当前状态" />
          <el-table-column prop="username" label="用户名" />
          <el-table-column prop="descriptions" label="描述信息" />
          <el-table-column prop="appendix" label="附件信息" width="200">
            <template v-if="scope.row.appendix !== null" slot-scope="scope">
              <div v-for="item in scope.row.appendix.split(',')" :key="item">
                <div v-if="item !== ''">
                  <span>{{ item }}</span>
                  <el-button
                    style="margin-left: 15px"
                    size="mini"
                    type="text"
                    icon="el-icon-download"
                    @click="download(item)"
                  />
                </div>
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="job_number" label="job 序列号" width="90" />
          <el-table-column prop="approval_name" label="所属工单" width="120" />
          <el-table-column label="操作" width="60">
            <template slot-scope="scope">
              <el-button
                v-if="
                  scope.row.status === '驳回' && scope.row.username === username
                "
                type="text"
                @click="restart()"
              >
                启动
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-card>
      <div class="sub-title">实例表单</div>
      <el-card>
        <el-table :data="services" style="width: 100%">
          <el-table-column
            v-for="key in tableKeys"
            :key="key"
            :prop="key"
            :label="key"
            width="150"
          />
          <el-table-column label="操作" width="60">
            <template slot-scope="scope">
              <el-button
                v-if="trigger_btn_show === 1"
                type="text"
                @click="build_one(scope.row)"
                >触发</el-button
              >
            </template>
          </el-table-column>
        </el-table>
      </el-card>
      <div id="operator_btn" class="sub-title">实例流程</div>
      <el-card v-if="!old_version">
        <el-table
          :data="task_list"
          border
          style="width: 100%"
          :row-class-name="tableRowClassName"
        >
          <el-table-column prop="name" label="节点名称" width="120" />
          <el-table-column prop="type" label="类型" />
          <el-table-column prop="status" label="状态" />
          <el-table-column prop="start_time" label="开始时间" width="160" />
          <el-table-column prop="end_time" label="结束时间" width="160" />
          <el-table-column prop="comment" label="评论信息" width="140" />
          <el-table-column prop="username" label="操作人员" width="160">
            <template slot-scope="scope">
              {{
                scope.row.name === "开始"
                  ? sp_person
                  : scope.row.name === "结束"
                  ? sp_person
                  : scope.row.name === "经办人"
                  ? "系统"
                  : scope.row.username
              }}
              <el-button
                v-if="
                  scope.row.start_time !== '' &&
                  scope.row.end_time === '' &&
                  scope.row.username == username &&
                  is_refuted == false
                "
                type="primary"
                size="small"
                @click="approval_btn(scope.row.id)"
              >
                审批
              </el-button>
            </template>
          </el-table-column>
        </el-table>
        <el-dialog title="审批操作" :visible.sync="dialogVisible" width="100%">
          <el-form ref="approval" :model="approval" label-width="80px">
            <el-form-item label="审批">
              <el-select
                v-model="approval.status"
                placeholder="请选择"
                style="width: 100%"
              >
                <el-option label="同意" value="DONE" />
                <el-option label="拒绝" value="REJECTED" />
                <el-option label="转办" value="TRANSFERRED" />
                <el-option label="驳回" value="REFUTED" />
              </el-select>
            </el-form-item>
            <el-form-item
              v-if="approval.status === 'TRANSFERRED'"
              label="操作人"
            >
              <el-select
                v-model="approval.user_id"
                filterable
                clearable
                placeholder="请选择"
                style="width: 100%"
              >
                <el-option
                  v-for="user in user_list"
                  :key="user.id"
                  :label="user.username"
                  :value="user.id"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="评论">
              <el-input v-model="approval.descriptions" type="textarea" />
            </el-form-item>
            <el-form-item>
              <el-button @click="dialog_cancel">取 消</el-button>
              <el-button type="primary" @click="dialog_confirm">
                确 定
              </el-button>
            </el-form-item>
          </el-form>
        </el-dialog>
      </el-card>
      <el-card v-else style="text-align: center; color: #999">
        老版本不展示审批流程
      </el-card>
      <div class="sub-title">时间线</div>
      <el-card>
        <el-row>
          <el-col v-if="!old_version" :span="24">
            <el-timeline>
              <el-timeline-item
                v-for="(activity, index) in activities"
                :key="index"
                style="position: relative"
                type="primary"
                :color="activity.status === '完成' ? '#0bbd87' : 'red'"
                size="large"
                :timestamp="activity.start_time"
              >
                <span>
                  {{ "[" + activity.name + " 节点] " + activity.comment }}
                </span>
                <span style="float: right">
                  {{
                    activity.name === "开始"
                      ? sp_person
                      : activity.name === "结束"
                      ? sp_person
                      : activity.name === "经办人"
                      ? "系统"
                      : activity.username
                  }}
                </span>
              </el-timeline-item>
            </el-timeline>
          </el-col>
          <el-col v-else :span="24" style="text-align: center; color: #999">
            老版本不展示审批流程
          </el-col>
        </el-row>
      </el-card>
    </div>
  </div>
</template>
<script>
import { getToken } from "@/utils/auth";
import jwtDecode from "jwt-decode";
import { usernames_get, user_get } from "@/api/user";
import axios from "axios";
import {
  instance_get,
  instance_cud,
  fiform_get,
  tasks_get,
  task_cud,
} from "@/api/approval/instance";
import { btn_check } from "@/api/btn";
import { deploy_jenkins } from "@/api/approval/approval";
import config from "@/utils/config";
import { ApprovalDetailsPage } from "@/utils/auth";
function get_date_time(dt) {
  var date = new Date(dt);
  var year = date.getFullYear();
  var month = date.getMonth() + 1;
  month = month < 10 ? "0" + month : month;
  var day = date.getDate();
  day = day < 10 ? "0" + day : day;
  var hour = date.getHours();
  hour = hour < 10 ? "0" + hour : hour;
  var minute = date.getMinutes();
  minute = minute < 10 ? "0" + minute : minute;
  var second = date.getSeconds();
  second = second < 10 ? "0" + second : second;
  return (
    year + "-" + month + "-" + day + " " + hour + ":" + minute + ":" + second
  );
}
export default {
  name: "InstanceDetail",
  data() {
    return {
      sp_person: "",
      instance_id: "",
      instance_list: [],
      forms: [],
      task_list: [],
      activities: [],
      dialogVisible: false,
      approval: {
        id: "",
        status: "",
        descriptions: "",
        user_id: "",
      },
      username: "",
      user_list: [],
      is_mobile: config.isMobile,
      services: [],
      tableKeys: [],
      is_refuted: false,
      old_version: 1,
      trigger_btn_show: 0,
      ApprovalDetailsPage: ApprovalDetailsPage,
    };
  },
  mounted() {
    this.init_data();
    this.showDetails("operator_btn");
  },
  methods: {
    build_one(row) {
      row["instance_id"] = this.instance_id;
      deploy_jenkins(row).then((response) => {
        if (response.code === 200) {
          this.$message({
            message: response.msg,
            type: "success",
          });
        }
      });
    },
    response_refresh_func(response) {
      if (response.code == 200) {
        this.dialogVisible = false;
        this.$message({
          message: response.msg,
          type: "success",
        });
        setTimeout(() => {
          location.reload(0);
        }, 500);
      } else {
        this.$message({
          message: response.msg,
          type: "warning",
        });
      }
    },
    showDetails(id) {
      var el = document.getElementById(id);
      this.$nextTick(function () {
        window.scrollTo({ behavior: "smooth", top: el && el.offsetTop });
      });
    },
    async init_data() {
      btn_check(this.ApprovalDetailsPage).then((response) => {
        if (response.code == 200) {
          this.trigger_btn_show = response.data.conn;
        }
      });
      const userObj = jwtDecode(getToken());
      this.username = userObj.username;
      const obj2 = await user_get({ name: this.username });
      if (obj2.code === 200) {
        var username = obj2.data.users[0]["zh_name"].replace(/\s*/g, "");
        if (username.length > 0) {
          this.username = username;
        }
      }
      const obj = await usernames_get();
      this.user_list = obj.data.users;
      this.instance_id = this.$route.query.instance_id;
      if (this.instance_id === undefined) {
        this.$message({
          message: "必须提供 instance_id 参数",
          type: "error",
        });
      } else {
        const result = await instance_get({ id: this.instance_id });
        if (result.code === 200) {
          this.instance_list = result.data.instances;
          this.sp_person = this.instance_list[0]["username"];
          this.old_version = this.instance_list[0]["old_version"];
        }
        const result2 = await fiform_get({ instance_id: this.instance_id });
        if (result2.code === 200) {
          console.log(result2);
          var services = result2.data;
          for(var index in services){
            var newDate = services[index]["期望发布时间"]
            services[index]["期望发布时间"] = get_date_time(newDate)
          }
          this.services = services;
          this.tableKeys = Object.keys(this.services[0]);
        }
        const result3 = await tasks_get({ instance_id: this.instance_id });
        if (result3.code === 200) {
          this.task_list = result3.data;
          for (var i = 0; i < this.task_list.length; i++) {
            const item = this.task_list[i];
            if (item.start_time !== "") {
              this.activities.push(item);
            }
          }
        }
        console.log(this.sp_person);
      }
    },
    restart() {
      const data = {
        id: this.instance_id,
        status: "PENDING",
        descriptions: "重新启动工单",
      };
      instance_cud("put", data).then((response) => {
        if (response.code === 200) {
          this.$message({
            message: response.msg,
            type: "success",
          });
          location.reload(0);
        }
      });
    },
    tableRowClassName({ row, rowIndex }) {
      if (row.status === "拒绝") {
        return "error-row";
      } else {
        if (row.end_time === "") {
          if (row.start_time !== "") {
            return "warning-row";
          } else {
            return "";
          }
        } else {
          return "success-row";
        }
      }
    },
    approval_btn(id) {
      this.dialogVisible = true;
      this.approval["id"] = id;
    },
    dialog_cancel() {
      this.approval = {
        id: "",
        status: "",
        descriptions: "",
      };
      this.dialogVisible = false;
    },
    dialog_confirm() {
      if (this.approval.status !== "TRANSFERRED") {
        this.approval.user_id = "";
      }
      task_cud("post", this.approval).then((response) => {
        this.response_refresh_func(response);
      });
    },
    download(file_name) {
      const url = process.env.VUE_APP_BASE_API + "/approval/v2/download-file";
      let data = {
        file_name: file_name,
      };
      axios({
        method: "post",
        url: url,
        data: data,
        responseType: "blob", // 指定获取数据的类型为blob
      })
        .then(function (response) {
          data = response.data; // 创建a标签并点击， 即触发下载
          let disposition_data = response.headers["content-disposition"]; // 注意这里headers中的参数名是小写的
          disposition_data = disposition_data.split(";");
          let download_file_name = "Untitled"; // 从 content-dispostition 中解析出文件名
          for (let i = 0; i < disposition_data.length; i++) {
            const k_v = disposition_data[i].split("=");
            if (k_v.length == 2) {
              const k = k_v[0].trim();
              const v = k_v[1].trim();
              if (k == "filename") {
                download_file_name = v;
              }
            }
          }
          const url = window.URL.createObjectURL(new Blob([data]));
          const link = document.createElement("a");
          link.style.display = "none";
          link.href = url;
          link.setAttribute("download", download_file_name); // 下载到本地的文件名

          document.body.appendChild(link);
          link.click();
        })
        .catch(function (err) {
          console.log(err);
        });
    },
  },
};
</script>
<style lang="scss">
.instance-detail {
  min-height: 450px;
  width: 98%;
  margin: 15px auto;
  background-color: #fff;
  padding: 15px;
  .sub-title {
    padding-top: 20px;
    text-align: center;
    font-size: 24px;
    font-weight: bold;
    padding-bottom: 10px;
    border-bottom: 2px solid #ccc;
  }
  .warning-row {
    background: oldlace;
  }

  .success-row {
    background: #f0f9eb;
  }
  .error-row {
    background: rgb(218, 183, 183);
  }
}
</style>
