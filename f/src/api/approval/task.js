import request from "@/utils/request";

export function task_user_cud(data) {
  return request({
    url: "/approval/v2/task-modify-user",
    method: "post",
    data: data,
  });
}
export function task_get(query) {
  return request({
    url: "/approval/v2/tasks",
    method: "get",
    params: query,
  });
}
