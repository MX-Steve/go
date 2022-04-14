<template>
  <div class="cdn">
    <div v-if="!is_mobile">
      <div class="sub-title">CDN 信息</div>
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
      <el-table v-loading="loading" :data="cdn_list" style="width: 100%">
        <el-table-column prop="id" label="ID" width="55" />
        <el-table-column prop="domain_name" label="域名" />
        <el-table-column prop="cdn_name" label="加速域名" />
        <el-table-column prop="coverage" label="加速区域" width="100" />
        <el-table-column prop="description" label="描述" />
        <el-table-column prop="domain_status" label="加速状态" width="80">
          <template slot-scope="scope">
            {{ get_name(scope.row.domain_status) }}
          </template>
        </el-table-column>
        <el-table-column prop="cdn_type" label="业务类型" width="80" />
        <el-table-column label="其他信息" width="80">
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
          <el-col
            v-if="
              k.indexOf('time') != -1 ||
                k.indexOf('gmt_created') != -1 ||
                k.indexOf('gmt_modified') != -1
            "
            :span="18"
            style="color: green"
          >
            : {{ v.replace("T", " ").replace("Z", " ") }}
          </el-col>
          <el-col v-else-if="k != 'source'" :span="18" style="color: green">
            : {{ v }}
          </el-col>
          <el-col v-else-if="k == 'source'" :span="18" style="color: green">
            <el-row v-for="(item, k2) in v" :key="k2">
              <el-col :span="6">{{ k2 }}</el-col>
              <el-col :span="18">: {{ item }}</el-col>
            </el-row>
          </el-col>
        </el-row>
      </el-dialog>
    </div>
    <div v-if="is_mobile">
      <div class="sub-title">CDN 信息</div>
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
      <el-table v-loading="loading" :data="cdn_list" style="width: 100%">
        <el-table-column prop="id" label="ID" width="40" fixed />
        <el-table-column prop="domain_name" label="域名" fixed width="250" />
        <el-table-column prop="cdn_name" label="加速域名" width="350" />
        <el-table-column prop="coverage" label="加速区域" />
        <el-table-column prop="description" label="描述" width="200" />
        <el-table-column prop="domain_status" label="加速域名状态" width="120">
          <template slot-scope="scope">
            {{ get_name(scope.row.domain_status) }}
          </template>
        </el-table-column>
        <el-table-column prop="cdn_type" label="业务类型" />
        <el-table-column label="其他信息" width="100">
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
          <el-col
            v-if="
              k.indexOf('time') != -1 ||
                k.indexOf('gmt_created') != -1 ||
                k.indexOf('gmt_modified') != -1
            "
            :span="18"
            style="color: green"
          >
            : {{ v.replace("T", " ").replace("Z", " ") }}
          </el-col>
          <el-col v-else-if="k != 'source'" :span="18" style="color: green">
            : {{ v }}
          </el-col>
          <el-col v-else-if="k == 'source'" :span="18" style="color: green">
            <el-row v-for="(item, k2) in v" :key="k2">
              <el-col :span="6">{{ k2 }}</el-col>
              <el-col :span="18">: {{ item }}</el-col>
            </el-row>
          </el-col>
        </el-row>
      </el-dialog>
    </div>
  </div>
</template>
<script>
import { cdn_get } from "@/api/assets-list/basic";
import { device_status_get } from "@/api/assets/basic";
import { btn_check } from "@/api/btn";
import { kvs, get_item_name } from "@/utils/assets";
import config from "@/utils/config";
import { NetworkPage } from "@/utils/auth";
export default {
  name: "CDN",
  data() {
    return {
      dialogVisible: false,
      cdn_list: [],
      page_no: 1,
      page_size: 10,
      total: 1,
      cdn: {},
      specs: {},
      kvs: kvs,
      status_list: [],
      down_btn_show: 0,
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
      const cdnResObj = await cdn_get({
        type: "get_all_cdn",
        page_no: 1,
        page_size: 999,
      });
      const cdns = cdnResObj.data.cdn_infos;
      for (var i = 0; i < cdns.length; i++) {
        cdns[i]["domain_status"] = this.get_name(
          cdns[i]["domain_status"],
          this.status_list
        );
      }
      const cdn = cdns[1];
      const headers = Object.keys(cdn);
      const data = this.formatJson(headers, cdns);
      import("@/vendor/Export2Excel").then((excel) => {
        const tHeader = headers;
        excel.export_json_to_excel({
          header: tHeader, // 表头 必填
          data, // 具体数据 必填
          filename: "cdns", // 非必填
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
        type: "get_all_cdn",
        page_no: this.page_no,
        page_size: this.page_size,
      };
      cdn_get(query).then((response) => {
        if (response.code === 200) {
          this.cdn_list = response.data.cdn_infos;
          this.total = response.data.total;
          this.loading = false;
        }
      });
    },
    show_more(id) {
      cdn_get({ type: "cdn_details", id: id }).then((response) => {
        if (response.code === 200) {
          this.cdn = response.data[0];
          var source = JSON.parse(this.cdn.source)[0];
          this.specs = {
            resource_groupId: this.cdn.resource_group_id,
            ssl_protocol: this.cdn.ssl_protocol,
            gmt_created: this.cdn.gmt_created,
            gmt_modified: this.cdn.gmt_modified,
            source: source,
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
          }
        }
      );
      this.page_data();
      device_status_get().then((response) => {
        if (response.code === 200) {
          this.status_list = response.data;
        }
      });
    },
    get_name(id) {
      return get_item_name(id, this.status_list);
    },
  },
};
</script>

<style lang="scss" scoped>
.cdn {
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
