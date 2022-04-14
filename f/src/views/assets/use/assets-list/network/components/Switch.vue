<template>
  <div class="switch">
    <div v-if="!is_mobile">
      <div class="sub-title">Switch 信息</div>
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
      <el-table v-loading="loading" :data="switch_list" style="width: 100%">
        <el-table-column prop="id" label="ID" width="55" />
        <el-table-column prop="name" label="名称" width="150" />
        <el-table-column prop="description" label="描述" />
        <el-table-column prop="cidr_block" label="ipv4网段" width="150" />
        <el-table-column prop="ip_address_no" label="可用IP数量" width="120" />
        <el-table-column prop="status" label="状态" width="80">
          <template slot-scope="scope">
            {{ get_name(scope.row.status, status_list) }}
          </template>
        </el-table-column>
        <el-table-column label="其他信息" width="100">
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
      <el-dialog title="更多信息" :visible.sync="dialogVisible" width="45%">
        <el-row v-for="(v, k) in specs" :key="k" class="text item">
          <el-col :span="6">{{ kvs[k] }}</el-col>
          <el-col
            v-if="k.indexOf('time') != -1"
            :span="18"
            style="color: green"
          >
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
        <el-form
          ref="switch2"
          :rules="switch2_roles"
          :model="switch2"
          label-width="100px"
        >
          <el-form-item label="交换机名" prop="name">
            <el-input v-model="switch2.name" placeholder="请输入交换机名字" />
          </el-form-item>
          <el-form-item label="所属状态">
            <el-select
              v-model="switch2.status"
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
          <el-form-item label="所属机房">
            <el-select
              v-model="switch2.idc_id"
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
          <el-form-item label="所属网段">
            <el-input v-model="switch2.cidr_block" placeholder="请输入网段" />
          </el-form-item>
          <el-form-item label="可用IP数量">
            <el-input
              v-model.number="switch2.ip_address_no"
              placeholder="请输入可用IP数量"
            />
          </el-form-item>
          <el-form-item label="描述信息">
            <el-input
              v-model="switch2.description"
              placeholder="请输入描述信息"
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
      <div class="sub-title">Switch 信息</div>
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
      <el-table v-loading="loading" :data="switch_list" style="width: 100%">
        <el-table-column prop="id" label="ID" width="55" fixed />
        <el-table-column prop="switch_id" label="交换机ID" width="220" />
        <el-table-column prop="name" label="名称" width="150" fixed />
        <el-table-column prop="description" label="描述" />
        <el-table-column prop="cidr_block" label="ipv4网段" width="150" />
        <el-table-column prop="ip_address_no" label="可用IP数量" width="150" />
        <el-table-column prop="status" label="状态" width="150">
          <template slot-scope="scope">
            {{ get_name(scope.row.status, status_list) }}
          </template>
        </el-table-column>
        <el-table-column label="其他信息" width="100">
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
          <el-col :span="6">{{ kvs[k] }}</el-col>
          <el-col
            v-if="k.indexOf('time') != -1"
            :span="18"
            style="color: green"
          >
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
        <el-form
          ref="switch2"
          :rules="switch2_roles"
          :model="switch2"
          label-width="100px"
        >
          <el-form-item label="交换机名" prop="name">
            <el-input v-model="switch2.name" placeholder="请输入交换机名字" />
          </el-form-item>
          <el-form-item label="所属状态">
            <el-select
              v-model="switch2.status"
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
          <el-form-item label="所属机房">
            <el-select
              v-model="switch2.idc_id"
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
          <el-form-item label="所属网段">
            <el-input v-model="switch2.cidr_block" placeholder="请输入网段" />
          </el-form-item>
          <el-form-item label="可用IP数量">
            <el-input
              v-model.number="switch2.ip_address_no"
              placeholder="请输入可用IP数量"
            />
          </el-form-item>
          <el-form-item label="描述信息">
            <el-input
              v-model="switch2.description"
              placeholder="请输入描述信息"
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
import { idc_get } from "@/api/assets-manage/assets-input";
import { switch_get, switch_cud } from "@/api/assets-list/basic";
import { device_status_get, device_type_get } from "@/api/assets/basic";
import { btn_check } from "@/api/btn";
import { kvs, get_item_name } from "@/utils/assets";
import config from "@/utils/config";
import { NetworkPage } from "@/utils/auth";
export default {
  name: "Switch2",
  data() {
    return {
      dialogVisible: false,
      switch_list: [],
      page_no: 1,
      page_size: 10,
      total: 1,
      switch2: {},
      switch2_roles: {
        name: [
          { required: true, message: "请输入交换机名称", trigger: "blur" },
        ],
      },
      specs: {},
      kvs: kvs,
      status_list: [],
      idc_list: [],
      down_btn_show: 0,
      add_btn_show: 0,
      edit_btn_show: 0,
      del_btn_show: 0,
      cudDialogVisible: false,
      cudDialogTitle: "新增交换机",
      is_mobile: config.isMobile,
      NetworkPage: NetworkPage,
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
      const idcObj = await idc_get({ type: "get_all_idcs" });
      this.idc_list = idcObj.data;
      const ResObj = await switch_get({
        type: "get_all_switch",
        page_no: 1,
        page_size: 999,
      });
      const switchs = ResObj.data.switch_infos;
      for (var i = 0; i < switchs.length; i++) {
        switchs[i]["status"] = this.get_name(
          switchs[i]["status"],
          this.status_list
        );
        switchs[i]["idc_id"] = this.get_name(
          switchs[i]["idc_id"],
          this.idc_list
        );
      }
      const sw = switchs[1];
      const headers = Object.keys(sw);
      const data = this.formatJson(headers, switchs);
      import("@/vendor/Export2Excel").then((excel) => {
        const tHeader = headers;
        excel.export_json_to_excel({
          header: tHeader, // 表头 必填
          data, // 具体数据 必填
          filename: "switchs", // 非必填
          autoWidth: true, // 非必填
          bookType: "xlsx", // 非必填
        });
      });
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
    page_data() {
      const query = {
        type: "get_all_switch",
        page_no: this.page_no,
        page_size: this.page_size,
      };
      switch_get(query).then((response) => {
        if (response.code === 200) {
          this.switch_list = response.data.switch_infos;
          this.total = response.data.total;
          this.loading = false;
        }
      });
    },
    show_more(id) {
      switch_get({ type: "switch_details", id: id }).then((response) => {
        if (response.code === 200) {
          this.switch2 = response.data[0];
          this.specs = {
            resource_group_id: this.switch2.resource_group_id,
            vpc_id: this.switch2.vpc_id,
            is_default: this.switch2.is_default,
            network_acl_id: this.switch2.network_acl_id,
            create_time: this.switch2.create_time,
            ipv6_cidr_block: this.switch2.ipv6_cidr_block,
            u_time: this.switch2.u_time,
          };
          this.dialogVisible = true;
        }
      });
    },
    handleCurrentChange(val) {
      this.page_no = val;
      this.page_data();
    },
    init_data() {
      btn_check(this.NetworkPage).then(
        (response) => {
          if (response.code == 200) {
            this.down_btn_show = response.data.down;
            this.add_btn_show = response.data.add;
            this.edit_btn_show = response.data.edit;
            this.del_btn_show = response.data.del;
          }
        }
      );
      this.page_data();
      device_type_get({ name: "交换机" }).then((response) => {
        if (response.code == 200) {
          const type_id = response.data[0]["id"];
          device_status_get({ type_id: type_id }).then((response) => {
            if (response.code === 200) {
              this.status_list = response.data;
            }
          });
        }
      });
      idc_get({ type: "get_all_idcs" }).then((response) => {
        if (response.code === 200) {
          this.idc_list = response.data;
        }
      });
    },
    get_name(id, list) {
      return get_item_name(id, list);
    },
    add_btn() {
      this.cudDialogTitle = "添加交换机";
      this.cudDialogVisible = true;
    },
    edit_btn(row) {
      this.switch2 = row;
      this.cudDialogTitle = "更新交换机";
      this.cudDialogVisible = true;
    },
    del_btn(row) {
      if (row.switch_id.indexOf("vsw-") !== -1) {
        this.$message({
          message: "非本地交换机不能删除",
          type: "error",
        });
      } else {
        switch_cud("delete", row).then((response) => {
          this.response_refresh_func(response);
        });
      }
    },
    dialog_confirm() {
      if (this.cudDialogTitle.indexOf("添加") !== -1) {
        this.switch2["type"] = "switch_add";
        switch_cud("post", this.switch2).then((response) => {
          this.response_refresh_func(response);
          this.switch2 = {};
          this.cudDialogVisible = false;
        });
      } else {
        if (this.switch2.switch_id.indexOf("vsw-") !== -1) {
          this.$message({
            message: "非本地交换机不能更新",
            type: "error",
          });
        } else {
          switch_cud("put", this.switch2).then((response) => {
            this.response_func(response);
            this.switch2 = {};
            this.cudDialogVisible = false;
          });
        }
      }
    },
    dialog_cancel() {
      this.cudDialogVisible = false;
      this.switch2 = {};
    },
  },
};
</script>

<style lang="scss" scoped>
.switch {
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
