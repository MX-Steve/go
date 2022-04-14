import request from "@/utils/request";

export function p_ssh(data) {
  return request({
    url: "/vssh-chat/p-host",
    method: "post",
    data,
  });
}

export function g_ssh(query) {
  return request({
    url: "/vssh-chat/g-host",
    method: "get",
    params: query
  });
}
