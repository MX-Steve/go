<template>
  <div class="approval-manage">
    <div v-if="!is_mobile" class="web">
      <el-steps :active="step" finish-status="success" simple align-center>
        <el-step title="基础信息" @click.native="change_step(0)" />
        <el-step title="表单设计" @click.native="change_step(1)" />
        <el-step title="流程设计" @click.native="change_step(2)" />
        <el-step
          v-if="add_btn_show === 1"
          title="激活使用"
          @click.native="change_step(3)"
        />
      </el-steps>
      <el-card v-if="step === 0" shadow="hover" class="approval-box">
        <el-form
          ref="approval_basic"
          :model="approval_basic"
          :rules="approval_basic_rule"
          label-width="120px"
        >
          <el-form-item label="审批名称" prop="approval_name">
            <el-input
              v-if="edit_btn_show"
              v-model="approval_basic.approval_name"
            />
            <span v-else>{{ approval_basic.approval_name }}</span>
          </el-form-item>
          <el-form-item label="Jenkins Job" prop="job_name">
            <el-input v-if="edit_btn_show" v-model="approval_basic.job_name" />
            <span v-else>{{ approval_basic.job_name }}</span>
          </el-form-item>
          <el-form-item label="描述信息">
            <el-input
              v-if="edit_btn_show"
              v-model="approval_basic.descriptions"
            />
            <span v-else>{{ approval_basic.descriptions }}</span>
          </el-form-item>
          <el-form-item v-if="edit_btn_show">
            <el-button type="success" @click="basic_confirm">
              下一步
            </el-button>
            <el-button @click="basic_cancel"> 取消 </el-button>
          </el-form-item>
          <el-form-item>
            <el-button v-if="approval_id === 0" type="text" @click="go_old">
              返回老版本
            </el-button>
          </el-form-item>
        </el-form>
      </el-card>
      <el-card v-if="step === 1" shadow="hover" class="approval-box">
        <el-row>
          <el-col :span="6" style="border-right: 2px solid #ccc">
            <h3>控件类型</h3>
            <draggable
              class="dragArea list-group"
              :list="controls"
              :group="{ name: 'people', pull: 'clone', put: false }"
              :clone="cloneItem"
              item-key="id"
              @change="log"
            >
              <div
                v-for="element in controls"
                :key="element.id"
                class="list-group-item"
              >
                <div style="cursor: move; height: 36px">
                  <div v-if="element.type === '单行文本'" class="input-control">
                    {{ element.type }}
                    <i
                      style="position: absolute; right: 5px"
                      class="el-icon-suitcase"
                    />
                  </div>
                  <div v-if="element.type === '多行文本'" class="input-control">
                    {{ element.type }}
                    <i
                      style="position: absolute; right: 5px"
                      class="el-icon-suitcase-1"
                    />
                  </div>
                  <div v-if="element.type === '数字'" class="input-control">
                    {{ element.type }}
                    <i style="position: absolute; right: 5px">123</i>
                  </div>
                  <div v-if="element.type === '日期'" class="input-control">
                    {{ element.type }}
                    <i
                      style="position: absolute; right: 5px"
                      class="el-icon-date"
                    />
                  </div>
                  <div v-if="element.type === '单选项'" class="input-control">
                    {{ element.type }}
                    <i
                      style="position: absolute; right: 5px"
                      class="el-icon-circle-check"
                    />
                  </div>
                </div>
              </div>
            </draggable>
          </el-col>
          <el-col
            :span="12"
            style="
              padding: 0 15px;
              border-right: 2px solid #ccc;
              min-height: 200px;
            "
          >
            <h3>表单设计</h3>
            <draggable
              class="dragArea list-group"
              :list="forms"
              group="people"
              item-key="id"
              @change="log"
            >
              <div
                v-for="(element, index) in forms"
                :key="element.id"
                label="控件名称"
                class="list-group-item"
                style="margin-bottom: 10px"
              >
                <div
                  v-if="element.type === '单行文本'"
                  :class="
                    form_item_active == index
                      ? 'input-control active'
                      : 'input-control'
                  "
                  @click="item_name(element, index)"
                >
                  {{ element.type }}[{{ element.name ? element.name : "?" }}]
                  <i
                    style="position: absolute; right: 5px"
                    class="el-icon-suitcase"
                  />
                  <el-button
                    v-if="del_btn_show === 1"
                    type="text"
                    style="position: absolute; right: -35px; top: -5px"
                    @click="remove_form(index, forms)"
                  >
                    删除
                  </el-button>
                </div>
                <div
                  v-if="element.type === '多行文本'"
                  :class="
                    form_item_active == index
                      ? 'input-control active'
                      : 'input-control'
                  "
                  @click="item_name(element, index)"
                >
                  {{ element.type }}[{{ element.name ? element.name : "?" }}]
                  <i
                    style="position: absolute; right: 5px"
                    class="el-icon-suitcase-1"
                  />
                  <el-button
                    v-if="del_btn_show === 1"
                    type="text"
                    style="position: absolute; right: -35px; top: -5px"
                    @click="remove_form(index, forms)"
                  >
                    删除
                  </el-button>
                </div>
                <div
                  v-if="element.type === '数字'"
                  :class="
                    form_item_active == index
                      ? 'input-control active'
                      : 'input-control'
                  "
                  @click="item_name(element, index)"
                >
                  {{ element.type }}[{{ element.name ? element.name : "?" }}]
                  <i style="position: absolute; right: 5px">123</i>
                  <el-button
                    v-if="del_btn_show === 1"
                    type="text"
                    style="position: absolute; right: -35px; top: -5px"
                    @click="remove_form(index, forms)"
                  >
                    删除
                  </el-button>
                </div>
                <div
                  v-if="element.type === '日期'"
                  :class="
                    form_item_active == index
                      ? 'input-control active'
                      : 'input-control'
                  "
                  @click="item_name(element, index)"
                >
                  {{ element.type }}[{{ element.name ? element.name : "?" }}]
                  <i
                    style="position: absolute; right: 5px"
                    class="el-icon-date"
                  />
                  <el-button
                    v-if="del_btn_show === 1"
                    type="text"
                    style="position: absolute; right: -35px; top: -5px"
                    @click="remove_form(index, forms)"
                  >
                    删除
                  </el-button>
                </div>
                <div
                  v-if="element.type === '单选项'"
                  :class="
                    form_item_active == index
                      ? 'input-control active'
                      : 'input-control'
                  "
                  @click="item_name(element, index)"
                >
                  {{ element.type }}[{{ element.name ? element.name : "?" }}]
                  <i
                    style="position: absolute; right: 5px"
                    class="el-icon-circle-check"
                  />
                  <el-button
                    v-if="del_btn_show === 1"
                    type="text"
                    style="position: absolute; right: -35px; top: -5px"
                    @click="remove_form(index, forms)"
                  >
                    删除
                  </el-button>
                </div>
              </div>
            </draggable>
          </el-col>
          <el-col v-if="forms.length > 0" :span="6" style="padding-left: 15px">
            <h3>{{ forms[form_item_active]["type"] }}</h3>
            <el-form v-if="forms[form_item_active]['type']" label-width="100px">
              <el-form-item label="控件中文名称">
                <el-input v-model="forms[form_item_active]['name']" />
              </el-form-item>
              <el-form-item label="控件英文名称">
                <el-input v-model="forms[form_item_active]['en_name']" />
              </el-form-item>
              <el-form-item style="color: rgb(183, 194, 172); font-size: 12px">
                英文名称与pipeline参数一致
              </el-form-item>
            </el-form>
            <div v-if="forms[form_item_active]['type'] === '单选项'">
              <h3 style="position: relative">
                <span>已有选项</span>
                <el-button
                  v-if="
                    edit_btn_show &&
                    (forms[form_item_active]['en_name'] === 'SERVICE_TYPE' ||
                      forms[form_item_active]['en_name'] === 'PROJECT' ||
                      forms[form_item_active]['en_name'] === 'SERVICE')
                  "
                  style="float: right; padding-bottom: 5px"
                  type="text"
                  @click="option_update()"
                  >同步更新</el-button
                >
              </h3>
              <div
                v-if="forms[form_item_active]['options'].length > 0"
                style="height: 100px; overflow-x: hidden; overflow-y: scroll"
              >
                <div
                  v-for="(item, index) in forms[form_item_active]['options']"
                  :key="item.id"
                  style="
                    line-height: 24px;
                    padding: 5px;
                    position: relative;
                    border-bottom: 1px solid #ccc;
                  "
                >
                  {{ item }}
                  <el-button
                    v-if="del_btn_show === 1"
                    type="text"
                    style="position: absolute; right: 35px; top: 5px"
                    @click="
                      deleteRow(index, forms[form_item_active]['options'])
                    "
                  >
                    删除
                  </el-button>
                </div>
              </div>
              <el-row
                style="margin-top: 20px"
                v-if="
                  forms[form_item_active]['en_name'] !== 'SERVICE_TYPE' &&
                  forms[form_item_active]['en_name'] !== 'PROJECT' &&
                  forms[form_item_active]['en_name'] !== 'SERVICE' &&
                  add_btn_show === 1
                "
              >
                <el-col :span="16">
                  <el-input v-model="option_value" />
                </el-col>
                <el-col :span="6" :offset="1">
                  <el-button type="success" @click="add_option">
                    添加选项
                  </el-button>
                </el-col>
              </el-row>
            </div>
          </el-col>
        </el-row>
        <el-row v-if="edit_btn_show">
          <el-col :span="24" style="text-align: center">
            <el-button type="success" @click="form_confirm"> 确认 </el-button>
            <el-button @click="form_cancel"> 取消 </el-button>
          </el-col>
        </el-row>
      </el-card>
      <el-card v-if="step === 2" shadow="hover" class="approval-box">
        <el-table :data="nodes" style="width: 100%">
          <el-table-column label="节点名称" prop="name" />
          <el-table-column label="是否发起人自选" prop="need_approver">
            <template slot-scope="scope">
              {{ scope.row.need_approver ? "是" : "否" }}
            </template>
          </el-table-column>
          <el-table-column label="审批方式" prop="node_type" />
          <el-table-column label="操作方式" prop="approval_type" />
          <el-table-column label="审批人">
            <template slot-scope="scope">
              {{
                scope.row.username === "未知"
                  ? "审批人自选"
                  : scope.row.username
              }}
            </template>
          </el-table-column>
          <el-table-column v-if="edit_btn_show === 1" label="操作" width="200">
            <template slot-scope="scope">
              <el-button
                v-if="
                  scope.row.name !== '开始' &&
                  scope.row.name !== '结束' &&
                  del_btn_show === 1
                "
                type="danger"
                icon="el-icon-remove-outline"
                size="small"
                @click.native.prevent="deleteRow(scope.$index, nodes)"
              />
              <el-button
                v-if="scope.row.name !== '结束' && add_btn_show === 1"
                type="primary"
                icon="el-icon-circle-plus-outline"
                size="small"
                @click="node_add_btn(scope.$index)"
              />
              <el-button
                v-if="
                  scope.row.name !== '开始' &&
                  scope.row.name !== '结束' &&
                  edit_btn_show === 1
                "
                type="primary"
                icon="el-icon-edit"
                size="small"
                @click="node_edit_btn(scope.$index, scope.row)"
              />
            </template>
          </el-table-column>
        </el-table>
        <el-row style="margin-top: 10px">
          <el-col :span="24" style="text-align: center">
            <el-button type="success" @click="node_confirm"> 确认 </el-button>
          </el-col>
        </el-row>
        <el-dialog
          :title="dialogTitle"
          :visible.sync="dialogVisible"
          width="35%"
          :before-close="node_dialog_cancel"
        >
          <el-form ref="node" :model="node" label-width="120px">
            <el-form-item label="节点名称">
              <el-input v-model="node.name" placeholder="请输入节点名称" />
            </el-form-item>
            <el-form-item label="是否发起人自选">
              <el-select
                v-model="node.need_approver"
                placeholder="请选择"
                style="width: 100%"
              >
                <el-option label="是" :value="true" />
                <el-option label="否" :value="false" />
              </el-select>
            </el-form-item>
            <el-form-item label="审批方式">
              <el-select
                v-model="node.node_type"
                placeholder="选择审批方式"
                style="width: 100%"
              >
                <el-option label="会签" value="会签" />
                <el-option label="或签" value="或签" />
                <el-option label="抄送" value="抄送" />
              </el-select>
            </el-form-item>
            <el-form-item label="审批人" v-if="!node.need_approver">
              <el-select
                v-model="node.user_id"
                filterable
                clearable
                placeholder="选择审批人"
                style="width: 100%"
              >
                <el-option
                  v-for="user in user_list"
                  :key="user.id"
                  :label="user.username"
                  :value="user.id"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="操作方式">
              <el-select
                v-model="node.approval_type"
                filterable
                clearable
                placeholder="选择操作方式"
                style="width: 100%"
              >
                <el-option
                  v-for="item in approval_type_choices"
                  :key="item"
                  :label="item"
                  :value="item"
                />
              </el-select>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="node_dialog_confirm">
                {{ dialogTitle.indexOf("添加") !== -1 ? "添加" : "更新" }}
              </el-button>
              <el-button @click="node_dialog_cancel"> 取消 </el-button>
            </el-form-item>
          </el-form>
        </el-dialog>
      </el-card>
      <el-card
        v-if="step === 3 && add_btn_show === 1"
        shadow="hover"
        class="approval-box"
      >
        <div style="text-align: center">
          <el-button
            v-if="approval_basic.subscribe === 0"
            type="success"
            @click="subscribe_btn"
          >
            激活
          </el-button>
          <el-button
            v-if="approval_basic.subscribe === 1"
            type="primary"
            @click="unsubscribe_btn"
          >
            禁用
          </el-button>
        </div>
        <div
          v-if="!percentageVisible"
          style="text-align: center; margin-top: 15px"
        >
          {{ approval_basic.subscribe ? "已经激活" : "尚未激活" }}
        </div>
        <div
          v-if="percentageVisible"
          style="text-align: center; margin-top: 15px"
        >
          <el-progress
            type="dashboard"
            :percentage="percentage"
            :color="colors"
          />
        </div>
      </el-card>
    </div>
    <div v-if="is_mobile" class="mobile">
      <el-steps :active="step" finish-status="success" align-center>
        <el-step title="基础信息" @click.native="change_step(0)" />
        <el-step title="表单设计" @click.native="change_step(1)" />
        <el-step title="流程设计" @click.native="change_step(2)" />
        <el-step
          v-if="add_btn_show === 1"
          title="激活使用"
          @click.native="change_step(3)"
        />
      </el-steps>
      <el-card v-if="step === 0" shadow="hover" class="approval-box">
        <el-form
          ref="approval_basic"
          :model="approval_basic"
          :rules="approval_basic_rule"
          label-width="100px"
        >
          <el-form-item label="审批名称" prop="approval_name">
            <el-input
              v-if="edit_btn_show === 1"
              v-model="approval_basic.approval_name"
            />
            <span v-else>{{ approval_basic.approval_name }}</span>
          </el-form-item>
          <el-form-item label="Job Name" prop="job_name">
            <el-input
              v-if="edit_btn_show === 1"
              v-model="approval_basic.job_name"
            />
            <span v-else>{{ approval_basic.approval_name }}</span>
          </el-form-item>
          <el-form-item label="描述信息">
            <el-input
              v-if="edit_btn_show === 1"
              v-model="approval_basic.descriptions"
            />
            <span v-else>{{ approval_basic.descriptions }}</span>
          </el-form-item>
          <el-form-item v-if="edit_btn_show === 1">
            <el-button type="success" @click="basic_confirm">
              下一步
            </el-button>
            <el-button @click="basic_cancel"> 取消 </el-button>
          </el-form-item>
          <el-form-item>
            <el-button v-if="approval_id === 0" type="text" @click="go_old">
              返回老版本
            </el-button>
          </el-form-item>
        </el-form>
      </el-card>
      <el-card v-if="step === 1" shadow="hover" class="approval-box">
        <el-row>
          <el-col v-if="edit_btn_show === 1" :span="24">
            <h3>控件类型</h3>
            <draggable
              class="dragArea list-group"
              :list="controls"
              :group="{ name: 'people', pull: 'clone', put: false }"
              :clone="cloneItem"
              item-key="id"
              @change="log"
            >
              <div
                v-for="element in controls"
                :key="element.id"
                class="list-group-item"
              >
                <div style="cursor: move; height: 36px">
                  <div v-if="element.type === '单行文本'" class="input-control">
                    {{ element.type }}
                    <i
                      style="position: absolute; right: 5px"
                      class="el-icon-suitcase"
                    />
                  </div>
                  <div v-if="element.type === '多行文本'" class="input-control">
                    {{ element.type }}
                    <i
                      style="position: absolute; right: 5px"
                      class="el-icon-suitcase-1"
                    />
                  </div>
                  <div v-if="element.type === '数字'" class="input-control">
                    {{ element.type }}
                    <i style="position: absolute; right: 5px">123</i>
                  </div>
                  <div v-if="element.type === '日期'" class="input-control">
                    {{ element.type }}
                    <i
                      style="position: absolute; right: 5px"
                      class="el-icon-date"
                    />
                  </div>
                  <div v-if="element.type === '单选项'" class="input-control">
                    {{ element.type }}
                    <i
                      style="position: absolute; right: 5px"
                      class="el-icon-circle-check"
                    />
                  </div>
                </div>
              </div>
            </draggable>
          </el-col>
          <el-col :span="24" style="min-height: 150px">
            <h3>表单设计</h3>
            <div v-if="forms.length === 0" class="wait">
              表单托放出 <i class="el-icon-bottom" />
            </div>
            <draggable
              class="dragArea list-group"
              :list="forms"
              group="people"
              item-key="id"
              @change="log"
            >
              <div
                v-for="(element, index) in forms"
                :key="element.id"
                label="控件名称"
                class="list-group-item"
                style="margin-bottom: 10px"
              >
                <div
                  v-if="element.type === '单行文本'"
                  :class="
                    form_item_active == index
                      ? 'input-control active'
                      : 'input-control'
                  "
                  @click="item_name(element, index)"
                >
                  {{ element.type }}[{{ element.name ? element.name : "?" }}]
                  <i
                    style="position: absolute; right: 5px"
                    class="el-icon-suitcase"
                  />
                  <el-button
                    v-if="del_btn_show === 1"
                    type="text"
                    style="position: absolute; right: -35px; top: -5px"
                    @click="remove_form(index, forms)"
                  >
                    删除
                  </el-button>
                </div>
                <div
                  v-if="element.type === '多行文本'"
                  :class="
                    form_item_active == index
                      ? 'input-control active'
                      : 'input-control'
                  "
                  @click="item_name(element, index)"
                >
                  {{ element.type }}[{{ element.name ? element.name : "?" }}]
                  <i
                    style="position: absolute; right: 5px"
                    class="el-icon-suitcase-1"
                  />
                  <el-button
                    v-if="del_btn_show === 1"
                    type="text"
                    style="position: absolute; right: -35px; top: -5px"
                    @click="remove_form(index, forms)"
                  >
                    删除
                  </el-button>
                </div>
                <div
                  v-if="element.type === '数字'"
                  :class="
                    form_item_active == index
                      ? 'input-control active'
                      : 'input-control'
                  "
                  @click="item_name(element, index)"
                >
                  {{ element.type }}[{{ element.name ? element.name : "?" }}]
                  <i style="position: absolute; right: 5px">123</i>
                  <el-button
                    v-if="del_btn_show === 1"
                    type="text"
                    style="position: absolute; right: -35px; top: -5px"
                    @click="remove_form(index, forms)"
                  >
                    删除
                  </el-button>
                </div>
                <div
                  v-if="element.type === '日期'"
                  :class="
                    form_item_active == index
                      ? 'input-control active'
                      : 'input-control'
                  "
                  @click="item_name(element, index)"
                >
                  {{ element.type }}[{{ element.name ? element.name : "?" }}]
                  <i
                    style="position: absolute; right: 5px"
                    class="el-icon-date"
                  />
                  <el-button
                    v-if="del_btn_show === 1"
                    type="text"
                    style="position: absolute; right: -35px; top: -5px"
                    @click="remove_form(index, forms)"
                  >
                    删除
                  </el-button>
                </div>
                <div
                  v-if="element.type === '单选项'"
                  :class="
                    form_item_active == index
                      ? 'input-control active'
                      : 'input-control'
                  "
                  @click="item_name(element, index)"
                >
                  {{ element.type }}[{{ element.name ? element.name : "?" }}]
                  <i
                    style="position: absolute; right: 5px"
                    class="el-icon-circle-check"
                  />
                  <el-button
                    v-if="del_btn_show === 1"
                    type="text"
                    style="position: absolute; right: -35px; top: -5px"
                    @click="remove_form(index, forms)"
                  >
                    删除
                  </el-button>
                </div>
              </div>
            </draggable>
          </el-col>
          <el-col v-if="forms.length > 0" :span="24">
            <h3>{{ forms[form_item_active]["type"] }}</h3>
            <el-form v-if="forms[form_item_active]['type']" label-width="80px">
              <el-form-item label="中文名称">
                <el-input v-model="forms[form_item_active]['name']" />
              </el-form-item>
              <el-form-item label="英文名称">
                <el-input v-model="forms[form_item_active]['en_name']" />
              </el-form-item>
              <el-form-item style="color: rgb(183, 194, 172); font-size: 12px">
                英文名称与pipeline参数一致
              </el-form-item>
            </el-form>
            <div v-if="forms[form_item_active]['type'] === '单选项'">
              <h3 style="position: relative">
                <span>已有选项</span>
                <el-button
                  v-if="
                    edit_btn_show === 1 &&
                    (forms[form_item_active]['en_name'] === 'SERVICE_TYPE' ||
                      forms[form_item_active]['en_name'] === 'PROJECT' ||
                      forms[form_item_active]['en_name'] === 'SERVICE')
                  "
                  style="float: right; padding-bottom: 5px"
                  type="text"
                  @click="option_update()"
                  >同步更新</el-button
                >
              </h3>
              <div
                v-if="forms[form_item_active]['options'].length > 0"
                style="height: 100px; overflow-x: hidden; overflow-y: scroll"
              >
                <div
                  v-for="(item, index) in forms[form_item_active]['options']"
                  :key="item.id"
                  style="
                    line-height: 24px;
                    padding: 5px;
                    position: relative;
                    border-bottom: 1px solid #ccc;
                  "
                >
                  {{ item }}
                  <el-button
                    v-if="del_btn_show === 1"
                    type="text"
                    style="position: absolute; right: 35px; top: 5px"
                    @click="
                      deleteRow(index, forms[form_item_active]['options'])
                    "
                  >
                    删除
                  </el-button>
                </div>
              </div>
              <el-row
                style="margin: 20px 0"
                v-if="
                  forms[form_item_active]['en_name'] !== 'SERVICE_TYPE' &&
                  forms[form_item_active]['en_name'] !== 'PROJECT' &&
                  forms[form_item_active]['en_name'] !== 'SERVICE' &&
                  add_btn_show === 1
                "
              >
                <el-col :span="16">
                  <el-input v-model="option_value" />
                </el-col>
                <el-col :span="6" :offset="1">
                  <el-button type="success" @click="add_option">
                    添加
                  </el-button>
                </el-col>
              </el-row>
            </div>
          </el-col>
        </el-row>
        <el-row>
          <el-col
            v-if="edit_btn_show === 1"
            :span="24"
            style="text-align: center"
          >
            <el-button type="success" @click="form_confirm"> 确认 </el-button>
            <el-button @click="form_cancel"> 取消 </el-button>
          </el-col>
        </el-row>
      </el-card>
      <el-card v-if="step === 2" shadow="hover" class="approval-box">
        <el-table :data="nodes" style="width: 100%">
          <el-table-column label="节点名称" prop="name" fixed width="90" />
          <el-table-column label="是否发起人自选" prop="need_approver">
            <template slot-scope="scope">
              {{ scope.row.need_approver ? "是" : "否" }}
            </template>
          </el-table-column>
          <el-table-column label="审批方式" prop="node_type" />
          <el-table-column label="操作方式" prop="approval_type" />
          <el-table-column label="审批人">
            <template slot-scope="scope">
              {{
                scope.row.username == "未知" ? "审批人自选" : scope.row.username
              }}
            </template>
          </el-table-column>
          <el-table-column v-if="edit_btn_show === 1" label="操作" width="180">
            <template slot-scope="scope">
              <el-button
                v-if="
                  scope.row.name !== '开始' &&
                  scope.row.name !== '结束' &&
                  del_btn_show === 1
                "
                type="danger"
                icon="el-icon-remove-outline"
                size="mini"
                @click.native.prevent="deleteRow(scope.$index, nodes)"
              />
              <el-button
                v-if="scope.row.name !== '结束' && add_btn_show === 1"
                type="primary"
                icon="el-icon-circle-plus-outline"
                size="mini"
                @click="node_add_btn(scope.$index)"
              />
              <el-button
                v-if="
                  scope.row.name !== '开始' &&
                  scope.row.name !== '结束' &&
                  edit_btn_show === 1
                "
                type="primary"
                icon="el-icon-edit"
                size="mini"
                @click="node_edit_btn(scope.$index, scope.row)"
              />
            </template>
          </el-table-column>
        </el-table>
        <el-row style="margin-top: 10px" v-if="edit_btn_show === 1">
          <el-col :span="24" style="text-align: center">
            <el-button type="success" @click="node_confirm"> 确认 </el-button>
          </el-col>
        </el-row>
        <el-dialog
          :title="dialogTitle"
          :visible.sync="dialogVisible"
          width="100%"
          :before-close="node_dialog_cancel"
        >
          <el-form ref="node" :model="node" label-width="100px">
            <el-form-item label="节点名称">
              <el-input v-model="node.name" placeholder="请输入节点名称" />
            </el-form-item>
            <el-form-item label="发起人自选">
              <el-select
                v-model="node.need_approver"
                placeholder="请选择"
                style="width: 100%"
              >
                <el-option label="是" :value="true" />
                <el-option label="否" :value="false" />
              </el-select>
            </el-form-item>
            <el-form-item label="审批方式">
              <el-select
                v-model="node.node_type"
                placeholder="选择审批方式"
                style="width: 100%"
              >
                <el-option label="会签" value="会签" />
                <el-option label="或签" value="或签" />
                <el-option label="抄送" value="抄送" />
              </el-select>
            </el-form-item>
            <el-form-item label="审批人" v-if="!node.need_approver">
              <el-select
                v-model="node.user_id"
                filterable
                clearable
                placeholder="选择审批人"
                style="width: 100%"
              >
                <el-option
                  v-for="user in user_list"
                  :key="user.id"
                  :label="user.username"
                  :value="user.id"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="操作方式">
              <el-select
                v-model="node.approval_type"
                filterable
                clearable
                placeholder="选择操作方式"
                style="width: 100%"
              >
                <el-option
                  v-for="item in approval_type_choices"
                  :key="item"
                  :label="item"
                  :value="item"
                />
              </el-select>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="node_dialog_confirm">
                {{ dialogTitle.indexOf("添加") !== -1 ? "添加" : "更新" }}
              </el-button>
              <el-button @click="node_dialog_cancel"> 取消 </el-button>
            </el-form-item>
          </el-form>
        </el-dialog>
      </el-card>
      <el-card
        v-if="step === 3 && add_btn_show === 1"
        shadow="hover"
        class="approval-box"
      >
        <div style="text-align: center">
          <el-button
            v-if="approval_basic.subscribe === 0"
            type="success"
            @click="subscribe_btn"
          >
            激活
          </el-button>
          <el-button
            v-if="approval_basic.subscribe === 1"
            type="primary"
            @click="unsubscribe_btn"
          >
            禁用
          </el-button>
        </div>
        <div
          v-if="!percentageVisible"
          style="text-align: center; margin-top: 15px"
        >
          {{ approval_basic.subscribe ? "已经激活" : "尚未激活" }}
        </div>
        <div
          v-if="percentageVisible"
          style="text-align: center; margin-top: 15px"
        >
          <el-progress
            type="dashboard"
            :percentage="percentage"
            :color="colors"
          />
        </div>
      </el-card>
    </div>
  </div>
