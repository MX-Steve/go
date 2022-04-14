<template>
  <div class="slb">
    <div v-if="!is_mobile">
      <div class="sub-title">SLB 信息</div>
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
      <el-table v-loading="loading" :data="slb_list" style="width: 100%">
        <el-table-column prop="id" label="ID" width="55" />
        <el-table-column prop="lb_name" label="实例名" width="200" />
        <el-table-column prop="address_type" label="网络类型" />
        <el-table-column prop="address_ip_version" label="IP版本" />
        <el-table-column prop="bandwidth" label="带宽" />
        <el-table-column prop="zone_id" label="所属区域" width="135">
          <template slot-scope="scope">
            {{ get_name(scope.row.zone_id, zone_list) }}
          </template>
        </el-table-column>
        <el-table-column
          prop="master_zone_id"
          label="实例的主可用区ID"
          width="135"
        >
          <template slot-scope="scope">
            {{ get_name(scope.row.master_zone_id, idc_list) }}
          </template>
        </el-table-column>
        <el-table-column
          prop="slave_zone_id"
          label="实例的备可用区ID"
          width="135"
        >
          <template slot-scope="scope">
            {{ get_name(scope.row.slave_zone_id, idc_list) }}
          </template>
        </el-table-column>
        <el-table-column prop="lb_status" label="状态">
          <template slot-scope="scope">
            {{ get_name(scope.row.lb_status, status_list) }}
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
      <div class="sub-title">SLB 信息</div>
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
      <el-table v-loading="loading" :data="slb_list" style="width: 100%">
        <el-table-column prop="id" label="ID" width="55" fixed />
        <el-table-column prop="lb_id" label="负载均衡实例ID" width="200" />
        <el-table-column
          prop="lb_name"
          label="负载均衡实例名"
          width="200"
          fixed
        />
        <el-table-column prop="address_type" label="负载均衡实例网络类型" />
        <el-table-column prop="address_ip_version" label="IP版本" />
        <el-table-column prop="bandwidth" label="带宽" />
        <el-table-column prop="zone_id" label="所属区域" width="135">
          <template slot-scope="scope">
            {{ get_name(scope.row.zone_id, zone_list) }}
          </template>
        </el-table-column>
        <el-table-column
          prop="master_zone_id"
          label="实例的主可用区ID"
          width="135"
        >
          <template slot-scope="scope">
            {{ get_name(scope.row.master_zone_id, idc_list) }}
          </template>
        </el-table-column>
        <el-table-column
          prop="slave_zone_id"
          label="实例的备可用区ID"
          width="135"
        >
          <template slot-scope="scope">
            {{ get_name(scope.row.slave_zone_id, idc_list) }}
          </template>
        </el-table-column>
        <el-table-column prop="lb_status" label="状态">
          <template slot-scope="scope">
            {{ get_name(scope.row.lb_status, status_list) }}
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
import { slb_get } from "@/api/assets-list/basic";
import { device_status_get, zone_info_get } from "@/api/assets/basic";
import { idc_get } from "@/api/assets-manage/assets-input";
import { btn_check } from "@/api/btn";
import { kvs, get_item_name } from "@/utils/assets";
import config from "@/utils/config";
import { NetworkPage } from "@/utils/auth";
export default {
  name: "Slb",
  data() {
    return {
      dialogVisible: false,
      slb_list: [],
      slb: {},
      specs: {},
      page_no: 1,
      page_size: 10,
      total: 1,
      kvs: kvs,
      status_list: [],
      zone_list: [],
      idc_list: [],
      down_btn_show: 0,
      is_mobile: config.isMobile,
      NetworkPage: NetworkPage,
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
      const slbResObj = await slb_get({
        type: "get_all_slb",
        page_no: 1,
        page_size: 999,
      });
      const slbs = slbResObj.data.slb_infos;
      for (var i = 0; i < slbs.length; i++) {
        slbs[i]["master_zone_id"] = this.get_name(
          slbs[i]["master_zone_id"],
          this.zone_list
        );
        slbs[i]["slave_zone_id"] = this.get_name(
          slbs[i]["slave_zone_id"],
          this.zone_list
        );
        slbs[i]["zone_id"] = this.get_name(slbs[i]["zone_id"], this.zone_list);
        slbs[i]["lb_status"] = this.get_name(
          slbs[i]["lb_status"],
          this.status_list
        );
      }
      const slb = slbs[1];
      const headers = Object.keys(slb);
      const data = this.formatJson(headers, slbs);
      import("@/vendor/Export2Excel").then((excel) => {
        const tHeader = headers;
        excel.export_json_to_excel({
          header: tHeader, // 表头 必填
          data, // 具体数据 必填
          filename: "slbs", // 非必填
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
        type: "get_all_slb",
        page_no: this.page_no,
        page_size: this.page_size,
      };
      slb_get(query).then((response) => {
        if (response.code === 200) {
          this.slb_list = response.data.slb_infos;
          this.total = response.data.total;
          this.loading = false;
        }
      });
    },
    show_more(id) {
      slb_get({ type: "slb_details", id: id }).then((response) => {
        if (response.code === 200) {
          this.slb = response.data[0];
          this.specs = {
            resource_group_id: this.slb.resource_group_id,
            create_time: this.slb.create_time,
            pay_type: this.slb.pay_type,
            network_type: this.slb.network_type,
            internet_charge_type: this.slb.internet_charge_type,
            u_time: this.slb.u_time,
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
      btn_check(this.NetworkPage).then((response) => {
        if (response.code == 200) {
          this.down_btn_show = response.data.down;
        }
      });
      this.page_data();
      device_status_get().then((response) => {
        if (response.code === 200) {
          this.status_list = response.data;
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
        }
      });
    },
    get_name(id, list) {
      return get_item_name(id, list);
    },
  },
};
</script>
<style lang="scss" scoped>
.slb {
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

