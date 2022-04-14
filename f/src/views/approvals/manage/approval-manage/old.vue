<template>
  <div class="old-approval">
    <div v-if="!is_mobile">
      <el-steps :active="step" finish-status="success" simple align-center>
        <el-step title="基础配置" @click.native="change_step(0)" />
        <el-step title="表单信息" @click.native="change_step(1)" />
        <el-step title="流程信息" @click.native="change_step(2)" />
      </el-steps>
      <el-card v-if="step === 0" shadow="hover" class="approval-box">
        <el-form
          ref="approval_basic"
          :model="approval_basic"
          :rules="approval_basic_rule"
          label-width="120px"
        >
          <el-form-item label="审批名称" prop="approval_name">
            <el-input v-model="approval_basic.approval_name" />
          </el-form-item>
          <el-form-item label="审批编码" prop="approval_code">
            <el-input v-model="approval_basic.approval_code" />
          </el-form-item>
          <el-form-item label="Jenkins Job" prop="job_name">
            <el-input v-model="approval_basic.job_name" />
          </el-form-item>
          <el-form-item label="是否订阅" prop="subscribe">
            <el-select
              v-model="approval_basic.subscribe"
              style="width: 100%"
              placeholder="是否订阅"
            >
              <el-option label="是" :value="1" />
              <el-option label="否" :value="0" />
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button type="success" @click="basic_confirm"> 确认 </el-button>
            <el-button @click="basic_cancel"> 取消 </el-button>
          </el-form-item>
        </el-form>
      </el-card>
      <el-card v-if="step === 1" shadow="hover" class="approval-box">
        <div style="text-align: center">
          表单信息飞书提供，无需更新或查看
        </div>
      </el-card>
      <el-card v-if="step === 2" shadow="hover" class="approval-box">
        <div style="text-align: center">
          流程信息飞书提供，无需更新或查看
        </div>
      </el-card>
    </div>
    <div v-if="is_mobile">
      <el-steps :active="step" finish-status="success" align-center>
        <el-step title="基础配置" @click.native="change_step(0)" />
        <el-step title="表单信息" @click.native="change_step(1)" />
        <el-step title="流程信息" @click.native="change_step(2)" />
      </el-steps>
      <el-card v-if="step === 0" shadow="hover" class="approval-box">
        <el-form
          ref="approval_basic"
          :model="approval_basic"
          :rules="approval_basic_rule"
          label-width="120px"
        >
          <el-form-item label="审批名称" prop="approval_name">
            <el-input v-model="approval_basic.approval_name" />
          </el-form-item>
          <el-form-item label="审批编码" prop="approval_code">
            <el-input v-model="approval_basic.approval_code" />
          </el-form-item>
          <el-form-item label="Jenkins Job" prop="job_name">
            <el-input v-model="approval_basic.job_name" />
          </el-form-item>
          <el-form-item label="是否订阅" prop="subscribe">
            <el-select
              v-model="approval_basic.subscribe"
              style="width: 100%"
              placeholder="是否订阅"
            >
              <el-option label="是" :value="1" />
              <el-option label="否" :value="0" />
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button type="success" @click="basic_confirm"> 确认 </el-button>
            <el-button @click="basic_cancel"> 取消 </el-button>
          </el-form-item>
        </el-form>
      </el-card>
      <el-card v-if="step === 1" shadow="hover" class="approval-box">
        <div style="text-align: center">
          表单信息飞书提供，无需更新或查看
        </div>
      </el-card>
      <el-card v-if="step === 2" shadow="hover" class="approval-box">
        <div style="text-align: center">
          流程信息飞书提供，无需更新或查看
        </div>
      </el-card>
    </div>
  </div>
</template>
<script>
import config from "@/utils/config";
import { subscribe_v1_cud } from "@/api/approval/v1";
import { approval_get } from "@/api/approval/approval";
export default {
  name: "OldApproval",
  data() {
    return {
      is_mobile: config.isMobile,
      step: 0,
      approval_basic: {
        approval_name: "",
        approval_code: "",
        job_name: "",
        subscribe: 0,
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
        approval_code: [
          { required: true, message: "请输入审批编码", trigger: "blur" },
        ],
        subscribe: [{ required: true, message: "请选择", trigger: "change" }],
      },
    };
  },
  mounted() {
    this.init_data();
  },
  methods: {
    init_data() {
      const id = this.$route.query.id;
      if (id != undefined) {
        this.approval_id = id;
        approval_get({ id: this.approval_id }).then((response) => {
          if (response.code === 200) {
            this.approval_basic = response.data[0];
          }
        });
      }
    },
    change_step(num) {
      this.step = num;
    },
    basic_confirm() {
      subscribe_v1_cud("post", this.approval_basic).then((response) => {
        if (response !== undefined && response.code === 200) {
          this.$message({
            message: response.msg,
            type: "success",
          });
        }
      });
    },
    basic_cancel() {
      this.approval_basic = {
        approval_name: "",
        approval_code: "",
        job_name: "",
        subscribe: 0,
      };
    },
  },
};
</script>
<style lang="scss" scoped>
.old-approval {
  min-height: 450px;
  width: 98%;
  margin: 15px auto;
  background-color: #fff;
  padding: 15px;
}
</style> >
