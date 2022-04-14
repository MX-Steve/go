import request from "@/utils/request";

export function login(data) {
  return request({
    url: "/users/v1/login",
    method: "post",
    data,
  });
}

export function getInfo() {
  return request({
    url: "/users/v1/userinfo",
    method: "get",
  });
}

export function logout() {
  return request({
    url: "/users/v1/logout",
    method: "post",
  });
}

export function users_get(query) {
  return request({
    url: "/users/v1/users",
    method: "get",
    params: query
  });
}

export function user_get(query) {
  return request({
    url: "/users/v1/users",
    method: "get",
    params: query
  });
}
export function usernames_get(query) {
  return request({
    url: "/users/v1/usernames",
    method: "get",
    params: query
  });
}

export function user_login_get(query) {
  return request({
    url: "/users/v1/dashboard",
    method: "get",
    params: query
  });
}

export function users_cud(data) {
  return request({
    url: "/users/v1/users",
    method: "post",
    data: data,
  });
}

export function robot_msg(data) {
  return request({
    url: "/users/v1/robot-msg",
    method: "post",
    data: data
  });
}
