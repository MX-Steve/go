<template>
  <div class="show-host">
    <!-- <el-row style="margin: 15px auto">
      <el-button size="mini"> <i class="el-icon-refresh"></i> 刷新 </el-button>
      <el-button size="mini">
        <i class="el-icon-refresh-right"></i> 重置密码
      </el-button>
      <el-button size="mini">
        <i class="el-icon-video-play"></i> 开机
      </el-button>
      <el-button size="mini">
        <i class="el-icon-video-pause"></i> 关机
      </el-button>
      <el-button size="mini"> <i class="el-icon-loading"></i> 重启 </el-button>
    </el-row> -->
    <div class="web">
      <el-row style="margin: 15px auto">
        <el-badge
          v-if="first_list.length > 0"
          :value="first_list.length"
          class="item"
          style="margin-right: 20px"
        >
          <el-button
            size="small"
            :type="activeName == 'first' ? 'primary' : ''"
            @click="tab_change('first')"
          >
            物理机
          </el-button>
        </el-badge>
        <el-badge
          v-if="second_list.length > 0"
          :value="second_list.length"
          class="item"
          style="margin-right: 20px"
        >
          <el-button
            size="small"
            :type="activeName == 'second' ? 'primary' : ''"
            @click="tab_change('second')"
          >
            云主机
          </el-button>
        </el-badge>
        <el-badge
          v-if="third_list.length > 0"
          :value="third_list.length"
          class="item"
        >
          <el-button
            size="small"
            :type="activeName == 'third' ? 'primary' : ''"
            @click="tab_change('third')"
          >
            虚拟机
          </el-button>
        </el-badge>
      </el-row>
      <div style="border: 2px solid #eee; margin: 15px auto" />
      <el-table
        ref="multipleTable"
        tooltip-effect="dark"
        :data="
          activeName == 'first'
            ? first_list
            : activeName == 'second'
              ? second_list
              : third_list
        "
        style="width: 100%"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column prop="id" label="ID" width="55" />
        <el-table-column prop="instance_name" label="实例名称" width="180" />
        <el-table-column prop="ip_address" label="IP地址" width="180" />
        <el-table-column label="区域">
          <template slot-scope="scope">
            <div>{{ get_zone_name(scope.row.zone_id) }}</div>
          </template>
        </el-table-column>
        <el-table-column show-overflow-tooltip label="机房">
          <template slot-scope="scope">
            <div>{{ get_idc_name(scope.row.idc_id) }}</div>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button
              v-if="conn_btn_show"
              size="mini"
              type="primary"
              @click="conn_btn(scope.row)"
            >
              连接
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
    </div>
    <div class="mobile">
      <el-row style="margin: 15px auto">
        <el-badge
          v-if="first_list.length > 0"
          :value="first_list.length"
          class="item"
          style="margin-right: 20px"
        >
          <el-button
            size="small"
            :type="activeName == 'first' ? 'primary' : ''"
            @click="tab_change('first')"
          >
            物理机
          </el-button>
        </el-badge>
        <el-badge
          v-if="second_list.length > 0"
          :value="second_list.length"
          class="item"
          style="margin-right: 20px"
        >
          <el-button
            size="small"
            :type="activeName == 'second' ? 'primary' : ''"
            @click="tab_change('second')"
          >
            云主机
          </el-button>
        </el-badge>
        <el-badge
          v-if="third_list.length > 0"
          :value="third_list.length"
          class="item"
        >
          <el-button
            size="small"
            :type="activeName == 'third' ? 'primary' : ''"
            @click="tab_change('third')"
          >
            虚拟机
          </el-button>
        </el-badge>
      </el-row>
      <div style="border: 2px solid #eee; margin: 15px auto" />
      <el-table
        ref="multipleTable"
        tooltip-effect="dark"
        :data="
          activeName == 'first'
            ? first_list
            : activeName == 'second'
              ? second_list
              : third_list
        "
        style="width: 100%"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" fixed />
        <el-table-column prop="id" label="ID" width="55" />
        <el-table-column prop="instance_name" label="实例名称" width="180" />
        <el-table-column prop="ip_address" label="IP地址" width="180" fixed />
        <el-table-column label="区域">
          <template slot-scope="scope">
            <div>{{ get_zone_name(scope.row.zone_id) }}</div>
          </template>
        </el-table-column>
        <el-table-column show-overflow-tooltip label="机房">
          <template slot-scope="scope">
            <div>{{ get_idc_name(scope.row.idc_id) }}</div>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button
              v-if="conn_btn_show"
              size="mini"
              type="primary"
              @click="conn_btn(scope.row)"
            >
              连接
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
    </div>
  </div>
