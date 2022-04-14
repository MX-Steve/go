<template>
  <div class="dashboard-container">
    <div v-if="!is_mobile" class="web">
      <div class="head">
        <el-row>
          <el-col :span="6">
            <el-card shadow="hover">
              <div slot="header" class="clearfix card-head">
                <span> <i class="el-icon-s-home" /> 总服务器统计 </span>
              </div>
              <div class="card-box">
                <div class="number">{{ header.total }}</div>
                <div class="info">总数展示</div>
              </div>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card shadow="hover">
              <div slot="header" class="clearfix card-head">
                <span> <i class="el-icon-magic-stick" /> 在线服务器 </span>
              </div>
              <div class="card-box">
                <div class="number online">{{ header.online }}</div>
                <div class="info">在线服务器</div>
              </div>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card shadow="hover">
              <div slot="header" class="clearfix card-head">
                <span> <i class="el-icon-shopping-cart-full" /> 项目 </span>
              </div>
              <div class="card-box">
                <div class="number">{{ header.projects }}</div>
                <div class="info">现有项目个数</div>
              </div>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card shadow="hover">
              <div slot="header" class="clearfix card-head">
                <span> <i class="el-icon-coin" /> 服务 </span>
              </div>
              <div class="card-box">
                <div class="number">{{ header.services }}</div>
                <div class="info">现有服务个数</div>
              </div>
            </el-card>
          </el-col>
        </el-row>
      </div>
      <div class="body">
        <el-row>
          <el-col
            id="idc_table"
            ref="idc_table"
            :span="12"
            class="idc_table"
            style="min-height: 340px"
          >
            <el-card shadow="hover" style="min-height: 340px">
              <div slot="header" class="clearfix">
                <span>机房 IDC</span>
              </div>
              <div class="card-box">
                <div class="title">服务器资产分布情况</div>

                <el-table :data="idc_total_list" style="width: 100%">
                  <el-table-column prop="name" label="机房名" />
                  <el-table-column label="所占比例">
                    <template slot-scope="scope">
                      <el-progress
                        :percentage="scope.row.percentage"
                      />
                    </template>
                  </el-table-column>
                  <el-table-column prop="count" label="数量" />
                </el-table>
              </div>
            </el-card>
          </el-col>
          <el-col :span="12" style="min-height: 340px">
            <el-card shadow="hover" style="min-height: 340px">
              <div slot="header" class="clearfix">
                <span>服务器类型 Assets Type</span>
              </div>
              <div class="card-box">
                <el-table :data="assets_type_list" style="width: 100%">
                  <el-table-column prop="name" label="服务器类型" />
                  <el-table-column label="所占比例">
                    <template slot-scope="scope">
                      <el-progress
                        :percentage="Math.round((scope.row.value/header.total)*100)"
                      />
                    </template>
                  </el-table-column>
                  <el-table-column prop="value" label="数量" />
                </el-table>
                <!-- <div id="chartPie" style="width: 100%; height: 220px" /> -->
              </div>
            </el-card>
          </el-col>
        </el-row>
        <el-row style="margin-top: 15px">
          <el-col :span="12" style="min-height: 250px">
            <el-card shadow="hover" style="min-height: 340px">
              <div slot="header" class="clearfix">
                <span>工单申请 top10</span>
              </div>
              <div class="card-box">
                <div id="main" style="width: 100%; height: 340px" />
              </div>
            </el-card>
          </el-col>
          <el-col :span="12" style="min-height: 250px">
            <el-card shadow="hover" style="min-height: 340px">
              <div slot="header" class="clearfix">
                <span>最近30天登录</span>
              </div>
              <div class="card-box">
                <el-table
                  ref="table_users"
                  :data="table_users"
                  style="width: 100%; height: 315px"
                >
                  <el-table-column
                    label="上次登录时间"
                    prop="last_login"
                  />
                  <el-table-column
                    label="用户名"
                    prop="username"
                  />
                  <el-table-column label="邮箱" prop="email" />
                </el-table>
                <el-pagination
                  small
                  layout="prev, pager, next"
                  :page-size="page_size"
                  :total="total"
                  @current-change="handleCurrentChange"
                />
              </div>
            </el-card>
          </el-col>
        </el-row>
      </div>
    </div>
    <div v-if="is_mobile" class="mobile">
      <div class="head">
        <el-row>
          <el-col :span="12">
            <el-card shadow="hover">
              <div slot="header" class="clearfix card-head">
                <span> <i class="el-icon-s-home" /> 总服务器统计 </span>
              </div>
              <div class="card-box">
                <div class="number">{{ header.total }}</div>
                <div class="info">总数展示</div>
              </div>
            </el-card>
          </el-col>
          <el-col :span="12">
            <el-card shadow="hover">
              <div slot="header" class="clearfix card-head">
                <span> <i class="el-icon-magic-stick" /> 在线服务器 </span>
              </div>
              <div class="card-box">
                <div class="number online">{{ header.online }}</div>
                <div class="info">在线服务器</div>
              </div>
            </el-card>
          </el-col>
          <el-col :span="12">
            <el-card shadow="hover">
              <div slot="header" class="clearfix card-head">
                <span> <i class="el-icon-shopping-cart-full" /> 项目 </span>
              </div>
              <div class="card-box">
                <div class="number">{{ header.projects }}</div>
                <div class="info">现有项目个数</div>
              </div>
            </el-card>
          </el-col>
          <el-col :span="12">
            <el-card shadow="hover">
              <div slot="header" class="clearfix card-head">
                <span> <i class="el-icon-coin" /> 服务 </span>
              </div>
              <div class="card-box">
                <div class="number">{{ header.services }}</div>
                <div class="info">现有服务个数</div>
              </div>
            </el-card>
          </el-col>
        </el-row>
      </div>
      <div class="body">
        <el-row>
          <el-col
            id="idc_table"
            ref="idc_table"
            :span="24"
            class="idc_table"
            style="min-height: 340px"
          >
            <el-card shadow="hover" style="min-height: 340px">
              <div slot="header" class="clearfix">
                <span>机房 IDC</span>
              </div>
              <div class="card-box">
                <div class="title">服务器资产分布情况</div>

                <el-table :data="idc_total_list" style="width: 100%">
                  <el-table-column prop="name" label="机房名" />
                  <el-table-column label="所占比例">
                    <template slot-scope="scope">
                      <el-progress
                        :percentage="scope.row.percentage"
                      />
                    </template>
                  </el-table-column>
                  <el-table-column prop="count" label="数量" />
                </el-table>
              </div>
            </el-card>
          </el-col>
          <el-col :span="24" style="min-height: 340px">
            <el-card shadow="hover" style="min-height: 340px">
              <div slot="header" class="clearfix">
                <span>服务器类型 Assets Type</span>
              </div>
              <div class="card-box">
                <el-table :data="assets_type_list" style="width: 100%">
                  <el-table-column prop="name" label="服务器类型" />
                  <el-table-column label="所占比例">
                    <template slot-scope="scope">
                      <el-progress
                        :percentage="Math.round((scope.row.value/header.total)*100)"
                      />
                    </template>
                  </el-table-column>
                  <el-table-column prop="value" label="数量" />
                </el-table>
                <!-- <div id="chartPie" style="width: 100%; height: 220px" /> -->
              </div>
            </el-card>
          </el-col>
        </el-row>
        <el-row style="margin-top: 15px">
          <el-col :span="24" style="min-height: 250px">
            <el-card shadow="hover" style="min-height: 340px">
              <div slot="header" class="clearfix">
                <span>工单申请 top10</span>
              </div>
              <div class="card-box">
                <div id="main" style="width: 100%; height: 340px" />
              </div>
            </el-card>
          </el-col>
          <el-col :span="24" style="min-height: 250px">
            <el-card shadow="hover" style="min-height: 340px">
              <div slot="header" class="clearfix">
                <span>最近30天登录</span>
              </div>
              <div class="card-box">
                <el-table
                  ref="table_users"
                  :data="table_users"
                  style="width: 100%; height: 315px"
                >
                  <el-table-column
                    label="上次登录时间"
                    prop="last_login"
                  />
                  <el-table-column
                    label="用户名"
                    prop="username"
                  />
                  <el-table-column label="邮箱" prop="email" />
                </el-table>
                <el-pagination
                  small
                  layout="prev, pager, next"
                  :page-size="page_size"
                  :total="total"
                  @current-change="handleCurrentChange"
                />
              </div>
            </el-card>
          </el-col>
        </el-row>
      </div>
    </div>
  </div>
