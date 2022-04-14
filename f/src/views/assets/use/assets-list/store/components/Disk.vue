<template>
  <div class="disk">
    <div v-if="!is_mobile">
      <div class="sub-title">Disk 信息</div>
      <el-row style="margin-top: 15px">
        <el-col :span="2" :offset="21">
          <el-button
            v-if="down_btn_show"
            type="success"
            size="mini"
            icon="el-icon-download"
            @click="download_excel"
          />
          <el-button
            v-if="add_btn_show"
            size="mini"
            style="display: none"
            @click="add_btn"
          >
            添加
          </el-button>
        </el-col>
      </el-row>
      <el-row style="margin: 25px auto">
        <el-col :span="3">
          <el-select
            v-model="selects.zone_id"
            placeholder="请选择区域"
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
            v-model="selects.status"
            placeholder="请选择状态"
            @change="server_status_change()"
          >
            <el-option
              v-for="item in status_list"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-col>
      </el-row>
      <el-table v-loading="loading" :data="disk_list" style="width: 100%">
        <el-table-column prop="id" label="ID" width="55" />
        <el-table-column prop="disk_name" label="磁盘名称" />
        <el-table-column prop="disk_type" label="磁盘类型" width="80" />
        <el-table-column prop="category" label="磁盘类别" />
        <el-table-column label="区域" width="120">
          <template slot-scope="scope">
            {{ get_name(scope.row.zone_id, zone_list) }}
          </template>
        </el-table-column>
        <el-table-column label="机房" width="120">
          <template slot-scope="scope">
            {{ get_name(scope.row.idc_id, idc_list) }}
          </template>
        </el-table-column>
        <el-table-column prop="disk_size" label="大小(G)" width="80" />

        <el-table-column prop="run_status" label="运行状态">
          <template slot-scope="scope">
            {{ get_name(scope.row.run_status, status_list) }}
          </template>
        </el-table-column>
        <el-table-column label="过期时间" width="150">
          <template slot-scope="scope">
            {{ scope.row.expired_time.replace("T", " ").replace("Z", "") }}
          </template>
        </el-table-column>
        <el-table-column label="更新时间" width="150">
          <template slot-scope="scope">
            {{
              scope.row.u_time.replace("T", " ").replace("Z", "").split(".")[0]
            }}
          </template>
        </el-table-column>
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
              icon="el-icon-edit"
              size="mini"
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
        <el-row class="text item">
          <el-col :span="6">关联的实例ID</el-col>
          <el-col :span="18" style="color: green">
            : {{ disk.instance_id }}
            <el-button
              type="text"
              @click="go_to_machine_page(disk.instance_id)"
            >
              前往
            </el-button>
          </el-col>
        </el-row>
        <el-row class="text item">
          <el-col :span="6">tags</el-col>
          <el-col :span="18" style="color: green">
            <template v-if="disk.tags">
              <el-row
                v-for="item in disk.tags"
                :key="item['TagKey']"
                style="height: 20px; line-height: 20px"
              >
                <el-col :span="6"> {{ item["TagKey"] }} </el-col>
                <el-col :span="18"> : {{ item["TagValue"] }} </el-col>
              </el-row>
            </template>
          </el-col>
        </el-row>
        <el-row class="text item">
          <el-col :span="6">磁盘创建时间</el-col>
          <el-col :span="18" style="color: green">
            :
            {{
              disk.create_time
                ? disk.create_time.replace("T", " ").replace("Z", "")
                : ""
            }}
          </el-col>
        </el-row>
      </el-dialog>
      <el-dialog
        :title="cudDialogTitle"
        :visible.sync="cudDialogVisible"
        width="45%"
      >
        <el-form
          ref="disk"
          :model="disk"
          :rules="disk_roles"
          label-width="100px"
        >
          <el-form-item label="磁盘名" prop="disk_name">
            <el-input v-model="disk.disk_name" placeholder="请输入磁盘名" />
          </el-form-item>
          <el-form-item label="使用类型">
            <el-select
              v-model="disk.disk_type"
              placeholder="请选择使用类型"
              style="width: 100%"
            >
              <el-option
                v-for="item in disk_type_choices"
                :key="item.id"
                :value="item.id"
                :label="item.name"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="磁盘类型">
            <el-select
              v-model="disk.category"
              placeholder="请选择磁盘类型"
              style="width: 100%"
            >
              <el-option
                v-for="item in category_choices"
                :key="item.id"
                :value="item.id"
                :label="item.name"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="磁盘大小" prop="disk_size">
            <el-input
              v-model.number="disk.disk_size"
              placeholder="请输入磁盘大小"
              style="width: 100%"
            />
          </el-form-item>
          <el-form-item label="所属状态">
            <el-select
              v-model="disk.run_status"
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
              v-model="disk.zone_id"
              placeholder="请选择所属区域"
              style="width: 100%"
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
              v-model="disk.idc_id"
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
          <el-form-item label="所挂机器">
            <el-select
              v-model="disk.instance_id"
              placeholder="请选择挂载机器"
              style="width: 100%"
            >
              <el-option
                v-for="item in machine_list"
                :key="item.instance_id"
                :label="item.instance_name"
                :value="item.instance_id"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="描述信息">
            <el-input v-model="disk.description" placeholder="请输入描述信息" />
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
      <div class="sub-title">Disk 信息</div>
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
      <el-row style="margin: 25px auto">
        <el-col :span="24">
          <el-select
            v-model="selects.zone_id"
            placeholder="请选择区域"
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
            v-model="selects.status"
            placeholder="请选择状态"
            style="width: 100%"
            @change="server_status_change()"
          >
            <el-option
              v-for="item in status_list"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-col>
      </el-row>
      <el-table v-loading="loading" :data="disk_list" style="width: 100%">
        <el-table-column prop="id" label="ID" width="55" />
        <el-table-column prop="disk_id" label="磁盘编号" />
        <el-table-column prop="disk_name" label="磁盘名称" />
        <el-table-column prop="disk_type" label="磁盘类型" />
        <el-table-column prop="category" label="磁盘类别" />
        <el-table-column label="区域" width="120">
          <template slot-scope="scope">
            {{ get_name(scope.row.zone_id, zone_list) }}
          </template>
        </el-table-column>
        <el-table-column label="机房" width="120">
          <template slot-scope="scope">
            {{ get_name(scope.row.idc_id, idc_list) }}
          </template>
        </el-table-column>
        <el-table-column prop="disk_size" label="大小(G)" width="80" />

        <el-table-column prop="run_status" label="运行状态">
          <template slot-scope="scope">
            {{ get_name(scope.row.run_status, status_list) }}
          </template>
        </el-table-column>
        <el-table-column label="过期时间" width="150">
          <template slot-scope="scope">
            {{ scope.row.expired_time.replace("T", " ").replace("Z", "") }}
          </template>
        </el-table-column>
        <el-table-column label="更新时间" width="150">
          <template slot-scope="scope">
            {{
              scope.row.u_time.replace("T", " ").replace("Z", "").split(".")[0]
            }}
          </template>
        </el-table-column>
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
              icon="el-icon-edit"
              size="mini"
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
        <el-row class="text item">
          <el-col :span="6">关联的实例ID</el-col>
          <el-col :span="18" style="color: green">
            : {{ disk.instance_id }}
            <el-button
              type="text"
              @click="go_to_machine_page(disk.instance_id)"
            >
              前往
            </el-button>
          </el-col>
        </el-row>
        <el-row class="text item">
          <el-col :span="6">tags</el-col>
          <el-col :span="18" style="color: green">
            <template v-if="disk.tags">
              <el-row
                v-for="item in disk.tags"
                :key="item['TagKey']"
                style="height: 20px; line-height: 20px"
              >
                <el-col :span="6"> {{ item["TagKey"] }} </el-col>
                <el-col :span="18"> : {{ item["TagValue"] }} </el-col>
              </el-row>
            </template>
          </el-col>
        </el-row>
        <el-row class="text item">
          <el-col :span="6">磁盘创建时间</el-col>
          <el-col :span="18" style="color: green">
            :
            {{
              disk.create_time
                ? disk.create_time.replace("T", " ").replace("Z", "")
                : ""
            }}
          </el-col>
        </el-row>
      </el-dialog>
      <el-dialog
        :title="cudDialogTitle"
        :visible.sync="cudDialogVisible"
        width="45%"
      >
        <el-form
          ref="disk"
          :model="disk"
          :rules="disk_roles"
          label-width="100px"
        >
          <el-form-item label="磁盘名" prop="disk_name">
            <el-input v-model="disk.disk_name" placeholder="请输入磁盘名" />
          </el-form-item>
          <el-form-item label="使用类型">
            <el-select
              v-model="disk.disk_type"
              placeholder="请选择使用类型"
              style="width: 100%"
            >
              <el-option
                v-for="item in disk_type_choices"
                :key="item.id"
                :value="item.id"
                :label="item.name"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="磁盘类型">
            <el-select
              v-model="disk.category"
              placeholder="请选择磁盘类型"
              style="width: 100%"
            >
              <el-option
                v-for="item in category_choices"
                :key="item.id"
                :value="item.id"
                :label="item.name"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="磁盘大小" prop="disk_size">
            <el-input
              v-model.number="disk.disk_size"
              placeholder="请输入磁盘大小"
              style="width: 100%"
            />
          </el-form-item>
          <el-form-item label="所属状态">
            <el-select
              v-model="disk.run_status"
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
              v-model="disk.zone_id"
              placeholder="请选择所属区域"
              style="width: 100%"
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
              v-model="disk.idc_id"
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
          <el-form-item label="所挂机器">
            <el-select
              v-model="disk.instance_id"
              placeholder="请选择挂载机器"
              style="width: 100%"
            >
              <el-option
                v-for="item in machine_list"
                :key="item.instance_id"
                :label="item.instance_name"
                :value="item.instance_id"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="描述信息">
            <el-input v-model="disk.description" placeholder="请输入描述信息" />
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
import { disk_get, disk_cud } from "@/api/assets-list/basic";
import {
  device_status_get,
  device_type_get,
  zone_info_get,
} from "@/api/assets/basic";
import { idc_get, machine_get } from "@/api/assets-manage/assets-input";
import { btn_check } from "@/api/btn";
import { get_item_name } from "@/utils/assets";
import config from "@/utils/config";
import { StorePage } from "@/utils/auth";
export default {
  name: "Disk",
  data() {
    return {
      dialogVisible: false,
      disk_list: [],
      zone_list: [],
      idc_list: [],
      machine_list: [],
      status_list: [],
      page_no: 1,
      page_size: 10,
      total: 0,
      disk: {},
      disk_roles: {
        disk_size: [
          { required: true, message: "请输入磁盘大小", trigger: "blur" },
        ],
      },
      selects: {
        zone_id: "",
        idc_id: "",
        status_id: "",
      },
      down_btn_show: 0,
      add_btn_show: 0,
      edit_btn_show: 0,
      del_btn_show: 0,
      cudDialogVisible: false,
      cudDialogTitle: "新增磁盘",
      disk_type_choices: [
        { id: "system", name: "系统盘" },
        { id: "data", name: "数据盘" },
      ],
      category_choices: [
        { id: "local_ssd", name: "本地固态盘" },
        { id: "local_hdd", name: "本地机械硬盘" },
      ],
      is_mobile: config.isMobile,
      StorePage: StorePage,
      loading: true,
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
      const ResObj = await disk_get({
        type: "get_all_disk",
        page_no: 1,
        page_size: 999,
      });
      const disks = ResObj.data.disk_infos;
      for (var i = 0; i < disks.length; i++) {
        disks[i]["run_status"] = this.get_name(
          disks[i]["run_status"],
          this.status_list
        );
        disks[i]["idc_id"] = this.get_name(disks[i]["idc_id"], this.idc_list);
        disks[i]["zone_id"] = this.get_name(
          disks[i]["zone_id"],
          this.zone_list
        );
      }
      const disk = disks[1];
      const headers = Object.keys(disk);
      const data = this.formatJson(headers, disks);
      import("@/vendor/Export2Excel").then((excel) => {
        const tHeader = headers;
        excel.export_json_to_excel({
          header: tHeader, // 表头 必填
          data, // 具体数据 必填
          filename: "disks", // 非必填
          autoWidth: true, // 非必填
          bookType: "xlsx", // 非必填
        });
      });
    },
    go_to_machine_page(instance_id) {
      this.$router.push({
        path: "/assets/use/list/machine",
        query: { instance_id: instance_id },
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
      device_type_get({ name: "磁盘" }).then((response) => {
        if (response.code == 200) {
          if (response.data.length > 0) {
            const type_id = response.data[0]["id"];
            device_status_get({ type_id: type_id }).then((response) => {
              if (response.code === 200) {
                this.status_list = response.data;
              }
            });
          }
        }
      });
      zone_info_get().then((response) => {
        if (response.code === 200) {
          this.zone_list = response.data;
        }
      });
      idc_get({ type: "get_all_idcs" }).then((response) => {
        if (response.code === 200) {
          this.idc_list = response.data;
        } else {
          this.idc_list = [];
        }
      });
      machine_get({
        type: "get_all_machines",
        page_no: 1,
        page_size: 10000,
      }).then((response) => {
        if (response.code === 200) {
          const machine_list = response.data.machines;
          for (var i = 0; i < machine_list.length; i++) {
            const item = machine_list[i];
            if (item.optional != null) {
              const optional_instance_id = JSON.parse(item.optional)[
                "instance_id"
              ];
              machine_list[i]["instance_id"] = optional_instance_id;
            } else {
              machine_list[i]["instance_id"] = "";
            }
          }
          this.machine_list = machine_list;
        }
      });
    },
    page_data() {
      const query = {
        type: "get_all_disk",
        page_no: this.page_no,
        page_size: this.page_size,
      };
      disk_get(query).then((response) => {
        if (response.code === 200) {
          this.disk_list = response.data.disk_infos;
          this.total = response.data.total;
          this.loading = false;
        }
      });
    },
    show_more(id) {
      disk_get({ type: "disk_details", id: id }).then((response) => {
        if (response.code === 200) {
          this.disk = response.data[0];
          this.dialogVisible = true;
        }
      });
    },
    handleCurrentChange(val) {
      this.page_no = val;
      this.page_data();
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
      disk_get({
        type: "get_all_disk",
        idc_id: this.selects.idc_id,
        page_no: this.page_no,
        page_size: this.page_size,
      }).then((response) => {
        if (response.code === 200) {
          this.disk_list = response.data.disk_infos;
          this.total = response.data.total;
        }
      });
    },
    server_status_change() {
      disk_get({
        type: "get_all_disk",
        idc_id: this.selects.idc_id,
        status_id: this.selects.status,
        page_no: this.page_no,
        page_size: this.page_size,
      }).then((response) => {
        if (response.code === 200) {
          this.disk_list = response.data.disk_infos;
          this.total = response.data.total;
        }
      });
    },
    add_btn() {
      this.cudDialogTitle = "添加磁盘";
      this.cudDialogVisible = true;
    },
    edit_btn(row) {
      disk_get({ type: "disk_details", id: row.id }).then((response) => {
        if (response.code === 200) {
          this.disk = response.data[0];
        }
      });
      this.disk = row;
      this.cudDialogTitle = "更新磁盘";
      this.cudDialogVisible = true;
    },
    del_btn(row) {
      const choices = ",local_ssd,local_hdd,";
      if (choices.indexOf("," + row.category + ",") === -1) {
        this.$message({
          message: "非本地磁盘不能删除",
          type: "error",
        });
      } else {
        disk_cud("delete", row).then((response) => {
          this.response_refresh_func(response);
        });
      }
    },
    dialog_confirm() {
      if (this.cudDialogTitle.indexOf("添加") !== -1) {
        disk_cud("post", this.disk).then((response) => {
          this.response_refresh_func(response);
          this.disk = {};
        });
      } else {
        const choices = ",local_ssd,local_hdd,";
        if (choices.indexOf("," + this.disk.category + ",") === -1) {
          this.$message({
            message: "非本地磁盘不能更新",
            type: "error",
          });
        } else {
          disk_cud("put", this.disk).then((response) => {
            this.response_func(response);
            this.disk = {};
          });
        }
      }
    },
    dialog_cancel() {
      this.cudDialogVisible = false;
      this.disk = {};
    },
  },
};
</script>

<style lang="scss" scoped>
.disk {
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
