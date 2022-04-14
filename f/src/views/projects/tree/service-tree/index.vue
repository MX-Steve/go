<template>
  <div class="service-tree">
    <div class="web">
      <el-row>
        <el-col :span="6">
          <SideTree
            ref="side"
            :tree_data="tree_data"
            @get_current_tree="get_current_tree"
            @get_id="get_id"
          />
        </el-col>
        <el-col :span="18">
          <ShowHost :tree="currentTree" :ids="currentId" />
        </el-col>
      </el-row>
    </div>
    <div class="mobile">
      <el-row>
        <el-col :span="10">
          <SideTree
            ref="side"
            :tree_data="tree_data"
            @get_current_tree="get_current_tree"
            @get_id="get_id"
          />
        </el-col>
        <el-col :span="14">
          <ShowHost :tree="currentTree" :ids="currentId" />
        </el-col>
      </el-row>
    </div>
  </div>
</template>
<script>
import { tree_data_get } from "@/api/tree/index";
import { SideTree, ShowHost } from "./components";
export default {
  name: "ServiceTree",
  components: {
    SideTree,
    ShowHost,
  },
  data() {
    return {
      title: "业务树",
      currentTree: {},
      tree_data: [],
      currentId: {},
    };
  },
  beforeMount() {
    this.init_data();
  },
  methods: {
    init_data() {
      this.get_tree();
      this.tree_data = [
        {
          id: 1,
          label: "占位项目",
          children: [
            {
              id: 2,
              label: "占位服务",
              children: [
                {
                  id: 4,
                  label: "开发环境",
                },
              ],
            },
          ],
        },
      ];
      this.currentTree = {
        id: this.tree_data[0].id,
        label: this.tree_data[0].label,
      };
    },
    get_tree() {
      tree_data_get().then((response) => {
        if (response.code === 200) {
          this.tree_data = response.data;
        }
      });
    },
    get_current_tree(val) {
      this.currentTree = val;
    },
    get_id(val) {
      this.currentId = val;
    },
  },
};
</script>
<style lang='scss' scoped>
.service-tree {
  padding: 10px;
  @media screen and (min-width: 1000px) {
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
