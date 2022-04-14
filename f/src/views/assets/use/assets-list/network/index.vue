<template>
  <div class="assets-network-list">
    <div v-if="!is_mobile">
      <el-tabs v-model="activeName" type="card">
        <el-tab-pane label="CDN信息" name="cdn">
          <CDN v-if="activeName == 'cdn'" />
        </el-tab-pane>
        <el-tab-pane label="Domain信息" name="domain">
          <Domain v-if="activeName == 'domain'" />
        </el-tab-pane>
        <el-tab-pane label="SLB信息" name="slb">
          <Slb v-if="activeName == 'slb'" />
        </el-tab-pane>
        <el-tab-pane label="VPC信息" name="vpc">
          <Vpc v-if="activeName == 'vpc'" />
        </el-tab-pane>
        <el-tab-pane label="Switch信息" name="switch">
          <Switch2 v-if="activeName == 'switch'" />
        </el-tab-pane>
      </el-tabs>
    </div>
    <div v-if="is_mobile">
      <el-tabs v-model="activeName" type="card">
        <el-tab-pane label="CDN" name="cdn">
          <CDN v-if="activeName == 'cdn'" />
        </el-tab-pane>
        <el-tab-pane label="Domain" name="domain">
          <Domain v-if="activeName == 'domain'" />
        </el-tab-pane>
        <el-tab-pane label="SLB" name="slb">
          <Slb v-if="activeName == 'slb'" />
        </el-tab-pane>
        <el-tab-pane label="VPC" name="vpc">
          <Vpc v-if="activeName == 'vpc'" />
        </el-tab-pane>
        <el-tab-pane label="Switch" name="switch">
          <Switch2 v-if="activeName == 'switch'" />
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>
<script>
import { Slb, Domain, Vpc, CDN, Switch2 } from "./components";
import config from "@/utils/config";
import { NetworkPage } from "@/utils/auth";
export default {
  name: "AssetsNetworkList",
  components: {
    Slb,
    Domain,
    Vpc,
    CDN,
    Switch2,
  },
  data() {
    return {
      activeName: "cdn",
      is_mobile: config.isMobile,
      NetworkPage: NetworkPage
    };
  },
  methods: {
  },
};
</script>
<style lang='scss'>
.assets-network-list {
  width: 99%;
  min-height: 680px;
  margin: 0 auto;
  background-color: #fff;
}
</style>
