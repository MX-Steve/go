import request from "@/utils/request";

export function audit_get(query) {
  return request({
    url: "/audit/v1/audit",
    method: "get",
    params: query,
  });
}
