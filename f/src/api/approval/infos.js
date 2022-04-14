import request from "@/utils/request";

export function infos_get(query) {
  return request({
    url: "/approval/v2/infos",
    method: "get",
    params: query,
  });
}
