<template>
  <div :class="classObj" class="app-wrapper">
    <div
      v-if="device === 'mobile' && sidebar.opened"
      class="drawer-bg"
      @click="handleClickOutside"
    />
    <sidebar class="sidebar-container" />
    <div :class="{ hasTagsView: needTagsView }" class="main-container">
      <div v-if="showHeader" :class="{ 'fixed-header': fixedHeader }">
        <navbar />
        <tags-view v-if="needTagsView" />
      </div>
      <app-main />
      <right-panel v-if="showSettings">
        <settings />
      </right-panel>
    </div>
  </div>
</template>

<script>
import Cookies from "js-cookie";
import RightPanel from "@/components/RightPanel";
import { AppMain, Navbar, Settings, Sidebar, TagsView } from "./components";
import ResizeMixin from "./mixin/ResizeHandler";
import { mapState } from "vuex";
import { infos_get } from "@/api/approval/infos";

export default {
  name: "Layout",
  components: {
    AppMain,
    Navbar,
    RightPanel,
    Settings,
    Sidebar,
    TagsView,
  },
  mixins: [ResizeMixin],
  computed: {
    ...mapState({
      sidebar: (state) => state.app.sidebar,
      device: (state) => state.app.device,
      showSettings: (state) => state.settings.showSettings,
      needTagsView: (state) => state.settings.tagsView,
      fixedHeader: (state) => state.settings.fixedHeader,
      showHeader: (state) => state.settings.showHeader,
    }),
    classObj() {
      return {
        hideSidebar: !this.sidebar.opened,
        openSidebar: this.sidebar.opened,
        withoutAnimation: this.sidebar.withoutAnimation,
        mobile: this.device === "mobile",
      };
    },
  },
  mounted() {
    this.message_info();
    const superuser = Cookies.get("superuser");
    if (superuser == "false") {
      document.oncontextmenu = function() {
        // 右键禁用
        return false;
      };
      document.addEventListener("keydown", function(e) {
        // 按键禁用
        try {
          e = e || event || window.event;
          var keycode = e.keyCode || e.which || e.charCode;
          var ctrlKey = e.ctrlKey || e.metaKey;
          if (keycode == 123) {
            // F12
            e.preventDefault();
            window.event.returnValue = false;
          } else if (keycode == 8) {
            var act = document.activeElement.tagName.toLowerCase();
            if (act.indexOf("body") != -1 || act.indexOf("html") != -1) {
              window.history.back();
              return false;
            }
            return true;
          } else if (ctrlKey && keycode == 83) {
            // ctrl+S
            e.preventDefault();
            window.event.returnValue = false;
          } else if (ctrlKey && keycode == 87) {
            // ctrl+W
            e.preventDefault();
            window.event.returnValue = false;
          } else if (ctrlKey && keycode == 107) {
            // ctrl++
            e.preventDefault();
            window.event.returnValue = false;
          } else if (ctrlKey && keycode == 109) {
            // ctrl+-
            e.preventDefault();
            window.event.returnValue = false;
          } else if (keycode == 122) {
            // F11
            $("#switchFullScreenCloseBtn").trigger("click");
          }
        } catch (e) {
          console.error(e);
        }
      });
    }
  },
  methods: {
    handleClickOutside() {
      this.$store.dispatch("app/closeSideBar", { withoutAnimation: false });
    },
    message_info() {
      const protocol = window.location.protocol;
      const host = window.location.host;
      let url_basic = `${protocol}//${host}/approvals/use/detail?instance_id=`;
      infos_get().then((response) => {
        if (response.code === 200) {
          const data = response.data;
          for (var i in data) {
            const approval_name = data[i]["approval_name"];
            const instance_ids = data[i]["instance_ids"];
            for (let i = 0; i < instance_ids.length; i++) {
              let url = url_basic + instance_ids[i];
              this.$notify.info({
                title: `服务发布[${approval_name}]`,
                dangerouslyUseHTMLString: true,
                message: url + `<a href='${url}' style='color: green;font-weight: bold;' target='_blank'> 前往 </a>`,
                duration: 0,
                position: "bottom-right",
              });
            }
          }
        }
      });
    },
  },
};
</script>

<style lang="scss" scoped>
@import "~@/styles/mixin.scss";
@import "~@/styles/variables.scss";

.app-wrapper {
  background-color: #eee;
  @include clearfix;
  position: relative;
  height: 100%;
  width: 100%;

  &.mobile.openSidebar {
    position: fixed;
    top: 0;
  }
}

.drawer-bg {
  background: #000;
  opacity: 0.3;
  width: 100%;
  top: 0;
  height: 100%;
  position: absolute;
  z-index: 999;
}

.fixed-header {
  position: fixed;
  top: 0;
  right: 0;
  z-index: 9;
  width: calc(100% - #{$sideBarWidth});
  transition: width 0.28s;
}

.hideSidebar .fixed-header {
  width: calc(100% - 54px);
}

.mobile .fixed-header {
  width: 100%;
}
</style>
