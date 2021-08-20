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
    generate_conf()
    supervisor_update = 'supervisorctl update'
    start_workeragent = 'supervisorctl start workeragent'
    commands = [supervisor_update, start_workeragent]
    for command in commands:
        res = subprocess.call(command, shell=True)
        print(res)




if __name__ == '__main__':
    main()