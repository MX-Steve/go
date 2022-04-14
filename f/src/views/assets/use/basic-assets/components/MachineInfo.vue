<template>
  <div class="machine-info">
    <div v-if="!is_mobile">
      <div class="sub-title">区域信息</div>
      <div class="plus-button-box">
        <el-button
          v-if="add_btn_show"
          size="small"
          @click="go_to('/assets/manage/basic', 'first')"
        >
          添加
        </el-button>
      </div>
      <el-table :data="zones" style="width: 100%">
        <el-table-column prop="id" label="ID" width="180" />
        <el-table-column prop="name" label="区域名称" />
        <el-table-column prop="description" label="区域描述" />
        <el-table-column label="操作" width="180">
          <template slot-scope="scope">
            <el-button
              v-if="edit_btn_show"
              type="primary"
              size="mini"
              icon="el-icon-edit"
              @click="edit_btn('zone', scope.row)"
            />
            <el-button
              v-if="del_btn_show"
              type="danger"
              size="mini"
              icon="el-icon-delete"
              @click="delete_btn('zone', scope.row)"
            />
          </template>
        </el-table-column>
      </el-table>
      <div class="sub-title">机房信息</div>
      <div class="plus-button-box">
        <el-button
          v-if="add_btn_show"
          size="small"
          @click="go_to('/assets/manage/basic', 'second')"
        >
          添加
        </el-button>
      </div>
      <el-table
        :data="
          idc.slice(
            (idc_page_no - 1) * idc_page_size,
            idc_page_no * idc_page_size
          )
        "
        style="width: 100%"
      >
        <el-table-column prop="id" label="ID" width="180" />
        <el-table-column label="区域">
          <template slot-scope="scope">
            <div>{{ get_zone_name(scope.row.zone_id) }}</div>
          </template>
        </el-table-column>
        <el-table-column prop="name" label="机房名称" />
        <el-table-column prop="provider" label="运营商" />
        <el-table-column prop="manager" label="联系人" />
        <el-table-column prop="phone" label="联系电话" />
        <el-table-column prop="address" label="机房地址" />
        <el-table-column prop="bandwidth" label="机房带宽" />
        <el-table-column prop="ip_range" label="IP地址段" />
        <el-table-column prop="description" label="描述信息" />
        <el-table-column label="操作" width="180">
          <template slot-scope="scope">
            <el-button
              v-if="edit_btn_show"
              type="primary"
              size="mini"
              icon="el-icon-edit"
              @click="edit_btn('idc', scope.row)"
            />
            <el-button
              v-if="del_btn_show"
              type="danger"
              size="mini"
              icon="el-icon-delete"
              @click="delete_btn('idc', scope.row)"
            />
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        small
        layout="prev, pager, next"
        :page-size="idc_page_size"
        :total="idc_total"
        @current-change="handleCurrentChange"
      />
      <div class="sub-title">机房闲置资产信息汇总</div>
      <div class="plus-button-box">
        <el-button
          v-if="add_btn_show"
          size="small"
          @click="create_btn('闲置资产 idle')"
        >
          添加
        </el-button>
      </div>
      <el-table :data="assets" style="width: 100%">
        <el-table-column prop="id" label="ID" width="180" />
        <el-table-column prop="idc" label="所在机房">
          <template slot-scope="scope">
            {{ get_idc_name(scope.row.idc) }}
          </template>
        </el-table-column>
        <el-table-column prop="name" label="资产名称" />
        <el-table-column prop="count" label="数量" />
        <el-table-column prop="recorder" label="登记员" />
        <el-table-column prop="description" label="备注" />
        <el-table-column prop="update_time" label="记录时间" />
        <el-table-column label="操作" width="180">
          <template slot-scope="scope">
            <el-button
              v-if="edit_btn_show"
              type="primary"
              size="mini"
              icon="el-icon-edit"
              @click="edit_btn('闲置资产 idle', scope.row)"
            />
            <el-button
              v-if="del_btn_show"
              type="danger"
              size="mini"
              icon="el-icon-delete"
              @click="delete_btn('idle', scope.row)"
            />
          </template>
        </el-table-column>
      </el-table>
      <el-dialog
        :title="dialogTitle"
        :visible.sync="dialogVisible"
        width="30%"
        :before-close="handleClose"
        center
      >
        <el-form ref="form" :model="form" label-width="100px">
          <div v-for="(item, index) in form" :key="index">
            <el-form-item
              v-if="
                index !== 'id' &&
                  index !== 'del_tag' &&
                  index !== 'recorder' &&
                  index !== 'update_time' &&
                  index !== 'zone_id'
              "
              :label="index == 'name' ? '名称: ' : kvs[index]"
            >
              <el-select
                v-if="
                  index == 'name' &&
                    (dialogTitle == '开始创建 闲置资产 idle' ||
                      dialogTitle == '开始修改 闲置资产 idle')
                "
                v-model="form[index]"
                placeholder="请选择资产名称"
                style="width: 100%"
              >
                <el-option
                  v-for="(item, index) in tags"
                  :key="index"
                  :label="item.name"
                  :value="item.name"
                />
              </el-select>
              <el-select
                v-else-if="index == 'idc'"
                v-model="form[index]"
                placeholder="请选择机房"
                style="width: 100%"
              >
                <el-option
                  v-for="(item, index) in idc"
                  :key="index"
                  :label="item.name"
                  :value="item.id"
                />
              </el-select>
              <el-input v-else v-model="form[index]" />
            </el-form-item>
          </div>
          <el-form-item>
            <el-button type="primary" @click="dialog_confirm()">确定</el-button>
            <el-button @click="dialog_cancel()">取消</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
    <div v-if="is_mobile">
      <div class="sub-title">区域信息</div>
      <div class="plus-button-box">
        <el-button
          v-if="add_btn_show"
          size="small"
          @click="go_to('/assets/manage/basic', 'first')"
        >
          添加
        </el-button>
      </div>
      <el-table :data="zones" style="width: 100%">
        <el-table-column prop="id" label="ID" width="40" />
        <el-table-column prop="name" label="区域名称" width="120" />
        <el-table-column prop="description" label="区域描述" />
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button
              v-if="edit_btn_show"
              type="primary"
              size="mini"
              icon="el-icon-edit"
              @click="edit_btn('zone', scope.row)"
            />
            <el-button
              v-if="del_btn_show"
              type="danger"
              size="mini"
              icon="el-icon-delete"
              @click="delete_btn('zone', scope.row)"
            />
          </template>
        </el-table-column>
      </el-table>
      <div class="sub-title">机房信息</div>
      <div class="plus-button-box">
        <el-button
          v-if="add_btn_show"
          size="small"
          @click="go_to('/assets/manage/basic', 'second')"
        >
          添加
        </el-button>
      </div>
      <el-table
        :data="
          idc.slice(
            (idc_page_no - 1) * idc_page_size,
            idc_page_no * idc_page_size
          )
        "
        style="width: 100%"
      >
        <el-table-column fixed prop="id" label="ID" width="40" />
        <el-table-column fixed label="区域" width="120">
          <template slot-scope="scope">
            <div>{{ get_zone_name(scope.row.zone_id) }}</div>
          </template>
        </el-table-column>
        <el-table-column prop="name" label="机房名称" />
        <el-table-column prop="provider" label="运营商" />
        <el-table-column prop="manager" label="联系人" />
        <el-table-column prop="phone" label="联系电话" />
        <el-table-column prop="address" label="机房地址" />
        <el-table-column prop="bandwidth" label="机房带宽" />
        <el-table-column prop="ip_range" label="IP地址段" />
        <el-table-column prop="description" label="描述信息" />
        <el-table-column label="操作" fixed="right" width="120">
          <template slot-scope="scope">
            <el-button
              v-if="edit_btn_show"
              type="primary"
              size="mini"
              icon="el-icon-edit"
              @click="edit_btn('idc', scope.row)"
            />
            <el-button
              v-if="del_btn_show"
              type="danger"
              size="mini"
              icon="el-icon-delete"
              @click="delete_btn('idc', scope.row)"
            />
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        small
        layout="prev, pager, next"
        :page-size="idc_page_size"
        :total="idc_total"
        @current-change="handleCurrentChange"
      />
      <div class="sub-title">机房闲置资产信息汇总</div>
      <div class="plus-button-box">
        <el-button
          v-if="add_btn_show"
          size="small"
          @click="create_btn('闲置资产 idle')"
        >
          添加
        </el-button>
      </div>
      <el-table :data="assets" style="width: 100%">
        <el-table-column prop="id" fixed label="ID" width="40" />
        <el-table-column prop="idc" label="所在机房">
          <template slot-scope="scope">
            {{ get_idc_name(scope.row.idc) }}
          </template>
        </el-table-column>
        <el-table-column prop="name" label="资产名称" />
        <el-table-column prop="count" label="数量" />
        <el-table-column prop="recorder" label="登记员" />
        <el-table-column prop="description" label="备注" />
        <el-table-column prop="update_time" label="记录时间" />
        <el-table-column fixed="right" label="操作" width="120">
          <template slot-scope="scope">
            <el-button
              v-if="edit_btn_show"
              type="primary"
              size="mini"
              icon="el-icon-edit"
              @click="edit_btn('闲置资产 idle', scope.row)"
            />
            <el-button
              v-if="del_btn_show"
              type="danger"
              size="mini"
              icon="el-icon-delete"
              @click="delete_btn('idle', scope.row)"
            />
          </template>
        </el-table-column>
      </el-table>
      <el-dialog
        :title="dialogTitle"
        :visible.sync="dialogVisible"
        width="100%"
        :before-close="handleClose"
        center
      >
        <el-form ref="form" :model="form" label-width="100px">
          <div v-for="(item, index) in form" :key="index">
            <el-form-item
              v-if="
                index !== 'id' &&
                  index !== 'del_tag' &&
                  index !== 'recorder' &&
                  index !== 'update_time' &&
                  index !== 'zone_id'
              "
              :label="index == 'name' ? '名称: ' : kvs[index]"
            >
              <el-select
                v-if="
                  index == 'name' &&
                    (dialogTitle == '开始创建 闲置资产 idle' ||
                      dialogTitle == '开始修改 闲置资产 idle')
                "
                v-model="form[index]"
                placeholder="请选择资产名称"
                style="width: 100%"
              >
                <el-option
                  v-for="(item, index) in tags"
                  :key="index"
                  :label="item.name"
                  :value="item.name"
                />
              </el-select>
              <el-select
                v-else-if="index == 'idc'"
                v-model="form[index]"
                placeholder="请选择机房"
                style="width: 100%"
              >
                <el-option
                  v-for="(item, index) in idc"
                  :key="index"
                  :label="item.name"
                  :value="item.id"
                />
              </el-select>
              <el-input v-else v-model="form[index]" />
            </el-form-item>
          </div>
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
  zone_info_get,
  zone_info_cud,
  idle_assets_get,
  idle_assets_cud,
  device_type_get,
} from "@/api/assets/basic";
import { idc_get, idc_cud } from "@/api/assets-manage/assets-input";
import { btn_check } from "@/api/btn";
import { kvs } from "@/utils/assets";
import { BasicPage } from "@/utils/auth";
import config from "@/utils/config";
export default {
  name: "MachineInfo",
  data() {
    return {
      kvs: kvs,
      tags: [],
      dialogVisible: false,
      dialogTitle: "",
      form: {},
      zones: [],
      idc: [],
      assets: [],
      idc_page_no: 1,
      idc_page_size: 5,
      idc_total: 1,
      add_btn_show: 0,
      edit_btn_show: 0,
      del_btn_show: 0,
      is_mobile: config.isMobile,
      BasicPage: BasicPage,
    };
  },
  mounted() {
    this.init_data();
  },
  methods: {
    handleCurrentChange(val) {
      this.idc_page_no = val;
    },
    go_to(url, name) {
      this.$router.push({
        path: url,
        query: {
          activeName: name,
        },
      });
    },
    get_zone_name(id) {
      var name = "";
      for (var i = 0; i < this.zones.length; i++) {
        var item = this.zones[i];
        if (item.id === id) {
          name = item.name;
          break;
        }
      }
      return name;
    },
    get_idc_name(id) {
      var name = "";
      for (var i = 0; i < this.idc.length; i++) {
        var item = this.idc[i];
        if (item.id === id) {
          name = item.name;
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
        location.replace(location.href.split("?")[0] + "?tab=first");
      } else {
        this.$message({
          message: response.msg,
          type: "warning",
        });
      }
    },
    init_data() {
      btn_check(this.BasicPage).then((response) => {
        if (response.code == 200) {
          this.add_btn_show = response.data.add;
          this.edit_btn_show = response.data.edit;
          this.del_btn_show = response.data.del;
        }
      });
      zone_info_get().then((response) => {
        if (response.code === 200) {
          this.zones = response.data;
        } else {
          this.zones = [];
        }
      });
      idle_assets_get().then((response) => {
        if (response.code === 200) {
          this.assets = response.data;
        } else {
          this.assets = [];
        }
      });
      device_type_get().then((response) => {
        if (response.code === 200) {
          this.tags = response.data;
        } else {
          this.tags = [];
        }
      });
      idc_get({ type: "get_all_idcs" }).then((response) => {
        if (response.code === 200) {
          this.idc = response.data;
          this.idc_total = this.idc.length;
        } else {
          this.idc = [];
        }
      });
    },
    create_btn(type) {
      this.form = {
        idc: "",
        name: "",
        count: 0,
        description: "",
      };
      this.dialogTitle = "开始创建 " + type;
      this.dialogVisible = true;
    },
    edit_btn(type, row) {
      this.dialogVisible = true;
      this.dialogTitle = "开始修改 " + type;
      this.form = row;
    },
    delete_btn(type, row) {
      if (type === "zone") {
        zone_info_cud("delete", row).then((response) => {
          this.response_refresh_func(response);
        });
      } else if (type === "idc") {
        idc_cud("delete", row).then((response) => {
          this.response_refresh_func(response);
        });
      } else {
        idle_assets_cud("delete", row).then((response) => {
          this.response_refresh_func(response);
        });
      }
    },
    handleClose(done) {
      this.$confirm("确认关闭？")
        .then((_) => {
          done();
        })
        .catch((_) => {});
    },
    dialog_confirm() {
      if (this.dialogTitle.indexOf("zone") !== -1) {
        zone_info_cud("put", this.form).then((response) => {
          this.response_func(response);
        });
      } else if (this.dialogTitle.indexOf("idc") !== -1) {
        idc_cud("put", this.form).then((response) => {
          this.response_func(response);
        });
      } else if (this.dialogTitle.indexOf("修改 闲置资产 idle") !== -1) {
        idle_assets_cud("put", this.form).then((response) => {
          this.response_func(response);
        });
      } else if (this.dialogTitle.indexOf("创建 闲置资产 idle") !== -1) {
        idle_assets_cud("post", this.form).then((response) => {
          this.response_refresh_func(response);
        });
      } else {
        this.$message({
          message: "所选类型不存在",
          type: "warning",
        });
        console.log("类型错误");
      }
    },
    dialog_cancel() {
      this.form = {};
      this.dialogVisible = false;
    },
  },
};
</script>

<style lang="scss" scoped>
.machine-info {
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
