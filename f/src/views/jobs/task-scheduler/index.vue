<template>
  <div class="task-scheduler">
    <div v-if="!is_mobile">
      <div class="task-header">
        <el-card shadow="hover">
          <el-form :inline="true" :model="formInline" class="demo-form-inline">
            <el-form-item label="周期类型">
              <el-select
                v-model="formInline.type_id"
                clearable
                placeholder="请选择类型"
              >
                <template v-for="item in type_list">
                  <el-option
                    v-if="item.table == 'interval'"
                    :key="item.id + item.table"
                    :value="item.table + '_' + item.id"
                    :label="
                      '每过 ' + item.every + ' ' + item.period + ' 执行一次'
                    "
                    @click.native="option_click_handler(item)"
                  />
                  <el-option
                    v-else
                    :key="item.id + item.table"
                    :value="item.table + '_' + item.id"
                    :label="
                      'Crontab: ' +
                        item.minute +
                        ' ' +
                        item.hour +
                        ' ' +
                        item.day_of_month +
                        ' ' +
                        item.month_of_year +
                        ' ' +
                        item.day_of_week
                    "
                    @click.native="option_click_handler(item)"
                  />
                </template>
              </el-select>
            </el-form-item>
            <el-form-item label="任务名称">
              <el-input
                v-model="formInline.job_name"
                placeholder="请输入任务名称"
              />
            </el-form-item>
            <el-form-item>
              <el-button @click="job_changed"> 搜索 </el-button>
            </el-form-item>
          </el-form>
        </el-card>
      </div>
      <div class="task-box">
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>任务列表</span>
            <el-button-group style="float: right">
              <el-button
                size="mini"
                :type="active_btn == 'all' ? 'primary' : ''"
                @click="search_range('all')"
              >
                全部
              </el-button>
              <el-button
                size="mini"
                :type="active_btn == 'on' ? 'primary' : ''"
                @click="search_range('on')"
              >
                已激活
              </el-button>
              <el-button
                size="mini"
                :type="active_btn == 'off' ? 'primary' : ''"
                @click="search_range('off')"
              >
                未激活
              </el-button>
            </el-button-group>
            <el-button
              v-if="add_btn_show"
              style="float: right; margin-right: 25px"
              type="primary"
              size="mini"
              @click="create_btn"
            >
              + 新建
            </el-button>
          </div>
          <el-table v-loading="loading" :data="table_data" style="width: 100%">
            <el-table-column prop="id" label="ID" width="55" />
            <el-table-column prop="name" label="任务名称" width="220" />
            <el-table-column prop="task" label="任务对象" />
            <el-table-column label="最新状态">
              <template slot-scope="scope">
                <el-button
                  v-if="scope.row.enabled == 1"
                  style="color: green"
                  size="mini"
                  @click="enabled_changed(scope.row)"
                >
                  已激活
                </el-button>
                <el-button
                  v-else
                  style="color: red"
                  size="mini"
                  @click="enabled_changed(scope.row)"
                >
                  未激活
                </el-button>
              </template>
            </el-table-column>
            <el-table-column label="更新时间">
              <template slot-scope="scope">
                {{ scope.row.date_changed.replace("T", " ").split(".")[0] }}
              </template>
            </el-table-column>
            <el-table-column label="操作">
              <template slot-scope="scope">
                <el-button
                  type="text"
                  style="margin-right: 15px"
                  @click="more_btn(scope.row)"
                >
                  详情
                </el-button>
                <el-button
                  v-if="edit_btn_show"
                  type="text"
                  style="margin-right: 15px"
                  @click="edit_btn(scope.row)"
                >
                  编辑
                </el-button>
                <el-dropdown>
                  <span class="el-dropdown-link">
                    更多 <i class="el-icon-arrow-down el-icon--right" />
                  </span>
                  <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item
                      v-if="run_btn_show === 1"
                      @click.native="run_one(scope.row)"
                    >
                      执行一次
                    </el-dropdown-item>
                    <el-dropdown-item @click.native="history_btn(scope.row)">
                      历史记录
                    </el-dropdown-item>
                    <el-dropdown-item
                      v-if="del_btn_show"
                      style="color: red"
                      @click.native="del_btn(scope.row)"
                    >
                      删除
                    </el-dropdown-item>
                  </el-dropdown-menu>
                </el-dropdown>
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
        </el-card>
      </div>
      <el-dialog
        v-loading="loading"
        :title="dialogTitle"
        :visible.sync="dialogVisible"
        width="45%"
        :before-close="handleClose"
      >
        <el-form
          ref="task"
          :model="task"
          :rules="taskRules"
          label-width="180px"
        >
          <el-form-item label="任务类型" prop="type_id">
            <el-select
              v-model="task.type_id"
              placeholder="请选择任务类型"
              style="width: 80%"
            >
              <template v-for="item in type_list">
                <el-option
                  v-if="item.table == 'interval'"
                  :key="item.table + '2_' + item.id"
                  :value="item.table + '_' + item.id"
                  :label="
                    '每过 ' + item.every + ' ' + item.period + ' 执行一次'
                  "
                  @click.native="option_click_handler2(item)"
                />
                <el-option
                  v-else
                  :key="item.table + '2_' + item.id"
                  :value="item.table + '_' + item.id"
                  :label="
                    'Crontab: ' +
                      item.minute +
                      ' ' +
                      item.hour +
                      ' ' +
                      item.day_of_month +
                      ' ' +
                      item.month_of_year +
                      ' ' +
                      item.day_of_week
                  "
                  @click.native="option_click_handler2(item)"
                />
              </template>
            </el-select>
            <el-button style="margin-left: 20px" @click="add_type_btn">
              添加
            </el-button>
          </el-form-item>
          <el-form-item label="任务名称" prop="name">
            <el-input v-model="task.name" placeholder="请输入任务名称" />
          </el-form-item>
          <el-form-item label="任务内容" prop="task">
            <el-input
              v-model="task.task"
              type="textarea"
              :autosize="{ minRows: 2, maxRows: 4 }"
              placeholder="请输入内容"
            />
          </el-form-item>
          <el-form-item label="描述信息" prop="description">
            <el-input v-model="task.description" placeholder="请输入描述信息" />
          </el-form-item>
          <el-form-item label="一次性任务" prop="one_off">
            <el-checkbox v-model="task.one_off" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="dialog_confirm()">确定</el-button>
            <el-button @click="dialog_cancel()">取消</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
      <el-dialog
        :title="dialogTitle2"
        :visible.sync="dialogVisible2"
        width="45%"
        :before-close="handleClose2"
      >
        <el-row v-for="(v, k) in task" :key="k" class="text item">
          <el-col :span="6"> {{ k }} </el-col>
          <el-col :span="18" style="color: green"> : {{ v }} </el-col>
        </el-row>
      </el-dialog>
      <el-dialog
        title="任务执行历史记录"
        :visible.sync="dialogVisible3"
        width="80%"
        :before-close="handleClose3"
      >
        <el-table :data="history_tasks" style="width: 100%">
          <el-table-column prop="id" label="ID" width="75" />
          <el-table-column prop="task_name" label="任务名称" width="200" />
          <el-table-column prop="status" label="任务状态" width="90">
            <template slot-scope="scope">
              <span
                :style="{
                  color: scope.row.status == 'SUCCESS' ? 'green' : 'red',
                }"
              >
                {{ scope.row.status }}
              </span>
            </template>
          </el-table-column>
          <el-table-column prop="worker" label="工作节点" width="180" />
          <el-table-column prop="result" label="执行结果">
            <template slot-scope="scope">
              {{ JSON.parse(scope.row.result) }}
            </template>
          </el-table-column>
          <el-table-column prop="date_created" label="执行时间" width="150">
            <template slot-scope="scope">
              {{ scope.row.date_created.replace("T", " ").split(".")[0] }}
            </template>
          </el-table-column>
        </el-table>
      </el-dialog>
      <el-dialog title="添加类型" :visible.sync="add_type_check" width="45%">
        <el-form ref="type_dict" :model="type_dict" label-width="180px">
          <el-form-item label="选择类型" prop="type">
            <el-select v-model="type_dict.type" placeholder="请选择类型">
              <el-option label="时间间隔" value="interval" />
              <el-option label="周期性计划任务" value="crontab" />
            </el-select>
          </el-form-item>
          <el-form-item v-if="type_dict.type == 'interval'" label="长度">
            <el-input
              v-model="type_dict.every"
              type="number"
              placeholder="请输入间隔长度(s)"
            />
          </el-form-item>
          <template v-if="type_dict.type == 'crontab'">
            <el-form-item label="分: " prop="minute">
              <el-input v-model="type_dict.minute" placeholder="请输入分钟" />
            </el-form-item>
            <el-form-item label="时: " prop="hour">
              <el-input v-model="type_dict.hour" placeholder="请输入小时" />
            </el-form-item>
            <el-form-item label="日: " prop="day_of_month">
              <el-input
                v-model="type_dict.day_of_month"
                placeholder="请输入日"
              />
            </el-form-item>
            <el-form-item label="月: " prop="month_of_year">
              <el-input
                v-model="type_dict.month_of_year"
                placeholder="请输入月"
              />
            </el-form-item>
            <el-form-item label="周: " prop="day_of_week">
              <el-input
                v-model="type_dict.day_of_week"
                placeholder="请输入周"
              />
            </el-form-item>
          </template>
          <el-form-item>
            <el-button @click="dialog_type_confirm"> 添加 </el-button>
            <el-button @click="dialog_type_cancel"> 取消 </el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
    <div v-if="is_mobile">
      <div class="task-header">
        <el-card shadow="hover">
          <el-form :inline="true" :model="formInline" class="demo-form-inline">
            <el-form-item label="周期类型">
              <el-select
                v-model="formInline.type_id"
                clearable
                style="width: 100%"
                placeholder="请选择类型"
              >
                <template v-for="item in type_list">
                  <el-option
                    v-if="item.table == 'interval'"
                    :key="item.id + item.table"
                    :value="item.table + '_' + item.id"
                    :label="
                      '每过 ' + item.every + ' ' + item.period + ' 执行一次'
                    "
                    @click.native="option_click_handler(item)"
                  />
                  <el-option
                    v-else
                    :key="item.id + item.table"
                    :value="item.table + '_' + item.id"
                    :label="
                      'Crontab: ' +
                        item.minute +
                        ' ' +
                        item.hour +
                        ' ' +
                        item.day_of_month +
                        ' ' +
                        item.month_of_year +
                        ' ' +
                        item.day_of_week
                    "
                    @click.native="option_click_handler(item)"
                  />
                </template>
              </el-select>
            </el-form-item>
            <el-form-item label="任务名称">
              <el-input
                v-model="formInline.job_name"
                style="width: 100%"
                placeholder="请输入任务名称"
              />
            </el-form-item>
            <el-form-item>
              <el-button @click="job_changed"> 搜索 </el-button>
            </el-form-item>
          </el-form>
        </el-card>
      </div>
      <div class="task-box">
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>任务列表</span>
            <el-button
              v-if="add_btn_show"
              style="float: right"
              type="primary"
              size="mini"
              @click="create_btn"
            >
              + 新建
            </el-button>
          </div>
          <div>
            <el-button-group style="float: right">
              <el-button
                size="mini"
                :type="active_btn == 'all' ? 'primary' : ''"
                @click="search_range('all')"
              >
                全部
              </el-button>
              <el-button
                size="mini"
                :type="active_btn == 'on' ? 'primary' : ''"
                @click="search_range('on')"
              >
                已激活
              </el-button>
              <el-button
                size="mini"
                :type="active_btn == 'off' ? 'primary' : ''"
                @click="search_range('off')"
              >
                未激活
              </el-button>
            </el-button-group>
          </div>
          <el-table v-loading="loading" :data="table_data" style="width: 100%">
            <el-table-column prop="id" label="ID" width="35" fixed />
            <el-table-column prop="name" label="任务名称" width="150" fixed />
            <el-table-column
              prop="task"
              label="任务对象"
              width="200"
            />
            <el-table-column label="最新状态">
              <template slot-scope="scope">
                <el-button
                  v-if="scope.row.enabled == 1"
                  style="color: green"
                  size="mini"
                  @click="enabled_changed(scope.row)"
                >
                  已激活
                </el-button>
                <el-button
                  v-else
                  style="color: red"
                  size="mini"
                  @click="enabled_changed(scope.row)"
                >
                  未激活
                </el-button>
              </template>
            </el-table-column>
            <el-table-column label="更新时间" width="150">
              <template slot-scope="scope">
                {{ scope.row.date_changed.replace("T", " ").split(".")[0] }}
              </template>
            </el-table-column>
            <el-table-column label="操作" width="180">
              <template slot-scope="scope">
                <el-button type="text" @click="more_btn(scope.row)">
                  详情
                </el-button>
                <el-button
                  v-if="edit_btn_show"
                  type="text"
                  style="margin-right: 15px"
                  @click="edit_btn(scope.row)"
                >
                  编辑
                </el-button>
                <el-dropdown>
                  <span class="el-dropdown-link">
                    更多 <i class="el-icon-arrow-down el-icon--right" />
                  </span>
                  <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item
                      v-if="run_btn_show === 1"
                      @click.native="run_one(scope.row)"
                    >
                      执行一次
                    </el-dropdown-item>
                    <el-dropdown-item @click.native="history_btn(scope.row)">
                      历史记录
                    </el-dropdown-item>
                    <el-dropdown-item
                      v-if="del_btn_show"
                      style="color: red"
                      @click.native="del_btn(scope.row)"
                    >
                      删除
                    </el-dropdown-item>
                  </el-dropdown-menu>
                </el-dropdown>
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
        </el-card>
      </div>
      <el-dialog
        v-loading="loading"
        :title="dialogTitle"
        :visible.sync="dialogVisible"
        width="100%"
        :before-close="handleClose"
      >
        <el-form ref="task" :model="task" :rules="taskRules" label-width="85px">
          <el-form-item label="任务类型" prop="type_id">
            <el-select
              v-model="task.type_id"
              placeholder="请选择任务类型"
              style="width: 100%"
            >
              <template v-for="item in type_list">
                <el-option
                  v-if="item.table == 'interval'"
                  :key="item.table + '2_' + item.id"
                  :value="item.table + '_' + item.id"
                  :label="
                    '每过 ' + item.every + ' ' + item.period + ' 执行一次'
                  "
                  @click.native="option_click_handler2(item)"
                />
                <el-option
                  v-else
                  :key="item.table + '2_' + item.id"
                  :value="item.table + '_' + item.id"
                  :label="
                    'Crontab: ' +
                      item.minute +
                      ' ' +
                      item.hour +
                      ' ' +
                      item.day_of_month +
                      ' ' +
                      item.month_of_year +
                      ' ' +
                      item.day_of_week
                  "
                  @click.native="option_click_handler2(item)"
                />
              </template>
            </el-select>
            <el-button @click="add_type_btn"> 添加 </el-button>
          </el-form-item>
          <el-form-item label="任务名称" prop="name">
            <el-input v-model="task.name" placeholder="请输入任务名称" />
          </el-form-item>
          <el-form-item label="任务内容" prop="task">
            <el-input
              v-model="task.task"
              type="textarea"
              :autosize="{ minRows: 2, maxRows: 4 }"
              placeholder="请输入内容"
            />
          </el-form-item>
          <el-form-item label="描述信息" prop="description">
            <el-input v-model="task.description" placeholder="请输入描述信息" />
          </el-form-item>
          <el-form-item label="一次性任务" prop="one_off">
            <el-checkbox v-model="task.one_off" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="dialog_confirm()">确定</el-button>
            <el-button @click="dialog_cancel()">取消</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
      <el-dialog
        :title="dialogTitle2"
        :visible.sync="dialogVisible2"
        width="100%"
        :before-close="handleClose2"
      >
        <el-row v-for="(v, k) in task" :key="k" class="text item">
          <el-col :span="6"> {{ k }} </el-col>
          <el-col :span="18" style="color: green"> : {{ v }} </el-col>
        </el-row>
      </el-dialog>
      <el-dialog
        title="任务执行历史记录"
        :visible.sync="dialogVisible3"
        width="100%"
        :before-close="handleClose3"
      >
        <el-table :data="history_tasks" style="width: 100%">
          <el-table-column prop="id" label="ID" width="60" fixed />
          <el-table-column prop="task_name" label="任务名称" width="200" />
          <el-table-column prop="status" label="任务状态" width="90" fixed>
            <template slot-scope="scope">
              <span
                :style="{
                  color: scope.row.status == 'SUCCESS' ? 'green' : 'red',
                }"
              >
                {{ scope.row.status }}
              </span>
            </template>
          </el-table-column>
          <el-table-column prop="worker" label="工作节点" width="180" />
          <el-table-column prop="result" label="执行结果" width="300">
            <template slot-scope="scope">
              {{ JSON.parse(scope.row.result) }}
            </template>
          </el-table-column>
          <el-table-column prop="date_created" label="执行时间" width="150">
            <template slot-scope="scope">
              {{ scope.row.date_created.replace("T", " ").split(".")[0] }}
            </template>
          </el-table-column>
        </el-table>
      </el-dialog>
      <el-dialog title="添加类型" :visible.sync="add_type_check" width="100%">
        <el-form ref="type_dict" :model="type_dict" label-width="80px">
          <el-form-item label="选择类型" prop="type">
            <el-select
              v-model="type_dict.type"
              style="width: 100%"
              placeholder="请选择类型"
            >
              <el-option label="时间间隔" value="interval" />
              <el-option label="周期性计划任务" value="crontab" />
            </el-select>
          </el-form-item>
          <el-form-item v-if="type_dict.type == 'interval'" label="长度">
            <el-input
              v-model="type_dict.every"
              type="number"
              placeholder="请输入间隔长度(s)"
            />
          </el-form-item>
          <template v-if="type_dict.type == 'crontab'">
            <el-form-item label="分: " prop="minute">
              <el-input v-model="type_dict.minute" placeholder="请输入分钟" />
            </el-form-item>
            <el-form-item label="时: " prop="hour">
              <el-input v-model="type_dict.hour" placeholder="请输入小时" />
            </el-form-item>
            <el-form-item label="日: " prop="day_of_month">
              <el-input
                v-model="type_dict.day_of_month"
                placeholder="请输入日"
              />
            </el-form-item>
            <el-form-item label="月: " prop="month_of_year">
              <el-input
                v-model="type_dict.month_of_year"
                placeholder="请输入月"
              />
            </el-form-item>
            <el-form-item label="周: " prop="day_of_week">
              <el-input
                v-model="type_dict.day_of_week"
                placeholder="请输入周"
              />
            </el-form-item>
          </template>
          <el-form-item>
            <el-button @click="dialog_type_confirm"> 添加 </el-button>
            <el-button @click="dialog_type_cancel"> 取消 </el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
  </div>
