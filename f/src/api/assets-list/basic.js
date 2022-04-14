import request from "@/utils/request";

// 网络部分接口
export function cdn_get(query) {
  return request({
    url: "/assets/v1/cdn",
    method: "get",
    params: query,
  });
}

export function cdn_cud(method, data) {
  return request({
    url: "/assets/v1/cdn",
    method: method,
    data: data,
  });
}

export function domain_get(query) {
  return request({
    url: "/assets/v1/domain",
    method: "get",
    params: query,
  });
}

export function domain_cud(method, data) {
  return request({
    url: "/assets/v1/domain",
    method: method,
    data: data,
  });
}

export function vpc_get(query) {
  return request({
    url: "/assets/v1/vpc",
    method: "get",
    params: query,
  });
}

export function vpc_cud(method, data) {
  return request({
    url: "/assets/v1/vpc",
    method: method,
    data: data,
  });
}

// 存储部分接口
export function rds_get(query) {
  return request({
    url: "/assets/v1/rds",
    method: "get",
    params: query,
  });
}

export function rds_cud(method, data) {
  return request({
    url: "/assets/v1/rds",
    method: method,
    data: data,
  });
}

export function oss_get(query) {
  return request({
    url: "/assets/v1/oss",
    method: "get",
    params: query,
  });
}

export function oss_cud(method, data) {
  return request({
    url: "/assets/v1/oss",
    method: method,
    data: data,
  });
}
export function disk_get(query) {
  return request({
    url: "/assets/v1/disk",
    method: "get",
    params: query,
  });
}

export function disk_cud(method, data) {
  return request({
    url: "/assets/v1/disk",
    method: method,
    data: data,
  });
}
export function slb_get(query) {
  return request({
    url: "/assets/v1/slb",
    method: "get",
    params: query,
  });
}

export function slb_cud(method, data) {
  return request({
    url: "/assets/v1/slb",
    method: method,
    data: data,
  });
}
export function switch_get(query) {
  return request({
    url: "/assets/v1/switch",
    method: "get",
    params: query,
  });
}

export function switch_cud(method, data) {
  return request({
    url: "/assets/v1/switch",
    method: method,
    data: data,
  });
}
