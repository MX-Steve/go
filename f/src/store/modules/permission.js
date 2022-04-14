import { asyncRoutes, constantRoutes } from "@/router";
import jwtDecode from "jwt-decode";
import { getToken } from "@/utils/auth";

/**
 * Use meta.role to determine if the current user has permission
 * @param roles
 * @param route
 */

function hasPermission(roles, route) {
  const username = jwtDecode(getToken()).username;
  if (username === "admin") {
    return true;
  }
  if (route.meta && route.meta.tag) {
    var page_tag = route.meta.tag;
    var perm = false;
    for (var i = 0; i < roles.length; i++) {
      var item = roles[i];
      if (item.indexOf(page_tag) !== -1) {
        perm = true;
        break;
      }
    }
    return perm;
    // return roles.some(role => route.meta.roles.includes(role))
  }
  return true;
}

/**
 * Filter asynchronous routing tables by recursion
 * @param routes asyncRoutes
 * @param roles
 */
export function filterAsyncRoutes(routes, roles) {
  const res = [];

  routes.forEach((route) => {
    const tmp = { ...route };
    if (hasPermission(roles, tmp)) {
      if (tmp.children) {
        tmp.children = filterAsyncRoutes(tmp.children, roles);
      }
      res.push(tmp);
    }
  });

  return res;
}

const state = {
  routes: [],
  addRoutes: [],
};

const mutations = {
  SET_ROUTES: (state, routes) => {
    state.addRoutes = routes;
    state.routes = constantRoutes.concat(routes);
  },
};

const actions = {
  generateRoutes({ commit }, roles) {
    return new Promise((resolve) => {
      let accessedRoutes;
      if (roles.includes("admin")) {
        accessedRoutes = asyncRoutes || [];
      } else {
        accessedRoutes = filterAsyncRoutes(asyncRoutes, roles);
      }
      commit("SET_ROUTES", accessedRoutes);
      resolve(accessedRoutes);
    });
  },
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
};
