<template>
  <div class="side-tree">
    <el-input
      v-model="filterText"
      placeholder="搜索业务树"
      style="margin-bottom: 15px"
    />
    <el-tree
      ref="tree"
      class="filter-tree"
      :data="tree_data"
      node-key="id"
      :props="defaultProps"
      :filter-node-method="filterNode"
      :expand-on-click-node="false"
      @node-click="handleNodeClick"
    />
  </div>
</template>
<script>
export default {
  name: "SideTree",
  props: ["tree_data"],
  data() {
    return {
      filterText: "",
      currentTree: "",
      defaultProps: {
        children: "children",
        label: "label",
      },
      env_id: "",
      service_id: "",
      project_id: "",
    };
  },
  watch: {
    filterText(val) {
      this.$refs.tree.filter(val);
    },
  },
  mounted() {
    this.currentTree = {
      id: this.tree_data[0].id,
      label: this.tree_data[0].label,
    };
  },
  methods: {
    handleNodeClick(data, e) {
      if (e.level === 3) {
        this.env_id = data.id;
        this.service_id = e.parent.data.id;
        this.project_id = e.parent.parent.data.id;
      } else if (e.level === 2) {
        this.service_id = data.id;
        this.project_id = e.parent.data.id;
        this.env_id = "";
      } else {
        this.project_id = data.id;
        this.service_id = "";
        this.env_id = "";
      }
      this.currentTree = data;
      var ids = {
        env_id: this.env_id,
        service_id: this.service_id,
        project_id: this.project_id,
      };
      this.$emit("get_current_tree", this.currentTree);
      this.$emit("get_id", ids);
    },
    filterNode(value, data) {
      if (!value) return true;
      return data.label.indexOf(value) !== -1;
    },
  },
};
</script>
<style lang="scss" scoped>
.side-tree {
  padding: 5px;
  min-height: 650px;
  background-color: #fff;
  border-right: 2px solid #eee;
  /*tree组件背景色修改 */
  .el-tree {
    background-color: #fff;
    color: #333;
    .is-current {
      background-color: #333 !important;
      color: #fff !important;
    }
  }
}
</style>
