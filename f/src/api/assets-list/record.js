import request from "@/utils/request";

// 网络部分接口
export function record_get(query) {
  return request({
    url: "/assets/v1/record",
    method: "get",
    params: query,
  });
}

export function record_cud(method, data) {
  return request({
    url: "/assets/v1/record",
    method: method,
    data: data,
  });
}
