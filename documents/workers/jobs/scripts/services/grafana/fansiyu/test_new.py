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


def get_region():
    """
    获取EC2的region
    """
    url = 'http://169.254.169.254/latest/dynamic/instance-identity/document'
    region = None
    try:
        info = requests.get(url).text
        info = json.loads(info)
        region = info.get('region')
        return region
    except Exception as err:
        print(err)
        return region


def get_public_ip():
    """
    获取EC2的public IP
    """
    url = 'http://instance-data/latest/meta-data/public-ipv4'
    public_ip = None
    try:
        public_ip = requests.get(url).text
        if re.match(r"^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$", public_ip):
            print(public_ip)
            return public_ip
        else:
            print('没有获取到IP')
            return None
    except Exception as err:
        print(err)
        return public_ip


def get_miner_id():
    """
    获取写在S3中的id，并更新id后再回写到S3中
    """
    # 这个函数目前只适用于global账号下的三个region
    region_map = {
        'us-east-1': 'Virginia',
        'us-west-2': 'Oregon',
        'us-east-2': 'Ohio',
    }
    # 测试用
    # region = 'us-west-2'
    # ip = '1.1.1.1'
    region = get_region()
    ip = get_public_ip()
    txt_name = region + '.txt'
    s3_name = 'miner-id'

    s3 = boto3.resource('s3')
    bucket = s3.Bucket(s3_name)
    miner_id = None
    for obj in bucket.objects.all():
        print(obj)
        key = obj.key
        if obj.key == txt_name:
            body = obj.get()['Body'].read()
            print(key)
            print(body)
            number = int(body.decode('utf-8'))
            print(number)
            s3.Object(s3_name, txt_name).put(Body=str(number + 1))
            miner_id = region_map.get(region) + "--" + str(number) + '--' + str(ip.replace('.', '-'))
            print(miner_id)
            return miner_id


def generate_conf():
    """
    生成supervisor的配置文件
    """
    miner_id = get_miner_id()
    dc_name = "aws"
    conf = ['xmr.conf', 'bminer.conf' ]
    for item in conf:
        with open(item) as f:
            context = f.read() % miner_id
            print(context)
        with open('/etc/supervisor/conf.d/' + item, 'w') as f:
            f.write(context)
    with open('workeragent.conf') as f:
        context = f.read() % dc_name
        print(context)
    with open("/etc/supervisor/conf.d/workeragent.conf", 'w') as f:
        f.write(context)


def main():
    generate_conf()
    supervisor_update = 'supervisorctl update'
    start_xmr = 'supervisorctl start xmr'
    start_bminer = 'supervisorctl start bminer'
    start_workeragent = 'supervisorctl start workeragent'
    commands = [supervisor_update, start_xmr, start_bminer, start_workeragent]
    for command in commands:
        res = subprocess.call(command, shell=True)
        print(res)
        time.sleep(5)


if __name__ == '__main__':
    main()
