import pymysql

conn = pymysql.connect(
    host="3.142.141.226",
    user="root",password="Q1W2E3R4T5#$%qwe",
    database="miner",
    charset="utf8"
)
cursor = conn.cursor()
select_sql = """select id from Oregon where id=(select max(id) from Oregon)"""
insert_sql = """insert into Oregon(id) values(null)"""
cursor.execute(select_sql)
lates_id = cursor.fetchall()
cursor.execute(insert_sql)
res = cursor.fetchall()
conn.commit()
cursor.close()
conn.close()
print(lates_id[0][0])
