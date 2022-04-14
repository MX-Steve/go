<template>
  <div class="machine-edit">
    <div v-if="!is_mobile">
      <el-dialog
        title="机器信息更新"
        :visible.sync="dialogVisible2"
        width="60%"
        :before-close="dialog_cancel"
      >
        <el-form
          ref="sub_row"
          :model="machine"
          :rules="machine_rules"
          label-width="120px"
          style="width: 85%; margin: 15px auto"
          class="demo-ruleForm"
        >
          <el-row>
            <el-col :span="8">
              <el-form-item label="区域名称">
                <el-select
                  v-model="machine.zone_id"
                  clearable
                  placeholder="请选择区域"
                  :disabled="machine.server_type == 2"
                  @change="get_idc_by_zone(machine.zone_id)"
                >
                  <el-option
                    v-for="item in zone_list"
                    :key="item.id"
                    :label="item.name"
                    :value="item.id"
                  />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label="机房名称">
                <el-select
                  v-model="machine.idc_id"
                  style="width: 100%"
                  placeholder="请选择机房"
                  :disabled="machine.server_type == 2"
                >
                  <el-option
                    v-for="item in idc_list"
                    :key="item.id"
                    :label="item.name"
                    :value="item.id"
                  />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label="设备状态">
                <el-select
                  v-model="machine.status_id"
                  placeholder="请选择设备状态"
                  :disabled="machine.server_type == 2"
                >
                  <el-option
                    v-for="item in device_status"
                    :key="item.id"
                    :label="item.name"
                    :value="item.id"
                  />
                </el-select>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="系统类型">
                <el-select
                  v-model="machine.os_type"
                  placeholder="请选择系统类型"
                  style="width: 100%"
                  :disabled="machine.server_type == 2"
                >
                  <el-option
                    v-for="item in os_types"
                    :key="item.id"
                    :label="item.name"
                    :value="item.id"
                  />
                </el-select>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="主机名称" prop="instance_name">
                <el-input
                  v-model="machine.instance_name"
                  :disabled="machine.server_type == 2"
                />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="ip地址" prop="ip_address">
                <el-input
                  v-model="machine.ip_address"
                  :disabled="machine.server_type == 2"
                />
              </el-form-item>
            </el-col>
          </el-row>
          <el-form-item label="用户名" prop="username">
            <el-input v-model="machine.username" />
          </el-form-item>
          <el-row>
            <el-col :span="12">
              <el-form-item label="密码类型">
                <el-select
                  v-model="machine.authentication_type"
                  placeholder="请选择密码类型"
                >
                  <el-option
                    v-for="item in authentication_types"
                    :key="item.id"
                    :label="item.name"
                    :value="item.id"
                  />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item
                v-if="machine.authentication_type == 1"
                label="用户密码"
                prop="password"
              >
                <el-input v-model="machine.password" show-password />
              </el-form-item>
              <el-form-item v-else label="密钥位置" prop="password">
                <el-input
                  v-model="machine.password"
                  placeholder="~/.ssh/id_rsa.pub"
                />
              </el-form-item>
            </el-col>
          </el-row>
          <el-form-item>
            <el-button type="primary" @click="dialog_confirm()">
              更新
            </el-button>
            <el-button @click="dialog_cancel()">取消</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
    <div v-if="is_mobile">
      <el-dialog
        title="机器信息更新"
        :visible.sync="dialogVisible2"
        width="100%"
      >
        <el-form
          ref="sub_row"
          :model="machine"
          :rules="machine_rules"
          label-width="120px"
          style="width: 85%; margin: 15px auto"
          class="demo-ruleForm"
        >
          <el-row>
            <el-col :span="24">
              <el-form-item label="区域名称">
                <el-select
                  v-model="machine.zone_id"
                  clearable
                  placeholder="请选择区域"
                  style="width: 100%"
                  :disabled="machine.server_type == 2"
                  @change="get_idc_by_zone(machine.zone_id)"
                >
                  <el-option
                    v-for="item in zone_list"
                    :key="item.id"
                    :label="item.name"
                    :value="item.id"
                  />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="机房名称">
                <el-select
                  v-model="machine.idc_id"
                  style="width: 100%"
                  placeholder="请选择机房"
                  :disabled="machine.server_type == 2"
                >
                  <el-option
                    v-for="item in idc_list"
                    :key="item.id"
                    :label="item.name"
                    :value="item.id"
                  />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="设备状态">
                <el-select
                  v-model="machine.status_id"
                  placeholder="请选择设备状态"
                  style="width: 100%"
                  :disabled="machine.server_type == 2"
                >
                  <el-option
                    v-for="item in device_status"
                    :key="item.id"
                    :label="item.name"
                    :value="item.id"
                  />
                </el-select>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="24">
              <el-form-item label="系统类型">
                <el-select
                  v-model="machine.os_type"
                  placeholder="请选择系统类型"
                  style="width: 100%"
                  :disabled="machine.server_type == 2"
                >
                  <el-option
                    v-for="item in os_types"
                    :key="item.id"
                    :label="item.name"
                    :value="item.id"
                  />
                </el-select>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="24">
              <el-form-item label="主机名称" prop="instance_name">
                <el-input
                  v-model="machine.instance_name"
                  :disabled="machine.server_type == 2"
                />
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="ip地址" prop="ip_address">
                <el-input
                  v-model="machine.ip_address"
                  :disabled="machine.server_type == 2"
                />
              </el-form-item>
            </el-col>
          </el-row>
          <el-form-item label="用户名" prop="username">
            <el-input v-model="machine.username" />
          </el-form-item>
          <el-row>
            <el-col :span="24">
              <el-form-item label="密码类型">
                <el-select
                  v-model="machine.authentication_type"
                  placeholder="请选择密码类型"
                  style="width: 100%"
                >
                  <el-option
                    v-for="item in authentication_types"
                    :key="item.id"
                    :label="item.name"
                    :value="item.id"
                  />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item
                v-if="machine.authentication_type == 1"
                label="用户密码"
                prop="password"
              >
                <el-input v-model="machine.password" />
              </el-form-item>
              <el-form-item v-else label="密钥位置" prop="password">
                <el-input
                  v-model="machine.password"
                  placeholder="~/.ssh/id_rsa.pub"
                />
              </el-form-item>
            </el-col>
          </el-row>
          <el-form-item>
            <el-button type="primary" @click="dialog_confirm()">
              更新
            </el-button>
            <el-button @click="dialog_cancel()">取消</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
  </div>
