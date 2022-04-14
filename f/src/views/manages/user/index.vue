<template>
  <div class="manages-user">
    <div v-if="!is_mobile">
      <el-card class="box-card" shadow="always">
        <el-form :inline="true" :model="formInline" class="demo-form-inline">
          <el-form-item label="用户名称">
            <el-input v-model="formInline.name" placeholder="请输入用户名称" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="search_user">查询</el-button>
          </el-form-item>
        </el-form>
      </el-card>
      <el-card class="box-card">
        <div slot="header" class="clearfix">
          <span>账户列表</span>
          <el-button-group style="float: right; padding: 3px 10px">
            <el-button
              :type="btn_status == 'all' ? 'primary' : ''"
              size="small"
              @click="status_btn('all')"
            >全部</el-button>
            <el-button
              :type="btn_status == 'on' ? 'primary' : ''"
              size="small"
              @click="status_btn('on')"
            >正常</el-button>
            <el-button
              :type="btn_status == 'off' ? 'primary' : ''"
              size="small"
              @click="status_btn('off')"
            >禁用</el-button>
          </el-button-group>
        </div>
        <el-table v-loading="loading" :data="userList" style="width: 100%">
          <el-table-column prop="username" label="登录名" width="200" />
          <el-table-column prop="zh_name" label="用户名" width="200" />
          <el-table-column label="状态" width="200">
            <template slot-scope="scope">
              <div v-if="scope.row.status === 'on'">
                <i class="el-icon-user-solid" style="color: green" />
                正常
              </div>
              <div v-else>
                <i class="el-icon-user-solid" />
                禁用
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="last_login" label="最近登录" />
          <el-table-column label="操作" width="250">
            <template slot-scope="scope">
              <el-button
                v-if="scope.row.status === 'on'"
                type="text"
                @click="forbidden_btn(scope.row)"
              >禁用</el-button>
              <el-button
                v-else
                type="text"
                @click="forbidden_btn(scope.row)"
              >启用</el-button>
              <el-button
                type="text"
                @click="edit_btn(scope.row)"
              >编辑</el-button>
            </template>
          </el-table-column>
        </el-table>
        <el-pagination
          small
          layout="prev, pager, next"
          :page-size="count_one_page"
          :total="total"
          @current-change="handleCurrentChange"
        />
      </el-card>
      <el-dialog
        title="编辑用户"
        :visible.sync="dialogVisible"
        width="45%"
        :before-close="handleClose"
      >
        <el-form ref="user" :model="user" label-width="100px">
          <el-form-item label="登录名">
            <el-input
              v-model="user.username"
              disabled
              placeholder="请输入登录名"
            />
          </el-form-item>
          <el-form-item label="姓名">
            <el-input
              v-model="user.zh_name"
              disabled
              placeholder="请输入用户名"
            />
          </el-form-item>
          <el-form-item label="绑定角色">
            <el-select v-model="user.roles" multiple placeholder="请选择">
              <el-option
                v-for="item in roles"
                :key="item.id"
                :label="item.name"
                :value="item.name"
              />
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button @click="dialog_cancel()">取 消</el-button>
            <el-button
              type="primary"
              @click="dialog_confirm()"
            >确 定</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
    <div v-if="is_mobile">
      <el-card class="box-card" shadow="always">
        <el-form :inline="true" :model="formInline" class="demo-form-inline">
          <el-form-item label="用户名称">
            <el-input v-model="formInline.name" placeholder="请输入用户名称" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="search_user">查询</el-button>
          </el-form-item>
        </el-form>
      </el-card>
      <el-card class="box-card">
        <div slot="header" class="clearfix">
          <span>账户列表</span>
          <el-button-group style="float: right; padding: 3px 10px">
            <el-button
              :type="btn_status == 'all' ? 'primary' : ''"
              size="small"
              @click="status_btn('all')"
            >全部</el-button>
            <el-button
              :type="btn_status == 'on' ? 'primary' : ''"
              size="small"
              @click="status_btn('on')"
            >正常</el-button>
            <el-button
              :type="btn_status == 'off' ? 'primary' : ''"
              size="small"
              @click="status_btn('off')"
            >禁用</el-button>
          </el-button-group>
        </div>
        <el-table v-loading="loading" :data="userList" style="width: 100%">
          <el-table-column prop="username" label="登录名" width="120" fixed />
          <el-table-column prop="zh_name" label="用户名" width="200" />
          <el-table-column label="状态" width="80" fixed>
            <template slot-scope="scope">
              <div v-if="scope.row.status === 'on'">
                <i class="el-icon-user-solid" style="color: green" />
                正常
              </div>
              <div v-else>
                <i class="el-icon-user-solid" />
                禁用
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="last_login" label="最近登录" width="200" />
          <el-table-column label="操作" width="250">
            <template slot-scope="scope">
              <el-button
                v-if="scope.row.status === 'on'"
                type="text"
                @click="forbidden_btn(scope.row)"
              >禁用</el-button>
              <el-button
                v-else
                type="text"
                @click="forbidden_btn(scope.row)"
              >启用</el-button>
              <el-button
                type="text"
                @click="edit_btn(scope.row)"
              >编辑</el-button>
            </template>
          </el-table-column>
        </el-table>
        <el-pagination
          small
          layout="prev, pager, next"
          :page-size="count_one_page"
          :total="total"
          @current-change="handleCurrentChange"
        />
      </el-card>
      <el-dialog
        title="编辑用户"
        :visible.sync="dialogVisible"
        width="100%"
        :before-close="handleClose"
      >
        <el-form ref="user" :model="user" label-width="100px">
          <el-form-item label="登录名">
            <el-input
              v-model="user.username"
              disabled
              placeholder="请输入登录名"
            />
          </el-form-item>
          <el-form-item label="姓名">
            <el-input
              v-model="user.zh_name"
              disabled
              placeholder="请输入用户名"
            />
          </el-form-item>
          <el-form-item label="绑定角色">
            <el-select v-model="user.roles" multiple placeholder="请选择">
              <el-option
                v-for="item in roles"
                :key="item.id"
                :label="item.name"
                :value="item.name"
              />
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button size="mini" @click="dialog_cancel()">取 消</el-button>
            <el-button
              size="mini"
              type="primary"
              @click="dialog_confirm()"
            >确 定</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
  </div>
