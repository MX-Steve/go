<template>
  <div class="vpc">
    <div v-if="!is_mobile" class="web">
      <div class="sub-title">VPC 信息</div>
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
      <el-table v-loading="loading" :data="vpc_list" style="width: 100%">
        <el-table-column prop="id" label="ID" width="55" />
        <el-table-column prop="vpc_name" label="名称" width="200" />
        <el-table-column prop="zone_id" label="区域">
          <template slot-scope="scope">
            {{ get_name(scope.row.zone_id) }}
          </template>
        </el-table-column>
        <el-table-column prop="vpc_status" label="状态" width="80">
          <template slot-scope="scope">
            {{ get_status_name(scope.row.vpc_status) }}
          </template>
        </el-table-column>
        <el-table-column prop="cidr_block" label="网段" width="120" />
        <el-table-column prop="description" label="描述信息" />
        <el-table-column label="更新时间" width="150">
          <template slot-scope="scope">
            {{
              scope.row.u_time.replace("T", " ").replace("Z", "").split(".")[0]
            }}
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
            v-if="k.indexOf('time') != -1"
            :span="18"
            style="color: green"
          >
            : {{ v.replace("T", " ").replace("Z", " ") }}
          </el-col>
          <el-col v-else :span="18" style="color: green">: {{ v }}</el-col>
        </el-row>
      </el-dialog>
    </div>
    <div v-if="is_mobile" class="mobile">
      <div class="sub-title">VPC 信息</div>
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
      <el-table v-loading="loading" :data="vpc_list" style="width: 100%">
        <el-table-column prop="id" label="ID" width="55" fixed />
        <el-table-column prop="vpc_id" label="VpcID" width="230" />
        <el-table-column prop="vpc_name" label="名称" width="150" fixed />
        <el-table-column prop="zone_id" label="区域">
          <template slot-scope="scope">
            {{ get_name(scope.row.zone_id) }}
          </template>
        </el-table-column>
        <el-table-column prop="vpc_status" label="状态">
          <template slot-scope="scope">
            {{ get_status_name(scope.row.vpc_status) }}
          </template>
        </el-table-column>
        <el-table-column prop="cidr_block" label="网段" width="150" />
        <el-table-column prop="description" label="描述信息" width="250" />
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
            v-if="k.indexOf('time') != -1"
            :span="18"
            style="color: green"
          >
            : {{ v.replace("T", " ").replace("Z", " ") }}
          </el-col>
          <el-col v-else :span="18" style="color: green">: {{ v }}</el-col>
        </el-row>
      </el-dialog>
    </div>
  </div>
</template>
<script>
import { zone_info_get, device_status_get } from "@/api/assets/basic";
import { vpc_get } from "@/api/assets-list/basic";
import { btn_check } from "@/api/btn";
import { kvs, get_item_name } from "@/utils/assets";
import config from "@/utils/config";
import { NetworkPage } from "@/utils/auth";
export default {
  name: "Vpc",
  data() {
    return {
      dialogVisible: false,
      vpc_list: [],
      vpc: {},
      specs: {},
      kvs: kvs,
      page_no: 1,
      page_size: 10,
      total: 1,
      zone_list: [],
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
      const ResObj = await vpc_get({
        type: "get_all_vpc",
        page_no: 1,
        page_size: 999,
      });
      const vpcs = ResObj.data.vpc_infos;
      for (var i = 0; i < vpcs.length; i++) {
        vpcs[i]["zone_id"] = this.get_name(vpcs[i]["zone_id"], this.zone_list);
        vpcs[i]["vpc_status"] = this.get_name(
          vpcs[i]["vpc_status"],
          this.status_list
        );
      }
      const vpc = vpcs[1];
      const headers = Object.keys(vpc);
      const data = this.formatJson(headers, vpcs);
      import("@/vendor/Export2Excel").then((excel) => {
        const tHeader = headers;
        excel.export_json_to_excel({
          header: tHeader, // 表头 必填
          data, // 具体数据 必填
          filename: "vpcs", // 非必填
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
        type: "get_all_vpc",
        page_no: this.page_no,
        page_size: this.page_size,
      };
      vpc_get(query).then((response) => {
        if (response.code === 200) {
          this.vpc_list = response.data.vpc_infos;
          this.total = response.data.total;
          this.loading = false;
        }
      });
    },
    show_more(id) {
      vpc_get({ type: "vpc_details", id: id }).then((response) => {
        if (response.code === 200) {
          this.vpc = response.data[0];
          this.specs = {
            secondary_cidr_blocks: this.vpc.secondary_cidr_blocks,
            creation_time: this.vpc.creation_time,
            vswitch_ids: this.vpc.vswitch_ids,
            resource_group_id: this.vpc.resource_group_id,
            vrouter_id: this.vpc.vrouter_id,
            router_table_ids: this.vpc.router_table_ids,
          };
          this.dialogVisible = true;
        }
      });
    },
    handleCurrentChange(val) {
      this.page_no = val;
      this.page_data();
    },
    get_name(id) {
      return get_item_name(id, this.zone_list);
    },
    get_status_name(id) {
      return get_item_name(id, this.status_list);
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
      zone_info_get().then((response) => {
        if (response.code == 200) {
          this.zone_list = response.data;
        }
      });
      device_status_get().then((response) => {
        if (response.code == 200) {
          this.status_list = response.data;
        }
      });
    },
  },
};
</script>
<style lang="scss" scoped>
.vpc {
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

