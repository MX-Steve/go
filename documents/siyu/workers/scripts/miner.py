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
    instance-test  us-central1-a  n1-standard-1 True/(space) 10.128.0.15  34.133.224.201  RUNNING
    """
    cmd = "gcloud compute instances list | grep " + ip
    output = subprocess.check_output(cmd, shell=True)
    res = output.decode()
    print(res)
    info = re.split(r"[ ]+", res)
    return info

def gcp_get_region(ip):
    """
    获取GCP instance的region
    gcp region include AvailabilityZones
    """
    info = gcp_get_info_by_ip(ip)
    if(len(info) >= 2):
        return info[1]
    else:
        return None


def gcp_get_old_miner_id(ip):
    """
    get original instance_name
    """
    info = gcp_get_info_by_ip(ip)
    if(len(info) >= 1):
        return info[0]
    else:
        return None

def gcp_get_miner_id():
    """
    get id from db
    """
    region_map = {
        'asia-east1': 'Taiwan',
        'asia-east2': 'HongKong',
        'asia-northeast1': 'Tokyo',
        'asia-northeast2': 'Osaka',
        'asia-northeast3': 'Seoul',
        'asia-south1': 'Mumbai',
        'asia-south2': 'Delhi',
        'asia-southeast1': 'Singapore',
        'asia-southeast2': 'Jakarta',
        'australia-southeast1': 'Sydney',
        'australia-southeast2': 'Melbourne',
        'europe-central2': 'Warsaw',
        'europe-north1': 'Hamina',
        'europe-west1': 'Belgium',
        'europe-west2': 'London',
        'europe-west3': 'Frankfurt',
        'europe-west4': 'Netherlands',
        'europe-west6': 'Zurich',
        'northamerica-northeast1': 'Montreal',
        'southamerica-east1': 'Osasco',
        'us-central1': 'Iowa',
        'us-east1': 'Carolina',
        'us-east4': 'Virginia',
        'us-west1': 'Oregon',
        'us-west2': 'California',
        'us-west3': 'Utah',
        'us-west4': 'Nevada',
    }
    ip = gcp_get_public_ip()
    latest_id = '00000'                 #  means error here
    old_miner_id = gcp_get_old_miner_id(ip)

    ori_region = gcp_get_region(ip)                  #  us-central1-a
    az = ori_region[-1]                              #  a
    region = ori_region[0:len(ori_region)-2]         #  us-central1
    table_name = region_map.get(region, 'unset')     #  us-central1: Iowa, not found: unset

    time.sleep(random.randint(0,100))
    try:
        latest_id = search_db(table_name)
    except Exception:
        try:
          latest_id = search_db(table_name)
        except Exception as err:
            pass  
    # A: AWS global; C: AWS China region; GPC: Google Cloud Platform
    miner_id = 'GPC-' + str(region_map.get(region)[0:3]).upper() + "-" + az[-1] + '-' + str(latest_id) + '-' + str(ip.replace('.', '-'))
    # GCP: Must be a match of regex '(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?)'
    miner_id = miner_id.lower()
    
    return old_miner_id, miner_id, ori_region

def gcp_create_tag(new_miner_id, old_miner_id, zone):
    """
    create ec2 name with miner_id
    GCP: gcloud beta compute instances set-name INSTANCE_NAME --new-name=NEW_INSTANCE_NAME
    """
    create_tag_cmd = "gcloud compute instances add-labels " + old_miner_id + " --labels=name=" + new_miner_id + " --zone=" + zone
    print(create_tag_cmd)
    #gcloud compute instances list --filter='labels.k1:v2' #query
    os.system(create_tag_cmd)

    # - This instance is not stopped and cannot be renamed. 只能在停止时改名
    #set_name_cmd = "gcloud beta compute instances set-name " + old_miner_id + " --new-name=" + new_miner_id + " --zone=" + zone
    #print(set_name_cmd)
    #os.system(set_name_cmd)


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
    get id from db
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
    latest_id = '00000'  # means error here
    time.sleep(random.randint(0,100))
    try:
        latest_id = search_db(table_name)
    except Exception:
        try:
          latest_id = search_db(table_name)
        except Exception as err:
            pass  
    # A: AWS global; C: AWS China region; GPC: Google Cloud Platform
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
    is_gcp = True
    miner_id = None
    if(is_gcp):         #GCP
        old_miner_id = None
        zone = None
        old_miner_id, miner_id, zone = gcp_get_miner_id()
        gcp_create_tag(miner_id, old_miner_id, zone)
        generate_conf(miner_id)

    else:               #AWS
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