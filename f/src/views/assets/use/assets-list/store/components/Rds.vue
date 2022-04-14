<template>
  <div class="rds">
    <div v-if="!is_mobile">
      <div class="sub-title">Rds 信息</div>
      <el-row style="margin-top: 15px">
        <el-col :span="3" :offset="21">
          <el-button
            v-if="down_btn_show"
            type="success"
            size="mini"
            icon="el-icon-download"
            @click="download_excel"
          />
          <el-button v-if="add_btn_show" size="mini" @click="add_btn">
            添加
          </el-button>
        </el-col>
      </el-row>
      <el-table v-loading="loading" :data="rds_list" style="width: 100%">
        <el-table-column prop="id" label="ID" width="50" />
        <el-table-column prop="rds_name" label="名称" width="100" />
        <el-table-column label="状态" width="80">
          <template slot-scope="scope">
            {{ get_name(parseInt(scope.row.run_status), status_list) }}
          </template>
        </el-table-column>
        <el-table-column prop="engine" label="引擎" width="80" />
        <el-table-column prop="engine_version" label="版本" width="60" />
        <el-table-column prop="port" label="端口" width="60" />
        <el-table-column label="可用区" width="100">
          <template slot-scope="scope">
            {{ get_name(parseInt(scope.row.zone_id), zone_list) }}
          </template>
        </el-table-column>
        <el-table-column label="创建时间" width="150">
          <template slot-scope="scope">
            {{ scope.row.create_time.replace("T", " ").replace("Z", "") }}
          </template>
        </el-table-column>
        <el-table-column label="过期时间" width="150">
          <template slot-scope="scope">
            {{ scope.row.expiration_time.replace("T", " ").replace("Z", "") }}
          </template>
        </el-table-column>
        <el-table-column label="更新时间" width="150">
          <template slot-scope="scope">
            {{
              scope.row.u_time.replace("T", " ").replace("Z", "").split(".")[0]
            }}
          </template>
        </el-table-column>
        <el-table-column prop="conn_address" label="连接地址" />
        <el-table-column label="其他" width="80">
          <template slot-scope="scope">
            <el-button
              size="mini"
              icon="el-icon-more"
              @click="show_more(scope.row.id)"
            />
          </template>
        </el-table-column>
        <el-table-column label="操作" width="120">
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
              @click="del_btn(scope.row)"
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
      <el-dialog title="更多信息" :visible.sync="dialogVisible" width="45%">
        <el-row v-for="(v, k) in specs" :key="k" class="text item">
          <el-col :span="6">{{ k }}</el-col>
          <el-col v-if="k.indexOf('time') != -1" :span="18" style="color: green">
            : {{ v.replace("T", " ").replace("Z", " ") }}
          </el-col>
          <el-col v-else :span="18" style="color: green">: {{ v }}</el-col>
        </el-row>
      </el-dialog>
      <el-dialog
        :title="cudDialogTitle"
        :visible.sync="cudDialogVisible"
        width="45%"
      >
        <el-form ref="rds" :rules="rds_role" :model="rds" label-width="100px">
          <el-form-item label="数据库名" prop="rds_name">
            <el-input v-model="rds.rds_name" />
          </el-form-item>
          <el-form-item label="所属状态">
            <el-select
              v-model="rds.run_status"
              placeholder="请选择状态"
              style="width: 100%"
            >
              <el-option
                v-for="item in status_list"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="所属区域">
            <el-select
              v-model="rds.zone_id"
              placeholder="请选择区域"
              style="width: 100%"
              @change="zone_change"
            >
              <el-option
                v-for="item in zone_list"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="所属机房">
            <el-select
              v-model="rds.idc_id"
              placeholder="请选择机房"
              style="width: 100%"
            >
              <el-option
                v-for="item in idc_list"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="数据库引擎">
            <el-select
              v-model="rds.engine"
              placeholder="请选择引擎"
              style="width: 100%"
            >
              <el-option
                v-for="item in engine_list"
                :key="item"
                :label="item"
                :value="item"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="数据库版本">
            <el-select
              v-model="rds.engine_version"
              placeholder="请选择引擎"
              style="width: 100%"
            >
              <el-option
                v-for="item in version_list"
                :key="item"
                :label="item"
                :value="item"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="端口号">
            <el-input
              v-model.number="rds.port"
              placeholder="请输入端口号"
              style="width: 100%"
            />
          </el-form-item>
          <el-form-item label="连接地址">
            <el-input
              v-model="rds.conn_address"
              placeholder="请输入连接地址"
              style="width: 100%"
            />
          </el-form-item>
          <el-form-item label="备注标记">
            <el-input
              v-model.number="rds.remark"
              placeholder="请输入备注信息"
              style="width: 100%"
            />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="dialog_confirm">
              {{
                cudDialogTitle.indexOf("添加") !== -1 ? "立即创建" : "立即更新"
              }}
            </el-button>
            <el-button @click="dialog_cancel"> 取消 </el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
    <div v-if="is_mobile">
      <div class="sub-title">Rds 信息</div>
      <el-row style="margin-top: 15px">
        <el-col :span="8" :offset="16">
          <el-button
            v-if="down_btn_show"
            type="success"
            size="mini"
            icon="el-icon-download"
            @click="download_excel"
          />
          <el-button v-if="add_btn_show" size="mini" @click="add_btn">
            添加
          </el-button>
        </el-col>
      </el-row>
      <el-table v-loading="loading" :data="rds_list" style="width: 100%">
        <el-table-column prop="id" label="ID" width="55" fixed />
        <el-table-column prop="rds_name" label="实例名称" width="120" fixed />
        <el-table-column label="运行状态" width="80">
          <template slot-scope="scope">
            {{ get_name(parseInt(scope.row.run_status), status_list) }}
          </template>
        </el-table-column>
        <el-table-column prop="engine" label="引擎" width="80" />
        <el-table-column prop="engine_version" label="引擎版本" width="80" />
        <el-table-column prop="port" label="端口" width="80" />
        <el-table-column label="可用区" width="120">
          <template slot-scope="scope">
            {{ get_name(parseInt(scope.row.zone_id), zone_list) }}
          </template>
        </el-table-column>
        <el-table-column label="创建时间" width="150">
          <template slot-scope="scope">
            {{ scope.row.create_time.replace("T", " ").replace("Z", "") }}
          </template>
        </el-table-column>
        <el-table-column label="过期时间" width="150">
          <template slot-scope="scope">
            {{ scope.row.expiration_time.replace("T", " ").replace("Z", "") }}
          </template>
        </el-table-column>
        <el-table-column label="更新时间" width="150">
          <template slot-scope="scope">
            {{
              scope.row.u_time.replace("T", " ").replace("Z", "").split(".")[0]
            }}
          </template>
        </el-table-column>
        <el-table-column prop="conn_address" label="连接地址" />
        <el-table-column label="其他信息">
          <template slot-scope="scope">
            <el-button
              size="mini"
              icon="el-icon-more"
              @click="show_more(scope.row.id)"
            />
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200">
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
              @click="del_btn(scope.row)"
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
      <el-dialog title="更多信息" :visible.sync="dialogVisible" width="100%">
        <el-row v-for="(v, k) in specs" :key="k" class="text item">
          <el-col :span="6">{{ k }}</el-col>
          <el-col v-if="k.indexOf('time') != -1" :span="18" style="color: green">
            : {{ v.replace("T", " ").replace("Z", " ") }}
          </el-col>
          <el-col v-else :span="18" style="color: green">: {{ v }}</el-col>
        </el-row>
      </el-dialog>
      <el-dialog
        :title="cudDialogTitle"
        :visible.sync="cudDialogVisible"
        width="100%"
      >
        <el-form ref="rds" :rules="rds_role" :model="rds" label-width="100px">
          <el-form-item label="数据库名" prop="rds_name">
            <el-input v-model="rds.rds_name" />
          </el-form-item>
          <el-form-item label="所属状态">
            <el-select
              v-model="rds.run_status"
              placeholder="请选择状态"
              style="width: 100%"
            >
              <el-option
                v-for="item in status_list"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="所属区域">
            <el-select
              v-model="rds.zone_id"
              placeholder="请选择区域"
              style="width: 100%"
              @change="zone_change"
            >
              <el-option
                v-for="item in zone_list"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="所属机房">
            <el-select
              v-model="rds.idc_id"
              placeholder="请选择机房"
              style="width: 100%"
            >
              <el-option
                v-for="item in idc_list"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="数据库引擎">
            <el-select
              v-model="rds.engine"
              placeholder="请选择引擎"
              style="width: 100%"
            >
              <el-option
                v-for="item in engine_list"
                :key="item"
                :label="item"
                :value="item"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="数据库版本">
            <el-select
              v-model="rds.engine_version"
              placeholder="请选择引擎"
              style="width: 100%"
            >
              <el-option
                v-for="item in version_list"
                :key="item"
                :label="item"
                :value="item"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="端口号">
            <el-input
              v-model.number="rds.port"
              placeholder="请输入端口号"
              style="width: 100%"
            />
          </el-form-item>
          <el-form-item label="连接地址">
            <el-input
              v-model="rds.conn_address"
              placeholder="请输入连接地址"
              style="width: 100%"
            />
          </el-form-item>
          <el-form-item label="备注标记">
            <el-input
              v-model.number="rds.remark"
              placeholder="请输入备注信息"
              style="width: 100%"
            />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="dialog_confirm">
              {{
                cudDialogTitle.indexOf("添加") !== -1 ? "立即创建" : "立即更新"
              }}
            </el-button>
            <el-button @click="dialog_cancel"> 取消 </el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
  </div>
