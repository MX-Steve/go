import request from "@/utils/request";

export function btn_check(data) {
  return request({
    url: "/users/v1/btn-check",
    method: "post",
    data,
  });
}
