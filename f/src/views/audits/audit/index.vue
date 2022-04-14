<template>
  <div class="audit">
    <div v-if="!is_mobile">
      <div class="sub-title">审计日志</div>
      <el-row style="margin: 25px auto">
        <el-col :span="3" style="overflow: hidden">
          <el-date-picker
            v-model="formInline.start"
            type="datetime"
            placeholder="选择开始时间"
            value-format="yyyy-MM-dd HH:mm:ss"
          />
        </el-col>
        <el-col :span="3" style="overflow: hidden">
          <el-date-picker
            v-model="formInline.end"
            type="datetime"
            placeholder="选择结束时间"
            value-format="yyyy-MM-dd HH:mm:ss"
          />
        </el-col>
        <el-col :span="3">
          <el-select
            v-model="formInline.operater"
            placeholder="请选择操作用户"
            clearable
            @change="operater_change()"
          >
            <el-option
              v-for="item in user_list"
              :key="item.id"
              :label="item.username"
              :value="item.username"
            />
          </el-select>
        </el-col>
        <el-col :span="3">
          <el-select
            v-model="formInline.req_method"
            placeholder="请选择请求方法"
            clearable
            @change="method_change()"
          >
            <el-option
              v-for="item in req_methods"
              :key="item"
              :label="item"
              :value="item"
            />
          </el-select>
        </el-col>
        <el-col :span="5">
          <el-row>
            <el-col :span="19">
              <el-input
                v-model="formInline.req_url"
                placeholder="请输入请求地址"
              />
            </el-col>
            <el-col :span="4" :offset="1">
              <el-button type="primary" @click="url_change()"> 搜索 </el-button>
            </el-col>
          </el-row>
        </el-col>
      </el-row>
      <el-table
        v-loading="loading"
        tooltip-effect="dark"
        :data="audit_list"
        style="width: 100%"
      >
        <el-table-column prop="id" label="ID" width="55" />
        <el-table-column prop="operater" label="操作员" width="100" />
        <el-table-column prop="req_method" label="操作" width="180" />
        <el-table-column prop="req_url" label="操作对象" width="200">
          <template slot-scope="scope">
            {{
              scope.row.req_url.split("/")[
                scope.row.req_url.split("/").length - 1
              ]
            }}
          </template>
        </el-table-column>
        <el-table-column prop="req_data" label="参数列表">
          <template slot-scope="scope">
            <el-tree :data="scope.row.req_data" :props="defaultProps"></el-tree>
          </template>
        </el-table-column>
        <el-table-column prop="res_data" label="响应列表">
          <template slot-scope="scope">
            <el-tree :data="scope.row.res_data" :props="defaultProps"></el-tree>
          </template>
        </el-table-column>
        <el-table-column prop="operate_time" label="操作时间" width="150">
          <template slot-scope="scope">
            {{ scope.row.operate_time.replace("T", " ").split("+")[0] }}
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
    <div v-if="is_mobile">
      <div class="sub-title">审计日志</div>
      <el-row style="margin: 25px auto">
        <el-col :span="24" style="overflow: hidden">
          <el-date-picker
            v-model="formInline.start"
            type="datetime"
            style="width: 100%"
            placeholder="选择开始时间"
            value-format="yyyy-MM-dd HH:mm:ss"
          />
        </el-col>
        <el-col :span="24" style="overflow: hidden">
          <el-date-picker
            v-model="formInline.end"
            type="datetime"
            style="width: 100%"
            placeholder="选择结束时间"
            value-format="yyyy-MM-dd HH:mm:ss"
          />
        </el-col>
        <el-col :span="24">
          <el-select
            v-model="formInline.operater"
            placeholder="请选择操作用户"
            clearable
            style="width: 100%"
            @change="operater_change()"
          >
            <el-option
              v-for="item in user_list"
              :key="item.id"
              :label="item.username"
              :value="item.username"
            />
          </el-select>
        </el-col>
        <el-col :span="24">
          <el-select
            v-model="formInline.req_method"
            placeholder="请选择请求方法"
            clearable
            style="width: 100%"
            @change="method_change()"
          >
            <el-option
              v-for="item in req_methods"
              :key="item"
              :label="item"
              :value="item"
            />
          </el-select>
        </el-col>
        <el-col :span="24">
          <el-row>
            <el-col :span="19">
              <el-input
                v-model="formInline.req_url"
                placeholder="请输入请求地址"
              />
            </el-col>
            <el-col :span="4" :offset="1" style="margin-top: 5px">
              <el-button size="mini" type="primary" @click="url_change()">
                搜索
              </el-button>
            </el-col>
          </el-row>
        </el-col>
      </el-row>
      <el-table
        v-loading="loading"
        tooltip-effect="dark"
        :data="audit_list"
        style="width: 100%"
      >
        <el-table-column prop="id" label="ID" width="55" fixed />
        <el-table-column prop="operater" label="操作员" width="80" />
        <el-table-column prop="req_method" label="操作" width="80" />
        <el-table-column prop="req_url" label="操作对象" width="200">
          <template slot-scope="scope">
            {{
              scope.row.req_url.split("/")[
                scope.row.req_url.split("/").length - 1
              ]
            }}
          </template>
        </el-table-column>
        <el-table-column prop="req_data" label="参数列表" width="220">
          <template slot-scope="scope">
            <el-tree :data="scope.row.req_data" :props="defaultProps"></el-tree>
          </template>
        </el-table-column>
        <el-table-column prop="res_data" label="响应列表" width="220">
          <template slot-scope="scope">
            <el-tree :data="scope.row.res_data" :props="defaultProps"></el-tree>
          </template>
        </el-table-column>
        <el-table-column prop="operate_time" label="操作时间" width="180">
          <template slot-scope="scope">
            {{ scope.row.operate_time.replace("T", " ").replace("+08:00", "") }}
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
import { user_get } from "@/api/user";
import { audit_get } from "@/api/audit/index";
import config from "@/utils/config";
export default {
  name: "Audit",
  data() {
    return {
      user_list: [],
      audit_list: [],
      page_no: 1,
      page_size: 10,
      total: 1,
      formInline: {
        operater: "",
        req_url: "",
        req_method: "",
        start: "",
        end: "",
      },
      req_methods: ["POST", "PUT", "DELETE"],
      is_mobile: config.isMobile,
      loading: true,
      defaultProps: {
        children: "children",
        label: "label",
      },
    };
  },
  mounted() {
    this.init_data();
  },
  methods: {
    init_data() {
      user_get().then((response) => {
        if (response.code == 200) {
          this.user_list = response.data.users;
        }
      });
      this.page_data(this.formInline);
    },
    page_data() {
      this.formInline["page_no"] = this.page_no;
      this.formInline["page_size"] = this.page_size;
      audit_get(this.formInline).then((response) => {
        if (response.code === 200) {
          var audit_list = response.data.audit;
          var audit_list2 = [];
          for (var i = 0; i < audit_list.length; i++) {
            var audit = audit_list[i];
            if (audit !== null) {
              var o = [];
              // 请求信息处理
              for (var key in audit["req_data"]) {
                o.push({
                  label: key,
                  children: [{ label: audit["req_data"][key] }],
                });
              }
              audit["req_data"] = [{ label: "request data", children: o }];
              // 响应信息处理
              var res_data = audit["res_data"]["data"];
              if (res_data instanceof Array) {
                var res_list = [{ label: "list", children: [] }];
                for (var m = 0; m < res_data.length; m++) {
                  var kvs = res_data[m];
                  var mm = [];
                  for (var k in kvs) {
                    mm.push({
                      label: k,
                      children: [{ label: JSON.stringify(kvs[k]) }],
                    });
                  }
                  res_list[0]["children"] = mm;
                }
              } else {
                var res_list = [{ label: "obj", children: [] }];
                var data = audit["res_data"]["data"];
                var kvs = [];
                if (Object.keys(data) > 0) {
                  for (var key in data) {
                    kvs.push({
                      label: key,
                      children: [{ label: JSON.stringify(data[key]) }],
                    });
                  }
                } else {
                  kvs = [{ label: "{}" }];
                }
                res_list[0]["children"] = kvs;
              }
              var res_data_labels = [];
              res_data_labels.push({
                label: "code",
                children: [{ label: audit["res_data"]["code"] }],
              });
              res_data_labels.push({
                label: "msg",
                children: [{ label: audit["res_data"]["msg"] }],
              });
              res_data_labels.push({ label: "data", children: res_list });
              audit["res_data"] = [
                {
                  label: "response data: " + audit["res_data"]["code"],
                  children: res_data_labels,
                },
              ];
              audit_list2.push(audit)
            }
          }
          this.audit_list = audit_list2;
          this.total = response.data.total;
          this.loading = false;
        }
      });
    },
    operater_change() {
      this.page_data();
    },
    method_change() {
      this.page_data();
    },
    url_change() {
      this.page_data();
    },
    handleCurrentChange(val) {
      this.page_no = val;
      this.formInline["page_no"] = this.page_no;
      this.formInline["page_size"] = this.page_size;
      this.page_data(this.formInline);
    },
  },
};
</script>
<style lang="scss" scoped>
.audit {
  width: 99%;
  min-height: 680px;
  margin: 20px auto;
  background-color: #fff;
  padding: 20px;
  .sub-title {
    padding-top: 20px;
    text-align: center;
    font-size: 24px;
    font-weight: bold;
    padding-bottom: 10px;
    border-bottom: 2px solid #ccc;
  }
}
</style>