</template>
<script>
import { btn_check } from "@/api/btn";
import { zone_info_get } from "@/api/assets/basic";
import { idc_get, go_ssh } from "@/api/assets-manage/assets-input";
import { service_get } from "@/api/tree/index";
export default {
  name: "ShowHost",
  props: ["tree", "ids"],
  data() {
    return {
      activeName: "first",
      first_list: [],
      second_list: [],
      third_list: [],
      multipleSelection: [],
      zones: [],
      idc: [],
      env_id: "",
      service_id: "",
      project_id: "",
      page_no: 1,
      page_size: 10,
      total: 1,
      conn_btn_show: 0,
    };
  },
  watch: {
    ids: function(newVal, oldVal) {
      this.env_id = newVal.env_id;
      this.service_id = newVal.service_id;
      this.project_id = newVal.project_id;
      this.first_list = [];
      this.second_list = [];
      this.third_list = [];
      this.get_host();
    },
  },
  mounted() {
    this.init_data();
  },
  methods: {
    conn_btn(row) {
      go_ssh({ id: row.id }).then((response) => {
        if (response.code === 200) {
          this.$router.push({
            path: "/service-tree/web-ssh",
            query: { id: row.id },
          });
        } else {
          this.$message({
            message: response.msg,
            type: "error",
          });
        }
      });
    },
    init_data() {
      btn_check({ keys: "tree.service-tree.conn" }).then((response) => {
        if (response.code === 200) {
          this.conn_btn_show = response.data.conn;
        }
      });
      zone_info_get().then((response) => {
        if (response.code === 200) {
          this.zones = response.data;
        } else {
          this.zones = [];
        }
      });
      idc_get({ type: "get_all_idcs" }).then((response) => {
        if (response.code === 200) {
          this.idc = response.data;
        } else {
          this.idc = [];
        }
      });
      this.get_host();
    },
    get_host() {
      service_get({
        project_id: this.project_id,
        service_id: this.service_id,
        env_id: this.env_id,
      }).then((response) => {
        if (response.code === 200) {
          const data = response.data;
          for (var i = 0; i < data.length; i++) {
            var item = data[i];
            if (item.server_type == 1) {
              this.first_list.push(item);
            } else if (item.server_type == 2) {
              this.second_list.push(item);
            } else if (item.server_type == 3) {
              this.third_list.push(item);
            }
          }
        }
        if (this.first_list.length > 0) {
          this.activeName = "first";
        } else if (this.second_list.length > 0) {
          this.activeName = "second";
        } else {
          this.activeName = "third";
        }
      });
    },
    tab_change(tab) {
      this.activeName = tab;
    },
    get_zone_name(id) {
      var name = "";
      for (var i = 0; i < this.zones.length; i++) {
        var item = this.zones[i];
        if (item.id === id) {
          name = item.name;
          break;
        }
      }
      return name;
    },
    get_idc_name(id) {
      var name = "";
      for (var i = 0; i < this.idc.length; i++) {
        var item = this.idc[i];
        if (item.id === id) {
          name = item.name;
          break;
        }
      }
      return name;
    },
    handleSelectionChange(val) {
      this.multipleSelection = val;
    },
    web_ssh(row) {
      console.log(row);
    },
    handleCurrentChange(val) {
      this.page_no = val;
    },
  },
};
</script>
<style lang="scss" scoped>
.show-host {
  padding: 5px;
  background-color: #fff;
  min-height: 650px;
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
