import request from "@/utils/request";

export function summary_machine_get() {
  return request({
    url: "/assets/v1/summary-machine",
    method: "get",
  });
}

export function summary_approval_get() {
  return request({
    url: "/approval/v2/summary-approval",
    method: "get"
  });
}