</template>
<script>
import {
  approval_get,
  approval_cud,
  form_get,
  form_cud,
  subscribe_cud,
  node_get,
  node_cud,
} from "@/api/approval/approval";
import { usernames_get } from "@/api/user";
import { choices_get } from "@/api/approval/choices";
import { project_get, service_get } from "@/api/project-list/index";
import draggable from "vuedraggable";
import config from "@/utils/config";
import { btn_check } from "@/api/btn";
import { ApprovalManagePage } from "@/utils/auth";
let idGlobal = 1;
export default {
  name: "ApprovalManage",
  components: { draggable },
  data() {
    return {
      admin_id: "",
      approval_id: 0,
      dialogVisible: false,
      dialogTitle: "",
      option_value: "",
      node_index: 0,
      step: 0,
      form_item_active: 0,
      approval_basic: {
        approval_name: "",
        job_name: "",
        descriptions: "",
      },
      approval_basic_rule: {
        approval_name: [
          { required: true, message: "请输入审批流程名", trigger: "blur" },
        ],
        job_name: [
          {
            required: true,
            message: "请输入jenkins job 名称",
            trigger: "blur",
          },
        ],
      },
      controls: [
        { id: 1, type: "单行文本", name: "" },
        { id: 2, type: "多行文本", name: "" },
        { id: 3, type: "数字", name: "" },
        { id: 4, type: "日期", name: "" },
        { id: 5, type: "单选项", name: "", options: [] },
      ],
      user_list: [],
      control: {},
      forms: [],
      nodes: [],
      node: {
        name: "",
        need_approver: false,
        node_type: "",
        user_id: "",
      },
      percentage: 0,
      percentageVisible: false,
      colors: [
        { color: "#f56c6c", percentage: 20 },
        { color: "#e6a23c", percentage: 40 },
        { color: "#5cb87a", percentage: 60 },
        { color: "#1989fa", percentage: 80 },
        { color: "#6f7ad3", percentage: 100 },
      ],
      approval_type_choices: [],
      is_mobile: config.isMobile,
      service_types: [],
      project_list: [],
      service_list: [],
      is_update: 0,
      ApprovalManagePage: ApprovalManagePage,
      add_btn_show: 0,
      del_btn_show: 0,
      edit_btn_show: 0,
      conn_btn_show: 0,
    };
  },
  mounted() {
    this.init_data();
  },
  methods: {
    update_info(is_update) {
      if (is_update === 1) {
        this.$notify.info({
          title: `表单更新`,
          dangerouslyUseHTMLString: true,
          message: "有待更新,别忘了确认更新,",
          duration: 0,
          position: "top-right",
        });
      }
    },
    option_update() {
      let t = this.forms[this.form_item_active]["en_name"];
      console.log(this.forms[this.form_item_active]);
      console.log(t);
      if (t === "SERVICE_TYPE") {
        this.is_update += 1;
        this.forms[this.form_item_active]["options"] = this.service_types;
      } else if (t === "PROJECT") {
        this.is_update += 1;
        let options = [];
        for (var i = 0; i < this.project_list.length; i++) {
          var item = this.project_list[i];
          options.push(item.name);
        }
        console.log(options);
        this.forms[this.form_item_active]["options"] = options;
      } else if (t === "SERVICE") {
        this.is_update += 1;
        let options = [];
        for (var i = 0; i < this.service_list.length; i++) {
          var item = this.service_list[i];
          options.push(item.name);
        }
        this.forms[this.form_item_active]["options"] = options;
      }
      this.update_info(this.is_update);
    },
    go_old() {
      this.$router.push({
        path: "/approvals/manage/manage-old",
      });
    },
    init_data() {
      btn_check(this.ApprovalManagePage).then((response) => {
        if (response.code == 200) {
          this.add_btn_show = response.data.add;
          this.edit_btn_show = response.data.edit;
          this.del_btn_show = response.data.del;
          this.conn_btn_show = response.data.conn;
        }
      });
      const id = this.$route.query.id;
      const step = this.$route.query.step;
      if (id !== undefined) {
        if (step !== undefined) {
          this.step = step;
        }
        this.approval_id = id;
        approval_get({ id: this.approval_id }).then((response) => {
          if (response.code === 200) {
            this.approval_basic = response.data[0];
          }
        });
      }
    },
    remove_form(index, rows) {
      this.is_update += 1;
      this.form_item_active = 0;
      this.update_info(this.is_update);
      rows.splice(index, 1);
    },
    deleteRow(index, rows) {
      rows.splice(index, 1);
    },
    add_option() {
      if (this.option_value.trim() === "") {
        this.$message({
          message: "选项不能为空",
          type: "warning",
        });
      } else {
        this.forms[this.form_item_active]["options"].push(this.option_value);
        this.option_value = "";
      }
    },
    async change_step(num) {
      if (this.approval_basic.approval_code === undefined && num > 0) {
        this.$message({
          message: "必须完成之前流程",
          type: "warning",
        });
      } else {
        this.step = num;
        if (this.step === 1) {
          let result = await form_get({ approval_id: this.approval_id });
          if (result.code === 200) {
            this.forms = result.data;
          }
          choices_get({
            table_name: "BusinessServices",
            column: "type_choices",
          }).then((response) => {
            if (response.code === 200) {
              this.service_types = response.data;
            }
          });
          project_get().then((response) => {
            if (response.code === 200) {
              this.project_list = response.data.project_infos;
            }
          });
          service_get().then((response) => {
            if (response.code === 200) {
              this.service_list = response.data.service_infos;
            }
          });
        } else if (this.step === 2) {
          const u = await usernames_get();
          this.user_list = u.data.users;
          for (var i = 0; i < this.user_list.length; i++) {
            const user = this.user_list[i];
            if (user.username === "admin") {
              this.admin_id = user.id;
            }
          }
          choices_get({
            table_name: "FANode",
            column: "approval_type_choices",
          }).then((response) => {
            if (response.code === 200) {
              this.approval_type_choices = response.data;
            }
          });
          node_get({ approval_id: this.approval_id }).then((response) => {
            if (response.code === 200) {
              if (response.data.length === 0) {
                this.nodes = [
                  {
                    name: "开始",
                    need_approver: 0,
                    node_type: "会签",
                    approval_type: "自动通过",
                    username: "admin",
                    user_id: this.admin_id,
                  },
                  {
                    name: "结束",
                    need_approver: 0,
                    node_type: "会签",
                    approval_type: "自动通过",
                    username: "admin",
                    user_id: this.admin_id,
                  },
                ];
              } else {
                this.nodes = response.data;
              }
            }
          });
        } else if (this.step === 3) {
          console.log("激活使用步骤");
        }
      }
    },
    basic_cancel() {
      this.approval_basic = {
        approval_name: "",
        job_name: "",
        descriptions: "",
      };
    },
    basic_confirm() {
      if (this.approval_basic["approval_code"] != undefined) {
        approval_cud("PUT", this.approval_basic).then((response) => {
          if (response.code === 200) {
            this.$message({
              message: response.msg,
              type: "success",
            });
            this.step = 1;
          }
        });
      } else {
        approval_cud("POST", this.approval_basic).then((response) => {
          if (response.code === 200) {
            this.$message({
              message: response.msg,
              type: "success",
            });
            approval_get({
              approval_name: this.approval_basic.approval_name,
            }).then((response) => {
              if (response.code === 200) {
                this.approval_basic = response.data[0];
                this.approval_id = this.approval_basic.id;
                this.step = 1;
              }
            });
          }
        });
      }
    },
    form_cancel() {
      form_get({ approval_id: this.approval_id }).then((response) => {
        if (response.code === 200) {
          this.forms = response.data;
        }
      });
    },
    form_confirm() {
      const result = this.forms;
      for (let i = 0; i < result.length; i++) {
        result[i]["serial_number"] = i + 1;
      }
      const data = {
        approval_id: this.approval_id,
        form: result,
      };
      form_cud("post", data).then((response) => {
        if (response.code === 200) {
          this.$message({
            message: response.msg,
            type: "success",
          });
          this.is_update = 0;
          this.$notify.closeAll();
        }
      });
    },
    log: function (evt) {
      this.is_update += 1;
      this.update_info(this.is_update);
      window.console.log(evt);
    },
    cloneItem({ id }) {
      const item = {
        id: idGlobal++,
        name: this.controls[id - 1]["name"],
        type: this.controls[id - 1]["type"],
      };
      if (item["type"] === "单选项") {
        item["options"] = [];
      }
      return item;
    },
    item_name(element, index) {
      this.control = element;
      this.form_item_active = this.forms.indexOf(element);
    },
    node_add_btn(index) {
      this.node_index = index;
      this.dialogTitle = "节点添加";
      this.dialogVisible = true;
    },
    node_edit_btn(index, row) {
      this.node_index = index;
      this.node = row;
      this.dialogTitle = "节点修改";
      this.dialogVisible = true;
    },
    node_dialog_confirm() {
      if (this.dialogTitle === "节点添加") {
        for (var i = 0; i < this.user_list.length; i++) {
          var user = this.user_list[i];
          if (user.id === this.node.user_id) {
            this.node["username"] = user.username;
          }
        }
        this.nodes.splice(this.node_index + 1, 0, this.node);
      } else {
        if (this.node["need_approver"] === true) {
          this.node["user_id"] = "";
          this.node["username"] = "审批人自选";
        } else {
          for (var i = 0; i < this.user_list.length; i++) {
            var user = this.user_list[i];
            if (this.node["user_id"] === user.id) {
              this.node["username"] = user.username;
            }
          }
        }
        this.nodes[this.node_index] = this.node;
      }
      this.dialogVisible = false;
      this.node = {
        name: "",
        need_approver: false,
        node_type: "",
      };
    },
    node_dialog_cancel() {
      this.node = {
        name: "",
        need_approver: false,
        node_type: "",
      };
      this.dialogVisible = false;
    },
    node_confirm() {
      const result = this.nodes;
      for (var i = 0; i < result.length; i++) {
        result[i]["serial_number"] = i + 1;
      }
      const data = {
        approval_id: this.approval_id,
        nodes: result,
      };
      node_cud("post", data).then((response) => {
        if (response.code === 200) {
          this.$message({
            message: response.msg,
            type: "success",
          });
        }
      });
    },
    subscribe_btn() {
      this.percentageVisible = true;
      const data = {
        approval_id: this.approval_id,
        type: "subscribe",
      };
      subscribe_cud("post", data).then((response) => {
        if (response.code === 200) {
          const _this = this;
          const timer = setInterval(() => {
            _this.percentage += 5;
            if (_this.percentage > 100) {
              clearInterval(timer);
              _this.$message({
                message: "激活成功",
                type: "success",
              });
              _this.percentageVisible = false;
              _this.approval_basic.subscribe = 1;
            }
          }, 200);
        }
      });
    },
    unsubscribe_btn() {
      this.percentage = 100;
      this.percentageVisible = true;
      const data = {
        approval_id: this.approval_id,
        type: "unsubscribe",
      };
      subscribe_cud("post", data).then((response) => {
        if (response.code === 200) {
          const _this = this;
          const timer = setInterval(() => {
            _this.percentage -= 5;
            if (_this.percentage < 0) {
              clearInterval(timer);
              _this.$message({
                message: "禁用成功",
                type: "success",
              });
              _this.percentageVisible = false;
              _this.approval_basic.subscribe = 0;
            }
          }, 200);
        }
      });
    },
  },
};
</script>
<style lang='scss' scoped>
.approval-manage {
  min-height: 450px;
  width: 98%;
  margin: 15px auto;
  background-color: #fff;
  padding: 15px;
  el-step {
    cursor: pointer;
  }
  .approval-box {
    min-height: 250px;
    width: 85%;
    margin: 35px auto;
  }
  .input-control {
    border: 1px solid #ccc;
    padding: 5px 15px;
    width: 85%;
    border-radius: 5px;
    position: relative;
    cursor: pointer;
  }
  .input-control.active {
    background: rgb(204, 219, 204);
  }
  .mobile {
    .wait {
      border: 1px solid #ddd;
      text-align: center;
      padding: 5px;
      color: #999;
      border-radius: 8px;
      margin: 15px 0;
    }
  }
}
</style>
