<template>
  <div class="web-ssh">
    <el-tabs
      v-model="editableTabsValue"
      type="card"
      editable
      @edit="handleTabsEdit"
    >
      <el-tab-pane
        v-for="(item, index) in editableTabs"
        :key="index"
        :label="item.title"
        :name="item.name"
      />
    </el-tabs>
    <SSH v-if="instance_name != ''" :machine="machine" :instance_name="instance_name" />
    <div
      v-else
      id="terminal"
      v-loading="loading"
      style="min-height: 450px;"
      class="console"
      element-loading-text="机器好像不存在"
    />
  </div>
</template>

<script>
import { machine_get } from "@/api/assets-manage/assets-input";
import { SSH } from "@/views/projects/tree/service-tree/components";
export default {
  name: "WebSsh",
  components: {
    SSH,
  },
  data() {
    return {
      loading: true,
      editableTabsValue: "1",
      editableTabs: [
        {
          title: "undefined",
          name: "1",
        },
      ],
      tabIndex: 2,
      machine: {},
      instance_name: ""
    };
  },
  mounted() {
    this.init_data();
  },
  methods: {
    handleTabsEdit(targetName, action) {
      if (action === "add") {
        const newTabName = ++this.tabIndex + "";
        this.editableTabs.push({
          title: "New Tab",
          name: newTabName,
          content: "New Tab content",
        });
        this.editableTabsValue = newTabName;
      }
      if (action === "remove") {
        const tabs = this.editableTabs;
        let activeName = this.editableTabsValue;
        if (activeName === targetName) {
          tabs.forEach((tab, index) => {
            if (tab.name === targetName) {
              const nextTab = tabs[index + 1] || tabs[index - 1];
              if (nextTab) {
                activeName = nextTab.name;
              }
            }
          });
        }

        this.editableTabsValue = activeName;
        this.editableTabs = tabs.filter((tab) => tab.name !== targetName);
      }
    },
    init_data() {
      machine_get({ type: "machine_details", id: this.$route.query.id }).then(
        (response) => {
          if (response.code == 200) {
            this.machine = response.data[0];
            this.editableTabs[0].title = this.machine.instance_name;
            this.instance_name = this.machine.instance_name;
          } else {
            this.$message({
              message: response.msg,
              type: "warning",
            });
          }
        }
      );
    },
  },
};
</script>

<style lang="scss">
.web-ssh {
  margin: 10px;
  padding: 5px;
  background: #fff;
  .title {
    font-size: 24px;
    font-weight: bold;
  }
}
</style>
