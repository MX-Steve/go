<template>
  <div class="service">
    <div v-if="!is_mobile">
      <div class="sub-title">服务列表</div>
      <div class="plus-button-box">
        <el-button v-if="add_btn_show" size="small" @click="create_btn()">
          添加
        </el-button>
      </div>
      <el-table v-loading="loading" :data="service_list" style="width: 100%">
        <el-table-column prop="id" label="ID" width="55" />
        <el-table-column prop="name" label="服务名" width="180" />
        <el-table-column prop="manager" label="负责人" width="180">
          <template slot-scope="scope">
            <span v-for="(item, index) in scope.row.manager" :key="item">
              {{ get_username(item) }}
              <i v-if="index != scope.row.manager.length - 1"> , </i>
            </span>
          </template>
        </el-table-column>
        <el-table-column prop="rel_project" label="所属项目" width="180">
          <template slot-scope="scope">
            {{ get_name(scope.row.rel_project, project_list) }}
          </template>
        </el-table-column>
        <el-table-column prop="service_type" label="所属类型" width="180" />
        <el-table-column prop="rel_machine" label="包含机器">
          <template slot-scope="scope">
            <span
              v-for="(item, index) in scope.row.rel_machine"
              :key="index"
              style="margin-right: 15px"
            >
              {{ get_ip_address(item) }} <br />
            </span>
          </template>
        </el-table-column>
        <el-table-column prop="remark" label="备注" width="180" />
        <el-table-column prop="u_time" label="更新时间">
          <template slot-scope="scope">
            {{ scope.row.u_time.split(".")[0] }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="180">
          <template slot-scope="scope">
            <el-button
              v-if="edit_btn_show"
              type="primary"
              size="mini"
              icon="el-icon-edit"
              @click="edit_btn(scope.row)"
            />
            <el-button
              v-if="del_btn_show"
              type="danger"
              size="mini"
              icon="el-icon-delete"
              @click="delete_btn(scope.row)"
            />
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        small
        layout="prev, pager, next"
        :page-size="page_size"
        :total="total"
        @current-change="handleCurrentChange"
      />
      <el-dialog
        :title="dialogTitle"
        :visible.sync="dialogVisible"
        width="45%"
        center
        :before-close="handleClose"
      >
        <el-form
          ref="service"
          :rules="serviceRule"
          :model="service"
          label-width="80px"
        >
          <el-form-item prop="name" label="服务名">
            <el-input v-model="service.name" />
          </el-form-item>
          <el-form-item label="备注">
            <el-input v-model="service.remark" />
          </el-form-item>
          <el-form-item prop="manager" label="负责人">
            <el-select
              v-model="service.manager"
              multiple
              filterable
              placeholder="请输入"
              style="width: 100%"
            >
              <el-option
                v-for="item in user_list"
                :key="item.id"
                :label="item.username"
                :value="item.id"
              />
            </el-select>
          </el-form-item>
          <el-form-item prop="rel_project" label="所属项目">
            <el-select
              v-model="service.rel_project"
              filterable
              placeholder="请选择"
              style="width: 100%"
            >
              <el-option
                v-for="item in project_list"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
          </el-form-item>
          <el-form-item prop="service_type" label="所属类型">
            <el-select
              v-model="service.service_type"
              filterable
              placeholder="请选择"
              style="width: 100%"
            >
              <el-option
                v-for="item in service_types"
                :key="item"
                :label="item"
                :value="item"
              />
            </el-select>
          </el-form-item>
          <el-form-item prop="env_id" label="环境">
            <el-select
              v-model="service.env_id"
              filterable
              placeholder="请选择"
              style="width: 100%"
              @change="env_change(service.env_id)"
            >
              <el-option
                v-for="item in env_list"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="包含机器">
            <el-transfer
              v-model="service.rel_machine"
              filterable
              :titles="['所有机器', '已选机器']"
              :data="machine_list"
              :props="{
                key: 'id',
                label: 'ip_address',
              }"
              @change="ba(service.env_id, service.rel_machine)"
            />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="dialog_confirm()">确定</el-button>
            <el-button @click="dialog_cancel()">取消</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
    <div v-if="is_mobile">
      <div class="sub-title">服务列表</div>
      <div class="plus-button-box">
        <el-button v-if="add_btn_show" size="small" @click="create_btn()">
          添加
        </el-button>
      </div>
      <el-table v-loading="loading" :data="service_list" style="width: 100%">
        <el-table-column prop="id" label="ID" width="55" fixed />
        <el-table-column prop="name" label="服务名" width="150" fixed />
        <el-table-column prop="manager" label="负责人" width="80">
          <template slot-scope="scope">
            <span v-for="(item, index) in scope.row.manager" :key="item">
              {{ get_username(item) }}
              <i v-if="index != scope.row.manager.length - 1"> , </i>
            </span>
          </template>
        </el-table-column>
        <el-table-column prop="rel_project" label="所属项目" width="180">
          <template slot-scope="scope">
            {{ get_name(scope.row.rel_project, project_list) }}
          </template>
        </el-table-column>
        <el-table-column prop="service_type" label="所属类型" width="180" />
        <el-table-column prop="rel_machine" width="210" label="包含机器">
          <template slot-scope="scope">
            <span
              v-for="(item, index) in scope.row.rel_machine"
              :key="index"
              style="margin-right: 15px"
            >
              {{ get_ip_address(item) }} <br />
            </span>
          </template>
        </el-table-column>
        <el-table-column prop="remark" label="备注" width="180" />
        <el-table-column prop="u_time" width="150" label="更新时间">
          <template slot-scope="scope">
            {{ scope.row.u_time.split(".")[0] }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="180">
          <template slot-scope="scope">
            <el-button
              v-if="edit_btn_show"
              type="primary"
              size="mini"
              icon="el-icon-edit"
              @click="edit_btn(scope.row)"
            />
            <el-button
              v-if="del_btn_show"
              type="danger"
              size="mini"
              icon="el-icon-delete"
              @click="delete_btn(scope.row)"
            />
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        small
        layout="prev, pager, next"
        :page-size="page_size"
        :total="total"
        @current-change="handleCurrentChange"
      />
      <el-dialog
        :title="dialogTitle"
        :visible.sync="dialogVisible"
        width="100%"
        center
        :before-close="handleClose"
      >
        <el-form
          ref="service"
          :rules="serviceRule"
          :model="service"
          label-width="80px"
        >
          <el-form-item prop="name" label="服务名">
            <el-input v-model="service.name" />
          </el-form-item>
          <el-form-item label="备注">
            <el-input v-model="service.remark" />
          </el-form-item>
          <el-form-item prop="manager" label="负责人">
            <el-select
              v-model="service.manager"
              multiple
              filterable
              placeholder="请输入"
              style="width: 100%"
            >
              <el-option
                v-for="item in user_list"
                :key="item.id"
                :label="item.username"
                :value="item.id"
              />
            </el-select>
          </el-form-item>
          <el-form-item prop="rel_project" label="所属项目">
            <el-select
              v-model="service.rel_project"
              filterable
              placeholder="请选择"
              style="width: 100%"
            >
              <el-option
                v-for="item in project_list"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
          </el-form-item>
          <el-form-item prop="service_type" label="所属类型">
            <el-select
              v-model="service.service_type"
              filterable
              placeholder="请选择"
              style="width: 100%"
            >
              <el-option
                v-for="item in service_types"
                :key="item"
                :label="item"
                :value="item"
              />
            </el-select>
          </el-form-item>
          <el-form-item prop="env_id" label="环境">
            <el-select
              v-model="service.env_id"
              filterable
              placeholder="请选择"
              style="width: 100%"
              @change="env_change(service.env_id)"
            >
              <el-option
                v-for="item in env_list"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="包含机器">
            <el-transfer
              v-model="service.rel_machine"
              filterable
              :titles="['所有机器', '已选机器']"
              :data="machine_list"
              :props="{
                key: 'id',
                label: 'ip_address',
              }"
              @change="ba(service.env_id, service.rel_machine)"
            />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="dialog_confirm()">确定</el-button>
            <el-button @click="dialog_cancel()">取消</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
  </div>
