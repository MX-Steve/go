<template>
  <div class="assets-manage">
    <div v-if="!is_mobile">
      <el-tabs v-model="activeName">
        <el-tab-pane label="区域信息录入" name="first">
          <el-form
            ref="zone"
            :model="zone"
            label-width="120px"
            :rules="rules"
            style="width: 60%"
            class="demo-ruleForm"
          >
            <el-form-item label="区域名称" prop="name">
              <el-input v-model="zone.name" />
            </el-form-item>
            <el-form-item label="描述信息">
              <el-input v-model="zone.description" />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="zone_tab_confirm()">
                立即创建
              </el-button>
              <el-button @click="zone_tab_cancel()">取消</el-button>
            </el-form-item>
          </el-form>
        </el-tab-pane>
        <el-tab-pane label="机房信息录入" name="second">
          <el-form
            ref="idc"
            :model="idc"
            label-width="120px"
            style="width: 60%"
            :rules="idc_rules"
          >
            <el-form-item label="区域名称">
              <el-select
                v-model="idc.zone_id"
                placeholder="请选择区域"
                style="width: 100%"
                @change="zone_change()"
              >
                <el-option
                  v-for="item in zone_list"
                  :key="item.id"
                  :label="item.name"
                  :value="item.id"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="机房名称" prop="name">
              <el-input v-model="idc.name" />
            </el-form-item>
            <el-form-item label="运营商" prop="provider">
              <el-input v-model="idc.provider" />
            </el-form-item>
            <el-form-item label="机房带宽" prop="bandwidth">
              <el-input v-model.number="idc.bandwidth" />
            </el-form-item>
            <el-form-item label="机房网段">
              <el-input v-model="idc.ip_range" />
            </el-form-item>
            <el-form-item label="机房位置">
              <el-input v-model="idc.address" />
            </el-form-item>
            <el-form-item label="机房联系人">
              <el-input v-model="idc.manager" />
            </el-form-item>
            <el-form-item label="联系人电话">
              <el-input v-model="idc.phone" />
            </el-form-item>
            <el-form-item label="备注">
              <el-input v-model="idc.description" />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="idc_confirm()">
                立即创建
              </el-button>
              <el-button @click="idc_cancel()">取消</el-button>
            </el-form-item>
          </el-form>
        </el-tab-pane>
      </el-tabs>
    </div>
    <div v-if="is_mobile">
      <el-tabs v-model="activeName">
        <el-tab-pane label="区域信息录入" name="first">
          <el-form
            ref="zone"
            :model="zone"
            label-width="120px"
            :rules="rules"
            style="width: 100%"
            class="demo-ruleForm"
          >
            <el-form-item label="区域名称" prop="name">
              <el-input v-model="zone.name" />
            </el-form-item>
            <el-form-item label="描述信息">
              <el-input v-model="zone.description" />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="zone_tab_confirm()">
                创建
              </el-button>
              <el-button @click="zone_tab_cancel()">取消</el-button>
            </el-form-item>
          </el-form>
        </el-tab-pane>
        <el-tab-pane label="机房信息录入" name="second">
          <el-form
            ref="idc"
            :model="idc"
            label-width="120px"
            style="width: 100%"
            :rules="idc_rules"
          >
            <el-form-item label="区域名称">
              <el-select
                v-model="idc.zone_id"
                placeholder="请选择区域"
                style="width: 100%"
                @change="zone_change()"
              >
                <el-option
                  v-for="item in zone_list"
                  :key="item.id"
                  :label="item.name"
                  :value="item.id"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="机房名称" prop="name">
              <el-input v-model="idc.name" />
            </el-form-item>
            <el-form-item label="运营商" prop="provider">
              <el-input v-model="idc.provider" />
            </el-form-item>
            <el-form-item label="机房带宽" prop="bandwidth">
              <el-input v-model.number="idc.bandwidth" />
            </el-form-item>
            <el-form-item label="机房网段">
              <el-input v-model="idc.ip_range" />
            </el-form-item>
            <el-form-item label="机房位置">
              <el-input v-model="idc.address" />
            </el-form-item>
            <el-form-item label="机房联系人">
              <el-input v-model="idc.manager" />
            </el-form-item>
            <el-form-item label="联系人电话">
              <el-input v-model="idc.phone" />
            </el-form-item>
            <el-form-item label="备注">
              <el-input v-model="idc.description" />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="idc_confirm()">
                创建
              </el-button>
              <el-button @click="idc_cancel()">取消</el-button>
            </el-form-item>
          </el-form>
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>
<script>
import { zone_info_get, zone_info_cud } from "@/api/assets/basic";
import { idc_get, idc_cud } from "@/api/assets-manage/assets-input";
import config from "@/utils/config";
export default {
  name: "AssetsManage",
  data() {
    var checkIdc = (rule, value, callback) => {
      if (!value) {
        return callback(new Error("区域名称不能为空"));
      }
      if (value.trim().length == 0) {
        return callback(new Error("名称不能为空"));
      }
      for (var i = 0; i < this.idc_list.length; i++) {
        var item = this.idc_list[i];
        if (item.name == value) {
          return callback(new Error("机房名称已存在"));
        }
      }
      callback();
    };
    var checkZone = (rule, value, callback) => {
      if (!value) {
        return callback(new Error("区域名称不能为空"));
      }
      if (value.trim().length == 0) {
        return callback(new Error("名称不能为空"));
      }
      for (var i = 0; i < this.zone_list.length; i++) {
        var item = this.zone_list[i];
        if (item.name == value) {
          return callback(new Error("区域名称已存在"));
        }
      }
      callback();
    };
    return {
      rules: {
        name: [{ required: true, validator: checkZone, trigger: "blur" }],
      },
      idc_rules: {
        name: [{ required: true, validator: checkIdc, trigger: "blur" }],
        provider: [
          { required: true, message: "请输入运营商名称", trigger: "blur" },
        ],
      },
      activeName: "first",
      zone: {
        name: "",
        description: "",
      },
      zone_list: [],
      idc: {
        zone_id: "",
        name: "",
        provider: "",
        bandwidth: 0,
        manager: "",
        phone: "",
        address: "",
        ip_range: "",
        description: "",
      },
      idc_list: [],
      is_mobile: config.isMobile,
    };
  },
  mounted() {
    this.activeName = this.$route.query.activeName || "first";
    this.init_data();
  },
  methods: {
    init_data() {
      zone_info_get().then((response) => {
        if (response.code == 200) {
          this.zone_list = response.data;
        } else {
          this.zone_list = [];
        }
      });
    },
    zone_change() {
      idc_get({ type: "idc_details", zone_id: this.idc.zone_id }).then(
        (response) => {
          if (response.code == 200) {
            this.idc_list = response.data;
          } else {
            this.idc_list = [];
          }
        }
      );
    },
    response_func(response) {
      if (response.code === 200) {
        this.dialogVisible = false;
        this.$message({
          message: response.msg,
          type: "success",
        });
      } else {
        this.$message({
          message: response.msg,
          type: "warning",
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
        setTimeout(() => {
          location.reload(0);
        }, 1000);
      } else {
        this.$message({
          message: response.msg,
          type: "warning",
        });
      }
    },
    zone_tab_confirm() {
      zone_info_cud("post", this.zone).then((response) => {
        this.response_refresh_func(response);
      });
    },
    zone_tab_cancel() {
      this.zone = {
        name: "",
        description: "",
      };
    },
    idc_confirm() {
      console.log(this.idc);
      idc_cud("post", this.idc).then((response) => {
        this.response_refresh_func(response);
      });
    },
    idc_cancel() {
      this.idc = {
        zone: "",
        name: "",
        provider: "",
        bandwidth: 0,
        manager: "",
        phone: "",
        address: "",
        ip_range: "",
        description: "",
      };
    },
  },
};
</script>
<style lang='scss' scoped>
.assets-manage {
  min-height: 450px;
  width: 98%;
  margin: 15px auto;
  background-color: #fff;
  padding: 15px;
  @media screen and (min-width: 1000px){
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
