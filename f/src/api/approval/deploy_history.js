import request from "@/utils/request";

export function deploy_history_cud(method, data) {
  return request({
    url: "/approval/v2/deploy-history",
    method: method,
    data: data,
  });
}
export function deploy_history_get(query) {
  return request({
    url: "/approval/v2/deploy-history",
    method: "get",
    params: query,
  });
}
