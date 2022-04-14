<template>
  <div class="manages-role">
    <div v-if="!is_mobile" class="web">
      <el-card class="box-card" shadow="always">
        <el-form :inline="true" :model="formInline" class="demo-form-inline">
          <el-form-item label="角色名称">
            <el-input v-model="formInline.name" placeholder="请输入角色名称" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="searchRole">查询</el-button>
          </el-form-item>
        </el-form>
      </el-card>
      <el-card class="box-card">
        <div slot="header" class="clearfix">
          <span>角色列表</span>
          <el-button
            v-if="add_btn_show"
            type="primary"
            style="float: right; padding: 3px 10px"
            @click="create_btn()"
          >
            <i class="el-icon-plus" />
            新建
          </el-button>
        </div>
        <el-table
          v-loading="loading"
          :data="role_list"
          border
          style="width: 100%"
        >
          <el-table-column prop="id" label="ID" width="55" />
          <el-table-column prop="name" label="角色名称" width="150" />
          <el-table-column prop="count" label="关联账户" width="150" />
          <el-table-column prop="desc" label="描述信息" />
          <el-table-column label="操作" width="250">
            <template slot-scope="scope">
              <el-button type="text" v-if="edit_btn_show" @click="edit_btn(scope.row)">
                编辑
              </el-button>
              <el-button type="text" v-if="edit_btn_show" @click="ability_btn(scope.row)">
                功能列表
              </el-button>
              <el-button
                v-if="del_btn_show"
                type="text"
                style="color: red"
                @click="del_btn(scope.row)"
              >
                删除
              </el-button>
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
        :title="dialogTitle"
        :visible.sync="dialogVisible"
        width="30%"
        :before-close="handleClose"
      >
        <el-form ref="role" :model="role" label-width="100px">
          <el-form-item label="角色名称">
            <el-input v-model="role.name" placeholder="请输入角色名称" />
          </el-form-item>
          <el-form-item label="角色描述">
            <el-input v-model="role.desc" placeholder="请输入角色描述" />
          </el-form-item>
          <el-form-item>
            <el-button @click="dialog_cancel()">取 消</el-button>
            <el-button type="primary" @click="dialog_confirm()">
              确 定
            </el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
    <div v-if="is_mobile" class="mobile">
      <el-card class="box-card" shadow="always">
        <el-form :inline="true" :model="formInline" class="demo-form-inline">
          <el-form-item label="角色名称">
            <el-input v-model="formInline.name" placeholder="请输入角色名称" />
          </el-form-item>
          <el-form-item>
            <el-button size="mini" type="primary" @click="searchRole">
              查询
            </el-button>
          </el-form-item>
        </el-form>
      </el-card>
      <el-card class="box-card">
        <div slot="header" class="clearfix">
          <span>角色列表</span>
          <el-button
            v-if="add_btn_show"
            type="primary"
            style="float: right; padding: 3px 10px"
            @click="create_btn()"
          >
            <i class="el-icon-plus" />
            新建
          </el-button>
        </div>
        <el-table
          v-loading="loading"
          :data="role_list"
          border
          style="width: 100%"
        >
          <el-table-column prop="id" label="ID" width="45" fixed />
          <el-table-column prop="name" label="角色名称" width="100" />
          <el-table-column prop="count" label="关联账户" width="80" />
          <el-table-column prop="desc" label="描述信息" width="120" />
          <el-table-column label="操作" width="250">
            <template slot-scope="scope">
              <el-button v-if="edit_btn_show" type="text" @click="edit_btn(scope.row)">
                编辑
              </el-button>
              <el-button v-if="edit_btn_show" type="text" @click="ability_btn(scope.row)">
                功能列表
              </el-button>
              <el-button
                v-if="del_btn_show"
                type="text"
                style="color: red"
                @click="del_btn(scope.row)"
              >
                删除
              </el-button>
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
        :title="dialogTitle"
        :visible.sync="dialogVisible"
        width="100%"
        :before-close="handleClose"
      >
        <el-form ref="role" :model="role" label-width="100px">
          <el-form-item label="角色名称">
            <el-input v-model="role.name" placeholder="请输入角色名称" />
          </el-form-item>
          <el-form-item label="角色描述">
            <el-input v-model="role.desc" placeholder="请输入角色描述" />
          </el-form-item>
          <el-form-item>
            <el-button size="mini" @click="dialog_cancel()">取 消</el-button>
            <el-button size="mini" type="primary" @click="dialog_confirm()">
              确 定
            </el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
    <AbilityList
      v-if="sub_dialogVisible"
      :sub_dialogVisible="sub_dialogVisible"
      :sub_row="sub_row"
    />
  </div>
