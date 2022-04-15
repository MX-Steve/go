import ldap
import json
import pymysql
import uuid
import requests
import time

db = pymysql.connect(host='192.168.174.115',
                     user='root',
                     password='YHhg9-5.*',
                     database='cmdb')
cursor = db.cursor()
sql = "select email,username,user_id from users_user;"
cursor.execute(sql)
data = cursor.fetchall()
for user in data:
    if user[2] is None and user[0] != "" and user[1] not in [
            "admin", 'review'
    ]:
        headers = {
            "Authorization":
            "Bearer t-ced5000ac8e2cb75a065655e3669c7a6c3f5e166",
            "Content-Type": "application/json; charset=utf-8"
        }
        res = requests.get(
            "https://open.feishu.cn/open-apis/user/v1/batch_get_id?emails=%s" %
            (user[0]),
            headers=headers)
        data = res.json()
        data = data["data"]
        if "email_users" in data:
            print(data)
            if user[0] in data["email_users"]:
                user_id = data["email_users"][user[0]][0]["user_id"]
                print(user[0], user_id)
                # sql = "update users_user set user_id='%s' where email='%s'"%(user_id,user[0])
                # cursor.execute(sql)
                # db.commit()