</template>

<script>
import echarts from "echarts";
import {
  summary_machine_get,
  summary_approval_get,
} from "@/api/dashboard/index";
import { user_login_get } from "@/api/user";
import config from "@/utils/config";
export function get_date_time(dt) {
  var date = new Date(dt);
  var year = date.getFullYear();
  var month = date.getMonth() + 1;
  month = month < 10 ? "0" + month : month;
  var day = date.getDate();
  day = day < 10 ? "0" + day : day;
  var hour = date.getHours();
  hour = hour < 10 ? "0" + hour : hour;
  var minute = date.getMinutes();
  minute = minute < 10 ? "0" + minute : minute;
  var second = date.getSeconds();
  second = second < 10 ? "0" + second : second;
  return (
    year + "-" + month + "-" + day + " " + hour + ":" + minute + ":" + second
  );
}
export default {
  name: "Dashboard",
  data() {
    return {
      table_height: 0,
      header: {},
      idc_total_list: [],
      assets_type_list: [],
      charts: "",
      approvalXData: [],
      legends: [],
      approvals: [],
      user_list: [],
      page_no: 1,
      page_size: 6,
      total: 1,
      table_users: [],
      is_mobile: config.isMobile,
    };
  },
  beforeCreate() {},
  mounted() {
    this.init_data();
    this.$nextTick(function() {
      this.drawLine("main");
    });
  },
  methods: {
    handleCurrentChange(val) {
      this.page_no = val;
      this.page_data();
    },
    page_data() {
      user_login_get({page_no:this.page_no, page_size:this.page_size}).then((response) => {
        if (response.code === 200) {
          var table_users = response.data.users;
          for(var i=0;i<table_users.length;i++){
            var item = table_users[i]
            var newDate = get_date_time(item.last_login);
            table_users[i]["last_login"] = newDate;
          }
          this.table_users = table_users;
          this.total = response.data.total;
        } else {
          this.table_users = [];
          this.total = 1;
        }
      });
    },
    init_data() {
      summary_machine_get().then((response) => {
        if (response.code == 200) {
          this.header = response.data.header;
          this.idc_total_list = response.data.idc_total_list;
          this.assets_type_list = response.data.assets_type_list;
        }
      });
      var height = document.getElementById("idc_table").clientHeight;
      this.table_height = height;
      this.page_data()
    },
    drawPieChart() {
      var legend_data = [];
      for (var i = 0; i < this.assets_type_list.length; i++) {
        var item = this.assets_type_list[i];
        legend_data.push(item.name);
      }
      this.chartPie = echarts.init(document.getElementById("chartPie"));
      this.chartPie.setOption({
        title: {
          text: "服务器类型",
          subtext: "服务器类型展示",
          x: "center",
        },
        tooltip: {
          trigger: "item",
          formatter: "{a} <br/>{b} : {c} ({d}%)",
        },
        legend: {
          orient: "vertical",
          left: "left",
          data: legend_data,
        },
        series: [
          {
            name: "访问来源",
            type: "pie",
            radius: "55%",
            center: ["50%", "60%"],
            data: this.assets_type_list,
            itemStyle: {
              emphasis: {
                shadowBlur: 10,
                shadowOffsetX: 0,
                shadowColor: "rgba(0, 0, 0, 0.5)",
              },
            },
          },
        ],
      });
    },
    drawCharts() {
      this.drawPieChart();
    },
    async drawLine(id) {
      const obj = await summary_approval_get();
      this.approvals = obj.data;
      const series = [];
      this.approvalXData = Object.keys(this.approvals[0].data);
      for (let i = 0; i < this.approvals.length; i++) {
        const item = this.approvals[i];
        series.push({
          name: item.approval_name,
          type: "line",
          smooth: true,
          itemStyle: {
            normal: {
              color: item.color,
            },
          },
          data: Object.values(item.data),
        });
        this.legends.push(item.approval_name);
      }
      this.charts = echarts.init(document.getElementById(id));
      this.charts.setOption({
        tooltip: {
          trigger: "axis",
        },
        legend: {
          data: this.legends,
        },
        grid: {
          left: "3%",
          right: "4%",
          bottom: "3%",
          containLabel: true,
        },
        toolbox: {
          show: true,
          feature: {
            saveAsImage: {},
          },
        },
        xAxis: {
          type: "category",
          boundaryGap: false,
          data: this.approvalXData,
        },
        yAxis: {
          type: "value",
        },
        series: series,
      });
    },
  },
};
</script>
<style lang="scss" scoped>
.dashboard-container {
  width: 99%;
  min-height: 680px;
  margin: 5px auto;
  // background-color: #fff;
  padding: 5px 0;
  .head {
    margin: 15px auto;
    .card-head {
      color: #aaa;
    }
    .card-box {
      .number {
        font-size: 38px;
        color: #666;
      }
      .online {
        color: green;
      }
      .danger {
        color: red;
      }
      .info {
        color: #999;
        font-size: 14px;
        margin-top: 10px;
      }
    }
  }
  .body {
    .title {
      font-size: 18px;
      color: #999;
    }
    .card-box {
      min-height: 150px;
    }
  }
}
</style>
