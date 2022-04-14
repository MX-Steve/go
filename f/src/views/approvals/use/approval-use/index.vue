<template>
  <div class="approval-use">
    <div v-if="!is_mobile">
      <el-tabs v-model="activeName">
        <el-tab-pane label="服务发布" name="first">
          <div style="width: 650px; margin: 0 auto">
            <el-form
              :model="approval"
              :rules="approval_rule"
              label-width="120px"
            >
              <el-form-item label="审批工单" prop="approval_id">
                <el-select
                  v-model="approval.approval_id"
                  placeholder="请选择工单"
                  style="width: 80%"
                  @change="approval_change"
                >
                  <el-option
                    v-for="item in approval_list"
                    :key="item.id"
                    :label="item.approval_name"
                    :value="item.id"
                  />
                </el-select>
                <el-button
                  v-if="add_btn_show === 1"
                  style="margin-left: 15px"
                  size="small"
                  type="text"
                  @click="approval_yes"
                >
                  {{ instance_id ? "添加服务" : "创建" }}
                </el-button>
              </el-form-item>
            </el-form>
            <el-table
              v-if="services.length > 0"
              :data="services"
              style="width: 100%"
            >
              <el-table-column
                v-for="key in tableKeys"
                :key="key"
                :prop="key"
                :label="key"
                width="150"
              />
              <el-table-column label="操作" fixed="right" width="100">
                <template slot-scope="scope">
                  <el-button type="text" @click="addService()">
                    添加
                  </el-button>
                  <el-button v-if="del_btn_show" type="text" @click="delService(scope.row)">
                    移除
                  </el-button>
                </template>
              </el-table-column>
            </el-table>
            <el-row v-show="!un_finish">
              <el-col :span="16" :offset="2">
                <el-row>
                  <el-col :span="8" style="margin-top: 5px"> 附件信息 </el-col>
                  <el-col :span="16">
                    <el-upload
                      accept=".xlsx, .xls"
                      :file-list="fileList"
                      :action="actionUrl"
                      :data="info"
                      :before-upload="beforeUploadFile"
                      :on-exceed="exceedFile"
                      :on-success="uploadSuccess"
                      :show-file-list="isShowFile"
                      multiple
                      :limit="limitNum"
                    >
                      <el-button
                        v-if="detailIndex === 0"
                        type="text"
                        icon="el-icon-upload2"
                        size="mini"
                      >
                        添加文档
                      </el-button>
                    </el-upload>
                  </el-col>
                </el-row>
                <el-row>
                  <el-col
                    v-for="item in files"
                    :key="item"
                    :span="20"
                    :offset="4"
                  >
                    <div v-if="item !== ''">
                      <span> {{ item }} </span>
                      <el-button
                        size="mini"
                        type="text"
                        style="margin-left: 45px"
                        @click="download(item)"
                      >
                        下载
                      </el-button>
                      <el-button
                        size="mini"
                        type="text"
                        style="margin-left: 15px"
                        @click="del_btn(item)"
                      >
                        删除
                      </el-button>
                    </div>
                  </el-col>
                </el-row>
              </el-col>
            </el-row>
            <el-row v-show="!un_finish">
              <el-col :span="16" :offset="2">
                <h4>审批流程</h4>
                <div>
                  <div v-show="un_finish" class="info">
                    必填信息填写完整后，将显示审批流程
                  </div>
                  <el-row v-show="!un_finish">
                    <el-col :span="8" :offset="8">
                      <div
                        v-if="task_list.length > 0"
                        :style="{ height: task_list.length * 60 + 'px' }"
                      >
                        <el-steps
                          direction="vertical"
                          :active="1"
                          style="cursor: pointer"
                        >
                          <el-step
                            v-for="item in task_list"
                            :key="item.id"
                            :title="item.name"
                            @click.native="step_click(item)"
                          >
                            <template slot="description">
                              <div>
                                <span v-if="item.username !== ''">
                                  {{
                                    item.name === "开始"
                                      ? sp_person
                                      : item.name === "结束"
                                      ? sp_person
                                      : item.username
                                  }}
                                </span>
                                <span v-else style="color: red">
                                  [审批人自选]
                                </span>
                              </div>
                            </template>
                          </el-step>
                        </el-steps>
                      </div>
                    </el-col>
                  </el-row>
                </div>
              </el-col>
            </el-row>
            <el-row v-show="!un_finish" style="margin-top: 15px">
              <el-col :span="12" :offset="2">
                <el-button type="text" @click="confirm()">确认</el-button>
                <el-button
                  v-if="is_confirm"
                  type="text"
                  @click="go_to_detail()"
                >
                  详情
                </el-button>
              </el-col>
            </el-row>
          </div>
          <el-dialog
            title="添加服务"
            :visible.sync="dialogVisible"
            width="45%"
            :before-close="handleClose"
          >
            <el-row
              v-for="item in form_list"
              :key="item.id"
              :label="item.name"
              style="margin-bottom: 5px"
            >
              <el-col
                :span="4"
                style="
                  text-align: right;
                  padding-right: 15px;
                  font-weight: bold;
                  line-height: 30px;
                "
              >
                {{ item.name }}
              </el-col>
              <el-col :span="20">
                <el-input
                  v-if="item.type === '单行文本'"
                  v-model="approval[item.name]"
                  placeholder="请输入"
                />
                <el-input
                  v-else-if="item.type === '多行文本'"
                  v-model="approval[item.name]"
                  type="textarea"
                  placeholder="请输入"
                />
                <el-select
                  v-else-if="item.type === '单选项'"
                  v-model="approval[item.name]"
                  style="width: 100%"
                  @change="type_change(item)"
                >
                  <el-option
                    v-for="jtem in item.options"
                    :key="jtem"
                    :value="jtem"
                  />
                </el-select>
                <el-date-picker
                  v-if="item.type === '日期'"
                  v-model="approval[item.name]"
                  format="yyyy-MM-dd HH:mm:ss"
                  value-format="yyyy-MM-dd HH:mm:ss"
                  style="width: 100%"
                  type="datetime"
                  placeholder="选择日期时间"
                />
                <el-input
                  v-else-if="item.type === '数字'"
                  v-model="approval[item.name]"
                  type="number"
                  placeholder="请输入"
                />
              </el-col>
            </el-row>
            <span slot="footer" class="dialog-footer">
              <el-button type="text" @click="dialogVisible = false">
                取 消
              </el-button>
              <el-button type="text" @click="dialogConfirm"> 确 定 </el-button>
            </span>
          </el-dialog>
          <el-dialog
            title="选择审批人"
            :visible.sync="dialogVisible2"
            width="45%"
          >
            <el-form :model="form_one" label-width="80px">
              <el-form-item label="审批标题">
                <el-input v-model="form_one.name" disabled></el-input>
              </el-form-item>
              <el-form-item label="审批人">
                <el-select
                  v-model="form_one.user_id"
                  style="width: 100%"
                  placeholder="请选择审批人"
                  filterable
                  clearable
                >
                  <el-option
                    v-for="item in user_list"
                    :key="item.id"
                    :label="item.username"
                    :value="item.id"
                  />
                </el-select>
              </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
              <el-button type="text" @click="dialogCancel2"> 取 消 </el-button>
              <el-button type="text" @click="dialogConfirm2"> 确 定 </el-button>
            </span>
          </el-dialog>
        </el-tab-pane>
      </el-tabs>
    </div>
    <div v-if="is_mobile">
      <el-tabs v-model="activeName">
        <el-tab-pane label="服务发布" name="first">
          <div style="width: 100%; margin: 0 auto">
            <el-form
              :model="approval"
              :rules="approval_rule"
              label-width="120px"
            >
              <el-form-item label="审批工单" prop="approval_id">
                <el-select
                  v-model="approval.approval_id"
                  placeholder="请选择工单"
                  style="width: 100%"
                  @change="approval_change"
                >
                  <el-option
                    v-for="item in approval_list"
                    :key="item.id"
                    :label="item.approval_name"
                    :value="item.id"
                  />
                </el-select>
                <el-button
                  v-if="add_btn_show === 1"
                  style="margin-left: 15px"
                  size="small"
                  type="text"
                  @click="approval_yes"
                >
                  {{ instance_id ? "添加服务" : "创建" }}
                </el-button>
              </el-form-item>
            </el-form>
            <el-table
              v-if="services.length > 0"
              :data="services"
              style="width: 100%"
            >
              <el-table-column
                v-for="key in tableKeys"
                :key="key"
                :prop="key"
                :label="key"
                width="150"
              />
              <el-table-column label="操作" fixed="right" width="100">
                <template slot-scope="scope">
                  <el-button type="text" @click="addService()">
                    添加
                  </el-button>
                  <el-button type="text" @click="delService(scope.$index)">
                    移除
                  </el-button>
                </template>
              </el-table-column>
            </el-table>
            <el-row v-show="!un_finish">
              <el-col :span="24">
                <el-row>
                  <el-col :span="8" style="margin-top: 5px"> 附件信息 </el-col>
                  <el-col :span="16">
                    <el-upload
                      accept=".xlsx, .xls"
                      :file-list="fileList"
                      :action="actionUrl"
                      :data="info"
                      :before-upload="beforeUploadFile"
                      :on-exceed="exceedFile"
                      :on-success="uploadSuccess"
                      :show-file-list="isShowFile"
                      multiple
                      :limit="limitNum"
                    >
                      <el-button
                        v-if="detailIndex === 0"
                        type="text"
                        icon="el-icon-upload2"
                        size="mini"
                      >
                        添加文档
                      </el-button>
                    </el-upload>
                  </el-col>
                </el-row>
                <el-row>
                  <el-col
                    v-for="item in files"
                    :key="item"
                    :span="20"
                    :offset="4"
                  >
                    <span> {{ item }} </span>
                    <el-button
                      size="mini"
                      type="text"
                      style="margin-left: 45px"
                      @click="download(item)"
                    >
                      下载
                    </el-button>
                    <el-button
                      size="mini"
                      type="text"
                      style="margin-left: 15px"
                      @click="del_btn(item)"
                    >
                      删除
                    </el-button>
                  </el-col>
                </el-row>
              </el-col>
            </el-row>
            <el-row v-show="!un_finish">
              <el-col :span="12" :offset="6">
                <h4>审批流程</h4>
                <div>
                  <div v-show="un_finish" class="info">
                    必填信息填写完整后，将显示审批流程
                  </div>
                  <el-row v-show="!un_finish">
                    <el-col :span="24">
                      <div
                        v-if="task_list.length > 0"
                        :style="{ height: task_list.length * 60 + 'px' }"
                      >
                        <el-steps
                          direction="vertical"
                          :active="1"
                          style="cursor: pointer"
                        >
                          <el-step
                            v-for="item in task_list"
                            :key="item.id"
                            :title="item.name"
                            @click.native="step_click(item)"
                          >
                            <template slot="description">
                              <div>
                                <span v-if="item.username !== ''">
                                  {{
                                    item.name === "开始"
                                      ? sp_person
                                      : item.name === "结束"
                                      ? sp_person
                                      : item.username
                                  }}
                                </span>
                                <span v-else style="color: red">
                                  [审批人自选]
                                </span>
                              </div>
                            </template>
                          </el-step>
                        </el-steps>
                      </div>
                    </el-col>
                  </el-row>
                </div>
              </el-col>
            </el-row>
            <el-row v-show="!un_finish" style="margin-top: 15px">
              <el-col :span="12" :offset="6">
                <el-button type="text" @click="confirm()">确认</el-button>
                <el-button
                  v-if="is_confirm"
                  type="text"
                  @click="go_to_detail()"
                >
                  详情
                </el-button>
              </el-col>
            </el-row>
          </div>
          <el-dialog
            title="添加服务"
            :visible.sync="dialogVisible"
            width="100%"
            :before-close="handleClose"
          >
            <el-row
              v-for="item in form_list"
              :key="item.id"
              :label="item.name"
              style="margin-bottom: 5px"
            >
              <el-col
                :span="4"
                style="
                  text-align: right;
                  padding-right: 15px;
                  font-weight: bold;
                  line-height: 30px;
                "
              >
                {{ item.name }}
              </el-col>
              <el-col :span="20">
                <el-input
                  v-if="item.type === '单行文本'"
                  v-model="approval[item.name]"
                  placeholder="请输入"
                />
                <el-input
                  v-else-if="item.type === '多行文本'"
                  v-model="approval[item.name]"
                  type="textarea"
                  placeholder="请输入"
                />
                <el-select
                  v-else-if="item.type === '单选项'"
                  v-model="approval[item.name]"
                  style="width: 100%"
                >
                  <el-option
                    v-for="jtem in item.options"
                    :key="jtem"
                    :value="jtem"
                  />
                </el-select>
                <el-date-picker
                  v-if="item.type === '日期'"
                  v-model="approval[item.name]"
                  format="yyyy-MM-dd HH:mm:ss"
                  value-format="yyyy-MM-dd HH:mm:ss"
                  style="width: 100%"
                  type="datetime"
                  placeholder="选择日期时间"
                />
                <el-input
                  v-else-if="item.type === '数字'"
                  v-model="approval[item.name]"
                  type="number"
                  placeholder="请输入"
                />
              </el-col>
            </el-row>
            <span slot="footer" class="dialog-footer">
              <el-button type="text" @click="dialogVisible = false">
                取 消
              </el-button>
              <el-button type="text" @click="dialogConfirm"> 确 定 </el-button>
            </span>
          </el-dialog>
          <el-dialog
            title="选择审批人"
            :visible.sync="dialogVisible2"
            width="100%"
          >
            <el-form :model="form_one" label-width="80px">
              <el-form-item label="审批标题">
                <el-input v-model="form_one.name" disabled></el-input>
              </el-form-item>
              <el-form-item label="审批人">
                <el-select
                  v-model="form_one.user_id"
                  style="width: 100%"
                  placeholder="请选择审批人"
                  filterable
                  clearable
                >
                  <el-option
                    v-for="item in user_list"
                    :key="item.id"
                    :label="item.username"
                    :value="item.id"
                  />
                </el-select>
              </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
              <el-button type="text" @click="dialogCancel2"> 取 消 </el-button>
              <el-button type="text" @click="dialogConfirm2"> 确 定 </el-button>
            </span>
          </el-dialog>
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>
<script>
import { getToken } from "@/utils/auth";
import jwtDecode from "jwt-decode";
import { approval_get, form_get, node_get } from "@/api/approval/approval";
import { service_get } from "@/api/project-list";
import { usernames_get, user_get } from "@/api/user";
import {
  fiform_cud,
  fiform_get,
  instance_cud,
  instance_get,
} from "@/api/approval/instance";
import { task_user_cud, task_get } from "@/api/approval/task";
import axios from "axios";
import config from "@/utils/config";
import get_date_time from "@/utils/assets";
import { btn_check } from "@/api/btn";
import { ApprovalStartPage } from "@/utils/auth";
export default {
  name: "ApprovalUse",
  data() {
    return {
      sp_person: "",
      un_finish: true,
      activeName: "first",
      approval: {
        approval_id: "",
      },
      approval_list: [],
      approval_rule: {
        approval_id: [
          { required: true, message: "请选择审批流", trigger: "change" },
        ],
      },
      form_list: [],
      form_one: {
        id: "",
        name: "",
        user_id: "",
      },
      node_list: [],
      approval_id: "",
      instance_id: "",
      is_mobile: config.isMobile,
      actionUrl: process.env.VUE_APP_BASE_API + "/approval/v2/upload-file",
      info: {
        file_name: "", // 上传的额外参数
      },
      fileList: [],
      files: [],
      limitNum: 5,
      isShowFile: false,
      detailIndex: 0,
      services: [],
      tableKeys: [],
      dialogVisible: false,
      dialogVisible2: false,
      service_num: 0,
      is_confirm: false,
      user_list: [],
      task_list: [],
      ApprovalStartPage: ApprovalStartPage,
      add_btn_show: 0,
      del_btn_show: 0,
      edit_btn_show: 0,
      conn_btn_show: 0,
    };
  },
  watch: {
    $route(to, from) {
      this.init_data();
    },
  },
  mounted() {
    this.init_data();
  },
  methods: {
    async step_click(data) {
      if (this.user_list.length === 0) {
        var result6 = await usernames_get();
        if (result6.code === 200) {
          this.user_list = result6.data.users;
        }
      }
      if (data.need_approver === true && data.user_id === "") {
        this.form_one = data;
        this.dialogVisible2 = true;
      }
    },
    type_change(item) {
      if (item.name === "服务类型" || item.name === "项目名称") {
        service_get({
          service_type: this.approval["服务类型"],
          project: this.approval["项目名称"],
          page_no: 1,
          page_size: 999,
        }).then((response) => {
          if (response.code === 200) {
            let service_list = response.data.service_infos;
            let options = [];
            for (var i = 0; i < service_list.length; i++) {
              options.push(service_list[i]["name"]);
            }
            for (var i = 0; i < this.form_list.length; i++) {
              let f = this.form_list[i];
              if (f.en_name === "SERVICE") {
                this.form_list[i]["options"] = options;
              }
            }
          }
        });
      }
    },
    async init_data() {
      this.approval_id = Number(this.$route.query.id);
      this.instance_id = Number(this.$route.query.instance_id);
      var result5 = await approval_get({ subscribe: 1, old_version: 0 });
      if (result5.code === 200) {
        this.approval_list = result5.data;
      }
      btn_check(this.ApprovalStartPage).then((response) => {
        if (response.code == 200) {
          this.add_btn_show = response.data.add;
          this.edit_btn_show = response.data.edit;
          this.del_btn_show = response.data.del;
          this.conn_btn_show = response.data.conn;
        }
      });
      if (this.approval_id && this.instance_id) {
        this.approval.approval_id = this.approval_id;
        var result6 = await usernames_get();
        if (result6.code === 200) {
          this.user_list = result6.data.users;
        }
        const userObj = jwtDecode(getToken());
        var obj = await user_get({ name: userObj.username });
        if (obj.code === 200) {
          var username = obj.data.users[0]["zh_name"];
          if (username.length > 0) {
            this.sp_person = username;
          }
        }
        var result = await form_get({ approval_id: this.approval_id });
        if (result.code === 200) {
          this.form_list = result.data;
        }
        var result2 = await fiform_get({ instance_id: this.instance_id });
        if (result2.code === 200) {
          this.services = result2.data;
        }
        var result3 = await node_get({ approval_id: this.approval_id });
        if (result3.code === 200) {
          this.node_list = result3.data;
        }
        var result4 = await task_get({ instance_id: this.instance_id });
        if (result4.code === 200) {
          this.task_list = result4.data;
        }
        this.un_finish = false;
        if (this.services.length > 0) {
          this.tableKeys = Object.keys(this.services[0]);
        }
      }
    },
    handleClose(done) {
      done();
    },
    async addService() {
      var result = await form_get({ approval_id: this.approval.approval_id });
      if (result.code === 200) {
        this.form_list = result.data;
        this.approval = {};
        this.dialogVisible = true;
        this.approval.approval_id = this.approval_id;
      }
    },
    delService(row) {
      fiform_cud("delete", {
        instance_id: this.instance_id,
        service: row["服务编号"],
      }).then((response) => {
        if (response.code === 200) {
          this.$message({
            message: response.msg,
            type: "success",
          });
          setTimeout(()=>{
            location.reload(0);
          }, 500)
        }
      });
    },
    dialogCancel2() {
      this.form_one = {};
      this.dialogVisible2 = false;
    },
    dialogConfirm2() {
      if (this.form_one.user_id === "") {
        this.$message({
          message: "还未选择审批人",
          type: "warning",
        });
      } else {
        this.dialogVisible2 = false;
        task_user_cud({
          id: this.form_one.id,
          user_id: this.form_one.user_id,
        }).then((response) => {
          if (response.code === 200) {
            this.$message({
              message: response.msg,
              type: "success",
            });
            location.reload(0);
          }
        });
      }
    },
    dialogConfirm() {
      if (Object.keys(this.approval).length < 6) {
        this.$message({
          message: "表单内容必须填写完整",
          type: "error",
        });
      } else {
        const result = {};
        const kvs = {};
        for (var i = 0; i < this.form_list.length; i++) {
          const item = this.form_list[i];
          if (kvs.hasOwnProperty(item.name) === false) {
            kvs[item.name] = item.en_name;
          }
        }
        for (var j in this.approval) {
          if (kvs.hasOwnProperty(j)) {
            result[kvs[j]] = this.approval[j];
          }
        }
        if (result["DATE"] === undefined || result["DATE"] === "") {
          result["DATE"] = get_date_time();
        }
        delete result["approval_id"];
        this.service_num = this.services.length + 1;
        const data = {
          approval_id: this.approval_id,
          instance_id: this.instance_id,
          service: this.service_num,
          form: result,
        };
        if (data["approval_id"] === undefined) {
          this.$message({
            message: "必须传递 approval_id",
            type: "error",
          });
        } else {
          fiform_cud("post", data).then((response) => {
            if (response.code === 200) {
              this.$message({
                message: response.msg,
                type: "success",
              });
              this.service_num += 1;
              this.dialogVisible = false;
              this.services = response.data;
              this.tableKeys = Object.keys(this.services[0]);
              this.$router.push({
                path: "/approvals/use/start",
                query: { id: this.approval_id, instance_id: this.instance_id },
              });
            }
          });
        }
      }
    },
    approval_change() {
      this.approval_id = this.approval.approval_id;
    },
    approval_yes() {
      if (this.approval_id === undefined) {
        this.$message({
          message: "请选择审批工单",
          type: "warning",
        });
      } else {
        if (!this.instance_id) {
          instance_cud("post", { approval_id: this.approval_id }).then(
            (response) => {
              if (response.code === 200) {
                const instance = response.data[0];
                this.instance_id = instance.id;
                this.$message({
                  message: "审批实例可以创建",
                  type: "success",
                });
                this.un_finish = false;
                node_get({ approval_id: this.approval_id }).then((response) => {
                  if (response.code === 200) {
                    this.node_list = response.data;
                  }
                });
                task_get({ instance_id: this.instance_id }).then((response) => {
                  if (response.code === 200) {
                    this.task_list = response.data;
                  }
                });
                form_get({ approval_id: this.approval.approval_id }).then(
                  (response) => {
                    if (response.code === 200) {
                      this.form_list = response.data;
                      this.dialogVisible = true;
                      this.$router.push({
                        path: "/approvals/use/start",
                        query: {
                          id: this.approval_id,
                          instance_id: this.instance_id,
                        },
                      });
                    }
                  }
                );
              } else {
                this.$message({
                  message: "审批实例无法创建",
                  type: "error",
                });
              }
            }
          );
        } else {
          form_get({ approval_id: this.approval_id }).then((response) => {
            if (response.code === 200) {
              this.form_list = response.data;
              this.dialogVisible = true;
            }
          });
        }
      }
    },
    response_func(response) {
      if (response.code === 200) {
        this.dialogVisible = false;
        this.$message({
          message: response.msg,
          type: "success",
        });
      } else {
        this.$message({
          message: response.msg,
          type: "warning",
        });
      }
    },
    response_refresh_func(response) {
      if (response.code === 200) {
        this.dialogVisible = false;
        this.$message({
          message: response.msg,
          type: "success",
        });
        setTimeout(() => {
          location.reload(0);
        }, 1000);
      } else {
        this.$message({
          message: response.msg,
          type: "warning",
        });
      }
    },
    confirm() {
      const data = {
        id: this.instance_id,
        status: "PENDING",
      };
      var unfi = false;
      for (var i = 0; i < this.task_list.length; i++) {
        var item = this.task_list[i];
        if (item.need_approver === true && item.user_id === "") {
          this.$message({
            message: "还未选择审批人",
            type: "warning",
          });
          unfi = true;
          break;
        }
      }
      if (unfi === false) {
        instance_cud("put", data).then((response) => {
          if (response.code === 200) {
            this.$message({
              message: response.msg,
              type: "success",
            });
            this.$router.push({
              path: "/approvals/use/detail",
              query: {
                instance_id: this.instance_id,
              },
            });
          }
        });
      }
    },
    go_to_detail() {
      this.$router.push({
        path: "/approvals/use/detail",
        query: {
          instance_id: this.instance_id,
        },
      });
    },
    // 文件上传之前的钩子
    beforeUploadFile(file) {
      this.info.file_name = file.name;
      instance_get({ id: this.instance_id }).then((response) => {
        if (response.code === 200) {
          const instance = response.data.instances[0];
          if (instance.appendix !== null) {
            this.files = instance.appendix.split(",");
          }
          this.files.push(this.info.file_name);
          let str = "";
          for (var i = 0; i < this.files.length; i++) {
            const item = this.files[i];
            if (item !== "") {
              str += "," + this.files[i];
            }
          }
          instance_cud("put", {
            id: this.instance_id,
            appendix: str,
          }).then((response) => {
            if (response.code === 200) {
              this.$message({
                message: response.msg,
                type: "success",
              });
            }
          });
        }
      });

      var FileExt = file.name.replace(/.+\./, "");
      console.log("文件名: ", FileExt);
    },
    // 文件超出个数时的钩instance_cud子
    exceedFile(files, fileList) {
      console.log("文件超出个数：", files);
      this.$message.warning(
        `只能选择 ${this.limitNum} 个文件，当前共选择了 ${
          files.length + fileList.length
        } 个`
      );
    },
    // 文件上传成功的钩子
    uploadSuccess(response, file, fileList) {
      if (response.code !== 200) {
        return this.$message.error(response.msg);
      } else {
        // 上传成功之后在这里写逻辑，比如：重新调用查询列表接口
        return this.$message.success(response.msg);
      }
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
    del_btn(file_name) {
      this.files.splice(this.files.indexOf(file_name), 1);
      let appendix = "";
      for (var i = 0; i < this.files.length; i++) {
        appendix += "," + this.files[i];
      }
      instance_cud("put", { id: this.instance_id, appendix: appendix }).then(
        (response) => {
          if (response.code === 200) {
            this.$message({
              message: response.msg,
              type: "success",
            });
          }
        }
      );
    },
  },
};
</script>
<style lang="scss" scoped>
.approval-use {
  min-height: 450px;
  width: 98%;
  margin: 15px auto;
  background-color: #fff;
  padding: 15px;
  .info {
    font-size: 14px;
    color: #999;
  }
}
</style>
