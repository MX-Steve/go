<template>
  <div class="task">
    <div v-if="!is_mobile">
      <el-tabs v-model="activeName" type="card">
        <el-tab-pane label="脚本执行" name="task">
          <div class="pane-box">
            <el-row>
              <el-col :span="14">
                <div class="sub-title">
                  <div style="color: red; font-weight: bold; float: left">
                    *
                  </div>
                  <div style="padding-left: 18px">脚本执行</div>
                </div>
              </el-col>
              <el-col :span="10" class="history"> </el-col>
            </el-row>
          </div>
        </el-tab-pane>
        <el-tab-pane label="文件分发" name="file">
          <div class="pane-box">文件分发</div>
        </el-tab-pane>
      </el-tabs>
    </div>
    <div v-if="is_mobile"></div>
  </div>
</template>

<script>
import { btn_check } from "@/api/btn";
import config from "@/utils/config";
import { ExecTaskPage } from "@/utils/auth";
export default {
  name: "Task",
  data() {
    return {
      activeName: "task",
      is_mobile: config.isMobile,
      add_btn_show: 0,
      del_btn_show: 0,
      edit_btn_show: 0,
      upload_btn_show: 0,
      down_btn_show: 0,
      ExecTaskPage: ExecTaskPage,
    };
  },
  mounted() {
    this.init_data();
  },
  methods: {
    async init_data() {
      var btn_results = await btn_check(this.ExecTaskPage);
      if (btn_results.code === 200) {
        this.add_btn_show = btn_results.data.add;
        this.del_btn_show = btn_results.data.del;
        this.edit_btn_show = btn_results.data.edit;
        this.upload_btn_show = btn_results.data.conn;
        this.down_btn_show = btn_results.data.down;
      }
    },
  },
};
</script>

<style lang="scss" scoped>
.task {
  width: 99%;
  min-height: 680px;
  margin: 0 auto;
  background-color: #fff;
  .pane-box {
    padding: 2px 20px;
    .sub-title {
      margin-top: 20px;
      font-size: 16px;
      font-weight: bold;
      padding-bottom: 10px;
      border-bottom: 2px solid rgb(226, 221, 221);
    }
  }
}
</style>
