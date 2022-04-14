import request from "@/utils/request";

export function idc_get(query) {
  return request({
    url: "/assets/v1/idc",
    method: "get",
    params: query,
  });
}

export function idc_cud(method, data) {
  return request({
    url: "/assets/v1/idc",
    method: method,
    data: data,
  });
}

export function machine_get(query) {
  return request({
    url: "/assets/v1/machine",
    method: "get",
    params: query,
  });
}

export function machine_cud(method, data) {
  return request({
    url: "/assets/v1/machine",
    method: method,
    data: data,
  });
}

export function go_ssh(query) {
  return request({
    url: "/assets/v1/can-go-ssh",
    method: "get",
    params: query,
  });
}
