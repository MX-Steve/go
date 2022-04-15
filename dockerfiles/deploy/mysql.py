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

connect = ldap.initialize("ldap://prod-ldap-svc.prod-ldap:389")
connect.set_option(ldap.OPT_REFERRALS, 0)
connect.simple_bind_s("cn=admin,dc=ecoplants,dc=com", "ecoplants!@#")
result = connect.search_s("ou=user,dc=ecoplants,dc=com", ldap.SCOPE_SUBTREE)
for item in result:
    dc = item[0]
    obj = item[1]
    keys = obj.keys()
    if "sn" in keys and \
        "userPassword" in keys and \
            "mail" in keys and \
                "mobile" in keys :
        sn = obj["sn"]
        pwd = obj["userPassword"]
        mail = obj["mail"]
        mobile = obj["mobile"]
        # print(dc)
        # print(str(sn[0], encoding="utf-8"))
        # print(str(pwd[0], encoding="utf-8"))
        # print(str(mail[0], encoding="utf-8"))
        # print(str(mobile[0], encoding="utf-8"))
        # print("------")
        cursor.execute("select * from users_user where email='%s'" % (str(mail[0], encoding="utf-8")))
        data = cursor.fetchone()
        if data:
            sn = str(sn[0], encoding="utf-8")
            pwd = str(pwd[0], encoding="utf-8")
            mail = str(mail[0], encoding="utf-8")
            mobile = str(mobile[0], encoding="utf-8")
            headers = {
                "Authorization":
                "Bearer t-b2e02f8c426580e992269ce829fd7e6e7011b7b6",
                "Content-Type": "application/json; charset=utf-8"
            }
            url = "https://open.feishu.cn/open-apis/user/v1/batch_get_id?emails=" + mail
            res = requests.get(url, headers=headers)
            data = res.json()
            print(data)
            if "email_users" in data:
                user_id = data["email_users"][mail][0]["user_id"]
                print(user_id)
                sql = "update users_user set user_id='%s' where email='%s'"%(user_id,mail)
                print(sql)
                # cursor.execute(sql)
                # db.commit()
                time.sleep(1)
        else:
            print(str(sn[0], encoding="utf-8"), "没有入库")
            mail = str(mail[0], encoding="utf-8")
            headers = {
                "Authorization":
                "Bearer t-b2e02f8c426580e992269ce829fd7e6e7011b7b6",
                "Content-Type": "application/json; charset=utf-8"
            }
            url = "https://open.feishu.cn/open-apis/user/v1/batch_get_id?emails=" + mail
            res = requests.get(url, headers=headers)
            data = res.json()
            print(data)
            print("___________")
            
db.close()

# de9e59gb