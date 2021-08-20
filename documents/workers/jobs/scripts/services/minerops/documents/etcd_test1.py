import requests
import json
from base64 import b64encode, b64decode

ETCD_URL = "http://192.168.226.130:2379/v3"

result = {
    "header":
        {
            "cluster_id": "14841639068965178418",
            "member_id": "10276657743932975437",
            "revision": "7",
            "raft_term": "3"
        },
    "kvs": [
        {
            "key": "Zm9v",
            "create_revision": "3",
            "mod_revision": "7",
            "version": "4",
            "value": "YmFy"
        }
    ],
    "count": "1"
}


def GetData(key):
    res = requests.post("%s/kv/range" % (ETCD_URL), data=json.dumps({"key": key}))
    if res.status_code == 200:
        result = res.json()
        return result['kvs'][0]['value']
    else:
        print("Request 接口请求失败，请检查")
        return 1


def SetData(key, value):
    res = requests.post("%s/kv/put" % (ETCD_URL), data=json.dumps({"key": key, "value": value}))
    print(res.text)


if __name__ == '__main__':
    key = "Zm9v"
    # value = GetData(key)
    # print(value)
    SetData(key, "AsmFy")
