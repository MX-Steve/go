<template>
  <div class="assets-store-list">
    <el-tabs v-model="activeName" type="card" @tab_click="handleClick">
      <el-tab-pane label="磁盘信息" name="disk">
        <Disk v-if="activeName == 'disk'" />
      </el-tab-pane>
      <el-tab-pane label="存储桶信息" name="oss">
        <Oss v-if="activeName == 'oss'" />
      </el-tab-pane>
      <el-tab-pane label="数据库信息" name="rds">
        <Rds v-if="activeName == 'rds'" />
      </el-tab-pane>
    </el-tabs>
  </div>
</template>
<script>
import { Disk, Oss, Rds } from "./components";
export default {
  name: "AssetsStoreList",
  components: {
    Disk,
    Oss,
    Rds,
  },
  data() {
    return {
      activeName: "disk",
    };
  },
  methods: {
    handleClick(tab, event) {
      console.log(tab, event);
    },
  },
};
</script>
<style lang='scss'>
.assets-store-list {
  width: 99%;
  min-height: 680px;
  margin: 0 auto;
  background-color: #fff;
}
</style>
