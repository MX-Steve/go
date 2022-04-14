import request from "@/utils/request";

export function manage_role_get(query) {
  return request({
    url: "/manages/v1/role",
    method: "get",
    params: query
  });
}

export function manage_role_cud(method, data) {
  return request({
    url: "manages/v1/role",
    method: method,
    data: data,
  });
}
