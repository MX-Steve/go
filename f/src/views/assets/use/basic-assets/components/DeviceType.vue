<template>
  <div class="info">
    <div v-if="!is_mobile">
      <div class="sub-title">资产标签</div>
      <div class="plus-button-box">
        <el-button
          v-if="add_btn_show && superuser == 'true'"
          size="small"
          @click="create_btn('资产标签 Tag')"
        >
          添加
        </el-button>
      </div>
      <el-table
        :data="
          tags.slice(
            type_page_size * (type_page_no - 1),
            type_page_size * type_page_no
          )
        "
        style="width: 100%"
      >
        <el-table-column prop="id" label="ID" width="180" />
        <el-table-column prop="name" label="标签名称" />
        <el-table-column label="操作" width="180">
          <template slot-scope="scope">
            <el-button
              v-if="edit_btn_show && superuser == 'true'"
              type="primary"
              size="mini"
              icon="el-icon-edit"
              @click="edit_btn('资产标签 tag', scope.row)"
            />
            <el-button
              v-if="del_btn_show && superuser == 'true'"
              type="danger"
              size="mini"
              icon="el-icon-delete"
              @click="delete_btn('资产标签 tag', scope.row)"
            />
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        small
        layout="prev, pager, next"
        :page-size="type_page_size"
        :total="type_total"
        @current-change="handleCurrentTypeChange"
      />
      <div class="sub-title">资产状态</div>
      <div class="plus-button-box">
        <el-button
          v-if="add_btn_show && superuser == 'true'"
          size="small"
          @click="create_btn('资产状态 status')"
        >
          添加
        </el-button>
      </div>
      <el-table
        :data="
          status.slice(
            status_page_size * (status_page_no - 1),
            status_page_size * status_page_no
          )
        "
        style="width: 100%"
      >
        <el-table-column prop="id" label="ID" width="180" />
        <el-table-column prop="name" label="状态名称" />
        <el-table-column prop="type_id" label="关联设备">
          <template slot-scope="scope">
            {{ get_name(scope.row.type_id, tags) }}
          </template>
        </el-table-column>
        <el-table-column prop="description" label="描述信息" />
        <el-table-column label="操作" width="180">
          <template slot-scope="scope">
            <el-button
              v-if="edit_btn_show && superuser == 'true'"
              type="primary"
              size="mini"
              icon="el-icon-edit"
              @click="edit_btn('资产状态 status', scope.row)"
            />
            <el-button
              v-if="del_btn_show && superuser == 'true'"
              type="danger"
              size="mini"
              icon="el-icon-delete"
              @click="delete_btn('资产状态 status', scope.row)"
            />
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        small
        layout="prev, pager, next"
        :page-size="status_page_size"
        :total="status_total"
        @current-change="handleCurrentStatusChange"
      />
      <el-dialog
        :title="dialogTitle"
        :visible.sync="dialogVisible"
        width="30%"
        :before-close="handleClose"
        center
      >
        <el-form ref="form" :model="form" label-width="80px">
          <el-form-item label="名称">
            <el-input v-model="form.name" />
          </el-form-item>
          <el-form-item
            v-if="dialogTitle.indexOf('status') !== -1"
            label="关联设备"
          >
            <el-select
              v-model="form.type_id"
              placeholder="选择设备"
              style="width: 100%"
            >
              <el-option
                v-for="item in tags"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
          </el-form-item>
          <el-form-item
            v-if="dialogTitle.indexOf('status') !== -1"
            label="描述信息"
          >
            <el-input v-model="form.description" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="dialog_confirm()">确定</el-button>
            <el-button @click="dialog_cancel()">取消</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
    <div v-if="is_mobile">
      <div class="sub-title">资产标签</div>
      <div class="plus-button-box">
        <el-button
          v-if="add_btn_show && superuser == 'true'"
          size="small"
          @click="create_btn('资产标签 Tag')"
        >
          添加
        </el-button>
      </div>
      <el-table
        :data="
          tags.slice(
            type_page_size * (type_page_no - 1),
            type_page_size * type_page_no
          )
        "
        style="width: 100%"
      >
        <el-table-column prop="id" label="ID" width="50" />
        <el-table-column prop="name" label="标签名称" />
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button
              v-if="edit_btn_show && superuser == 'true'"
              type="primary"
              size="mini"
              icon="el-icon-edit"
              @click="edit_btn('资产标签 tag', scope.row)"
            />
            <el-button
              v-if="del_btn_show && superuser == 'true'"
              type="danger"
              size="mini"
              icon="el-icon-delete"
              @click="delete_btn('资产标签 tag', scope.row)"
            />
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        small
        layout="prev, pager, next"
        :page-size="type_page_size"
        :total="type_total"
        @current-change="handleCurrentTypeChange"
      />
      <div class="sub-title">资产状态</div>
      <div class="plus-button-box">
        <el-button
          v-if="add_btn_show && superuser == 'true'"
          size="small"
          @click="create_btn('资产状态 status')"
        >
          添加
        </el-button>
      </div>
      <el-table
        :data="
          status.slice(
            status_page_size * (status_page_no - 1),
            status_page_size * status_page_no
          )
        "
        style="width: 100%"
      >
        <el-table-column prop="id" label="ID" width="40" />
        <el-table-column prop="name" label="状态名称" width="60" />
        <el-table-column prop="type_id" label="关联设备" width="60">
          <template slot-scope="scope">
            {{ get_name(scope.row.type_id, tags) }}
          </template>
        </el-table-column>
        <el-table-column prop="description" label="描述信息" />
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button
              v-if="edit_btn_show && superuser == 'true'"
              type="primary"
              size="mini"
              icon="el-icon-edit"
              @click="edit_btn('资产状态 status', scope.row)"
            />
            <el-button
              v-if="del_btn_show && superuser == 'true'"
              type="danger"
              size="mini"
              icon="el-icon-delete"
              @click="delete_btn('资产状态 status', scope.row)"
            />
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        small
        layout="prev, pager, next"
        :page-size="status_page_size"
        :total="status_total"
        @current-change="handleCurrentStatusChange"
      />
      <el-dialog
        :title="dialogTitle"
        :visible.sync="dialogVisible"
        width="100%"
        :before-close="handleClose"
        center
      >
        <el-form ref="form" :model="form" label-width="80px">
          <el-form-item label="名称">
            <el-input v-model="form.name" />
          </el-form-item>
          <el-form-item
            v-if="dialogTitle.indexOf('status') !== -1"
            label="关联设备"
          >
            <el-select
              v-model="form.type_id"
              placeholder="选择设备"
              style="width: 100%"
            >
              <el-option
                v-for="item in tags"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
          </el-form-item>
          <el-form-item
            v-if="dialogTitle.indexOf('status') !== -1"
            label="描述信息"
          >
            <el-input v-model="form.description" />
          </el-form-item>
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
import Cookies from "js-cookie";
import {
  device_type_get,
  device_type_cud,
  device_status_get,
  device_status_cud,
} from "@/api/assets/basic";
import { btn_check } from "@/api/btn";
import { get_item_name } from "@/utils/assets";
import { BasicPage } from "@/utils/auth";
import config from "@/utils/config";
export default {
  name: "DeviceType",
  data() {
    return {
      form: {
        name: "",
        type: "",
        id: "",
      },
      tags: [],
      status: [],
      dialogVisible: false,
      dialogTitle: "",
      add_btn_show: 0,
      edit_btn_show: 0,
      del_btn_show: 0,
      superuser: "false",
      type_page_no: 1,
      type_page_size: 5,
      type_total: 1,
      status_page_no: 1,
      status_page_size: 5,
      status_total: 1,
      is_mobile: config.isMobile,
      BasicPage: BasicPage,
    };
  },
  mounted() {
    this.init_data();
  },
  methods: {
    handleCurrentTypeChange(val) {
      this.type_page_no = val;
    },
    handleCurrentStatusChange(val) {
      this.status_page_no = val;
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
        location.replace(location.href.split("?")[0] + "?tab=second");
      }
    },
    init_data() {
      this.superuser = Cookies.get("superuser");
      btn_check(BasicPage).then((response) => {
        if (response.code == 200) {
          this.add_btn_show = response.data.add;
          this.edit_btn_show = response.data.edit;
          this.del_btn_show = response.data.del;
        }
      });
      device_type_get().then((response) => {
        if (response.code === 200) {
          this.tags = response.data;
          this.type_total = this.tags.length;
        } else {
          this.tags = [];
        }
      });
      device_status_get().then((response) => {
        if (response.code === 200) {
          this.status = response.data;
          this.status_total = this.status.length;
        } else {
          this.status = [];
        }
      });
    },
    create_btn(type) {
      this.form = {
        name: "",
        type: "",
        id: "",
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
      if (type.indexOf("tag") !== -1) {
        device_type_cud("delete", row).then((response) => {
          this.response_refresh_func(response);
        });
      } else {
        device_status_cud("delete", row).then((response) => {
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
      if (this.dialogTitle.indexOf("创建 资产标签") !== -1) {
        device_type_cud("post", this.form).then((response) => {
          this.response_refresh_func(response);
        });
      } else if (this.dialogTitle.indexOf("创建 资产状态") !== -1) {
        device_status_cud("post", this.form).then((response) => {
          this.response_refresh_func(response);
        });
      } else if (this.dialogTitle.indexOf("修改 资产标签") !== -1) {
        device_type_cud("put", this.form).then((response) => {
          this.response_func(response);
        });
      } else if (this.dialogTitle.indexOf("修改 资产状态") !== -1) {
        device_status_cud("put", this.form).then((response) => {
          this.response_func(response);
        });
      } else {
        this.$message({
          message: "所选类型不存在",
          type: "warning",
        });
      }
    },
    dialog_cancel() {
      this.form = {
        name: "",
        type: "",
        id: "",
      };
      this.dialogVisible = false;
    },
  },
};
</script>

<style lang="scss" scoped>
.info {
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