</template>
<script>
import {
  machine_rules,
  server_types,
  os_types,
  authentication_types,
} from "@/utils/assets";
import { zone_info_get, device_status_get } from "@/api/assets/basic";
import { idc_get } from "@/api/assets-manage/assets-input";
import { machine_get, machine_cud } from "@/api/assets-manage/assets-input";
import config from "@/utils/config";
import { MachinePage } from "@/utils/auth";
export default {
  name: "MachineEdit",
  props: ["id", "dialogVisible2"],
  data() {
    return {
      machine_rules: machine_rules,
      device_status_list: [],
      server_types: server_types,
      os_types: os_types,
      authentication_types: authentication_types,
      machine: {
        id: 1,
        zone_id: "",
        idc_id: "",
        status_id: "",
        sn_id: "",
        hostname: "",
        server_type: 1,
        ip_address: "",
        os_type: 1,
        username: "",
        authentication_type: 1,
        port: 0,
        password: "",
      },
      zone_list: [],
      idc_list: [],
      device_status: [],
      is_mobile: config.isMobile,
      MachinePage: MachinePage,
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
      if (response.code === 200) {
        this.dialogVisible = false;
        this.$message({
          message: response.msg,
          type: "success",
        });
        setTimeout(() => {
          location.reload(0);
        }, 1000);
      } else {
        this.$message({
          message: response.msg,
          type: "warning",
        });
      }
    },
    async init_data() {
      this.os_types = os_types;
      var obj = await machine_get({ type: "machine_details", id: this.id })
      if(obj.code === 200){
        this.machine = obj.data[0];
      } else {
        this.machine = {};
      }
      var obj2 = await zone_info_get()
      if(obj2.code === 200){
        this.zone_list = obj2.data;
      }else{
        this.zone_list = [];
      }
      var obj3 = await device_status_get({type_id: this.machine.server_type})
      if(obj3.code === 200){
        this.device_status = obj3.data;
        console.log(this.device_status)
      } else {
        this.device_status = [];
      }
      idc_get({ type: "get_all_idcs" }).then((response) => {
        if (response.code === 200) {
          this.idc_list = response.data;
        } else {
          this.idc_list = [];
        }
      });
    },
    get_idc_by_zone(zone_id) {
      console.log("开始获取idc，通过", zone_id);
      idc_get({ zone_id: zone_id, type: "idc_add" }).then((response) => {
        if (response.code == 200) {
          this.idc_list = response.data;
        } else {
          this.idc_list = [];
        }
      });
    },
    dialog_confirm() {
      machine_cud("put", this.machine).then((response) => {
        this.response_refresh_func(response);
      });
    },
    dialog_cancel() {
      this.dialogVisible2 = false;
      this.machine = {
        id: 1,
        zone_id: "",
        idc_id: "",
        status_id: "",
        sn_id: "",
        hostname: "",
        server_type: 1,
        ip_address: "",
        os_type: 1,
        username: "",
        authentication_type: 1,
        port: 0,
        password: "",
      };
      setTimeout(() => {
        location.reload(0);
      }, 50);
    },
  },
};
</script>

<style lang="scss" scoped>
.machine-edit {
}
</style>