</template>
<script>
import { manage_role_get } from "@/api/manages/role";
import { users_get, users_cud, user_get } from "@/api/user";
import config from "@/utils/config";
export default {
  name: "ManagesUser",
  data() {
    return {
      roles: [],
      formInline: {
        name: "",
      },
      user: {},
      userList: [],
      dialogVisible: false,
      btn_status: "all",
      now_page: 1,
      count_one_page: 5,
      total: 1,
      is_mobile: config.isMobile,
      loading: true
    };
  },
  mounted() {
    this.init_data();
    const that = this;
    document.onkeydown = function(e) {
      const key = window.event.keyCode;
      if (key === 13) {
        const query = {
          page_no: this.now_page,
          page_size: this.count_one_page,
          name: that.formInline.name,
        };
        user_get(query).then((response) => {
          if (response.code === 200) {
            that.userList = response.data;
          } else {
            that.userList = [];
          }
        });
      }
    };
  },
  methods: {
    response_refresh_func(response) {
      if (response.code == 200) {
        this.dialogVisible = false;
        this.$message({
          message: response.msg,
          type: "success",
        });
        location.reload(0);
      } else {
        this.$message({
          message: response.msg,
          type: "warning",
        });
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
    init_data() {
      const query = {
        page_no: 1,
        page_size: 9999,
      };
      manage_role_get(query).then((response) => {
        if (response.code === 200) {
          this.roles = response.data.roles;
        } else {
          this.roles = [];
        }
      });
      this.page_data();
    },
    page_data() {
      const query = {
        page_no: this.now_page,
        page_size: this.count_one_page,
      };
      users_get(query).then((response) => {
        if (response.code === 200) {
          this.userList = response.data.users;
          this.total = response.data.total;
        } else {
          this.userList = [];
        }
        this.loading = false;
      });
    },
    handleCurrentChange(val) {
      this.now_page = val;
      this.page_data();
    },
    handleClose(done) {
      this.$confirm("确认关闭？")
        .then((_) => {
          done();
        })
        .catch((_) => {});
    },
    search_user() {
      user_get(this.formInline).then((response) => {
        if (response.code === 200) {
          this.userList = response.data.users;
          this.total = response.data.total;
        } else {
          this.userList = [];
        }
      });
    },
    status_btn(status) {
      this.btn_status = status;
      user_get({ status: status }).then((response) => {
        if (response.code === 200) {
          this.userList = response.data.users;
          this.total = response.data.total;
        } else {
          this.userList = [];
          this.total = 1;
        }
      });
    },
    edit_btn(row) {
      this.user = row;
      this.dialogVisible = true;
    },
    forbidden_btn(row) {
      this.user = row;
      const data = {
        type: "status",
        id: row.id,
        value: row.status == "on" ? "off" : "on",
      };
      users_cud(data).then((response) => {
        this.response_func(response);
      });
      row.status = row.status == "on" ? "off" : "on";
    },
    dialog_confirm() {
      const roles = "[" + this.user.roles.join(",") + "]";
      const data = {
        type: "roles",
        id: this.user.id,
        value: roles,
      };
      users_cud(data).then((response) => {
        this.response_refresh_func(response);
      });
    },
    dialog_cancel() {
      this.role = {};
      this.dialogVisible = false;
    },
  },
};
</script>
<style lang='scss' scoped>
.manages-user {
  width: 99%;
  min-height: 680px;
  margin: 0 auto;
  .box-card {
    margin: 15px auto;
  }
}
</style>
