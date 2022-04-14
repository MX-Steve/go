import request from "@/utils/request";

export function tree_data_get(query) {
  return request({
    url: "/assets/v1/tree-data",
    method: "get",
    params: query,
  });
}

export function service_get(query) {
  return request({
    url: "/assets/v1/tree-machines",
    method: "get",
    params: query,
  });
}