</template>
<script>
import { manage_role_get, manage_role_cud } from "@/api/manages/role";
import { AbilityList } from "./components";
import config from "@/utils/config";
import { ManageRolePage } from "@/utils/auth";
import { btn_check } from "@/api/btn";
export default {
  name: "ManagesRole",
  components: {
    AbilityList,
  },
  data() {
    return {
      formInline: {
        name: "",
      },
      role_list: [],
      role: {
        id: 0,
        name: "",
        user_counts: 0,
        desc: "",
      },
      dialogVisible: false,
      dialogTitle: "",
      sub_dialogVisible: false,
      sub_role_id: 0,
      sub_row: {},
      now_page: 1,
      count_one_page: 10,
      total: 1,
      is_mobile: config.isMobile,
      loading: true,
      add_btn_show: 0,
      edit_btn_show: 0,
      del_btn_show: 0,
      conn_btn_show: 0,
      down_btn_show: 0,
      ManageRolePage: ManageRolePage
    };
  },
  mounted() {
    this.init_data();

    const that = this;
    // Vue.config.keyCodes.enter = 13;
    document.onkeydown = function (e) {
      // 这里把事件注册在document上就是因为防止聚焦
      const key = window.event.keyCode;
      if (key === 13) {
        manage_role_get(that.formInline).then((response) => {
          if (response.code === 200) {
            that.role_list = response.data;
          } else {
            that.role_list = [];
          }
        });
      }
    };
  },
  methods: {
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
    handleCurrentChange(val) {
      this.now_page = val;
      this.init_data();
    },
    handleClose(done) {
      this.$confirm("确认关闭？")
        .then((_) => {
          done();
        })
        .catch((_) => {});
    },
    init_data() {
      btn_check(this.ManageRolePage).then(response=>{
        if(response.code === 200){
          this.add_btn_show = response.data.add;
          this.edit_btn_show = response.data.edit;
          this.del_btn_show = response.data.del;
        }
      })
      const query = {
        page_no: this.now_page,
        page_size: this.count_one_page,
      };
      manage_role_get(query).then((response) => {
        if (response.code === 200) {
          this.role_list = response.data.roles;
          this.total = response.data.total;
        } else {
          this.role_list = [];
        }
        this.loading = false;
      });
    },
    searchRole() {
      const query = {
        name: this.formInline.name,
        page_no: 1,
        page_size: this.count_one_page,
      };
      manage_role_get(query).then((response) => {
        if (response.code === 200) {
          this.role_list = response.data;
        } else {
          this.role_list = [];
        }
      });
    },
    create_btn() {
      this.dialogTitle = "创建角色";
      this.role = {
        name: "",
        desc: "",
      };
      this.dialogVisible = true;
    },
    edit_btn(row) {
      this.role = row;
      this.dialogTitle = "更新角色";
      this.dialogVisible = true;
    },
    del_btn(row) {
      manage_role_cud("delete", row).then((response) => {
        this.response_refresh_func(response);
      });
    },
    dialog_confirm() {
      if (this.dialogTitle === "更新角色") {
        manage_role_cud("put", this.role).then((response) => {
          this.response_func(response);
        });
      } else if (this.dialogTitle === "创建角色") {
        manage_role_cud("post", this.role).then((response) => {
          this.response_refresh_func(response);
        });
      }
    },
    dialog_cancel() {
      this.role = {
        id: 0,
        name: "",
        user_counts: 0,
        desc: "",
      };
      this.dialogVisible = false;
    },
    ability_btn(row) {
      this.sub_row = row;
      this.sub_dialogVisible = true;
    },
  },
};
</script>
<style lang="scss" scoped>
.manages-role {
  width: 99%;
  min-height: 680px;
  margin: 0 auto;
  .box-card {
    margin: 15px auto;
  }
}
</style>
