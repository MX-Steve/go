import Cookies from "js-cookie";

const TokenKey = "Admin-Token";

export function getToken() {
  return Cookies.get(TokenKey);
}

export function setToken(token) {
  return Cookies.set(TokenKey, token);
}

export function removeToken() {
  return Cookies.remove(TokenKey);
}
export const BasicPage = { keys: "assets.basic" };
export const NetworkPage = { keys: "assets.network-list" };
export const StorePage = { keys: "assets.store-list" };
export const MachinePage = { keys: "assets.machine-list" };
export const ProjectsPage = { keys: "projects.edit" };
export const JobsPage = { keys: "jobs.list" };
export const DomainRecordPage = { keys: "assets.domain-record" };
export const DeployHistoryPage = { keys: "approvals.deploy-history" };
export const ApprovalDetailsPage = { keys: "approvals.detail" };
export const ApprovalListPage = { keys: "approvals.list" };
export const ApprovalManagePage = { keys: "approvals.manage" };
export const ApprovalFiListPage = { keys: "approvals.fi-list" };
export const ApprovalStartPage = { keys: "approvals.start" };
export const ManageRolePage = { keys: "manages.role" };
export const ExecTaskPage = { keys: "exec-task.task" };
export const ExecTemplatePage = { keys: "exec-task.template" };
export const ExecHistoryPage = { keys: "exec-task.history" };
