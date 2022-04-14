<template>
  <div class="oss">
    <div v-if="!is_mobile">
      <div class="sub-title">Oss 信息</div>
      <el-row style="margin-top: 15px">
        <el-col :span="2" :offset="22">
          <el-button
            v-if="down_btn_show"
            type="success"
            size="mini"
            icon="el-icon-download"
            @click="download_excel"
          />
        </el-col>
      </el-row>
      <el-table v-loading="loading" :data="oss_list" style="width: 100%">
        <el-table-column prop="id" label="ID" width="55" />
        <el-table-column prop="oss_name" label="oss名称" />
        <el-table-column prop="storage_class" label="存储类型" width="100" />
        <el-table-column
          prop="extranet_endpoint"
          label="外网域名"
          width="240"
        />
        <el-table-column
          prop="intranet_endpoint"
          label="内网域名"
          width="270"
        />
        <el-table-column prop="grant" label="ACL权限" width="100" />
        <el-table-column prop="location" label="地域" />
        <el-table-column prop="u_time" label="更新时间" width="150">
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
        <el-table-column label="操作" width="80" />
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
          <el-col v-if="k.indexOf('time') != -1" :span="18" style="color: green">
            : {{ v.replace("T", " ").replace("Z", " ").split(".")[0] }}
          </el-col>
          <el-col v-else :span="18" style="color: green">: {{ v }}</el-col>
        </el-row>
      </el-dialog>
    </div>
    <div v-if="is_mobile">
      <div class="sub-title">Oss 信息</div>
      <el-row style="margin-top: 15px">
        <el-col :span="2" :offset="21">
          <el-button
            v-if="down_btn_show"
            type="success"
            size="mini"
            icon="el-icon-download"
            @click="download_excel"
          />
        </el-col>
      </el-row>
      <el-table v-loading="loading" :data="oss_list" style="width: 100%">
        <el-table-column prop="id" label="ID" width="55" fixed />
        <el-table-column prop="oss_name" label="oss名称" fixed width="120" />
        <el-table-column prop="storage_class" label="存储类型" width="100" />
        <el-table-column
          prop="extranet_endpoint"
          label="外网域名"
          width="240"
        />
        <el-table-column
          prop="intranet_endpoint"
          label="内网域名"
          width="270"
        />
        <el-table-column prop="grant" label="ACL权限" width="100" />
        <el-table-column prop="location" label="地域" />
        <el-table-column prop="u_time" label="更新时间" width="150">
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
        <el-table-column label="操作" width="180" />
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
          <el-col v-if="k.indexOf('time') != -1" :span="18" style="color: green">
            : {{ v.replace("T", " ").replace("Z", " ").split(".")[0] }}
          </el-col>
          <el-col v-else :span="18" style="color: green">: {{ v }}</el-col>
        </el-row>
      </el-dialog>
    </div>
  </div>
</template>
<script>
import { oss_get } from "@/api/assets-list/basic";
import { btn_check } from "@/api/btn";
import { kvs } from "@/utils/assets";
import config from "@/utils/config";
import { StorePage } from "@/utils/auth";
export default {
  name: "Oss",
  data() {
    return {
      dialogVisible: false,
      oss_list: [],
      page_no: 1,
      page_size: 10,
      total: 0,
      oss: {},
      specs: {},
      kvs: kvs,
      down_btn_show: 0,
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
      const ResObj = await oss_get({
        type: "get_all_oss",
        page_no: 1,
        page_size: 999,
      });
      const ossAll = ResObj.data.oss_infos;
      const oss = ossAll[1];
      const headers = Object.keys(oss);
      const data = this.formatJson(headers, ossAll);
      import("@/vendor/Export2Excel").then((excel) => {
        const tHeader = headers;
        excel.export_json_to_excel({
          header: tHeader, // 表头 必填
          data, // 具体数据 必填
          filename: "oss", // 非必填
          autoWidth: true, // 非必填
          bookType: "xlsx", // 非必填
        });
      });
    },
    show_more(id) {
      oss_get({ type: "oss_details", id: id }).then((response) => {
        if (response.code === 200) {
          this.oss = response.data[0];
          this.specs = {
            create_time: this.oss.create_time,
            data_redundancy_type: this.oss.data_redundancy_type,
          };
          this.dialogVisible = true;
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
        }
      });
      this.page_data();
    },
    page_data() {
      const query = {
        type: "get_all_oss",
        page_no: this.page_no,
        page_size: this.page_size,
      };
      oss_get(query).then((response) => {
        if (response.code === 200) {
          this.oss_list = response.data.oss_infos;
          this.total = response.data.total;
          this.loading = false;
        }
      });
    },
    handleCurrentChange(val) {
      this.page_no = val;
      this.page_data();
    },
  },
};
</script>

<style lang="scss" scoped>
.oss {
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
