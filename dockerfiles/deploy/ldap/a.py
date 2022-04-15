import ldap
connect = ldap.initialize("ldap://prod-ldap-svc.prod-ldap:389")
connect.set_option(ldap.OPT_REFERRALS, 0)
connect.simple_bind_s("cn=admin,dc=ecoplants,dc=com", "ecoplants!@#")

searchFilter = "(cn=lihan)"
result = connect.search_s("ou=user,dc=ecoplants,dc=com", ldap.SCOPE_SUBTREE,
                          searchFilter, None)
print(result)
try:
    user_dn = result[0][0]
    connect.simple_bind_s(user_dn, "xwzsl-18.*")
    print("ok")
except ldap.LDAPError as e:
    print(e)