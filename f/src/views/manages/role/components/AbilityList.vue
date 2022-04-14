<template>
  <div class="ability-list">
    <div v-if="!is_mobile" class="web">
      <el-dialog
        title="功能权限设置"
        :visible.sync="sub_dialogVisible"
        width="60%"
        :before-close="handleClose"
      >
        <el-table :data="page_perms" style="width: 100%">
          <el-table-column prop="label" label="模块" width="260" />
          <el-table-column label="页面" width="260">
            <template slot-scope="scope">
              <template v-for="(item, index) in scope.row.pages">
                <div
                  :key="index"
                  :style="{
                    height: item.perms.length * 28 + 'px',
                    lineHeight: item.perms.length * 28 + 'px',
                    borderBottom: '2px solid #eee',
                  }"
                >
                  <el-checkbox
                    :key="index"
                    v-model="item.checked"
                    :checked="item.checked"
                    @change="handleCheckAllChange(item, index, scope)"
                  >{{ item.label }}</el-checkbox>
                </div>
              </template>
            </template>
          </el-table-column>
          <el-table-column label="功能">
            <template slot-scope="scope">
              <template v-for="(item, index) in scope.row.pages">
                <div :key="index">
                  <template v-for="(j_item, j_index) in item.perms">
                    <div
                      :key="j_index"
                      :style="{
                        height: '28px',
                        borderBottom:
                          j_index == item.perms.length - 1
                            ? '2px solid #eee'
                            : '',
                      }"
                    >
                      <el-checkbox
                        :key="index + '-' + j_index"
                        :value="j_item.checked"
                        @change="
                          handleCheckPermChange(
                            j_item.checked,
                            index,
                            j_index,
                            scope
                          )
                        "
                      >{{ j_item.label }}</el-checkbox>
                    </div>
                  </template>
                </div>
              </template>
            </template>
          </el-table-column>
        </el-table>
        <span slot="footer" class="dialog-footer">
          <el-button @click="dialog_cancel()">取 消</el-button>
          <el-button type="primary" @click="dialog_confirm()">确 定</el-button>
        </span>
      </el-dialog>
    </div>
    <div v-if="is_mobile" class="mobile">
      <el-dialog
        title="功能权限设置"
        :visible.sync="sub_dialogVisible"
        width="100%"
        :before-close="handleClose"
      >
        <el-table :data="page_perms" style="width: 100%">
          <el-table-column prop="label" label="模块" width="100" />
          <el-table-column label="页面" width="120">
            <template slot-scope="scope">
              <template v-for="(item, index) in scope.row.pages">
                <div
                  :key="index"
                  :style="{
                    height: item.perms.length * 28 + 'px',
                    lineHeight: item.perms.length * 28 + 'px',
                    borderBottom: '2px solid #eee',
                  }"
                >
                  <el-checkbox
                    :key="index"
                    v-model="item.checked"
                    :checked="item.checked"
                    @change="handleCheckAllChange(item, index, scope)"
                  >{{ item.label }}</el-checkbox>
                </div>
              </template>
            </template>
          </el-table-column>
          <el-table-column label="功能">
            <template slot-scope="scope">
              <template v-for="(item, index) in scope.row.pages">
                <div :key="index">
                  <template v-for="(j_item, j_index) in item.perms">
                    <div
                      :key="j_index"
                      :style="{
                        height: '28px',
                        borderBottom:
                          j_index == item.perms.length - 1
                            ? '2px solid #eee'
                            : '',
                      }"
                    >
                      <el-checkbox
                        :key="index + '-' + j_index"
                        :value="j_item.checked"
                        @change="
                          handleCheckPermChange(
                            j_item.checked,
                            index,
                            j_index,
                            scope
                          )
                        "
                      >{{ j_item.label }}</el-checkbox>
                    </div>
                  </template>
                </div>
              </template>
            </template>
          </el-table-column>
        </el-table>
        <span slot="footer" class="dialog-footer">
          <el-button size="mini" @click="dialog_cancel()">取 消</el-button>
          <el-button size="mini" type="primary" @click="dialog_confirm()">确 定</el-button>
        </span>
      </el-dialog>
    </div>
  </div>