</template>
<script>
import {
  service_get,
  service_cud,
  project_get,
  env_get,
} from "@/api/project-list/index";
import { machine_get } from "@/api/assets-manage/assets-input";
import { choices_get } from "@/api/approval/choices";
import { usernames_get } from "@/api/user";
import { btn_check } from "@/api/btn";
import { get_item_name } from "@/utils/assets";
import config from "@/utils/config";
import { ProjectsPage } from "@/utils/auth";
export default {
  name: "Service",
  data() {
    return {
      dialogVisible: false,
      dialogTitle: "",
      service_list: [],
      service: {
        id: "",
        name: "",
        remark: "",
        manager: "",
        rel_project: "",
        rel_machine: [],
        c_time: "",
        u_time: "",
        env_id: "",
        service_type: "",
        env_machines: []
      },
      serviceRule: {
        name: [{ required: true, message: "服务名不能为空", trigger: "blur" }],
        env_id: [{ required: true, message: "请选择环境", trigger: "change" }],
        manager: [
          { required: true, message: "请选择负责人", trigger: "change" },
        ],
        rel_project: [
          { required: true, message: "请选择所属项目", trigger: "change" },
        ],
        service_type: [
          { required: true, message: "请选择所属类型", trigger: "change" },
        ],
      },
      page_no: 1,
      page_size: 10,
      total: 1,
      project_list: [],
      machine_list: [],
      user_list: [],
      env_list: [],
      service_types: [],
      add_btn_show: 0,
      edit_btn_show: 0,
      del_btn_show: 0,
      is_mobile: config.isMobile,
      ProjectsPage: ProjectsPage,
      loading: true,
    };
  },
  mounted() {
    this.init_data();
  },
  methods: {
    ba(id, data) {
      var has_env = false;
      for (var i = 0; i < this.service.env_machines.length; i++) {
        const item = this.service.env_machines[i];
        if (item.env_id === id) {
          has_env = true;
          this.service.env_machines[i]["machine_id"] = data;
        }
      }
      if (has_env === false) {
        this.service.env_machines.push({ env_id: id, machine_id: data });
      }
      console.log(this.service.env_machines);
    },
    handleClose() {
      this.service = {
        id: "",
        name: "",
        remark: "",
        manager: "",
        rel_project: "",
        rel_machine: [],
        c_time: "",
        u_time: "",
        env_id: "",
        env_machines: [],
        service_type: "",
      };
      this.dialogVisible = false;
    },
    get_name(id, list) {
      return get_item_name(id, list);
    },
    get_ip_address(id) {
      var name = "";
      for (var i = 0; i < this.machine_list.length; i++) {
        var item = this.machine_list[i];
        if (id == item.id) {
          name = item.ip_address;
          break;
        }
      }
      return name;
    },
    response_func(response) {
      if (response.code === 200) {
        this.dialogVisible = false;
        this.$message({
          message: response.msg,
          type: "success",
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
        location.replace(location.href.split("?")[0] + "?tab=service");
      }
    },
    init_data() {
      btn_check(this.ProjectsPage).then((response) => {
        if (response.code == 200) {
          this.add_btn_show = response.data.add;
          this.edit_btn_show = response.data.edit;
          this.del_btn_show = response.data.del;
        }
      });
      env_get().then((response) => {
        if (response.code == 200) {
          this.env_list = response.data;
        }
      });
      machine_get({
        type: "get_all_machines",
        page_no: 1,
        page_size: 999,
      }).then((response) => {
        this.machine_list = response.data.machines;
      });
      usernames_get().then((response) => {
        this.user_list = response.data.users;
      });
      this.page_data();
      this.get_project();
      choices_get({
        table_name: "BusinessServices",
        column: "type_choices",
      }).then((response) => {
        if (response.code === 200) {
          this.service_types = response.data;
        }
      });
    },
    get_username(id) {
      var name = "";
      for (var i = 0; i < this.user_list.length; i++) {
        var item = this.user_list[i];
        if (id == item.id) {
          name = item.username;
          break;
        }
      }
      return name;
    },
    create_btn() {
      this.service["env_machines"] = [{ env_id: "", machine_id: [] }];
      this.dialogVisible = true;
      this.dialogTitle = "服务创建";
    },
    async edit_btn(row) {
      if (row.manager == "[]") {
        row.manager = "";
      }
      this.dialogTitle = "服务更新";
      // this.service = row;
      const envMachinesObj = await service_get({
        type: "get_machine_env",
        id: row.id,
      });
      const env_machines = envMachinesObj.data;
      row.env_machines = env_machines;
      if (env_machines.length > 0) {
        row.env_id = env_machines[0]["env_id"];
        row.rel_machine = env_machines[0]["machine_id"];
      } else {
        row.env_id = this.env_list[0]["id"];
        row.rel_machine = [];
      }
      this.service = row;
      this.dialogVisible = true;
    },
    env_change(env_id) {
      this.$forceUpdate();
      this.service.rel_machine = [];
      for (var i = 0; i < this.service.env_machines.length; i++) {
        const item = this.service.env_machines[i];
        if (item.env_id == env_id) {
          this.service.rel_machine = item.machine_id;
          break;
        }
      }
    },
    delete_btn(row) {
      service_cud("delete", row).then((response) => {
        this.response_refresh_func(response);
      });
    },
    dialog_confirm() {
      if (this.service.id === "") {
        service_cud("post", this.service).then((response) => {
          this.response_refresh_func(response);
        });
      } else {
        // delete this.service["env_machines"];
        service_cud("put", this.service).then((response) => {
          this.response_refresh_func(response);
        });
      }
    },
    dialog_cancel() {
      this.service = {
        id: "",
        name: "",
        remark: "",
        manager: "",
        rel_project: "",
        rel_machine: [],
        c_time: "",
        u_time: "",
        service_type: "",
      };
      this.dialogVisible = false;
    },
    page_data() {
      const query = {
        type: "get_all_services",
        page_no: this.page_no,
        page_size: this.page_size,
      };
      service_get(query).then((response) => {
        if (response.code == 200) {
          this.service_list = response.data.service_infos;
          this.total = response.data.total;
          this.loading = false;
        }
      });
    },
    get_project() {
      project_get({
        type: "get_all_projects",
        page_no: 1,
        page_size: 999,
      }).then((response) => {
        if (response.code == 200) {
          this.project_list = response.data.project_infos;
        }
      });
    },
    handleCurrentChange(val) {
      this.page_no = val;
      this.page_data();
    },
  },
};
</script>
<style lang="scss" scoped>
.service {
  .sub-title {
    margin-top: 20px;
    text-align: center;
    font-size: 24px;
    font-weight: bold;
    padding-bottom: 10px;
    border-bottom: 2px solid #ccc;
  }
  .plus-button-box {
    text-align: right;
    padding: 15px 5px;
  }
}
</style>
