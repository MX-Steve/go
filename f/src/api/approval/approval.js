import request from "@/utils/request";

export function approval_cud(method, data) {
  return request({
    url: "/approval/v2/approval",
    method: method,
    data: data,
  });
}
export function approval_get(query) {
  return request({
    url: "/approval/v2/approval",
    method: "get",
    params: query,
  });
}
export function form_cud(method, data) {
  return request({
    url: "/approval/v2/form",
    method: method,
    data: data,
  });
}
export function form_get(query) {
  return request({
    url: "/approval/v2/form",
    method: "get",
    params: query,
  });
}
export function subscribe_cud(method, data) {
  return request({
    url: "/approval/v2/subscribe",
    method: method,
    data: data,
  });
}
export function node_cud(method, data) {
  return request({
    url: "/approval/v2/node",
    method: method,
    data: data,
  });
}
export function node_get(query) {
  return request({
    url: "/approval/v2/node",
    method: "get",
    params: query,
  });
}
export function deploy_jenkins(data) {
  return request({
    url: "/approval/v2/direct-deploy",
    method: "post",
    data: data,
  });
}