</template>
<script>
import { manage_role_cud } from "@/api/manages/role";
import config from "@/utils/config";
export default {
  name: "AbilityList",
  props: {
    sub_dialogVisible: Boolean,
    sub_row: Object,
    required: true,
  },
  data() {
    return {
      page_perms: [],
      user_perms: [],
      isIndeterminate: true,
      dialogVisible: false,
      is_mobile: config.isMobile
    };
  },
  mounted() {
    this.init_data();
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
    init_data() {
      console.log(this.sub_dialogVisible);
      console.log(this.sub_row);
      var page_perms = require("./codes");
      this.page_perms = page_perms.default;
      this.check_init_data();
    },
    check_init_data() {
      if ("page_perms" in this.sub_row && this.sub_row.page_perms !== null) {
        this.user_perms = this.sub_row.page_perms
          .split("[")[1]
          .split("]")[0]
          .split(",");
      } else {
        this.user_perms = [];
      }
      const user_pages = {
        "host.host": 0,
        "dashboard.dashboard": 0,
      };
      for (var i = 0; i < this.user_perms.length; i++) {
        var perm = this.user_perms[i];
        var perms = perm.split(".");
        // model
        for (var j = 0; j < this.page_perms.length; j++) {
          var model = this.page_perms[j];
          if (model.key === perms[0]) {
            // page
            var pages = model.pages;
            for (var p = 0; p < pages.length; p++) {
              var page = pages[p];
              // perm
              if (page.key == perms[1]) {
                var pes = page.perms;
                for (var pe = 0; pe < pes.length; pe++) {
                  var per = pes[pe];
                  if (per.key === perms[2]) {
                    this.$set(
                      this.page_perms[j].pages[p].perms[pe],
                      "checked",
                      true
                    );
                    user_pages[model.key + "." + page.key] += 1;
                    if (user_pages[model.key + "." + page.key] == pes.length) {
                      this.$set(this.page_perms[j].pages[p], "checked", true);
                    }
                  }
                }
              }
            }
          }
        }
      }
    },
    handleClose(done) {
      this.$confirm("确认关闭？")
        .then((_) => {
          location.reload(0);
          done();
        })
        .catch((_) => {});
    },
    dialog_confirm() {
      const user_perms = [];
      for (let i = 0; i < this.page_perms.length; i++) {
        const model = this.page_perms[i];
        const pages = model.pages;
        for (let j = 0; j < pages.length; j++) {
          const page = pages[j];
          const perms = page.perms;
          for (let m = 0; m < perms.length; m++) {
            const perm = perms[m];
            if (perm.checked) {
              user_perms.push(model.key + "." + page.key + "." + perm.key);
            }
          }
        }
      }
      this.sub_dialogVisible = false;
      // location.reload(0)
      const user_perms_str = "[" + user_perms.join(",") + "]";
      this.sub_row.page_perms = user_perms_str;
      manage_role_cud("put", this.sub_row).then((response) => {
        this.response_refresh_func(response);
      });
    },
    dialog_cancel() {
      location.reload(0);
    },
    handleCheckAllChange(item, index, scope) {
      if (item.checked) {
        // item.checked = false;
        for (var i = 0; i < item.perms.length; i++) {
          this.$set(
            this.page_perms[scope.$index].pages[index].perms[i],
            "checked",
            true
          );
        }
      } else {
        for (var i = 0; i < item.perms.length; i++) {
          this.$set(
            this.page_perms[scope.$index].pages[index].perms[i],
            "checked",
            false
          );
        }
      }
    },
    handleCheckPermChange(value, index, j_index, scope) {
      this.$set(
        this.page_perms[scope.$index].pages[index].perms[j_index],
        "checked",
        !value
      );
      this.checkAll(this.page_perms[scope.$index].pages[index]);
    },
    checkAll(page) {
      const perms_length = page.perms.length;
      let checkedCount = 0;
      for (let ii = 0; ii < perms_length; ii++) {
        const i_item = page.perms[ii];
        if (i_item.checked) {
          checkedCount += 1;
        }
      }
      if (checkedCount == perms_length) {
        this.$set(page, "checked", true);
      } else {
        this.$set(page, "checked", false);
      }
    },
  },
};
</script>

<style lang="scss" scoped>
.ability-list {
  padding: 15px auto;
}
</style>
