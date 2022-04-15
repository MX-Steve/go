import ldap
import json
import pymysql
import uuid
import requests

db = pymysql.connect(host='192.168.174.115',
                     user='root',
                     password='YHhg9-5.*',
                     database='cmdb')
cursor = db.cursor()

connect = ldap.initialize("ldap://prod-ldap-svc.prod-ldap:389")
connect.set_option(ldap.OPT_REFERRALS, 0)
connect.simple_bind_s("cn=admin,dc=ecoplants,dc=com", "ecoplants!@#")
result = connect.search_s("ou=user,dc=ecoplants,dc=com", ldap.SCOPE_SUBTREE)
def is_all_chinese(strs):
    for _char in strs:
        if not '\u4e00' <= _char <= '\u9fa5':
            return False
    return True
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
        cursor.execute("select * from users_user where email='%s'" %
                       (str(mail[0], encoding="utf-8")))
        data = cursor.fetchone()
        if not data:
            sn = str(sn[0], encoding="utf-8")
            pwd = str(pwd[0], encoding="utf-8")
            mail = str(mail[0], encoding="utf-8")
            mobile = str(mobile[0], encoding="utf-8")
            headers = {
                "Authorization":
                "Bearer t-b2e02f8c426580e992269ce829fd7e6e7011b7b6",
                "Content-Type": "application/json; charset=utf-8"
            }
            res = requests.get(
                "https://open.feishu.cn/open-apis/user/v1/batch_get_id?emails=%s"
                % (mail),
                headers=headers)
            data = res.json()
            data = data["data"]
            if "email_users" in data:
                print(data)
                if mail in data["email_users"]:
                    user_id = data["email_users"][mail][0]["user_id"]
                    print(mail, user_id)
                    sql = "insert into users_user values('%s','%s',0,'%s','','','%s', 1, 1, '%s','%s','%s','','','%s')" % (
                        pwd,
                        '2022-01-01 00:00:00.000000000',
                        sn,
                        mail,
                        '2022-01-01 00:00:00.000000000',
                        str(uuid.uuid4()).replace("-","", 4),
                        '[2,]',
                        user_id
                    )
                    cursor.execute(sql)
                    db.commit()
        else:
            sn = str(sn[0], encoding="utf-8")
            pwd = str(pwd[0], encoding="utf-8")
            mail = str(mail[0], encoding="utf-8")
            mobile = str(mobile[0], encoding="utf-8")
            if is_all_chinese(sn):
                first_name = sn[1:]
                last_name = sn[0]
                username = mail.split("@")[0]
                sql = "update users_user set first_name='%s', last_name='%s', username='%s' where email='%s'"%(first_name, last_name, username, mail)
                cursor.execute(sql)
                db.commit()
db.close()