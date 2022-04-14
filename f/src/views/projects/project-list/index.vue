<template>
  <div class="project-list">
    <el-tabs v-model="activeName" type="card">
      <el-tab-pane label="环境配置" name="env">
        <Env v-if="activeName == 'env'" />
      </el-tab-pane>
      <el-tab-pane label="项目配置" name="project">
        <Project v-if="activeName == 'project'" />
      </el-tab-pane>
      <el-tab-pane label="服务配置" name="service">
        <Service v-if="activeName == 'service'" />
      </el-tab-pane>
    </el-tabs>
  </div>
</template>
<script>
import { Env, Project, Service } from "./components";
export default {
  name: "ProjectList",
  components: {
    Env,
    Project,
    Service,
  },
  data() {
    return {
      activeName: "env",
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
<style lang="scss" scoped>
.project-list {
  width: 99%;
  min-height: 680px;
  margin: 0 auto;
  background-color: #fff;
}
</style>
