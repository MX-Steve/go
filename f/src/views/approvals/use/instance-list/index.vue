<template>
  <div class="instance-list">
    <div v-if="!is_mobile">
      <div class="sub-title">审批实例列表</div>
      <el-row style="margin: 25px auto">
        <el-col :span="3">
          <el-select
            v-model="selects.f_approval_id"
            placeholder="请选择工单"
            clearable
            @change="approval_change"
          >
            <el-option
              v-for="item in approval_list"
              :key="item.id"
              :label="item.approval_name"
              :value="item.id"
            />
          </el-select>
        </el-col>
        <el-col :span="3">
          <el-select
            v-model="selects.status"
            placeholder="请选择状态"
            clearable
            @change="status_change"
          >
            <el-option
              v-for="item in status_list"
              :key="item"
              :label="item"
              :value="item"
            />
          </el-select>
        </el-col>
        <el-col :span="3">
          <el-select
            v-model="selects.user_id"
            placeholder="请选择申请人"
            clearable
            filterable
            @change="username_change"
          >
            <el-option
              v-for="item in user_list"
              :key="item.username"
              :label="item.username"
              :value="item.id"
            />
          </el-select>
        </el-col>
        <el-col :span="1" :offset="14">
          <el-button v-if="add_btn_show === 1" type="text" @click="go_to_create()"> 创建 </el-button>
        </el-col>
      </el-row>
      <el-table
        v-loading="loading"
        :data="instance_list"
        border
        style="width: 100%"
        :row-class-name="tableRowClassName"
      >
        <el-table-column prop="id" label="ID" width="55" />
        <el-table-column prop="instance_code" label="实例编码" width="330" />
        <el-table-column prop="status" label="当前状态" width="100" />
        <el-table-column prop="username" label="用户名" width="100" />
        <el-table-column prop="descriptions" label="描述信息" />
        <el-table-column prop="job_number" label="job 序列号" width="90" />
        <el-table-column prop="approval_name" label="所属工单" />
        <el-table-column v-if="conn_btn_show === 1" label="操作">
          <template slot-scope="scope">
            <el-button v-if="conn_btn_show === 1" type="text" @click="detail_btn(scope.row.id)"> 详情 </el-button>
            <el-button v-if="del_btn_show === 1" type="text" style="color: red;" @click="del_btn(scope.row)"> 移除 </el-button>
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
    </div>
    <div v-if="is_mobile">
      <div class="sub-title">审批实例列表</div>
      <el-row style="margin: 25px auto">
        <el-col :span="24">
          <el-select
            v-model="selects.f_approval_id"
            placeholder="请选择工单"
            clearable
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
        </el-col>
        <el-col :span="24">
          <el-select
            v-model="selects.status"
            placeholder="请选择状态"
            clearable
            style="width: 100%"
            @change="status_change"
          >
            <el-option
              v-for="item in status_list"
              :key="item"
              :label="item"
              :value="item"
            />
          </el-select>
        </el-col>
        <el-col :span="24">
          <el-select
            v-model="selects.user_id"
            placeholder="请选择申请人"
            clearable
            filterable
            style="width: 100%"
            @change="username_change"
          >
            <el-option
              v-for="item in user_list"
              :key="item.username"
              :label="item.username"
              :value="item.id"
            />
          </el-select>
        </el-col>
        <el-col v-if="add_btn_show === 1" :span="24" style="text-align:right">
          <el-button type="text" @click="go_to_create()"> 创建 </el-button>
        </el-col>
      </el-row>
      <el-table
        v-loading="loading"
        :data="instance_list"
        border
        style="width: 100%"
        :row-class-name="tableRowClassName"
      >
        <el-table-column prop="id" label="ID" width="45" fixed />
        <el-table-column prop="instance_code" label="实例编码" width="300" />
        <el-table-column prop="status" label="当前状态" />
        <el-table-column prop="username" label="用户名" />
        <el-table-column prop="descriptions" label="描述信息" />
        <el-table-column prop="job_number" label="job 序列号" width="90" />
        <el-table-column
          prop="approval_name"
          label="所属工单"
          width="120"
          fixed
        />
        <el-table-column v-if="conn_btn_show === 1" label="操作" width="120">
          <template slot-scope="scope">
            <el-button v-if="conn_btn_show === 1" type="text" @click="detail_btn(scope.row.id)"> 详情 </el-button>
            <el-button v-if="del_btn_show === 1" type="text" style="color: red;" @click="del_btn(scope.row)"> 移除 </el-button>
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
    </div>
  </div>
