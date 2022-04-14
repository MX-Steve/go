<template>
  <div class="machine">
    <div v-if="!is_mobile">
      <el-tabs v-model="activeName">
        <el-tab-pane label="单台机器导入" name="single">
          <el-form
            ref="machine"
            :model="machine"
            :rules="machine_rules"
            label-width="120px"
            style="width: 75%; margin: 15px auto"
            class="demo-ruleForm"
          >
            <el-row>
              <el-col :span="6">
                <el-form-item label="区域名称">
                  <el-select
                    v-model="machine.zone_id"
                    clearable
                    placeholder="请选择区域"
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
              <el-col :span="6">
                <el-form-item label="机房名称">
                  <el-select
                    v-model="machine.idc_id"
                    style="width: 100%"
                    placeholder="请选择机房"
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
              <el-col :span="6">
                <el-form-item label="设备类型">
                  <el-select
                    v-model="machine.server_type"
                    placeholder="请选择设备类型"
                    @change="type_change"
                  >
                    <el-option
                      v-for="item in server_types"
                      :key="item.id"
                      :label="item.name"
                      :value="item.id"
                    />
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="设备状态">
                  <el-select
                    v-model="machine.status_id"
                    placeholder="请选择设备状态"
                  >
                    <el-option
                      v-for="item in device_status_list"
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
              <el-col :span="12">
                <el-form-item label="sn编码" prop="sn_id">
                  <el-input v-model="machine.sn_id" />
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="12">
                <el-form-item label="主机名称" prop="instance_name">
                  <el-input v-model="machine.instance_name" />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="ip地址" prop="ip_address">
                  <el-input v-model="machine.ip_address" />
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="12">
                <el-form-item label="用户名" prop="username">
                  <el-input v-model="machine.username" />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="端口号" prop="port">
                  <el-input v-model="machine.port" />
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="12">
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
              <el-button type="primary" @click="machine_confirm()">
                立即创建
              </el-button>
              <el-button @click="machine_cancel()">取消</el-button>
              <el-button type="primary" @click="show_more()">
                附加信息
              </el-button>
            </el-form-item>
          </el-form>
          <el-dialog
            :title="dialogTitle"
            :visible.sync="dialogVisible"
            width="30%"
            :before-close="handleClose"
          >
            <el-form
              v-if="machine.server_type == 2"
              ref="cloud_ecs"
              :model="cloud_ecs"
              label-width="100px"
            >
              <el-form-item label="实例名称" prop="instance_name">
                <el-input
                  v-model="cloud_ecs.instance_name"
                  aria-placeholder="请输入实例名称"
                />
              </el-form-item>
              <el-form-item label="付费类型" prop="charge_type">
                <el-input
                  v-model="cloud_ecs.charge_type"
                  aria-placeholder="请输入付费类型"
                />
              </el-form-item>
              <el-form-item label="创建时间" prop="create_time">
                <el-input
                  v-model="cloud_ecs.create_time"
                  aria-placeholder="请输入创建时间"
                />
              </el-form-item>
              <el-form-item label="到期时间" prop="expiration_time">
                <el-input
                  v-model="cloud_ecs.expiration_time"
                  aria-placeholder="请输入到期时间"
                />
              </el-form-item>
              <el-form-item>
                <el-button @click="dialog_cancel(2)">取 消</el-button>
                <el-button
                  type="primary"
                  @click="dialog_confirm(2)"
                >确 定</el-button>
              </el-form-item>
            </el-form>
            <el-form
              v-if="machine.server_type == 1"
              ref="physical_server"
              :model="physical_server"
              label-width="100px"
            >
              <el-form-item label="选择机柜" prop="cabinet_id">
                <el-select
                  v-model="physical_server.cabinet_id"
                  placeholder="请选择机柜"
                >
                  <el-option
                    v-for="item in cabinet_list"
                    :key="item.id"
                    :label="item.name"
                    :value="item.id"
                  />
                </el-select>
              </el-form-item>
              <el-form-item>
                <el-button @click="dialog_cancel(1)">取 消</el-button>
                <el-button
                  type="primary"
                  @click="dialog_confirm(1)"
                >确 定</el-button>
              </el-form-item>
            </el-form>
            <el-form
              v-if="machine.server_type == 3"
              ref="virtual_machine"
              :model="virtual_machine"
              label-width="100px"
            >
              <el-form-item label="选择机柜" prop="cabinet_id">
                <el-select
                  v-model="virtual_machine.cabinet_id"
                  placeholder="请选择机柜"
                >
                  <el-option
                    v-for="item in cabinet_list"
                    :key="item.id"
                    :label="item.name"
                    :value="item.id"
                  />
                </el-select>
              </el-form-item>
              <el-form-item label="选择物理机" prop="cabinet_id">
                <el-select
                  v-model="virtual_machine.machine_id"
                  placeholder="请选择物理机"
                >
                  <el-option
                    v-for="item in machine_list"
                    :key="item.id"
                    :label="item.instance_name"
                    :value="item.id"
                  />
                </el-select>
              </el-form-item>
              <el-form-item>
                <el-button @click="dialog_cancel(3)">取 消</el-button>
                <el-button
                  type="primary"
                  @click="dialog_confirm(3)"
                >确 定</el-button>
              </el-form-item>
            </el-form>
          </el-dialog>
        </el-tab-pane>
        <el-tab-pane
          v-if="activeName == 'more'"
          label="Excel批量导入"
          name="more"
        >
          <upload-excel-component
            :on-success="handleSuccess"
            :before-upload="beforeUpload"
          />
        </el-tab-pane>
      </el-tabs>
    </div>
    <div v-if="is_mobile">
      <el-tabs v-model="activeName">
        <el-tab-pane label="单台机器导入" name="single">
          <el-form
            ref="machine"
            :model="machine"
            :rules="machine_rules"
            label-width="120px"
            style="width: 100%;"
            class="demo-ruleForm"
          >
            <el-row>
              <el-col :span="24">
                <el-form-item label="区域名称">
                  <el-select
                    v-model="machine.zone_id"
                    clearable
                    placeholder="请选择区域"
                    style="width: 100%;"
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
                <el-form-item label="设备类型">
                  <el-select
                    v-model="machine.server_type"
                    placeholder="请选择设备类型"
                    style="width: 100%;"
                    @change="type_change"
                  >
                    <el-option
                      v-for="item in server_types"
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
                  >
                    <el-option
                      v-for="item in device_status_list"
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
              <el-col :span="24">
                <el-form-item label="sn编码" prop="sn_id">
                  <el-input v-model="machine.sn_id" />
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="24">
                <el-form-item label="主机名称" prop="instance_name">
                  <el-input v-model="machine.instance_name" />
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item label="ip地址" prop="ip_address">
                  <el-input v-model="machine.ip_address" />
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="24">
                <el-form-item label="用户名" prop="username">
                  <el-input v-model="machine.username" />
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item label="端口号" prop="port">
                  <el-input v-model="machine.port" />
                </el-form-item>
              </el-col>
            </el-row>
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
              <el-button type="primary" @click="machine_confirm()">
                创建
              </el-button>
              <el-button @click="machine_cancel()">取消</el-button>
              <el-button type="primary" @click="show_more()">
                附加
              </el-button>
            </el-form-item>
          </el-form>
          <el-dialog
            :title="dialogTitle"
            :visible.sync="dialogVisible"
            width="100%"
            :before-close="handleClose"
          >
            <el-form
              v-if="machine.server_type == 2"
              ref="cloud_ecs"
              :model="cloud_ecs"
              label-width="100px"
            >
              <el-form-item label="实例名称" prop="instance_name">
                <el-input
                  v-model="cloud_ecs.instance_name"
                  aria-placeholder="请输入实例名称"
                />
              </el-form-item>
              <el-form-item label="付费类型" prop="charge_type">
                <el-input
                  v-model="cloud_ecs.charge_type"
                  aria-placeholder="请输入付费类型"
                />
              </el-form-item>
              <el-form-item label="创建时间" prop="create_time">
                <el-input
                  v-model="cloud_ecs.create_time"
                  aria-placeholder="请输入创建时间"
                />
              </el-form-item>
              <el-form-item label="到期时间" prop="expiration_time">
                <el-input
                  v-model="cloud_ecs.expiration_time"
                  aria-placeholder="请输入到期时间"
                />
              </el-form-item>
              <el-form-item>
                <el-button @click="dialog_cancel(2)">取 消</el-button>
                <el-button
                  type="primary"
                  @click="dialog_confirm(2)"
                >确 定</el-button>
              </el-form-item>
            </el-form>
            <el-form
              v-if="machine.server_type == 1"
              ref="physical_server"
              :model="physical_server"
              label-width="100px"
            >
              <el-form-item label="选择机柜" prop="cabinet_id">
                <el-select
                  v-model="physical_server.cabinet_id"
                  placeholder="请选择机柜"
                >
                  <el-option
                    v-for="item in cabinet_list"
                    :key="item.id"
                    :label="item.name"
                    :value="item.id"
                  />
                </el-select>
              </el-form-item>
              <el-form-item>
                <el-button @click="dialog_cancel(1)">取 消</el-button>
                <el-button
                  type="primary"
                  @click="dialog_confirm(1)"
                >确 定</el-button>
              </el-form-item>
            </el-form>
            <el-form
              v-if="machine.server_type == 3"
              ref="virtual_machine"
              :model="virtual_machine"
              label-width="100px"
            >
              <el-form-item label="选择机柜" prop="cabinet_id">
                <el-select
                  v-model="virtual_machine.cabinet_id"
                  placeholder="请选择机柜"
                >
                  <el-option
                    v-for="item in cabinet_list"
                    :key="item.id"
                    :label="item.name"
                    :value="item.id"
                  />
                </el-select>
              </el-form-item>
              <el-form-item label="选择物理机" prop="cabinet_id">
                <el-select
                  v-model="virtual_machine.machine_id"
                  placeholder="请选择物理机"
                >
                  <el-option
                    v-for="item in machine_list"
                    :key="item.id"
                    :label="item.instance_name"
                    :value="item.id"
                  />
                </el-select>
              </el-form-item>
              <el-form-item>
                <el-button @click="dialog_cancel(3)">取 消</el-button>
                <el-button
                  type="primary"
                  @click="dialog_confirm(3)"
                >确 定</el-button>
              </el-form-item>
            </el-form>
          </el-dialog>
        </el-tab-pane>
        <el-tab-pane
          v-if="activeName == 'more'"
          label="Excel批量导入"
          name="more"
        >
          <upload-excel-component
            :on-success="handleSuccess"
            :before-upload="beforeUpload"
          />
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>
<script>
import { zone_info_get, device_status_get } from "@/api/assets/basic";
import { idc_get } from "@/api/assets-manage/assets-input";
import { machine_get, machine_cud } from "@/api/assets-manage/assets-input";
import {
  machine_rules,
  server_types,
  os_types,
  authentication_types,
} from "@/utils/assets";
import config from "@/utils/config";
import UploadExcelComponent from "@/components/UploadExcel/index";
export default {
  name: "AssetsMachine",
  components: { UploadExcelComponent },
  data() {
    return {
      activeName: "single",
      dialogTitle: "",
      dialogVisible: false,
      machine_rules: machine_rules,
      server_types: server_types,
      os_types: os_types,
      authentication_types: authentication_types,
      zone_list: [],
      zone: {},
      idc_list: [],
      idc: {},
      device_status_list: [],
      machine_list: [],
      machine: {
        zone_id: "",
        idc_id: "",
        status_id: "",
        sn_id: "",
        instance_name: "",
        server_type: 1,
        ip_address: "",
        os_type: "",
        username: "",
        authentication_type: 1,
        port: 22,
        password: "",
      },
      cloud_ecs: {
        instance_name: "实例名称",
        charge_type: "付费类型",
        create_time: "创建时间",
        expiration_time: "到期时间",
      },
      physical_server: {
        cabinet_id: "机柜id",
      },
      virtual_machine: {
        cabinet_id: "机柜id",
        machine_id: "机器id",
      },
      cabinet_list: [
        { id: 1, name: "机柜1" },
        { id: 2, name: "机柜2" },
      ],
      is_mobile: config.isMobile,
    };
  },
  mounted() {
    this.init_data();
  },
  methods: {
    beforeUpload(file) {
      const isLt1M = file.size / 1024 / 1024 < 1; // 判断传入的数据是否在1M以内  返回布尔值
      if (isLt1M) {
        return true;
      }
      this.$message({
        message: "文件大于1M 请重新上传",
        type: "warning",
      });
      return false;
    }, // 你可以在上传之前做一些自己的特殊判断，如判断文件的大小是否大于 1 兆？若大于 1 兆则停止解析并提示错误信息。
    handleSuccess({ results, header }) {
      console.log("results", results);
      console.log("header", header);
    }, // 解析成功时候会触发的回调函数，它会返回表格的表头和内容。
    async show_more() {
      if (this.machine.server_type == 1) {
        this.dialogTitle = "物理机扩展部分";
      } else if (this.machine.server_type == 2) {
        this.dialogTitle = "云服务器扩展部分";
      } else if (this.machine.server_type == 3) {
        this.dialogTitle = "虚拟机扩展部分";
        var obj = await machine_get({type: "get_all_machines", server_type: 3})
        if(obj.code === 200){
          this.machine_list = obj.data.machines;
        }
      }
      this.dialogVisible = true;
    },
    handleClose(done) {
      this.$confirm("确认关闭？")
        .then((_) => {
          done();
        })
        .catch((_) => {});
    },
    init_data() {
      zone_info_get().then((response) => {
        if (response.code == 200) {
          this.zone_list = response.data;
        } else {
          this.zone_list = [];
        }
      });
      this.type_change();
    },
    type_change() {
      device_status_get({ type_id: this.machine.server_type }).then(
        (response) => {
          if (response.code == 200) {
            this.device_status_list = response.data;
          } else {
            this.device_status_list = [];
          }
        }
      );
    },
    get_idc_by_zone(zone_id) {
      idc_get({ zone_id: zone_id, type: "idc_details" }).then((response) => {
        if (response.code == 200) {
          this.idc_list = response.data;
        } else {
          this.idc_list = [];
        }
      });
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
    dialog_confirm(t) {
      if (t == 1) {
        this.machine.optional = this.physical_server;
      } else if (t == 2) {
        this.machine.optional = this.cloud_ecs;
      } else if (t == 3) {
        this.machine.optional = this.virtual_machine;
      }
      this.dialogVisible = false;
    },
    dialog_cancel(t) {
      if (t == 1) {
        this.physical_server = {
          cabinet_id: "机柜id",
        };
      } else if (t == 2) {
        this.cloud_ecs = {
          instance_name: "",
          charge_type: "",
          create_time: "",
          expiration_time: "",
        };
      } else if (t == 3) {
        this.virtual_machine = {
          cabinet_id: "",
          machine_id: "",
        };
      }
      this.dialogVisible = false;
    },
    machine_confirm() {
      this.machine["type"] = "machine_add";
      machine_cud("post", this.machine).then((response) => {
        this.response_refresh_func(response);
      });
    },
    machine_cancel() {
      this.machine = {
        zone_id: 0,
        idc_id: 0,
        sn_id: "",
        instance_name: "",
        server_type: "",
        ip_address: "",
        os_type: "",
        username: "",
        authentication_type: "",
        port: 0,
        password: "",
      };
    },
  },
};
</script>
<style lang='scss' scoped>
.machine {
  min-height: 450px;
  width: 98%;
  margin: 15px auto;
  background-color: #fff;
  padding: 15px;
  @media screen and (min-width: 1000px) {
    .web {
      display: block;
    }
    .mobile {
      display: none;
    }
  }
  @media screen and (max-width: 500px) {
    .web {
      display: none;
    }
    .mobile {
      display: block;
    }
  }
}
</style>
