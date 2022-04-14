import request from "@/utils/request";

export function device_type_get(query) {
  return request({
    url: "/assets/v1/device-type",
    method: "get",
    params: query,
  });
}

export function device_type_cud(method, data) {
  return request({
    url: "/assets/v1/device-type",
    method: method,
    data: data,
  });
}
export function device_status_get(query) {
  return request({
    url: "/assets/v1/device-status",
    method: "get",
    params: query,
  });
}

export function device_status_cud(method, data) {
  return request({
    url: "/assets/v1/device-status",
    method: method,
    data: data,
  });
}

export function zone_info_get() {
  return request({
    url: "/assets/v1/zones",
    method: "get",
  });
}

export function zone_info_cud(method, data) {
  return request({
    url: "/assets/v1/zones",
    method: method,
    data: data,
  });
}

export function idle_assets_get() {
  return request({
    url: "/assets/v1/idle-assets",
    method: "get",
  });
}

export function idle_assets_cud(method, data) {
  return request({
    url: "/assets/v1/idle-assets",
    method: method,
    data: data,
  });
}