</template>
<script>
import { approval_get } from "@/api/approval/approval";
import { instance_get, instance_cud } from "@/api/approval/instance";
import { choices_get } from "@/api/approval/choices";
import { usernames_get } from "@/api/user";
import { getToken } from "@/utils/auth";
import jwtDecode from "jwt-decode";
import config from "@/utils/config";
import { btn_check } from "@/api/btn";
import { ApprovalFiListPage } from "@/utils/auth";
export default {
  name: "InstanceList",
  data() {
    return {
      checked: "",
      selects: {
        status: "",
        f_approval_id: "",
        user_id: "",
      },
      instance_list: [],
      approval_list: [],
      status_list: [],
      page_no: 1,
      page_size: 10,
      total: 1,
      user_list: [],
      username: "",
      is_mobile: config.isMobile,
      loading: true,
      ApprovalFiListPage: ApprovalFiListPage,
      add_btn_show: 0,
      del_btn_show: 0,
      edit_btn_show: 0,
      conn_btn_show: 0,
    };
  },
  mounted() {
    this.init_data();
  },
  methods: {
    del_btn(row) {
      instance_cud("delete", { id: row.id }).then(response => {
        if (response.code === 200) {
          this.$message({
            message: response.msg,
            type: "success"
          });
          setTimeout(()=>{
            location.reload(0);
          }, 500)
        }
      });
    },
    go_to_create() {
      this.$router.push({
        path: "/approvals/use/start",
      });
    },
    async init_data() {
      const userObj = jwtDecode(getToken());
      this.username = userObj.username;
      const obj = await usernames_get();
      this.user_list = obj.data.users;
      this.page_data();
      this.get_approval();
      this.get_choices();
      btn_check(this.ApprovalFiListPage).then((response) => {
          if (response.code == 200) {
            this.add_btn_show = response.data.add;
            this.edit_btn_show = response.data.edit;
            this.del_btn_show = response.data.del;
            this.conn_btn_show = response.data.conn;
          }
        });
    },
    get_approval() {
      approval_get({ subscribe: 1 }).then((response) => {
        if (response.code === 200) {
          this.approval_list = response.data;
        }
      });
    },
    get_choices() {
      choices_get({
        table_name: "FInstanceValue",
        column: "status_choices",
      }).then((response) => {
        if (response.code === 200) {
          this.status_list = response.data;
        }
      });
    },
    approval_change() {
      this.page_no = 1;
      this.page_data();
    },
    status_change() {
      this.page_no = 1;
      this.page_data();
    },
    page_data() {
      const query = {
        page_no: this.page_no,
        page_size: this.page_size,
        f_approval_id: this.selects.f_approval_id,
        status: this.selects.status,
        user_id: this.selects.user_id,
      };
      instance_get(query).then((response) => {
        if (response.code === 200) {
          this.instance_list = response.data.instances;
          this.total = response.data.total;
        } else {
          this.instance_list = [];
          this.total = 1;
        }
        this.loading = false;
      });
    },
    handleCurrentChange(val) {
      this.page_no = val;
      this.init_data();
    },
    detail_btn(id) {
      this.$router.push({
        path: "/approvals/use/detail",
        query: {
          instance_id: id,
        },
      });
    },
    username_change() {
      this.page_no = 1;
      this.page_data();
    },
    tableRowClassName({ row, rowIndex }) {
      if (row.status === "拒绝") {
        return "error-row";
      } else if (row.status === "完成") {
        return "success-row";
      } else {
        return "warning-row";
      }
    },
  },
};
</script>
<style lang="scss">
.instance-list {
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
