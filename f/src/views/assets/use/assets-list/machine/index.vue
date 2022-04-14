<template>
  <div class="assets-machine-list">
    <div v-if="!is_mobile">
      <div class="sub-title">机器信息</div>
      <el-row style="margin: 25px auto">
        <el-col :span="3">
          <el-select
            v-model="selects.zone_id"
            placeholder="请选择区域"
            clearable
            @change="zone_change()"
          >
            <el-option
              v-for="item in zone_list"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-col>
        <el-col :span="3">
          <el-select
            v-model="selects.idc_id"
            placeholder="请选择机房"
            clearable
            @change="idc_change()"
          >
            <el-option
              v-for="item in idc_list"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-col>
        <el-col :span="3">
          <el-select
            v-model="selects.server_type"
            placeholder="请选择设备类型"
            clearable
            @change="server_type_change()"
          >
            <el-option
              v-for="item in server_types"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-col>
        <el-col :span="3">
          <el-select
            v-model="selects.status"
            placeholder="请选择状态"
            clearable
            @change="server_status_change()"
          >
            <el-option
              v-for="item in device_status"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-col>
        <el-col :span="7" :offset="5">
          <el-row>
            <el-col :span="19">
              <el-input
                v-model="selects.ip"
                placeholder="请输入内网IP"
              />
            </el-col>
            <el-col :span="4" :offset="1">
              <el-button type="primary" @click="ip_change()"> 搜索 </el-button>
            </el-col>
          </el-row>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="2" :offset="22">
          <el-button
            v-if="del_btn_show"
            type="danger"
            size="mini"
            icon="el-icon-delete"
            @click="delete_more()"
          />
          <el-button
            v-if="down_btn_show"
            type="success"
            size="mini"
            icon="el-icon-download"
            @click="download_excel"
          />
        </el-col>
      </el-row>
      <el-table
        ref="multipleTable"
        v-loading="loading"
        tooltip-effect="dark"
        :data="machine_list"
        style="width: 100%"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column prop="id" label="ID" width="55" />
        <el-table-column label="区域">
          <template slot-scope="scope">
            {{ get_name(scope.row.zone_id, zone_list) }}
          </template>
        </el-table-column>
        <el-table-column label="机房">
          <template slot-scope="scope">
            {{ get_name(scope.row.idc_id, idc_list) }}
          </template>
        </el-table-column>
        <el-table-column prop="instance_name" label="主机名" />
        <el-table-column label="设备类型">
          <template slot-scope="scope">
            {{ get_name(scope.row.server_type, server_types) }}
          </template>
        </el-table-column>
        <el-table-column prop="ip_address" label="IP地址" />
        <el-table-column label="操作系统">
          <template slot-scope="scope">
            {{ get_name(scope.row.os_type, os_types) || scope.row.os_type }}
          </template>
        </el-table-column>
        <el-table-column label="机器状态" show-overflow-tooltip>
          <template slot-scope="scope">
            {{ get_name(scope.row.status_id, device_status) }}
          </template>
        </el-table-column>
        <el-table-column label="其他信息" width="100">
          <template slot-scope="scope">
            <el-button
              size="mini"
              icon="el-icon-more"
              @click="show_more(scope.row)"
            />
          </template>
        </el-table-column>
        <el-table-column label="操作" width="220">
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
            <el-button
              v-if="conn_btn_show"
              type="success"
              size="mini"
              @click="conn_btn(scope.row)"
            >
              连接
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
      <el-dialog title="更多信息" :visible.sync="dialogVisible" width="45%">
        <el-row v-for="(v, k) in machine_extra" :key="k" class="text item">
          <el-col v-if="k != 'optional'" :span="6">{{ kvs[k] }}</el-col>
          <el-col v-if="k != 'optional'" :span="18" style="color: green">
            : {{ v }}
          </el-col>
        </el-row>
        <el-row
          v-for="(v1, k1) in machine_extra.optional"
          :key="k1"
          class="text item"
        >
          <el-col :span="6"> {{ kvs[k1] }} </el-col>
          <el-col :span="18"> : {{ v1 }} </el-col>
        </el-row>
      </el-dialog>
    </div>
    <div v-if="is_mobile">
      <div class="sub-title">机器信息</div>
      <el-row style="margin: 25px auto">
        <el-col :span="24">
          <el-select
            v-model="selects.zone_id"
            placeholder="请选择区域"
            clearable
            style="width: 100%"
            @change="zone_change()"
          >
            <el-option
              v-for="item in zone_list"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-col>
        <el-col :span="24">
          <el-select
            v-model="selects.idc_id"
            placeholder="请选择机房"
            clearable
            style="width: 100%"
            @change="idc_change()"
          >
            <el-option
              v-for="item in idc_list"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-col>
        <el-col :span="24">
          <el-select
            v-model="selects.server_type"
            placeholder="请选择设备类型"
            clearable
            style="width: 100%"
            @change="server_type_change()"
          >
            <el-option
              v-for="item in server_types"
              :key="item.id"
              :label="item.name"
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
            @change="server_status_change()"
          >
            <el-option
              v-for="item in device_status"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-col>
        <el-col :span="24">
          <el-row>
            <el-col :span="19">
              <el-input
                v-model="selects.ip"
                placeholder="请输入内网IP"
              />
            </el-col>
            <el-col :span="4" :offset="1">
              <el-button type="primary" size="small" @click="ip_change()">
                搜索
              </el-button>
            </el-col>
          </el-row>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="8" :offset="16">
          <el-button
            v-if="del_btn_show"
            type="danger"
            size="mini"
            icon="el-icon-delete"
            @click="delete_more()"
          />
          <el-button
            v-if="down_btn_show"
            type="success"
            size="mini"
            icon="el-icon-download"
            @click="download_excel"
          />
        </el-col>
      </el-row>
      <el-table
        ref="multipleTable"
        v-loading="loading"
        tooltip-effect="dark"
        :data="machine_list"
        style="width: 100%"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" fixed />
        <el-table-column prop="id" label="ID" width="55" />
        <el-table-column label="区域">
          <template slot-scope="scope">
            {{ get_name(scope.row.zone_id, zone_list) }}
          </template>
        </el-table-column>
        <el-table-column label="机房">
          <template slot-scope="scope">
            {{ get_name(scope.row.idc_id, idc_list) }}
          </template>
        </el-table-column>
        <el-table-column
          prop="instance_name"
          label="主机名"
          fixed
          width="120"
        />
        <el-table-column label="设备类型">
          <template slot-scope="scope">
            {{ get_name(scope.row.server_type, server_types) }}
          </template>
        </el-table-column>
        <el-table-column prop="ip_address" label="IP地址" />
        <el-table-column label="操作系统">
          <template slot-scope="scope">
            {{ get_name(scope.row.os_type, os_types) || scope.row.os_type }}
          </template>
        </el-table-column>
        <el-table-column label="机器状态" show-overflow-tooltip>
          <template slot-scope="scope">
            {{ get_name(scope.row.status_id, device_status) }}
          </template>
        </el-table-column>
        <el-table-column label="其他信息" width="100">
          <template slot-scope="scope">
            <el-button
              size="mini"
              icon="el-icon-more"
              @click="show_more(scope.row)"
            />
          </template>
        </el-table-column>
        <el-table-column label="操作" width="220">
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
            <el-button
              v-if="conn_btn_show"
              type="success"
              size="mini"
              @click="conn_btn(scope.row)"
            >
              连接
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
      <el-dialog title="更多信息" :visible.sync="dialogVisible" width="100%">
        <el-row v-for="(v, k) in machine_extra" :key="k" class="text item">
          <el-col v-if="k != 'optional'" :span="6">{{ kvs[k] }}</el-col>
          <el-col v-if="k != 'optional'" :span="18" style="color: green">
            : {{ v }}
          </el-col>
        </el-row>
        <el-row
          v-for="(v1, k1) in machine_extra.optional"
          :key="k1"
          class="text item"
        >
          <el-col :span="6"> {{ kvs[k1] }} </el-col>
          <el-col :span="18"> : {{ v1 }} </el-col>
        </el-row>
      </el-dialog>
    </div>
    <MachineEdit
      v-if="dialogVisible2"
      :id="machine_id"
      :dialog-visible2="dialogVisible2"
    />
  </div>
