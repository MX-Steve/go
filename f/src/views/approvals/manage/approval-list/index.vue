<template>
  <div class="approval-list">
    <div v-if="!is_mobile" class="web">
      <el-col :span="23" style="text-align: right">
        <el-button v-if="add_btn_show" type="text" @click="go_to_create()">
          创建
        </el-button>
      </el-col>
      <el-table v-loading="loading" :data="approval_list" style="width: 100%">
        <el-table-column prop="id" label="编号" width="55" />
        <el-table-column prop="approval_name" label="流程名称" width="180" />
        <el-table-column prop="approval_code" label="流程编号" width="320" />
        <el-table-column prop="job_name" label="jenkins job" width="250" />
        <el-table-column prop="descriptions" label="描述信息" />
        <el-table-column prop="subscribe" label="订阅" width="80">
          <template slot-scope="scope">
            {{ scope.row.subscribe === 0 ? "否" : "是" }}
          </template>
        </el-table-column>
        <el-table-column prop="old_version" label="V1版本">
          <template slot-scope="scope">
            {{ scope.row.old_version === 0 ? "否" : "是" }}
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button
              v-if="conn_btn_show"
              type="text"
              @click="detail_btn(scope.row)"
            >
              详情
            </el-button>
            <el-button
              v-if="del_btn_show"
              type="text"
              style="color: red"
              @click="del_btn(scope.row)"
            >
              移除
            </el-button>
            <el-button
              v-if="edit_btn_show"
              type="text"
              style="color: #fabb00"
              @click="
                scope.row.old_version === 0
                  ? down_v(scope.row.id)
                  : up_v(scope.row.id)
              "
            >
              {{ scope.row.old_version === 0 ? "降级" : "升级" }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div v-if="is_mobile" class="mobile">
      <el-col :span="24" style="text-align: right">
        <el-button v-if="add_btn_show" type="text" @click="go_to_create()">
          创建
        </el-button>
      </el-col>
      <el-table v-loading="loading" :data="approval_list" style="width: 100%">
        <el-table-column prop="id" label="ID" width="40" fixed />
        <el-table-column
          prop="approval_name"
          label="流程名称"
          width="100"
          fixed
        />
        <el-table-column prop="approval_code" label="流程编号" width="350" />
        <el-table-column prop="job_name" label="jenkins job" width="150" />
        <el-table-column prop="descriptions" label="描述信息" width="200" />
        <el-table-column prop="subscribe" label="是否订阅" width="100">
          <template slot-scope="scope">
            {{ scope.row.subscribe === 0 ? "否" : "是" }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="120">
          <template slot-scope="scope">
            <el-button
              v-if="conn_btn_show"
              type="text"
              @click="detail_btn(scope.row.id)"
            >
              详情
            </el-button>
            <el-button
              v-if="del_btn_show"
              type="text"
              style="color: red"
              @click="del_btn(scope.row)"
            >
              移除
            </el-button>
            <el-button
              v-if="edit_btn_show"
              type="text"
              style="color: #fabb00"
              @click="
                scope.row.old_version === 0
                  ? down_v(scope.row.id)
                  : up_v(scope.row.id)
              "
            >
              {{ scope.row.old_version === 0 ? "降级" : "升级" }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>
<script>
import { approval_get, approval_cud } from "@/api/approval/approval";
import { up_version, down_version } from "@/api/approval/v1";
import { instance_cud } from "@/api/approval/instance";
import { btn_check } from "@/api/btn";
import { ApprovalListPage } from "@/utils/auth";
import config from "@/utils/config";
export default {
  name: "ApprovalList",
  data() {
    return {
      approval_list: [],
      is_mobile: config.isMobile,
      loading: true,
      ApprovalListPage: ApprovalListPage,
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
    up_v(id) {
      up_version({ id: id }).then((response) => {
        if (response.code === 200) {
          this.$message({
            message: response.msg,
            type: "success",
          });
          location.reload(0);
        }
      });
    },
    down_v(id) {
      down_version({ id: id }).then((response) => {
        if (response.code === 200) {
          this.$message({
            message: response.msg,
            type: "success",
          });
          location.reload(0);
        }
      });
    },
    del_btn(row) {
      approval_cud("delete", { id: row.id }).then((response) => {
        if (response.code === 200) {
          this.$message({
            message: response.msg,
            type: "success",
          });
          location.reload(0);
        }
      });
    },
    go_to_create() {
      this.$router.push({
        path: "/approvals/manage/manage",
      });
    },
    init_data() {
      approval_get().then((response) => {
        if (response.code == 200) {
          this.approval_list = response.data;
          this.loading = false;
        }
      });
      btn_check(this.ApprovalListPage).then((response) => {
        if (response.code == 200) {
          this.add_btn_show = response.data.add;
          this.edit_btn_show = response.data.edit;
          this.del_btn_show = response.data.del;
          this.conn_btn_show = response.data.conn;
        }
      });
    },
    detail_btn(row) {
      if (row.old_version === 1) {
        this.$router.push({
          path: "/approvals/manage/manage-old",
          query: {
            id: row.id,
          },
        });
      } else {
        this.$router.push({
          path: "/approvals/manage/manage",
          query: {
            id: row.id,
          },
        });
      }
    },
    use_btn(id) {
      instance_cud("post", { approval_id: id }).then((response) => {
        if (response.code === 200) {
          const instance = response.data[0];
          this.$router.push({
            path: "/approvals/use/start",
            query: {
              id: instance.f_approval_id,
              instance_id: instance.id,
            },
          });
        } else {
          this.$message({
            message: "实例无法创建，请检查",
            type: "error",
          });
        }
      });
    },
  },
};
</script>
<style lang='scss' scoped>
.approval-list {
  min-height: 450px;
  width: 98%;
  margin: 15px auto;
  background-color: #fff;
  padding: 15px;
  .mobile {
    .wait {
      border: 1px solid #ddd;
      text-align: center;
      padding: 5px;
      color: #999;
      border-radius: 8px;
      margin: 15px 0;
    }
  }
}
</style>
