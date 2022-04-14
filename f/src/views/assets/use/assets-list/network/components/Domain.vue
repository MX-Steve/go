<template>
  <div class="domain">
    <div v-if="!is_mobile">
      <div class="sub-title">Domain 信息</div>
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
      <el-row style="margin: 15px auto">
        <el-col :span="7" :offset="17">
          <el-input
            v-model="domain_name"
            placeholder="请输入域名"
            style="width: 70%; margin-right: 5px"
          />
          <el-button type="primary" @click="domain_change()"> 搜索 </el-button>
        </el-col>
      </el-row>
      <el-table v-loading="loading" :data="domain_list" style="width: 100%">
        <el-table-column prop="id" label="ID" width="55" />
        <el-table-column prop="domain_name" label="域名" />
        <el-table-column label="域名注册类型">
          <template slot-scope="scope">
            {{ get_registrant_type(scope.row.registrant_type) }}
          </template>
        </el-table-column>
        <el-table-column prop="domain_type" label="域名类型" />
        <el-table-column prop="expiration_date" label="到期日期" />
        <el-table-column label="域名状态">
          <template slot-scope="scope">
            {{ get_domain_status(scope.row.domain_status) }}
          </template>
        </el-table-column>
        <el-table-column label="更新时间">
          <template slot-scope="scope">
            {{ scope.row.u_time.split(".")[0] }}
          </template>
        </el-table-column>
        <el-table-column label="其他信息">
          <template slot-scope="scope" width="100">
            <el-button
              size="mini"
              icon="el-icon-more"
              @click="show_more(scope.row.id)"
            />
          </template>
        </el-table-column>
        <el-table-column label="操作" width="180">
          <template slot-scope="scope">
            <el-button
              v-if="conn_btn_show"
              type="text"
              @click="record_btn(scope.row)"
            >
              解析
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
      <el-dialog title="更多信息" :visible.sync="dialogVisible" width="45%">
        <el-row class="text item">
          <el-col :span="6">域名ID</el-col>
          <el-col :span="18" style="color: green">
            {{ domain_extra["domain_id"] }}
          </el-col>
        </el-row>
        <el-row class="text item">
          <el-col :span="6">注册邮箱</el-col>
          <el-col :span="18" style="color: green">
            {{ domain_extra["email"] }}
          </el-col>
        </el-row>
        <el-row class="text item">
          <el-col :span="6">注册日期</el-col>
          <el-col :span="18" style="color: green">
            {{ domain_extra["registration_date"] }}
          </el-col>
        </el-row>
        <el-row class="text item">
          <el-col :span="6">域名持有者</el-col>
          <el-col :span="18" style="color: green">
            {{ domain_extra["domain_owner"] }}
          </el-col>
        </el-row>
      </el-dialog>
    </div>
    <div v-if="is_mobile">
      <div class="sub-title">Domain 信息</div>
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
      <el-row style="margin: 15px auto">
        <el-col :span="24">
          <el-input
            v-model="domain_name"
            placeholder="请输入域名"
            style="width: 70%; margin-right: 5px"
          />
          <el-button type="primary" @click="domain_change()"> 搜索 </el-button>
        </el-col>
      </el-row>
      <el-table v-loading="loading" :data="domain_list" style="width: 100%">
        <el-table-column prop="id" label="ID" width="55" fixed />
        <el-table-column prop="domain_name" label="域名" fixed width="120" />
        <el-table-column label="域名注册类型">
          <template slot-scope="scope">
            {{ get_registrant_type(scope.row.registrant_type) }}
          </template>
        </el-table-column>
        <el-table-column prop="domain_type" label="域名类型" />
        <el-table-column prop="expiration_date" label="到期日期" />
        <el-table-column label="域名状态">
          <template slot-scope="scope">
            {{ get_domain_status(scope.row.domain_status) }}
          </template>
        </el-table-column>
        <el-table-column label="更新时间">
          <template slot-scope="scope">
            {{ scope.row.u_time.split(".")[0] }}
          </template>
        </el-table-column>
        <el-table-column label="其他信息">
          <template slot-scope="scope" width="100">
            <el-button
              size="mini"
              icon="el-icon-more"
              @click="show_more(scope.row.id)"
            />
          </template>
        </el-table-column>
        <el-table-column label="操作" width="180">
          <template slot-scope="scope">
            <el-button
              v-if="conn_btn_show"
              type="text"
              @click="record_btn(scope.row)"
            >
              解析
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
      <el-dialog title="更多信息" :visible.sync="dialogVisible" width="100%">
        <el-row class="text item">
          <el-col :span="6">域名ID</el-col>
          <el-col :span="18" style="color: green">
            {{ domain_extra["domain_id"] }}
          </el-col>
        </el-row>
        <el-row class="text item">
          <el-col :span="6">注册邮箱</el-col>
          <el-col :span="18" style="color: green">
            {{ domain_extra["email"] }}
          </el-col>
        </el-row>
        <el-row class="text item">
          <el-col :span="6">注册日期</el-col>
          <el-col :span="18" style="color: green">
            {{ domain_extra["registration_date"] }}
          </el-col>
        </el-row>
        <el-row class="text item">
          <el-col :span="6">域名持有者</el-col>
          <el-col :span="18" style="color: green">
            {{ domain_extra["domain_owner"] }}
          </el-col>
        </el-row>
      </el-dialog>
    </div>
  </div>
