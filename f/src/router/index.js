import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

/* Layout */
import Layout from "@/layout";

/* Router Modules */
// import componentsRouter from './modules/components'
// import chartsRouter from './modules/charts'
// import tableRouter from './modules/table'
// import nestedRouter from './modules/nested'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'/'el-icon-x' the icon show in the sidebar
    noCache: true                if set true, the page will no be cached(default is false)
    affix: true                  if set true, the tag will affix in the tags-view
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  // {
  //   path: '/redirect',
  //   component: Layout,
  //   hidden: true,
  //   children: [
  //     {
  //       path: '/redirect/:path(.*)',
  //       component: () => import('@/views/redirect/index')
  //     }
  //   ]
  // },
  {
    path: "/login",
    component: () => import("@/views/login/index"),
    hidden: true,
  },
  {
    path: "/auth-redirect",
    component: () => import("@/views/login/auth-redirect"),
    hidden: true,
  },
  {
    path: "/404",
    component: () => import("@/views/error-page/404"),
    hidden: true,
  },
  {
    path: "/401",
    component: () => import("@/views/error-page/401"),
    hidden: true,
  },
];

/**
 * asyncRoutes
 * the routes that need to be dynamically loaded based on user roles
 */
export const asyncRoutes = [
  {
    path: "/",
    component: Layout,
    redirect: "/dashboard",
    children: [
      {
        path: "dashboard",
        component: () => import("@/views/dashboard/index"),
        name: "Dashboard",
        meta: {
          title: "首页",
          icon: "dashboard",
          tag: "dashboard.dashboard",
          affix: true,
        },
      },
    ],
  },
  {
    path: "/assets",
    component: Layout,
    redirect: "/assets/manage",
    meta: {
      title: "CMDB",
      icon: "documentation",
      affix: true,
      tag: "assets.basic",
    },
    children: [
      {
        path: "manage",
        component: () => import("@/views/assets/manage/index"),
        name: "AssetsManage",
        meta: {
          title: "资产管理",
          icon: "documentation",
          affix: true,
          tag: "assets.idc-input",
        },
        children: [
          {
            path: "basic",
            component: () =>
              import("@/views/assets/manage/assets-manage/basic/index"),
            name: "AssetsManageBasic",
            meta: {
              title: "区域机房录入",
              icon: "guide",
              affix: true,
              tag: "assets.idc-input",
            },
          },
          {
            path: "machine",
            component: () =>
              import("@/views/assets/manage/assets-manage/machine/index"),
            name: "AssetsManageMachine",
            meta: {
              title: "设备信息录入",
              icon: "guide",
              affix: true,
              tag: "assets.machine-input",
            },
          },
        ],
      },
      {
        path: "use",
        component: () => import("@/views/assets/use/index"),
        name: "AssetsUse",
        meta: {
          title: "资产使用",
          icon: "documentation",
          affix: true,
          tag: "assets.basic",
        },
        children: [
          {
            path: "basic",
            component: () => import("@/views/assets/use/basic-assets/index"),
            name: "AssetsUseBasic",
            meta: {
              title: "基础资产",
              icon: "documentation",
              affix: true,
              tag: "assets.basic",
            },
          },
          {
            path: "list",
            component: () => import("@/views/assets/use/assets-list/index"),
            name: "AssetsUseList",
            meta: {
              title: "资产列表",
              icon: "list",
              affix: true,
              tag: "assets.network-list",
            },
            children: [
              {
                path: "network",
                component: () =>
                  import("@/views/assets/use/assets-list/network/index"),
                name: "AssetsUseListNetwork",
                meta: {
                  title: "网络列表",
                  icon: "list",
                  affix: true,
                  tag: "assets.network-list",
                },
              },
              {
                path: "store",
                component: () =>
                  import("@/views/assets/use/assets-list/store/index"),
                name: "AssetsUseListStore",
                meta: {
                  title: "存储列表",
                  icon: "list",
                  affix: true,
                  tag: "assets.store-list",
                },
              },
              {
                path: "machine",
                component: () =>
                  import("@/views/assets/use/assets-list/machine/index"),
                name: "AssetsUseListMachine",
                meta: {
                  title: "机器列表",
                  icon: "list",
                  affix: true,
                  tag: "assets.machine-list",
                },
              },
              {
                path: "domain-parse",
                component: () =>
                  import(
                    "@/views/assets/use/assets-list/network/components/Domain-parse"
                  ),
                name: "AssetsDomainParse",
                meta: {
                  title: "域名解析",
                  icon: "list",
                  affix: true,
                  tag: "assets.domain-record",
                },
                hidden: true,
              },
            ],
          },
        ],
      },
    ],
  },
  {
    path: "/projects",
    component: Layout,
    redirect: "/projects/edit",
    name: "Projects",
    meta: {
      title: "项目管理",
      icon: "documentation",
      affix: true,
      tag: "projects.tree",
    },
    children: [
      {
        path: "edit",
        component: () => import("@/views/projects/project-list/index"),
        name: "ProjectsEdit",
        meta: {
          title: "配置中心",
          icon: "edit",
          affix: true,
          tag: "projects.edit",
        },
      },
      {
        path: "tree-main",
        component: () => import("@/views/projects/tree/index"),
        name: "ProjectsTreeMain",
        meta: {
          title: "业务管理",
          icon: "tree",
          affix: true,
          tag: "projects.tree",
        },
        children: [
          {
            path: "tree",
            component: () => import("@/views/projects/tree/service-tree/index"),
            name: "ProjectsTree",
            meta: {
              title: "业务树",
              icon: "tree",
              affix: true,
              tag: "projects.tree",
            },
          },
          {
            path: "ssh",
            component: () => import("@/views/projects/tree/ssh/Web-ssh"),
            name: "ProjectsSSH",
            meta: {
              title: "ssh终端",
              icon: "el-icon-monitor",
              affix: true,
              tag: "projects.ssh",
            },
            hidden: true,
          },
        ],
      },
    ],
  },
  {
    path: "/jobs",
    component: Layout,
    redirect: "/jobs/list",
    name: "Jobs",
    meta: {
      title: "作业管理",
      icon: "documentation",
      affix: true,
      tag: "jobs.list",
    },
    children: [
      {
        path: "/jobs/list",
        component: () => import("@/views/jobs/task-scheduler/index"),
        name: "JobsList",
        meta: {
          title: "周期性任务",
          icon: "list",
          affix: true,
          tag: "jobs.list",
        },
      },
      {
        path: "/jobs/exec",
        component: () => import("@/views/jobs/exec-task/index"),
        name: "ExecList",
        meta: {
          title: "批量执行",
          icon: "list",
          affix: true,
          tag: "jobs.list",
        },
        children: [
          {
            path: "task",
            component: () =>
              import("@/views/jobs/exec-task/task/index"),
            name: "ExecTask",
            meta: {
              title: "执行任务",
              icon: "edit",
              affix: true,
              tag: "jobs.list",
            },
          },
          {
            path: "template",
            component: () =>
              import("@/views/jobs/exec-task/template/index"),
            name: "ExecTemplate",
            meta: {
              title: "模板管理",
              icon: "edit",
              affix: true,
              tag: "jobs.list",
            },
          },
          {
            path: "history",
            component: () =>
              import("@/views/jobs/exec-task/history/index"),
            name: "ExecHistory",
            meta: {
              title: "执行历史",
              icon: "edit",
              affix: true,
              tag: "jobs.list",
            },
          },
        ]
      },
    ],
  },
  {
    path: "/approvals",
    component: Layout,
    redirect: "/approvals/manage",
    name: "Approvals",
    meta: {
      title: "审批管理",
      icon: "tree",
      affix: true,
      tag: "approvals.list",
    },
    children: [
      {
        path: "manage",
        component: () => import("@/views/approvals/manage/index"),
        name: "AManages",
        meta: {
          title: "审批管理",
          icon: "edit",
          affix: true,
          tag: "approvals.list",
        },
        children: [
          {
            path: "manage",
            component: () =>
              import("@/views/approvals/manage/approval-manage/index"),
            name: "AManage",
            meta: {
              title: "审批管理",
              icon: "edit",
              affix: true,
              tag: "approvals.manage",
            },
            hidden: true,
          },
          {
            path: "manage-old",
            component: () =>
              import("@/views/approvals/manage/approval-manage/old"),
            name: "AOldManage",
            meta: {
              title: "审批管理",
              icon: "edit",
              affix: true,
              tag: "approvals.manage",
            },
            hidden: true,
          },
          {
            path: "list",
            component: () =>
              import("@/views/approvals/manage/approval-list/index"),
            name: "ApprovalList",
            meta: {
              title: "任务创建",
              icon: "list",
              affix: true,
              tag: "approvals.list",
            },
          },
        ],
      },
      {
        path: "use",
        component: () => import("@/views/approvals/use/index"),
        name: "FIUses",
        meta: {
          title: "审批使用",
          icon: "guide",
          affix: true,
          tag: "approvals.fi-list",
        },
        children: [
          {
            path: "start",
            component: () => import("@/views/approvals/use/approval-use/index"),
            name: "FIUse",
            meta: {
              title: "创建工单",
              icon: "people",
              affix: true,
              tag: "approvals.start",
            },
            hidden: true,
          },
          {
            path: "list",
            component: () =>
              import("@/views/approvals/use/instance-list/index"),
            name: "FIList",
            meta: {
              title: "工单列表",
              icon: "list",
              affix: true,
              tag: "approvals.fi-list",
            },
          },
          {
            path: "detail",
            component: () =>
              import("@/views/approvals/use/instance-detail/index"),
            name: "FIDetail",
            meta: {
              title: "工单详情",
              icon: "documentation",
              affix: true,
              tag: "approvals.detail",
            },
            hidden: true,
          },
        ],
      },
      {
        path: "deploy",
        component: () => import("@/views/approvals/deploy/index"),
        name: "Deploy",
        meta: {
          title: "发布历史",
          icon: "documentation",
          affix: true,
          tag: "approvals.deploy-history",
        },
        children: [
          {
            path: "history",
            component: () => import("@/views/approvals/deploy/history/index"),
            name: "DeployHistory",
            meta: {
              title: "发布历史",
              icon: "list",
              affix: true,
              tag: "approvals.deploy-history"
            },
          },
        ],
      },
    ],
  },
  {
    path: "/audits",
    component: Layout,
    redirect: "/audits/list",
    name: "Audits",
    meta: {
      title: "审计管理",
      icon: "documentation",
      affix: true,
      tag: "audits.list",
    },
    children: [
      {
        path: "list",
        component: () => import("@/views/audits/audit/index"),
        name: "AuditsList",
        meta: {
          title: "审计日志",
          icon: "shopping",
          affix: true,
          tag: "audits.list",
        },
      },
    ],
  },
  {
    path: "manges",
    component: Layout,
    redirect: "/manages/role",
    name: "Manages",
    meta: {
      title: "系统管理",
      icon: "skill",
      tag: "manages.role",
      affix: true,
    },
    children: [
      {
        path: "/manages/role",
        component: () => import("@/views/manages/role/index"),
        name: "Manages-role",
        meta: {
          title: "角色管理",
          icon: "skill",
          tag: "manages.role",
          affix: true,
        },
      },
      {
        path: "/manages/user",
        component: () => import("@/views/manages/user/index"),
        name: "Manages-user",
        meta: {
          title: "用户管理",
          icon: "peoples",
          tag: "manages.user",
          affix: true,
        },
      },
    ],
  },
  {
    path: "personal",
    component: Layout,
    redirect: "/personal/info",
    name: "Personal",
    meta: {
      title: "个人中心",
      icon: "people",
      tag: "personal.user_center",
      affix: true,
    },
    children: [
      {
        path: "/personal/info",
        component: () => import("@/views/personal/index"),
        name: "PersonalInfo",
        meta: {
          title: "个人中心",
          icon: "people",
          tag: "personal.user_center",
          affix: true,
        },
      },
    ],
  },
  // 404 page must be placed at the end !!!
  { path: "*", redirect: "/404", hidden: true },
];

const createRouter = () =>
  new Router({
    mode: "history", // require service support
    scrollBehavior: () => ({ y: 0 }),
    routes: constantRoutes,
  });

const router = createRouter();

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter();
  router.matcher = newRouter.matcher; // reset router
}

export default router;
