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
                "Bearer t-ced5000ac8e2cb75a065655e3669c7a6c3f5e166",
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
db.close()

# searchFilter = "(cn=lihan)"
# result = connect.search_s("ou=user,dc=ecoplants,dc=com", ldap.SCOPE_SUBTREE,
#                           searchFilter, None)
# print(result)
# try:
#     user_dn = result[0][0]
#     connect.simple_bind_s(user_dn, "123456789")
#     print("ok")
# except ldap.LDAPError as e:
#     print(e)

