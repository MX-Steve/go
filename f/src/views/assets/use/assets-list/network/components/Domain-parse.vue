<template>
  <div class="domain-parse">
    <div v-if="!is_mobile">
      <el-row style="margin-top: 15px">
        <el-col :span="24">
          <el-button
            v-if="add_btn_show !== 0"
            style="float: left"
            type="success"
            @click="add_btn"
          >
            添加记录
          </el-button>
          <el-form style="float: right" :inline="true" :model="formInline">
            <el-form-item>
              <el-input v-model="formInline.RR" placeholder="请输入主机记录" />
            </el-form-item>
            <el-form-item>
              <el-button type="text" @click="RRSearch">查询</el-button>
            </el-form-item>
          </el-form>
        </el-col>
      </el-row>
      <el-table
        v-loading="loading"
        :data="record_list"
        style="width: 100%"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column prop="id" label="ID" width="55" />
        <el-table-column prop="RR" label="主机记录" width="240" />
        <el-table-column prop="type" label="解析类型" width="80" />
        <el-table-column prop="line" label="解析线路" width="100" />
        <el-table-column prop="value" label="记录值" width="250" />
        <el-table-column prop="ttl" label="TTL" width="100">
          <template slot-scope="scope">
            {{ scope.row.ttl }}
          </template>
        </el-table-column>
        <el-table-column prop="status_id" label="状态" width="100">
          <template slot-scope="scope">
            {{ get_name(scope.row.status_id, status_list) }}
          </template>
        </el-table-column>
        <el-table-column prop="remark" label="备注" />
        <el-table-column label="操作" width="180">
          <template slot-scope="scope">
            <el-button
              v-if="edit_btn_show !== 0"
              type="text"
              @click="edit_btn(scope.row)"
            >
              修改
            </el-button>
            <el-button
              v-if="del_btn_show !== 0"
              type="text"
              @click="del_btn(scope.row)"
            >
              删除
            </el-button>
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
      <el-dialog :title="dialogTitle" :visible.sync="dialogVisible" width="45%">
        <el-form
          ref="record"
          :model="record"
          :rules="record_rules"
          label-width="120px"
          style="width: 85%; margin: 15px auto"
          class="demo-ruleForm"
        >
          <el-form-item label="解析路线" prop="line">
            <el-select
              v-model="record.line"
              style="width: 100%"
              placeholder="请选择解析路线"
            >
              <el-option
                v-for="item in line_choices"
                :key="item"
                :label="item"
                :value="item"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="主机记录" prop="RR">
            <el-input v-model="record.RR" placeholder="请输入主机记录" />
          </el-form-item>
          <el-form-item label="记录值" prop="value">
            <el-input v-model="record.value" placeholder="请输入记录值" />
          </el-form-item>
          <el-form-item label="记录类型" prop="type">
            <el-select
              v-model="record.type"
              style="width: 100%"
              placeholder="请选择记录类型"
            >
              <el-option
                v-for="item in type_choices"
                :key="item"
                :label="item"
                :value="item"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="域名解析归属" prop="belong_to">
            <el-select
              v-model="record.belong_to"
              style="width: 100%"
              placeholder="请选择记录类型"
            >
              <el-option
                v-for="item in belongs"
                :key="item"
                :label="item"
                :value="item"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="状态" prop="status_id">
            <el-select
              v-model="record.status_id"
              style="width: 100%"
              placeholder="请选择解析路线"
            >
              <el-option
                v-for="item in status_list"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="dialog_confirm()">
              {{ dialogTitle.indexOf("添加") !== -1 ? "添加" : "更新" }}
            </el-button>
            <el-button @click="dialog_cancel()"> 取消 </el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
    <div v-if="is_mobile">
      <el-row style="margin-top: 15px">
        <el-col :span="24">
          <el-button
            v-if="add_btn_show !== 0"
            style="float: left"
            size="mini"
            type="success"
            @click="add_btn"
          >
            添加记录
          </el-button>
          <el-form style="float: right" :inline="true" :model="formInline">
            <el-form-item>
              <el-input
                v-model="formInline.RR"
                style="width: 100%"
                placeholder="请输入主机记录"
              />
            </el-form-item>
            <el-form-item>
              <el-button type="text" @click="RRSearch">查询</el-button>
            </el-form-item>
          </el-form>
        </el-col>
      </el-row>
      <el-table
        v-loading="loading"
        :data="record_list"
        style="width: 100%"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" fixed />
        <el-table-column prop="id" label="ID" width="55" fixed />
        <el-table-column prop="RR" label="主机记录" width="120" fixed />
        <el-table-column prop="type" label="解析类型" width="80" />
        <el-table-column prop="line" label="解析线路" width="100" />
        <el-table-column prop="value" label="记录值" width="250" />
        <el-table-column prop="ttl" label="TTL" width="100">
          <template slot-scope="scope"> {{ scope.row.ttl }} 分钟 </template>
        </el-table-column>
        <el-table-column prop="status_id" label="状态" width="100">
          <template slot-scope="scope">
            {{ get_name(scope.row.status_id, status_list) }}
          </template>
        </el-table-column>
        <el-table-column prop="remark" label="备注" />
        <el-table-column label="操作" width="180">
          <template slot-scope="scope">
            <el-button
              v-if="edit_btn_show !== 0"
              type="text"
              @click="edit_btn(scope.row)"
            >
              修改
            </el-button>
            <el-button
              v-if="del_btn_show !== 0"
              type="text"
              @click="del_btn(scope.row)"
            >
              删除
            </el-button>
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
      <el-dialog
        :title="dialogTitle"
        :visible.sync="dialogVisible"
        width="100%"
      >
        <el-form
          ref="record"
          :model="record"
          :rules="record_rules"
          label-width="80px"
          style="width: 100%; margin: 15px auto"
          class="demo-ruleForm"
        >
          <el-form-item label="解析路线" prop="line">
            <el-select
              v-model="record.line"
              style="width: 100%"
              placeholder="请选择解析路线"
            >
              <el-option
                v-for="item in line_choices"
                :key="item"
                :label="item"
                :value="item"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="主机记录" prop="RR">
            <el-input v-model="record.RR" placeholder="请输入主机记录" />
          </el-form-item>
          <el-form-item label="记录值" prop="value">
            <el-input v-model="record.value" placeholder="请输入记录值" />
          </el-form-item>
          <el-form-item label="均衡权重" prop="weight">
            <el-input v-model="record.weight" placeholder="请输入均衡权重" />
          </el-form-item>
          <el-form-item label="记录类型" prop="type">
            <el-select
              v-model="record.type"
              style="width: 100%"
              placeholder="请选择记录类型"
            >
              <el-option
                v-for="item in type_choices"
                :key="item"
                :label="item"
                :value="item"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="状态" prop="status_id">
            <el-select
              v-model="record.status_id"
              style="width: 100%"
              placeholder="请选择解析路线"
            >
              <el-option
                v-for="item in status_list"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="dialog_confirm()">
              {{ dialogTitle.indexOf("添加") !== -1 ? "添加" : "更新" }}
            </el-button>
            <el-button @click="dialog_cancel()"> 取消 </el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
  </div>
