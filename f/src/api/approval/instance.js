import request from "@/utils/request";

export function instance_cud(method, data) {
  return request({
    url: "/approval/v2/instance",
    method: method,
    data: data,
  });
}
export function instance_get(query) {
  return request({
    url: "/approval/v2/instance",
    method: "get",
    params: query,
  });
}
export function fiform_cud(method, data) {
  return request({
    url: "/approval/v2/fiform",
    method: method,
    data: data,
  });
}
export function fiform_get(query) {
  return request({
    url: "/approval/v2/fiform",
    method: "get",
    params: query,
  });
}
export function tasks_cud(method, data) {
  return request({
    url: "/approval/v2/tasks",
    method: method,
    data: data,
  });
}
export function tasks_get(query) {
  return request({
    url: "/approval/v2/tasks",
    method: "get",
    params: query,
  });
}
export function task_cud(method, data) {
  return request({
    url: "/approval/v2/tasks",
    method: method,
    data: data,
  });
}

