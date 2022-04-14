<template>
  <div class="project">
    <div v-if="!is_mobile">
      <div class="sub-title">项目列表</div>
      <div class="plus-button-box">
        <el-button v-if="add_btn_show" size="small" @click="create_btn()">
          添加
        </el-button>
      </div>
      <el-table v-loading="loading" :data="project_list" style="width: 100%">
        <el-table-column prop="id" label="ID" width="55" />
        <el-table-column prop="name" label="项目名" width="180" />
        <el-table-column prop="manager" label="负责人" width="180">
          <template slot-scope="scope">
            <span v-for="(item, index) in scope.row.manager" :key="item">
              {{ get_username(item) }}
              <i v-if="index != scope.row.manager.length - 1"> , </i>
            </span>
          </template>
        </el-table-column>
        <el-table-column prop="c_time" label="创建时间" width="180">
          <template slot-scope="scope">
            {{ scope.row.c_time.replace("T", " ").split(".")[0] }}
          </template>
        </el-table-column>
        <el-table-column prop="remark" label="备注" />
        <el-table-column prop="u_time" label="更新时间" width="180">
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
        :before-close="handleClose"
        @current-change="handleCurrentChange"
      />
      <el-dialog
        :title="dialogTitle"
        :visible.sync="dialogVisible"
        width="35%"
        center
      >
        <el-form ref="project" :model="project" label-width="80px">
          <el-form-item label="名称">
            <el-input v-model="project.name" />
          </el-form-item>
          <el-form-item label="备注">
            <el-input v-model="project.remark" />
          </el-form-item>
          <el-form-item label="负责人">
            <el-select style="width:100%" v-model="project.manager" filterable multiple placeholder="请输入">
              <el-option
                v-for="item in user_list"
                :key="item.id"
                :label="item.username"
                :value="item.id"
              />
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="dialog_confirm()">确定</el-button>
            <el-button @click="dialog_cancel()">取消</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
    <div v-if="is_mobile">
      <div class="sub-title">项目列表</div>
      <div class="plus-button-box">
        <el-button v-if="add_btn_show" size="small" @click="create_btn()">
          添加
        </el-button>
      </div>
      <el-table v-loading="loading" :data="project_list" style="width: 100%">
        <el-table-column prop="id" label="ID" width="55" fixed />
        <el-table-column prop="name" label="项目名" width="150" fixed />
        <el-table-column prop="manager" label="负责人" width="180">
          <template slot-scope="scope">
            <span v-for="(item, index) in scope.row.manager" :key="item">
              {{ get_username(item) }}
              <i v-if="index != scope.row.manager.length - 1"> , </i>
            </span>
          </template>
        </el-table-column>
        <el-table-column prop="c_time" label="创建时间" width="180">
          <template slot-scope="scope">
            {{ scope.row.c_time.replace("T", " ").split(".")[0] }}
          </template>
        </el-table-column>
        <el-table-column prop="remark" label="备注" />
        <el-table-column prop="u_time" label="更新时间" width="180">
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
        :before-close="handleClose"
        @current-change="handleCurrentChange"
      />
      <el-dialog
        :title="dialogTitle"
        :visible.sync="dialogVisible"
        width="100%"
        center
      >
        <el-form ref="project" :model="project" label-width="80px">
          <el-form-item label="名称">
            <el-input v-model="project.name" />
          </el-form-item>
          <el-form-item label="备注">
            <el-input v-model="project.remark" />
          </el-form-item>
          <el-form-item label="负责人">
            <el-select
              v-model="project.manager"
              style="width: 100%"
              multiple
              filterable
              placeholder="请输入"
            >
              <el-option
                v-for="item in user_list"
                :key="item.id"
                :label="item.username"
                :value="item.id"
              />
            </el-select>
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
import { project_get, project_cud } from "@/api/project-list/index";
import { usernames_get } from "@/api/user";
import { btn_check } from "@/api/btn";
import config from "@/utils/config";
import { ProjectsPage } from "@/utils/auth";
export default {
  name: "Project",
  data() {
    return {
      dialogVisible: false,
      dialogTitle: "",
      project_list: [],
      user_list: [],
      project: {
        id: "",
        name: "",
        remark: "",
        manager: "",
        c_time: "",
        u_time: "",
      },
      page_no: 1,
      page_size: 10,
      total: 1,
      add_btn_show: 0,
      edit_btn_show: 0,
      del_btn_show: 0,
      is_mobile: config.isMobile,
      ProjectsPage: ProjectsPage,
      loading: true
    };
  },
  mounted() {
    this.init_data();
  },
  methods: {
    handleClose() {
      this.project = {
        id: "",
        name: "",
        remark: "",
        manager: "",
        c_time: "",
        u_time: "",
      };
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
        location.replace(location.href.split("?")[0] + "?tab=project");
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
      usernames_get().then((response) => {
        this.user_list = response.data.users;
      });
      this.page_data();
    },
    create_btn() {
      this.dialogVisible = true;
      this.dialogTitle = "项目创建";
    },
    edit_btn(row) {
      this.dialogVisible = true;
      this.dialogTitle = "项目更新";
      this.project = row;
    },
    delete_btn(row) {
      project_cud("delete", row).then((response) => {
        this.response_refresh_func(response);
      });
    },
    dialog_confirm() {
      if (this.project.id === "") {
        project_cud("post", this.project).then((response) => {
          this.response_refresh_func(response);
        });
      } else {
        project_cud("put", this.project).then((response) => {
          this.response_func(response);
        });
      }
    },
    dialog_cancel() {
      this.project = {
        id: "",
        name: "",
        remark: "",
        manager: "",
        c_time: "",
        u_time: "",
      };
      this.dialogVisible = false;
    },
    page_data() {
      const query = {
        type: "get_all_projects",
        page_no: this.page_no,
        page_size: this.page_size,
      };
      project_get(query).then((response) => {
        if (response.code == 200) {
          this.project_list = response.data.project_infos;
          this.total = response.data.total;
          this.loading = false;
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
.project {
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