</template>
<script>
import Cookies from "js-cookie";
import {
  zone_info_get,
  device_type_get,
  device_status_get,
} from "@/api/assets/basic";
import {
  machine_get,
  machine_cud,
  idc_get,
  go_ssh,
} from "@/api/assets-manage/assets-input";
import { kvs, server_types, os_types, get_item_name } from "@/utils/assets";
import { btn_check } from "@/api/btn";
import { MachineEdit } from "./components";
import config from "@/utils/config";
import { MachinePage } from "@/utils/auth";
import { p_ssh, g_ssh } from "@/api/vssh";

export default {
  name: "AssetsMachineList",
  components: {
    MachineEdit,
  },
  data() {
    return {
      machine_id: 0,
      kvs: kvs,
      dialogVisible: false,
      dialogVisible2: false,
      selects: {
        zone_id: "",
        idc_id: "",
        server_type: 2,
        status: "",
        ip: "",
      },
      zone_list: [],
      idc_list: [],
      server_types: server_types,
      device_status: [],
      machine_list: [],
      machine_extra: {},
      multipleSelection: [],
      now_page: 1,
      count_one_page: 10,
      total: 1,
      sub_row: {},
      os_types: os_types,
      downloadLoading: false,
      add_btn_show: 0,
      edit_btn_show: 0,
      del_btn_show: 0,
      conn_btn_show: 0,
      down_btn_show: 0,
      instance_id: "",
      superuser: false,
      is_mobile: config.isMobile,
      MachinePage: MachinePage,
      loading: true,
    };
  },
  mounted() {
    this.instance_id = this.$route.query.instance_id;
    this.init_data();
  },
  methods: {
    conn_btn(row) {
      machine_get({ type: "machine_details", id: row.id }).then((response) => {
        if (response.code == 200) {
          const machine = response.data[0];
          const obj = {
            ip_address: machine.ip_address,
            password: machine.password,
            port: machine.port,
            username: machine.username,
          };
          if (obj.password === "" || obj.password === null) {
            this.$message({
              message: "密码没有配置",
              type: "warning",
            });
          } else if (obj.port === "" || obj.port === null) {
            this.$message({
              message: "端口没有配置",
              type: "warning",
            });
          } else if (obj.username === "" || obj.username === null) {
            this.$message({
              message: "用户名没有配置",
              type: "warning",
            });
          } else {
            g_ssh({ ip_address: machine.ip_address }).then((response) => {
              if (response.code === 200) {
                if (response.data.length > 0) {
                  const m = response.data[0];
                  window.open(
                    "http://ws.devops.epverse.net/?id=" + m.id,
                    "_blank"
                  );
                } else {
                  p_ssh(obj).then((response) => {
                    if (response.code === 200) {
                      const id = response.data.id;
                      window.open(
                        "http://ws.devops.epverse.net/?id=" + id,
                        "_blank"
                      );
                    }
                  });
                }
              }
            });
          }
        } else {
          this.$message({
            message: response.msg,
            type: "warning",
          });
        }
      });
      // go_ssh({ id: row.id }).then((response) => {
      //   if (response.code === 200) {
      //     this.$router.push({
      //       path: "/projects/tree-main/ssh",
      //       query: { id: row.id },
      //     });
      //   } else {
      //     this.$message({
      //       message: response.msg,
      //       type: "error",
      //     });
      //   }
      // });
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
        setTimeout(() => {
          location.reload(0);
        }, 1000);
      }
    },
    delete_more() {
      const ids = [];
      for (var i = 0; i < this.multipleSelection.length; i++) {
        ids.push(this.multipleSelection[i].id);
      }
      if (ids.length == 0) {
        this.$message({
          message: "未选择要删除的元素",
          type: "warning",
        });
      } else {
        machine_cud("delete", { id: ids }).then((response) => {
          this.response_refresh_func(response);
        });
      }
    },
    get_name(id, list) {
      return get_item_name(id, list);
    },
    init_data() {
      this.superuser = Cookies.get("superuser");
      btn_check(this.MachinePage).then((response) => {
        if (response.code == 200) {
          this.add_btn_show = response.data.add;
          this.edit_btn_show = response.data.edit;
          this.del_btn_show = response.data.del;
          this.down_btn_show = response.data.down;
          this.conn_btn_show = response.data.conn;
        }
      });
      zone_info_get().then((response) => {
        if (response.code === 200) {
          this.zone_list = response.data;
        } else {
          this.zone_list = [];
        }
      });
      device_status_get({ type_id: this.selects.server_type }).then(
        (response) => {
          if (response.code === 200) {
            this.device_status = response.data;
          } else {
            this.device_status = [];
          }
        }
      );
      idc_get({ type: "get_all_idcs" }).then((response) => {
        if (response.code === 200) {
          this.idc_list = response.data;
        } else {
          this.idc_list = [];
        }
      });
      if (this.instance_id) {
        machine_get({
          type: "ins_details",
          instance_id: this.instance_id,
        }).then((response) => {
          if (response.code === 200) {
            this.machine_list = response.data;
            this.total = 1;
            this.loading = false;
          } else {
            this.machine_list = [];
            this.loading = false;
          }
        });
      } else {
        const query = {
          type: "get_all_machines",
          page_no: 1,
          page_size: 10,
        };
        machine_get(query).then((response) => {
          if (response.code === 200) {
            this.machine_list = response.data.machines;
            this.total = response.data.total;
          } else {
            this.machine_list = [];
          }
          this.loading = false;
        });
      }
    },
    zone_change() {
      idc_get({ type: "idc_details", zone_id: this.selects.zone_id }).then(
        (response) => {
          if (response.code === 200) {
            this.idc_list = response.data;
          } else {
            this.idc_list = [];
          }
        }
      );
    },
    idc_change() {
      machine_get({
        type: "get_all_machines",
        idc_id: this.selects.idc_id,
        page_no: 1,
        page_size: 10,
      }).then((response) => {
        if (response.code === 200) {
          this.machine_list = response.data.machines;
          this.total = response.data.total;
        }
      });
    },
    server_type_change() {
      device_status_get({ type_id: this.selects.server_type }).then(
        (response) => {
          if (response.code === 200) {
            this.device_status = response.data;
          } else {
            this.device_status = [];
          }
        }
      );
      machine_get({
        type: "get_all_machines",
        idc_id: this.selects.idc_id,
        server_type: this.selects.server_type,
        page_no: 1,
        page_size: 10,
      }).then((response) => {
        if (response.code === 200) {
          this.machine_list = response.data.machines;
          this.total = response.data.total;
        }
      });
    },
    ip_change() {
      machine_get({
        type: "get_all_machines",
        ip_address: this.selects.ip,
        page_no: 1,
        page_size: 10,
      }).then((response) => {
        if (response.code === 200) {
          this.machine_list = response.data.machines;
          this.total = response.data.total;
        }
      });
    },
    server_status_change() {
      machine_get({
        type: "get_all_machines",
        idc_id: this.selects.idc_id,
        server_type: this.selects.server_type,
        status_id: this.selects.status,
        page_no: 1,
        page_size: 10,
      }).then((response) => {
        if (response.code === 200) {
          this.machine_list = response.data.machines;
          this.total = response.data.total;
        }
      });
    },
    toggleSelection(rows) {
      if (rows) {
        rows.forEach((row) => {
          this.$refs.multipleTable.toggleRowSelection(row);
        });
      } else {
        this.$refs.multipleTable.clearSelection();
      }
    },
    handleSelectionChange(val) {
      this.multipleSelection = val;
    },
    show_more(row) {
      machine_get({ type: "machine_details", id: row.id }).then((response) => {
        if (response.code === 200) {
          this.machine_extra = {
            sn_id: response.data[0].sn_id,
            u_time: response.data[0].u_time,
            port: response.data[0].port,
            optional: response.data[0].optional,
          };
          if (this.superuser === "true") {
            this.machine_extra["username"] = response.data[0].username;
            this.machine_extra["password"] = response.data[0].password;
            this.machine_extra["authentication_type"] =
              response.data[0].authentication_type + " (1: 密码;2: 密钥)";
          }
        }
      });
      this.dialogVisible = true;
    },
    page_data() {
      const query = {
        type: "get_all_machines",
        idc_id: this.selects.idc_id,
        server_type: this.selects.server_type,
        page_no: this.now_page,
        page_size: this.count_one_page,
      };
      machine_get(query).then((response) => {
        if (response.code === 200) {
          this.machine_list = response.data.machines;
          this.total = response.data.total;
        } else {
          this.machine_list = [];
        }
      });
    },
    handleCurrentChange(val) {
      this.now_page = val;
      this.page_data();
    },
    edit_btn(row) {
      console.log(row);
      this.machine_id = row.id;
      this.dialogVisible2 = true;
    },
    del_btn(row) {
      const ids = { id: [row.id] };
      machine_cud("delete", ids).then((response) => {
        this.response_refresh_func(response);
      });
    },
    formatJson(filterVal, jsonData) {
      return jsonData.map((v) => filterVal.map((j) => v[j]));
    },
    async download_excel() {
      const machineResObj = await machine_get({
        type: "get_all_machines",
        page_no: 1,
        page_size: 999,
      });
      const machines = machineResObj.data.machines;
      for (var i = 0; i < machines.length; i++) {
        var item = machines[i];
        var zone = this.get_name(item.zone_id, this.zone_list);
        delete machines[i]["zone_id"];
        machines[i]["zone"] = zone;
        var idc = this.get_name(item.idc_id, this.idc_list);
        delete machines[i]["idc_id"];
        machines[i]["idc"] = idc;
        var status = this.get_name(item.status_id, this.device_status);
        delete machines[i]["status_id"];
        machines[i]["status"] = status;
        machines[i]["server_type"] = this.get_name(
          item.server_type,
          this.server_types
        );
        machines[i]["os_type"] = this.get_name(item.os_type, this.os_types);
        delete machines[i]["username"];
        delete machines[i]["authentication_type"];
        delete machines[i]["password"];
        delete machines[i]["sudo_password"];
      }
      const machine = machines[1];
      const headers = Object.keys(machine);
      const data = this.formatJson(headers, machines);
      import("@/vendor/Export2Excel").then((excel) => {
        const tHeader = headers;
        excel.export_json_to_excel({
          header: tHeader, // 表头 必填
          data, // 具体数据 必填
          filename: "machines", // 非必填
          autoWidth: true, // 非必填
          bookType: "xlsx", // 非必填
        });
      });
    },
  },
};
</script>
<style lang='scss' scoped>
.assets-machine-list {
  width: 99%;
  min-height: 680px;
  margin: 20px auto;
  background-color: #fff;
  padding: 20px;
  .sub-title {
    padding-top: 20px;
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
