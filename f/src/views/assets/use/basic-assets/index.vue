<template>
  <div class="basic-assets">
    <el-tabs v-model="activeName" type="card">
      <el-tab-pane label="机房信息" name="first">
        <MachineInfo v-if="activeName == 'first'" />
      </el-tab-pane>
      <el-tab-pane label="标签分组" name="second">
        <DeviceType v-if="activeName == 'second'" />
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import { MachineInfo, DeviceType } from "./components";
export default {
  name: "BasicAssets",
  components: {
    MachineInfo,
    DeviceType,
  },
  data() {
    return {
      activeName: "first",
    };
  },
  mounted() {
    const tab = this.$route.query.tab;
    if (tab !== undefined) {
      this.activeName = tab;
    }
  },
  methods: {
  },
};
</script>
<style lang='scss' scoped>
.basic-assets {
  width: 99%;
  min-height: 680px;
  margin: 0 auto;
  background-color: #fff;
}
</style>