</template>
<script>
import { get_item_name } from "@/utils/assets";
import {
  tasks_get,
  tasks_cud,
  intervals_get,
  intervals_cud,
  results_get,
  crontab_get,
  crontab_cud,
  deploy_one,
} from "@/api/task-scheduler/task";
import { btn_check } from "@/api/btn";
import config from "@/utils/config";
import { JobsPage } from "@/utils/auth";
export default {
  name: "TaskScheduler",
  data() {
    return {
      add_type_check: false,
      type_dict: {
        every: "",
        type: "",
        minute: "*",
        hour: "*",
        day_of_week: "*",
        day_of_month: "*",
        month_of_year: "*",
      },
      dialogTitle: "新建任务",
      dialogVisible: false,
      dialogTitle2: "任务详情",
      dialogVisible2: false,
      dialogVisible3: false,
      type_list: [],
      task: {
        id: "",
        name: "",
        task: "",
        description: "",
        one_off: "",
        interval_id: "",
        total_run_count: 0,
        headers: "{}",
        enabled: 0,
        args: "[]",
        kwargs: "{}",
        type_id: "",
      },
      taskRules: {
        name: [{ required: true, message: "请输入任务名称", trigger: "blur" }],
        minute: [{ required: true, message: "请输入分钟", trigger: "blur" }],
        hour: [{ required: true, message: "请输入小时", trigger: "blur" }],
        day_of_week: [
          { required: true, message: "请输入星期", trigger: "blur" },
        ],
        day_of_month: [
          { required: true, message: "请输入几号", trigger: "blur" },
        ],
        month_of_year: [
          { required: true, message: "请输入月份", trigger: "blur" },
        ],
        task: [{ required: true, message: "请输入任务位置", trigger: "blur" }],
        type_id: [{ required: true, message: "请选择周期", trigger: "change" }],
      },
      table_data: [],
      formInline: {
        status_id: "",
        type_id: "",
        job_name: "",
      },
      page_no: 1,
      page_size: 10,
      total: 1,
      active_btn: "all",
      history_tasks: [],
      dialog_task_type: "",
      add_btn_show: 0,
      edit_btn_show: 0,
      del_btn_show: 0,
      run_btn_show: 0,
      is_mobile: config.isMobile,
      JobsPage: JobsPage,
      loading: true,
    };
  },
  beforeMount() {
    this.init_data();
  },
  mounted() {},
  methods: {
    run_one(data){
      deploy_one(data).then(response=>{
        if(response.code === 200){
          this.$message({
            message: response.msg,
            type: "success"
          })
        }
      })
    },
    option_click_handler(data) {
      const query = {
        page_no: this.page_no,
        page_size: this.page_size,
      };
      if (data.table == "interval") {
        query["interval_id"] = data.id;
      } else {
        query["crontab_id"] = data.id;
      }
      tasks_get(query).then((response) => {
        if (response.code == 200) {
          this.table_data = response.data.tasks;
          this.total = response.data.total;
        }
      });
    },
    option_click_handler2(data) {
      if (this.task.type_id.indexOf("interval") !== -1) {
        this.task.interval_id = this.task.type_id.split("_")[1];
        this.task.crontab_id = null;
      } else {
        this.task.interval_id = null;
        this.task.crontab_id = this.task.type_id.split("_")[1];
      }
    },
    run_test_btn(row) {
      console.log(row);
    },
    dialog_type_confirm() {
      console.log(this.type_dict);
      if (this.type_dict.type == "crontab") {
        delete this.type_dict["type"];
        delete this.type_dict["every"];
        crontab_cud("post", this.type_dict).then((response) => {
          this.response_refresh_func(response);
        });
      } else if (this.type_dict.type == "interval") {
        delete this.type_dict["type"];
        delete this.type_dict["minute"];
        delete this.type_dict["hour"];
        delete this.type_dict["day_of_week"];
        delete this.type_dict["day_of_month"];
        delete this.type_dict["month_of_year"];
        intervals_cud("post", this.type_dict).then((response) => {
          this.response_refresh_func(response);
        });
      } else {
        this.$message({
          message: "所选择类型不对",
          type: "warning",
        });
      }
    },
    dialog_type_cancel() {
      this.add_type_check = false;
      this.type_dict = {};
    },
    history_btn(row) {
      results_get({ task_name: row.task }).then((response) => {
        if (response.code === 200) {
          this.history_tasks = response.data;
          this.dialogVisible3 = true;
        }
      });
    },
    add_type_btn() {
      this.add_type_check = true;
    },
    more_btn(row) {
      this.task = row;
      this.dialogVisible2 = true;
    },
    del_btn(row) {
      this.loading = true;
      tasks_cud("delete", row).then((response) => {
        this.response_refresh_func(response);
      });
    },
    create_btn() {
      this.dialogTitle = "新建任务";
      this.dialogVisible = true;
    },
    response_func(response) {
      if (response.code === 200) {
        this.loading = false;
        this.dialogVisible = false;
        this.$message({
          message: response.msg,
          type: "success",
        });
      }
    },
    response_refresh_func(response) {
      if (response.code === 200) {
        this.loading = false;
        this.dialogVisible = false;
        this.$message({
          message: response.msg,
          type: "success",
        });
        setTimeout(() => {
          location.reload(0);
        }, 200);
      }
    },
    handleClose() {
      this.task = {};
      this.dialogVisible = false;
    },
    handleClose2() {
      this.task = {};
      this.dialogVisible2 = false;
    },
    handleClose3() {
      this.history_tasks = [];
      this.dialogVisible3 = false;
    },
    edit_btn(row) {
      this.task = row;
      if (this.task.interval_id) {
        this.task.type_id = "interval_" + this.task.interval_id;
      } else {
        this.task.type_id = "crontab_" + this.task.crontab_id;
      }
      this.dialogVisible = true;
      this.dialogTitle = "编辑任务";
    },
    async init_data() {
      btn_check(this.JobsPage).then((response) => {
        if (response.code == 200) {
          this.add_btn_show = response.data.add;
          this.edit_btn_show = response.data.edit;
          this.del_btn_show = response.data.del;
          this.run_btn_show = response.data.conn;
        }
      });
      this.page_data();
      let result = [];
      const intervals_result = await intervals_get();
      const crontab_result = await crontab_get();
      result = intervals_result.data;
      for (var i = 0; i < crontab_result.data.length; i++) {
        var item = crontab_result.data[i];
        result.push(item);
      }
      for (var i = 0; i < result.length; i++) {
        var item = result[i];
        result["type_id"] = result["table"] + "_" + result["id"];
      }
      this.type_list = result;
    },
    page_data() {
      const query = {
        page_no: this.page_no,
        page_size: this.page_size,
      };
      tasks_get(query).then((response) => {
        if (response.code == 200) {
          const table_data = response.data.tasks;
          for (var i = 0; i < table_data.length; i++) {
            const task = table_data[i];
            if (task.one_off === 1) {
              table_data[i]["one_off"] = true;
            } else {
              table_data[i]["one_off"] = false;
            }
          }
          this.table_data = table_data;
          this.total = response.data.total;
          this.loading = false;
        }
      });
    },
    get_name(id, list) {
      return get_item_name(id, list);
    },
    type_changed() {
      const query = {
        page_no: this.page_no,
        page_size: this.page_size,
        interval_id: this.formInline.type_id,
      };
      tasks_get(query).then((response) => {
        if (response.code == 200) {
          this.table_data = response.data.tasks;
          this.total = response.data.total;
        }
      });
    },
    job_changed() {
      tasks_get({ name: this.formInline.job_name }).then((response) => {
        if (response.code == 200) {
          this.table_data = response.data.tasks;
          this.total = response.data.total;
        }
      });
    },
    handleCurrentChange(val) {
      this.page_no = val;
    },
    enabled_changed(row) {
      const data = {
        enabled: !row.enabled,
        id: row.id,
        type: "enabled",
      };
      if (this.edit_btn_show !== 0) {
        this.loading = true;
        tasks_cud("put", data).then((response) => {
          if (response.code == 200) {
            this.loading = false;
            this.$message({
              message: response.msg,
              type: "success",
            });
            location.reload(0);
          }
        });
      }
    },
    search_range(type) {
      this.active_btn = type;
      const query = {
        page_no: this.page_no,
        page_size: this.page_size,
      };
      if (type == "on") {
        query["enabled"] = 1;
      } else if (type == "off") {
        query["enabled"] = 0;
      }
      tasks_get(query).then((response) => {
        if (response.code == 200) {
          this.table_data = response.data.tasks;
          this.total = response.data.total;
        }
      });
    },
    dialog_confirm() {
      if (this.dialogTitle.indexOf("新建") !== -1) {
        if (this.task.one_off) {
          this.task.one_off = 1;
        } else {
          this.task.one_off = 0;
        }
        tasks_cud("post", this.task).then((response) => {
          if (response.code == 200) {
            this.response_refresh_func(response);
          }
        });
      } else {
        this.task.type = "modify";
        this.loading = true;
        if (this.dialog_task_type == "crontab") {
          this.task.crontab_id = this.task.interval_id;
          this.task.interval_id = "";
        }
        tasks_cud("put", this.task).then((response) => {
          if (response.code == 200) {
            this.response_func(response);
          }
        });
      }
    },
    dialog_cancel() {
      this.task = {
        id: "",
        name: "",
        task: "",
        description: "",
        one_off: "",
        interval_id: "",
      };
      this.dialogVisible = false;
    },
  },
};
</script>
<style lang="scss" scoped>
.task-scheduler {
  padding: 5px;
  background-color: #fff;
  min-height: 650px;
  .task-header {
    width: 90%;
    margin: 15px auto;
  }
  .task-box {
    .clearfix:before,
    .clearfix:after {
      display: table;
      content: "";
    }
    .clearfix:after {
      clear: both;
    }
    .box-card {
      width: 90%;
      margin: 15px auto;
    }
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
