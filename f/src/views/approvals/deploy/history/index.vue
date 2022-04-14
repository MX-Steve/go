<template>
  <div class="deploy-history">
    <div v-if="!is_mobile">
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
      <el-row style="margin-bottom: 15px">
        <el-col :span="4" style="">
          <el-date-picker
            v-model="formInline.start"
            type="datetime"
            placeholder="选择开始时间"
            value-format="yyyy-MM-dd HH:mm:ss"
            @change="time_search()"
          />
        </el-col>
        <el-col :span="4" style="">
          <el-date-picker
            v-model="formInline.end"
            type="datetime"
            placeholder="选择结束时间"
            value-format="yyyy-MM-dd HH:mm:ss"
            @change="time_search()"
          />
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="3">
          <el-select
            v-model="formInline.env"
            placeholder="请选择环境"
            clearable
            filterable
            @change="env_search()"
          >
            <el-option
              v-for="item in env_list"
              :key="item.id"
              :label="item.name"
              :value="item.name"
            />
          </el-select>
        </el-col>
        <el-col :span="3">
          <el-select
            v-model="formInline.project"
            placeholder="请选择项目"
            clearable
            filterable
            @change="project_search()"
          >
            <el-option
              v-for="item in project_list"
              :key="item.id"
              :label="item.name"
              :value="item.name"
            />
          </el-select>
        </el-col>
        <el-col :span="3">
          <el-select
            v-model="formInline.service_type"
            placeholder="请选择服务类型"
            clearable
            filterable
            @change="type_search()"
          >
            <el-option
              v-for="item in service_types"
              :key="item"
              :label="item"
              :value="item"
            />
          </el-select>
        </el-col>
        <el-col :span="3">
          <el-select
            v-model="formInline.service"
            placeholder="请选择服务"
            clearable
            filterable
            @change="service_search()"
          >
            <el-option
              v-for="item in service_list"
              :key="item.id"
              :label="item.name"
              :value="item.name"
            />
          </el-select>
        </el-col>
        <el-col :span="3">
          <el-select
            v-model="formInline.result"
            placeholder="请选择结果"
            clearable
            filterable
            @change="result_search()"
          >
            <el-option
              v-for="item in result_list"
              :key="item"
              :label="item"
              :value="item"
            />
          </el-select>
        </el-col>
        <el-col :span="6" :offset="3">
          <el-row>
            <el-col :span="19">
              <el-input
                v-model="formInline.service"
                placeholder="请输入服务名"
              />
            </el-col>
            <el-col :span="4" :offset="1">
              <el-button type="primary" size="small" @click="service_search()">
                搜索
              </el-button>
            </el-col>
          </el-row>
        </el-col>
      </el-row>
      <el-row style="margin: 15px auto">
        <el-col :span="24">
          <el-button
            v-if="del_btn_show !== 0"
            type="danger"
            size="mini"
            @click="del_btn"
          >
            移除
          </el-button>
        </el-col>
      </el-row>
      <el-table
        v-loading="loading"
        tooltip-effect="dark"
        :data="deploy_list"
        style="width: 100%"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column prop="id" label="ID" width="55" />
        <el-table-column prop="env" label="环境" width="150" />
        <el-table-column prop="project" label="项目" width="150" />
        <el-table-column prop="service" label="服务" width="150" />
        <el-table-column prop="service_type" label="类型" width="80" />
        <el-table-column prop="BrName" label="版本" />
        <el-table-column prop="deploy_time" label="发布时间" width="150" />
        <el-table-column prop="result" label="部署结果" />
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
      <el-row>
        <el-col :span="24" style="">
          <el-date-picker
            v-model="formInline.start"
            style="width: 100%"
            type="datetime"
            placeholder="选择开始时间"
            value-format="yyyy-MM-dd HH:mm:ss"
            @change="time_search()"
          />
        </el-col>
        <el-col :span="24" style="">
          <el-date-picker
            v-model="formInline.end"
            style="width: 100%"
            type="datetime"
            placeholder="选择结束时间"
            value-format="yyyy-MM-dd HH:mm:ss"
            @change="time_search()"
          />
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="24">
          <el-select
            v-model="formInline.env"
            style="width: 100%"
            placeholder="请选择环境"
            clearable
            filterable
          >
            <el-option
              v-for="item in env_list"
              :key="item.id"
              :label="item.name"
              :value="item.name"
            />
          </el-select>
        </el-col>
        <el-col :span="24">
          <el-select
            v-model="formInline.project"
            style="width: 100%"
            placeholder="请选择项目"
            clearable
            filterable
          >
            <el-option
              v-for="item in project_list"
              :key="item.id"
              :label="item.name"
              :value="item.name"
            />
          </el-select>
        </el-col>
        <el-col :span="24">
          <el-select
            v-model="formInline.service_type"
            placeholder="请选择服务类型"
            clearable
            filterable
            @change="type_search()"
          >
            <el-option
              v-for="item in service_types"
              :key="item"
              :label="item"
              :value="item"
            />
          </el-select>
        </el-col>
        <el-col :span="24">
          <el-select
            v-model="formInline.service"
            style="width: 100%"
            placeholder="请选择服务"
            clearable
            filterable
          >
            <el-option
              v-for="item in service_list"
              :key="item.id"
              :label="item.name"
              :value="item.name"
            />
          </el-select>
        </el-col>
        <el-col :span="24">
          <el-select
            v-model="formInline.result"
            style="width: 100%"
            placeholder="请选择结果"
            clearable
            filterable
          >
            <el-option
              v-for="item in result_list"
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
                v-model="formInline.service"
                placeholder="请输入服务名"
              />
            </el-col>
            <el-col :span="4" :offset="1">
              <el-button type="primary" size="small"> 搜索 </el-button>
            </el-col>
          </el-row>
        </el-col>
      </el-row>
      <el-row style="margin: 15px auto">
        <el-col :span="24">
          <el-button
            v-if="del_btn_show !== 0"
            type="danger"
            size="mini"
            @click="del_btn"
          >
            移除
          </el-button>
          <el-button
            v-if="down_btn_show"
            type="success"
            size="mini"
            icon="el-icon-download"
            @click="download_excel"
          />
        </el-col>
      </el-row>
      <el-table
        tooltip-effect="dark"
        :data="deploy_list"
        style="width: 100%"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column prop="id" label="ID" width="55" fixed />
        <el-table-column prop="env" label="环境" fixed />
        <el-table-column prop="project" label="项" />
        <el-table-column prop="service" label="服务" />
        <el-table-column prop="service_type" label="类型" width="80" />
        <el-table-column prop="BrName" label="版本" />
        <el-table-column prop="deploy_time" label="发布时间" width="150" />
        <el-table-column prop="result" label="部署结果" width="120" />
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
import config from "@/utils/config";
import { env_get, project_get, service_get } from "@/api/project-list";
import {
  deploy_history_get,
  deploy_history_cud,
} from "@/api/approval/deploy_history";
import { get_date_time } from "@/utils/assets";
import { btn_check } from "@/api/btn";
import { DeployHistoryPage } from "@/utils/auth";
import { choices_get } from "@/api/approval/choices";
export default {
  name: "DeployHistory",
  data() {
    return {
      is_mobile: config.isMobile,
      loading: true,
      deploy_list: [],
      page_no: 1,
      page_size: 10,
      total: 1,
      env_list: [],
      project_list: [],
      service_list: [],
      result_list: ["SUCCESS", "FAILURE"],
      formInline: {
        start: "",
        end: "",
        env: "",
        project: "",
        service: "",
        service_type: "",
        result: "",
      },
      multipleSelection: [],
      ids: [],
      DeployHistoryPage: DeployHistoryPage,
      del_btn_show: 0,
      down_btn_show: 0,
      service_types: [],
    };
  },
  mounted() {
    this.init_data();
  },
  methods: {
    init_data() {
      this.formInline.start = get_date_time(-30);
      this.formInline.end = get_date_time(0);
      env_get().then((response) => {
        if (response.code === 200) {
          this.env_list = response.data;
        }
      });
      project_get({ page_no: 1, page_size: 999 }).then((response) => {
        if (response.code === 200) {
          this.project_list = response.data.project_infos;
        }
      });
      service_get({ page_no: 1, page_size: 999 }).then((response) => {
        if (response.code === 200) {
          this.service_list = response.data.service_infos;
        }
      });
      btn_check(this.DeployHistoryPage).then((response) => {
        if (response.code === 200) {
          this.del_btn_show = response.data.del;
          this.down_btn_show = response.data.down;
        }
      });
      this.page_data();
      choices_get({
        table_name: "BusinessServices",
        column: "type_choices",
      }).then((response) => {
        if (response.code === 200) {
          this.service_types = response.data;
        }
      });
    },
    formatJson(filterVal, jsonData) {
      return jsonData.map((v) => filterVal.map((j) => v[j]));
    },
    async download_excel() {
      const query = {
        page_no: 1,
        page_size: 999,
        env: this.formInline.env,
        start: this.formInline.start,
        end: this.formInline.end,
        project: this.formInline.project,
        service: this.formInline.service,
        service_type: this.formInline.service_type,
        result: this.formInline.result,
      };
      const historyResObj = await deploy_history_get(query);
      const histories = historyResObj.data.histories;
      const history = histories[0];
      const headers = Object.keys(history);
      const data = this.formatJson(headers, histories);
      import("@/vendor/Export2Excel").then((excel) => {
        const tHeader = headers;
        excel.export_json_to_excel({
          header: tHeader, // 表头 必填
          data, // 具体数据 必填
          filename: "histories", // 非必填
          autoWidth: true, // 非必填
          bookType: "xlsx", // 非必填
        });
      });
    },
    handleCurrentChange(val) {
      this.page_no = val;
      this.page_data();
    },
    page_data() {
      const query = {
        page_no: this.page_no,
        page_size: this.page_size,
        env: this.formInline.env,
        start: this.formInline.start,
        end: this.formInline.end,
        project: this.formInline.project,
        service: this.formInline.service,
        service_type: this.formInline.service_type,
        result: this.formInline.result,
      };
      deploy_history_get(query).then((response) => {
        if (response.code === 200) {
          this.deploy_list = response.data.histories;
          this.total = response.data.total;
          this.loading = false;
        }
      });
    },
    time_search() {
      this.total = 1;
      this.page_data();
    },
    env_search() {
      this.total = 1;
      this.page_data();
    },
    project_search() {
      this.total = 1;
      service_get({
        project: this.formInline.project,
        page_no: 1,
        page_size: 999,
      }).then((response) => {
        if (response.code === 200) {
          this.service_list = response.data.service_infos;
        }
      });
      this.page_data();
    },
    type_search() {
      this.total = 1;
      this.page_no = 1;
      this.page_data();
    },
    service_search() {
      this.total = 1;
      this.page_data();
    },
    result_search() {
      this.total = 1;
      this.page_data();
    },
    handleSelectionChange(val) {
      this.multipleSelection = val;
    },
    del_btn() {
      for (var i = 0; i < this.multipleSelection.length; i++) {
        this.ids.push(this.multipleSelection[i]["id"]);
      }
      deploy_history_cud("delete", { ids: this.ids }).then((response) => {
        if (response.code === 200) {
          this.$message({
            message: response.msg,
            type: "success",
          });
          location.reload(0);
        }
      });
    },
  },
};
</script>

<style lang="scss" scoped>
.deploy-history {
  min-height: 450px;
  width: 98%;
  margin: 15px auto;
  background-color: #fff;
  padding: 15px;
}
</style>