</template>
<script>
import { rds_get, rds_cud } from "@/api/assets-list/basic";
import {
  device_status_get,
  device_type_get,
  zone_info_get,
} from "@/api/assets/basic";
import { idc_get } from "@/api/assets-manage/assets-input";
import { btn_check } from "@/api/btn";
import { get_item_name } from "@/utils/assets";
import config from "@/utils/config";
import { StorePage } from "@/utils/auth";
export default {
  name: "Rds",
  data() {
    return {
      dialogVisible: false,
      rds_list: [],
      zone_list: [],
      page_no: 1,
      page_size: 10,
      total: 0,
      status_list: [],
      rds: {},
      rds_role: {
        rds_name: [
          { required: true, message: "请输入数据库名称", trigger: "blur" },
        ],
      },
      specs: {},
      down_btn_show: 0,
      add_btn_show: 0,
      edit_btn_show: 0,
      del_btn_show: 0,
      cudDialogVisible: false,
      cudDialogTitle: "数据库",
      engine_list: ["Mariadb", "MySQL"],
      version_list: ["8.0", "5.7", "5.6"],
      idc_list: [],
      is_mobile: config.isMobile,
      StorePage: StorePage,
      loading: true
    };
  },
  mounted() {
    this.init_data();
  },
  methods: {
    formatJson(filterVal, jsonData) {
      return jsonData.map((v) => filterVal.map((j) => v[j]));
    },
    async download_excel() {
      const ResObj = await rds_get({
        type: "get_all_rds",
        page_no: 1,
        page_size: 999,
      });
      const rdsAll = ResObj.data.rds_infos;
      for (var i = 0; i < rdsAll.length; i++) {
        rdsAll[i]["run_status"] = this.get_name(
          rdsAll[i]["run_status"],
          this.status_list
        );
        rdsAll[i]["zone_id"] = this.get_name(
          rdsAll[i]["zone_id"],
          this.zone_list
        );
      }
      const rds = rdsAll[1];
      const headers = Object.keys(rds);
      const data = this.formatJson(headers, rdsAll);
      import("@/vendor/Export2Excel").then((excel) => {
        const tHeader = headers;
        excel.export_json_to_excel({
          header: tHeader, // 表头 必填
          data, // 具体数据 必填
          filename: "rds", // 非必填
          autoWidth: true, // 非必填
          bookType: "xlsx", // 非必填
        });
      });
    },
    show_more(id) {
      rds_get({ type: "rds_details", id: id }).then((response) => {
        if (response.code === 200) {
          this.rds = response.data[0];
          this.dialogVisible = true;
          const specs = this.rds.specs
            .replace("{", "")
            .replace("}", "")
            .split(",");
          for (var i = 0; i < specs.length; i++) {
            var item = specs[i].trim().replaceAll("'", "");
            var kv = item.split(": ");
            this.specs[kv[0]] = kv[1];
          }
          this.specs["max_conn"] = parseInt(this.specs["max_conn"]);
          this.specs["account_max_quantity"] = parseInt(
            this.specs["account_max_quantity"]
          );
          this.specs["cpu"] = parseInt(this.specs["cpu"]);
          this.specs["max_iops"] = parseInt(this.specs["max_iops"]);
          this.specs["memory"] = parseInt(this.specs["memory"]);
        }
      });
    },
    handleCurrentChange(val) {
      this.page_no = val;
      this.page_data();
    },
    page_data() {
      const query = {
        type: "get_all_rds",
        page_no: this.page_no,
        page_size: this.page_size,
      };
      rds_get(query).then((response) => {
        if (response.code === 200) {
          this.rds_list = response.data.rds_infos;
          this.total = response.data.total;
          this.loading = false;
        }
      });
    },
    get_name(id, list) {
      return get_item_name(id, list);
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
        location.reload(0);
      }
    },
    init_data() {
      btn_check(this.StorePage).then((response) => {
        if (response.code == 200) {
          this.down_btn_show = response.data.down;
          this.add_btn_show = response.data.add;
          this.edit_btn_show = response.data.edit;
          this.del_btn_show = response.data.del;
        }
      });
      this.page_data();
      device_type_get({ name: "数据库" }).then((response) => {
        if (response.code == 200) {
          const type_id = response.data[0]["id"];
          device_status_get({ type_id: type_id }).then((response) => {
            if (response.code === 200) {
              this.status_list = response.data;
            }
          });
        }
      });
      zone_info_get().then((response) => {
        if (response.code === 200) {
          this.zone_list = response.data;
        }
      });
    },
    zone_change() {
      idc_get({ type: "idc_details", zone_id: this.rds.zone_id }).then(
        (response) => {
          if (response.code === 200) {
            this.idc_list = response.data;
          }
        }
      );
    },
    add_btn() {
      this.cudDialogTitle = "添加数据库";
      this.cudDialogVisible = true;
    },
    edit_btn(row) {
      this.rds = row;
      this.cudDialogTitle = "更新数据库";
      this.cudDialogVisible = true;
    },
    del_btn(row) {
      if (row.conn_address.indexOf("rm-") !== -1) {
        this.$message({
          message: "非本地数据库不能删除",
          type: "error",
        });
      } else {
        rds_cud("delete", row).then((response) => {
          this.response_refresh_func(response);
        });
      }
    },
    dialog_confirm() {
      if (this.cudDialogTitle.indexOf("添加") !== -1) {
        this.rds["type"] = "rds_add";
        rds_cud("post", this.rds).then((response) => {
          this.response_refresh_func(response);
        });
      } else {
        if (this.rds.conn_address.indexOf("rm-") !== -1) {
          this.$message({
            message: "非本地数据库不能更新",
            type: "error",
          });
        } else {
          rds_cud("put", this.rds).then((response) => {
            this.response_func(response);
            this.cudDialogVisible = true;
          });
        }
      }
    },
    dialog_cancel() {
      this.cudDialogVisible = false;
      this.rds = {};
    },
  },
};
</script>

<style lang="scss" scoped>
.rds {
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
  .text {
    font-size: 14px;
  }

  .item {
    margin-bottom: 18px;
    border-bottom: 1px solid #ccc;
  }
}
</style>
