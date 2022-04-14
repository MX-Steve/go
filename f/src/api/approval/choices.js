import request from "@/utils/request";

export function choices_get(query) {
  return request({
    url: "/approval/v2/choices",
    method: "get",
    params: query,
  });
}
