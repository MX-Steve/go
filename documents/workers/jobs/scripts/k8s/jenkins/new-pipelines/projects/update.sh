#!/bin/bash
ProjectName=$1
if [ -d /root/update/projects/$ProjectName ];then
  helm uninstall $ProjectName
  helm install $ProjectName /root/update/projects/${ProjectName}/
else
  echo "项目为新项目，请联系管理员手动部署"
  exit 1
fi