</template>
<script>
import config from "@/utils/config";
import { btn_check } from "@/api/btn";
import { DomainRecordPage } from "@/utils/auth";
import { device_status_get, device_type_get } from "@/api/assets/basic";
import { record_get, record_cud } from "@/api/assets-list/record";
import { choices_get } from "@/api/approval/choices";
import { timeAgo } from "@/filters";
export default {
  name: "DomainParse",
  data() {
    return {
      page_no: 1,
      page_size: 10,
      total: 1,
      loading: true,
      is_mobile: config.isMobile,
      status_list: [],
      dialogVisible: false,
      dialogTitle: "",
      DomainRecordPage: DomainRecordPage,
      add_btn_show: 0,
      edit_btn_show: 0,
      del_btn_show: 0,
      domain_id: "",
      record_list: [],
      line_choices: [],
      belongs: ["alibaba cloud computing", "CloudFlare"],
      type_choices: ["A", "TXT"],
      record: {
        line: "",
        locked: "",
        RR: "",
        status_id: "",
        remark: "",
        ttl: 0,
        type: "",
        value: "",
        weight: "",
        u_time: "",
        belong_to: "",
      },
      formInline: {
        RR: "",
      },
      record_rules: {
        line: [
          { required: true, message: "请选择解析路线", trigger: "change" },
        ],
        RR: [{ required: true, message: "请输入主机记录", trigger: "blur" }],
        status_id: [
          { required: true, message: "请选择状态", trigger: "change" },
        ],
        type: [
          { required: true, message: "请选择记录类型", trigger: "change" },
        ],
        belong_to: [
          { required: true, message: "请选择域名归属", trigger: "change" },
        ],
        value: [{ required: true, message: "请输入记录值", trigger: "blur" }],
        weight: [{ required: true, message: "请输入权重", trigger: "blur" }],
      },
      multipleSelection: [],
      old_row: {},
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
    get_name(id, list) {
      var name = "";
      for (var i = 0; i < list.length; i++) {
        const item = list[i];
        if (id === item.id) {
          name = item.name;
          break;
        }
      }
      return name;
    },
    async init_data() {
      this.domain_id = this.$route.query.domain_id;
      if (this.domain_id === undefined) {
        this.$message({
          message: "必须提供 domain_id",
          type: "error",
        });
      } else {
        btn_check(this.DomainRecordPage).then((response) => {
          if (response.code === 200) {
            this.add_btn_show = response.data.add;
            this.edit_btn_show = response.data.edit;
            this.del_btn_show = response.data.del;
          }
        });
        const obj = await device_type_get({ name: "Dns" });
        if (obj.code === 200) {
          const type_id = obj.data[0].id;
          const obj2 = await device_status_get({ type_id: type_id });
          if (obj2.code === 200) {
            this.status_list = obj2.data;
          }
        }
        this.page_data();
        choices_get({ table_name: "DnsRecords", column: "line_choices" }).then(
          (response) => {
            if (response.code === 200) {
              this.line_choices = response.data;
            }
          }
        );
      }
    },
    page_data() {
      const query = {
        domain_id: this.domain_id,
        RR: this.formInline.RR,
        page_no: this.page_no,
        page_size: this.page_size,
      };
      record_get(query).then((response) => {
        if (response.code === 200) {
          this.record_list = response.data.record_infos;
          this.total = response.data.total;
          this.loading = false;
        }
      });
    },
    RRSearch() {
      const query = {
        domain_id: this.domain_id,
        RR: this.formInline.RR,
        page_no: 1,
        page_size: this.page_size,
      };
      record_get(query).then((response) => {
        if (response.code === 200) {
          this.record_list = response.data.record_infos;
          this.total = response.data.total;
          this.loading = false;
        }
      });
    },
    handleCurrentChange(val) {
      this.page_no = val;
      this.page_data();
    },
    edit_btn(row) {
      this.record = row;
      this.dialogTitle = "更新记录解析";
      this.dialogVisible = true;
      this.old_row = JSON.parse(JSON.stringify(row));
    },
    del_btn(row) {
      record_cud("delete", { ids: [row.id], domain_id: this.domain_id }).then(
        (response) => {
          if (response.code === 200) {
            this.response_refresh_func(response);
          }
        }
      );
    },
    add_btn() {
      this.dialogTitle = "添加记录解析";
      this.dialogVisible = true;
    },
    handleSelectionChange(val) {
      this.multipleSelection = val;
    },

    dialog_confirm() {
      this.record["domain_id"] = this.domain_id;
      if (this.dialogTitle.indexOf("添加") !== -1) {
        this.record["op_type"] = "record_add";
        record_cud("post", this.record).then((response) => {
          if (response.code === 200) {
            this.response_refresh_func(response);
          }
        });
      } else {
        var result = {};
        for (var key in this.old_row) {
          var item = this.record[key];
          if (this.old_row[key] !== this.record[key]) {
            result[key] = item;
          }
        }
        result["id"] = this.old_row["id"];
        record_cud("put", result).then((response) => {
          if (response.code === 200) {
            this.response_refresh_func(response);
          }
        });
      }
    },
    dialog_cancel() {
      this.dialogVisible = false;
      this.record = {
        line: "",
        locked: "",
        RR: "",
        status_id: 0,
        remark: "",
        ttl: 0,
        type: "",
        value: "",
        weight: "",
        u_time: "",
      };
      this.old_row = {};
      location.reload(0);
    },
  },
};
</script>
<style lang="scss" scoped>
.domain-parse {
  padding: 5px;
  background-color: #fff;
  min-height: 650px;
}
</style>
