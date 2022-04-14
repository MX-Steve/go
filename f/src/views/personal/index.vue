<template>
  <div class="personal">
    <div v-if="!is_mobile">
      <el-card class="box-card">
        <div slot="header" class="clearfix">
          <span>个人主页</span>
        </div>
        <div class="card-box">
          <el-row>
            <el-col :span="12" style="padding-right: 10px">
              <div class="title">个人基础信息</div>
              <el-form ref="user" :model="user" label-width="120px">
                <el-form-item label="登录名: ">
                  {{ user.username }}
                </el-form-item>
                <el-form-item label="中文名: ">
                  {{ user.zh_name }}
                </el-form-item>
                <el-form-item label="邮箱: ">
                  {{ user.email ? user.email : "无信息" }}
                </el-form-item>
                <el-form-item label="角色: ">
                  <span v-for="item in role_names" :key="item">
                    {{ item }}
                  </span>
                </el-form-item>
                <el-form-item label="机器人地址: ">
                  {{ user.robot_url ? user.robot_url : "无信息" }}
                </el-form-item>
                <el-form-item label="机器人密钥: ">
                  {{ user.robot_secret ? user.robot_secret : "无信息" }}
                </el-form-item>
              </el-form>
            </el-col>
            <el-col :span="12" style="border-left: 2px solid #999">
              <div class="title">飞书机器人配置</div>
              <el-form ref="user" :model="user" label-width="120px">
                <el-form-item label="机器人地址:">
                  <el-input
                    v-model="user.robot_url"
                    placeholder="请输入机器人地址"
                  />
                </el-form-item>
                <el-form-item label="机器人密钥:">
                  <el-input
                    v-model="user.robot_secret"
                    placeholder="请输入机器人密钥"
                  />
                </el-form-item>
                <el-form-item>
                  <el-button type="primary" @click="add_robot_btn">
                    {{ user.robot_url ? "更新" : "添加" }}
                  </el-button>
                </el-form-item>
              </el-form>
              <div class="title">选择角色</div>
              <el-form :inline="true" :model="inlineForm">
                <el-form-item label="申请角色:" label-width="120px">
                  <el-select v-model="inlineForm.role" placeholder="选择角色">
                    <el-option
                      v-for="item in roles"
                      :key="item.id"
                      :label="item.desc"
                      :value="item.desc"
                    />
                  </el-select>
                </el-form-item>
                <el-form-item>
                  <el-button type="primary" @click="onSubmit"> 申请 </el-button>
                </el-form-item>
              </el-form>
              <div style="text-align: center">
                <el-tag type="info">
                  申请角色请等待开通，如果着急，请联系管理员开通
                </el-tag>
              </div>
            </el-col>
          </el-row>
        </div>
      </el-card>
    </div>
    <div v-if="is_mobile">
      <el-card class="box-card">
        <div slot="header" class="clearfix">
          <span>个人主页</span>
        </div>
        <div class="card-box">
          <el-row>
            <el-col :span="24">
              <div class="title">个人基础信息</div>
              <el-form ref="user" :model="user" label-width="100px">
                <el-form-item label="用户名: ">
                  {{ user.username }}
                </el-form-item>
                <el-form-item label="邮箱: ">
                  {{ user.email ? user.email : "无信息" }}
                </el-form-item>
                <el-form-item label="角色: ">
                  <span v-for="item in role_names" :key="item">
                    {{ item }}
                  </span>
                </el-form-item>
                <el-form-item label="机器人地址: ">
                  {{ user.robot_url ? user.robot_url : "无信息" }}
                </el-form-item>
                <el-form-item label="机器人密钥: ">
                  {{ user.robot_secret ? user.robot_secret : "无信息" }}
                </el-form-item>
              </el-form>
            </el-col>
            <el-col :span="24">
              <div class="title">飞书机器人配置</div>
              <el-form ref="user" :model="user" label-width="100px">
                <el-form-item label="机器人地址:">
                  <el-input
                    v-model="user.robot_url"
                    placeholder="请输入机器人地址"
                  />
                </el-form-item>
                <el-form-item label="机器人密钥:">
                  <el-input
                    v-model="user.robot_secret"
                    placeholder="请输入机器人密钥"
                  />
                </el-form-item>
                <el-form-item>
                  <el-button type="primary" size="mini" @click="add_robot_btn">
                    {{ user.robot_url ? "更新" : "添加" }}
                  </el-button>
                </el-form-item>
              </el-form>
              <div class="title">选择角色</div>
              <el-form ref="inlineForm" :model="inlineForm" label-width="100px">
                <el-form-item label="申请角色:">
                  <el-select v-model="inlineForm.role" style="width: 100%" placeholder="选择角色">
                    <el-option
                      v-for="item in roles"
                      :key="item.id"
                      :label="item.desc"
                      :value="item.desc"
                    />
                  </el-select>
                </el-form-item>
                <el-form-item>
                  <el-button type="primary" size="mini" @click="onSubmit"> 申请 </el-button>
                </el-form-item>
              </el-form>
              <div style="text-align: center; margin-bottom: 15px;">
                <el-tag type="info">
                  申请角色请等待开通，如果着急，请联系管理员开通
                </el-tag>
              </div>
            </el-col>
          </el-row>
        </div>
      </el-card>
    </div>
  </div>
</template>
<script>
import { getInfo, users_cud, robot_msg } from "@/api/user";
import { manage_role_get } from "@/api/manages/role";
import config from "@/utils/config";
export default {
  name: "Personal",
  data() {
    return {
      user: {},
      roles: [],
      inlineForm: {
        role: "",
      },
      role_names: [],
      is_mobile: config.isMobile
    };
  },
  mounted() {
    this.init_data();
  },
  methods: {
    onSubmit() {
      if (this.inlineForm.role.trim() === "") {
        this.$message({
          message: "角色必须选择",
          type: "warning",
        });
      } else {
        robot_msg(this.inlineForm).then((response) => {
          this.response_func(response);
        });
      }
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
    init_data() {
      manage_role_get().then((response) => {
        if (response.code === 200) {
          this.roles = response.data.roles;
        }
      });
      getInfo().then((response) => {
        if (response.code === 200) {
          this.user = response.data.user_info[0];
          let roles = this.user.roles;
          roles = roles.replace("[", "").replace("]", "").split(",");
          for (var i = 0; i < roles.length; i++) {
            var role_id = roles[i];
            if (role_id !== "") {
              for (var j = 0; j < this.roles.length; j++) {
                var item = this.roles[j];
                if (item.id == role_id) {
                  this.role_names.push(item.name);
                }
              }
            }
          }
        }
      });
    },
    add_robot_btn() {
      const obj = this.user;
      obj["type"] = "robot";
      obj["value"] = "value";
      if (
        obj["robot_url"].indexOf(
          "https://open.feishu.cn/open-apis/bot/v2/hook/"
        ) === -1
      ) {
        this.$message({
          message: "不合理的机器人地址",
          type: "error",
        });
      } else {
        users_cud(obj).then((response) => {
          this.response_func(response);
        });
      }
    },
  },
};
</script>
<style lang="scss" scoped>
.personal {
  width: 99%;
  min-height: 680px;
  margin: 20px auto;
  background-color: #fff;
  padding: 20px;
  .clearfix:after {
    clear: both;
  }

  .box-card {
    width: 99%;
    .card-box {
      .title {
        text-align: center;
        font-size: 18px;
        font-weight: bold;
        line-height: 36px;
      }
    }
  }
}
</style>
