import request from "@/utils/request";

// 环境部分接口
export function env_get(query) {
  return request({
    url: "/assets/v1/environment",
    method: "get",
    params: query,
  });
}

export function env_cud(method, data) {
  return request({
    url: "/assets/v1/environment",
    method: method,
    data: data,
  });
}

// 项目部分接口
export function project_get(query) {
  return request({
    url: "/assets/v1/project",
    method: "get",
    params: query,
  });
}

export function project_cud(method, data) {
  return request({
    url: "/assets/v1/project",
    method: method,
    data: data,
  });
}

// 服务部分接口
export function service_get(query) {
  return request({
    url: "/assets/v1/services",
    method: "get",
    params: query,
  });
}

export function service_cud(method, data) {
  return request({
    url: "/assets/v1/services",
    method: method,
    data: data,
  });
}