</template>
<script>
import { device_status_get } from "@/api/assets/basic";
import { btn_check } from "@/api/btn";
import { domain_get } from "@/api/assets-list/basic";
import { get_item_name, registrant_type_choices } from "@/utils/assets";
import config from "@/utils/config";
import { NetworkPage } from "@/utils/auth";
export default {
  name: "Domain",
  data() {
    return {
      dialogVisible: false,
      domain_list: [],
      page_no: 1,
      page_size: 10,
      total: 1,
      domain_extra: {},
      status_list: [],
      domain: {},
      domain_name: "",
      down_btn_show: 0,
      conn_btn_show: 0,
      is_mobile: config.isMobile,
      NetworkPage: NetworkPage,
      loading: true,
    };
  },
  mounted() {
    this.init_data();
    const that = this;
    // Vue.config.keyCodes.enter = 13;
    document.onkeydown = function (e) {
      // 这里把事件注册在document上就是因为防止聚焦
      const key = window.event.keyCode;
      if (key === 13) {
        that.domain_change();
      }
    };
  },
  methods: {
    formatJson(filterVal, jsonData) {
      return jsonData.map((v) => filterVal.map((j) => v[j]));
    },
    async download_excel() {
      const domainResObj = await domain_get({
        type: "get_all_domain",
        page_no: 1,
        page_size: 999,
      });
      const domains = domainResObj.data.domain_infos;
      for (var i = 0; i < domains.length; i++) {
        domains[i]["domain_status"] = this.get_name(
          domains[i]["domain_status"],
          this.status_list
        );
      }
      const domain = domains[1];
      const headers = Object.keys(domain);
      const data = this.formatJson(headers, domains);
      import("@/vendor/Export2Excel").then((excel) => {
        const tHeader = headers;
        excel.export_json_to_excel({
          header: tHeader, // 表头 必填
          data, // 具体数据 必填
          filename: "domains", // 非必填
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
    init_data() {
      btn_check(this.NetworkPage).then((response) => {
        if (response.code == 200) {
          this.down_btn_show = response.data.down;
          this.conn_btn_show = response.data.conn;
        }
      });
      this.page_data();
      device_status_get().then((response) => {
        if (response.code === 200) {
          this.status_list = response.data;
        }
      });
    },
    page_data() {
      const query = {
        type: "get_all_domain",
        page_no: this.page_no,
        page_size: this.page_size,
      };
      domain_get(query).then((response) => {
        if (response.code === 200) {
          this.domain_list = response.data.domain_infos;
          this.total = response.data.total;
          this.loading = false;
        }
      });
    },
    get_name(id, list) {
      return get_item_name(id, list);
    },
    get_registrant_type(id) {
      return get_item_name(Number(id), registrant_type_choices);
    },
    get_domain_status(id) {
      return get_item_name(id, this.status_list);
    },
    handleCurrentChange(val) {
      this.page_no = val;
      this.page_data();
    },
    show_more(id) {
      domain_get({ type: "domain_details", id: id }).then((response) => {
        if (response.code === 200) {
          this.domain = response.data[0];
          this.domain_extra = {
            domain_id: this.domain.domain_id,
            email: this.domain.email,
            registration_date: this.domain.registration_date,
            domain_owner: this.domain.domain_owner,
          };
          this.dialogVisible = true;
        }
      });
    },
    domain_change() {
      domain_get({
        type: "get_all_domain",
        domain_name: this.domain_name,
        page_no: 1,
        page_size: this.page_size,
      }).then((response) => {
        if (response.code === 200) {
          this.domain_list = response.data.domain_infos;
          this.total = response.data.total;
        }
      });
    },
    record_btn(row) {
      this.$router.push({
        path: "/assets/use/list/domain-parse",
        query: {
          domain_id: row.id,
        },
      });
    },
  },
};
</script>
<style lang="scss" scoped>
.domain {
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

