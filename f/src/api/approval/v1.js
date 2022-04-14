import request from "@/utils/request";

export function subscribe_v1_cud(method, data) {
  return request({
    url: "/approval/v1/subscribe",
    method: method,
    data: data,
  });
}

export function up_version(data){
  data["direct"] = "up"
  return request({
    url: "/approval/v2/change-version",
    method: "post",
    data: data,
  })
}

export function down_version(data){
  data["direct"] = "down"
  return request({
    url: "/approval/v2/change-version",
    method: "post",
    data: data,
  })
}

