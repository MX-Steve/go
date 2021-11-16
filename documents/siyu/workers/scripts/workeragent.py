import json
import os
import sys
import boto3
import requests
import re
import subprocess
import time


BASE_DIR = os.path.dirname(__file__)
sys.path.append(BASE_DIR)

def gcp_get_public_ip():
    """
    获取外网IP
    """
    url = "https://api.ipify.org/?format=json"
    ip = None
    try:
        info = requests.get(url).text
        info = json.loads(info)
        ip = info.get('ip')
        return ip
    except Exception as err:
        print(err)
        return ip

def gcp_get_info_by_ip(ip):
    """
    (base) root@instance-test:~# gcloud compute instances list | grep 34.133.224.201
    instance-test  us-central1-a  n1-standard-1 true/(space) 10.128.0.15  34.133.224.201  RUNNING
    """
    cmd = "gcloud compute instances list | grep " + ip
    output = subprocess.check_output(cmd, shell=True)
    res = output.decode()
    info = re.split(r"[ ]+", res)
    return info

def gcp_get_instance_type():
    """
    获取instanceType
    """
    info = gcp_get_info_by_ip(gcp_get_public_ip())
    if(len(info) >= 3):
        return info[2]
    else:
        return None

def gcp_generate_conf():
    """
    生成supervisor的配置文件
    """
    it = gcp_get_instance_type()
    dc_name = "GCP-%s"%(it)
    with open('workeragent.conf') as f:
        context = f.read() % dc_name
        print(context)
    with open("/etc/supervisor/conf.d/workeragent.conf", 'w') as f:
        f.write(context)

def get_instance_type():
    """
    获取EC2的instanceType
    """
    url = 'http://169.254.169.254/latest/dynamic/instance-identity/document'
    instanceType = None
    try:
        info = requests.get(url).text
        info = json.loads(info)
        instanceType = info.get('instanceType')
        return instanceType
    except Exception as err:
        print(err)
        return instanceType


def generate_conf():
    """
    生成supervisor的配置文件
    """
    it = get_instance_type()
    dc_name = "AWS-%s"%(it)
    with open('workeragent.conf') as f:
        context = f.read() % dc_name
        print(context)
    with open("/etc/supervisor/conf.d/workeragent.conf", 'w') as f:
        f.write(context)


def main():
    is_gcp = True
    if(is_gcp):    #GCP
        gcp_generate_conf()
    else:          #AWS
        generate_conf()
    supervisor_update = 'supervisorctl update'
    start_workeragent = 'supervisorctl start workeragent'
    commands = [supervisor_update, start_workeragent]
    for command in commands:
        res = subprocess.call(command, shell=True)
        print(res)




if __name__ == '__main__':
    main()