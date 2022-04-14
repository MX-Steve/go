<template>
  <div class="env">
    <div v-if="!is_mobile">
      <div class="sub-title">环境列表</div>
      <div class="plus-button-box">
        <el-button v-if="add_btn_show" size="small" @click="create_btn()">
          添加
        </el-button>
      </div>
      <el-table v-loading="loading" :data="env_list" style="width: 100%">
        <el-table-column prop="id" label="ID" width="55" />
        <el-table-column prop="name" label="环境名称" width="180" />
        <el-table-column prop="remark" label="备注" />
        <el-table-column prop="c_time" label="创建时间" width="180">
          <template slot-scope="scope">
            {{ scope.row.c_time.split(".")[0] }}
          </template>
        </el-table-column>
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
        @current-change="handleCurrentChange"
      />
      <el-dialog
        :title="dialogTitle"
        :visible.sync="dialogVisible"
        width="45%"
        center
        :before-close="handleClose"
      >
        <el-form ref="env" :model="env" label-width="80px">
          <el-form-item label="环境名称">
            <el-input v-model="env.name" />
          </el-form-item>
          <el-form-item label="备注">
            <el-input v-model="env.remark" />
          </el-form-item>
          <el-form-item label="job_name">
            <el-input v-model="env.job_name" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="dialog_confirm()">确定</el-button>
            <el-button @click="dialog_cancel()">取消</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
    <div v-if="is_mobile">
      <div class="sub-title">环境列表</div>
      <div class="plus-button-box">
        <el-button v-if="add_btn_show" size="small" @click="create_btn()">
          添加
        </el-button>
      </div>
      <el-table v-loading="loading" :data="env_list" style="width: 100%">
        <el-table-column prop="id" label="ID" width="40" fixed />
        <el-table-column prop="name" label="环境名称" width="150" fixed />
        <el-table-column prop="remark" label="备注" />
        <el-table-column prop="c_time" label="创建时间" width="180">
          <template slot-scope="scope">
            {{ scope.row.c_time.split(".")[0] }}
          </template>
        </el-table-column>
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
        @current-change="handleCurrentChange"
      />
      <el-dialog
        :title="dialogTitle"
        :visible.sync="dialogVisible"
        width="100%"
        center
        :before-close="handleClose"
      >
        <el-form ref="env" :model="env" label-width="80px">
          <el-form-item label="环境名称">
            <el-input v-model="env.name" />
          </el-form-item>
          <el-form-item label="备注">
            <el-input v-model="env.remark" />
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
import { env_get, env_cud } from "@/api/project-list/index";
import { btn_check } from "@/api/btn";
import config from "@/utils/config";
import { ProjectsPage } from "@/utils/auth";
export default {
  name: "Env",
  data() {
    return {
      dialogVisible: false,
      dialogTitle: "",
      env_list: [],
      page_no: 1,
      page_size: 10,
      total: 1,
      env: {
        id: "",
        name: "",
        remark: "",
        c_time: "",
        u_time: "",
        job_name: "",
      },
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
    handleClose() {
      this.env = {
        id: "",
        name: "",
        remark: "",
        c_time: "",
        u_time: "",
        job_name: "",
      };
      this.dialogVisible = false;
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
        location.replace(location.href.split("?")[0] + "?tab=env");
      }
    },
    init_data() {
      btn_check(this.ProjectsPage).then((response) => {
        if (response.code === 200) {
          this.add_btn_show = response.data.add;
          this.edit_btn_show = response.data.edit;
          this.del_btn_show = response.data.del;
        }
      });
      this.page_data();
    },
    create_btn() {
      this.dialogTitle = "环境创建";
      this.dialogVisible = true;
    },
    edit_btn(row) {
      this.dialogTitle = "环境更新";
      this.env = row;
      this.dialogVisible = true;
    },
    delete_btn(row) {
      env_cud("delete", row).then((response) => {
        if (response.code === 200) {
          this.response_refresh_func(response);
        }
      });
    },
    dialog_confirm() {
      if (this.env.id === "") {
        env_cud("post", this.env).then((response) => {
          this.response_refresh_func(response);
        });
      } else {
        env_cud("put", this.env).then((response) => {
          this.response_func(response);
        });
      }
    },
    dialog_cancel() {
      this.dialogVisible = false;
      this.env = {
        id: "",
        name: "",
        remark: "",
        c_time: "",
        u_time: "",
        job_name: "",
      };
    },
    page_data() {
      const query = {
        type: "get_all_envs",
        page_no: this.page_no,
        page_size: this.page_size,
      };
      env_get(query).then((response) => {
        if (response.code === 200) {
          this.env_list = response.data;
          this.total = 1;
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
.env {
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
