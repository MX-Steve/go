import request from "@/utils/request";

export function tasks_get(query) {
  return request({
    url: "/chat/v1/tasks",
    method: "get",
    params: query,
  });
}
export function tasks_cud(method, data) {
  return request({
    url: "/chat/v1/tasks",
    method: method,
    data: data,
  });
}

export function intervals_get(query) {
  return request({
    url: "/chat/v1/intervals",
    method: "get",
    params: query,
  });
}
export function intervals_cud(method, data) {
  return request({
    url: "/chat/v1/intervals",
    method: method,
    data: data,
  });
}

export function crontab_get(query) {
  return request({
    url: "/chat/v1/crontab",
    method: "get",
    params: query,
  });
}
export function crontab_cud(method, data) {
  return request({
    url: "/chat/v1/crontab",
    method: method,
    data: data,
  });
}

export function results_get(query) {
  return request({
    url: "/chat/v1/results",
    method: "get",
    params: query,
  });
}
export function results_cud(method, data) {
  return request({
    url: "/chat/v1/results",
    method: method,
    data: data,
  });
}

export function deploy_one(data) {
  return request({
    url: "/assets/v1/run-one",
    method: "post",
    data: data,
  })
}
