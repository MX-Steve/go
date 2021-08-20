import json
import os
import sys
import boto3
import requests
import re
import subprocess
import pymysql
import time
import random


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


def get_az():
    """
    get az info
    """
    region = get_region()
    client = boto3.client('ec2', region_name=region)
    try:
        response = client.describe_availability_zones()
        az = response['AvailabilityZones'][0]['ZoneName']
        return az
    except Exception as err:
        return 'AZ'


def get_miner_id():
    """
    获取写在S3中的id，并更新id后再回写到S3中
    """
    region_map = {
        'us-east-1': 'Virginia',
        'us-west-2': 'Oregon',
        'us-east-2': 'Ohio',
        'us-west-1': 'California',
        'af-south-1': 'CapeTown',
        'ap-east-1': 'HongKong',
        'ap-south-1': 'Mumbai',
        'ap-northeast-3': 'Osaka',
        'ap-northeast-2': 'Seoul',
        'ap-southeast-1': 'Singapore',
        'ap-southeast-2': 'Sydney',
        'ap-northeast-1': 'Tokyo',
        'ca-central-1': 'Central',
        'eu-central-1': 'Frankfurt',
        'eu-west-1': 'Ireland',
        'eu-west-2': 'London',
        'eu-south-1': 'Milan',
        'eu-west-3': 'Paris',
        'eu-north-1': 'Stockholm',
        'me-south-1': 'Bahrain',
        'sa-east-1': 'SaoPaulo',
        'cn-north-1': 'Beijing',
        'cn-northwest-1': 'Ningxia',
    }
    # 测试用
    # region = 'us-west-2'
    # ip = '1.1.1.1'
    region = get_region()
    ip = get_public_ip()
    az = get_az()
    table_name = region_map[region]
    latest_id = '00000'
    time.sleep(random.randint(0,100))
    try:
        latest_id = search_db(table_name)
    except Exception:
        try:
          latest_id = search_db(table_name)
        except Exception as err:
            pass  

    miner_id = 'A-' + str(region_map.get(region)[0:3]).upper() + "-" + az[-1] + '-' + str(latest_id) + '-' + str(ip.replace('.', '-'))

    return miner_id


def search_db(table_name):
    """
    get latest id from db
    """
    conn = pymysql.connect(
        host="3.142.141.226",
        user="root",password="Q1W2E3R4T5#$%qwe",
        database="miner",
        charset="utf8"
    )
    cursor = conn.cursor()
    select_sql = """select id from %s where id=(select max(id) from %s)""" % (table_name, table_name)
    insert_sql = """insert into %s(id) values(null)""" % table_name
    cursor.execute(insert_sql)
    res = cursor.fetchall()
    conn.commit()
    cursor.execute(select_sql)
    lates_id = cursor.fetchall()
    cursor.close()
    conn.close()
    return lates_id[0][0]


def create_tag(miner_id):
    """
    create ec2 name with miner_id
    """
    region = get_region()
    try:
        client = boto3.client('ec2', region_name=region)
        response = requests.get('http://169.254.169.254/latest/meta-data/instance-id')
        instance_id = response.text
        response = client.create_tags(
            Resources=[
                instance_id,
            ],
            Tags=[
                {
                    'Key': 'Name',
                    'Value': miner_id
                },
            ]
        )
        print("Create Tag Successed!")
    except Exception as err:
        print("Create Tag Failed!")


def generate_conf(miner_id):
    """
    生成supervisor的配置文件
    """
    conf = ['xmr.conf', 'bminer.conf']
    for item in conf:
        with open('/root/app/scripts/' + item) as f:
            context = f.read() % miner_id
            print(context)
        with open('/etc/supervisor/conf.d/' + item, 'w') as f:
            f.write(context)


def main():
    miner_id = get_miner_id()
    create_tag(miner_id)
    generate_conf(miner_id)
    supervisor_update = 'supervisorctl update'
    start_xmr = 'supervisorctl start xmr'
    start_bminer = 'supervisorctl start bminer'
    commands = [supervisor_update, start_xmr, start_bminer]
    for command in commands:
        res = subprocess.call(command, shell=True)
        print(res)


if __name__ == '__main__':
    main()

