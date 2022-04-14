#!/bin/bash
mysqldump -h devops-mysql-svc -uroot -pYHhg9-5.* cmdb > /tmp/c.sql
scp -P 53742 /tmp/c.sql root@121.43.41.139:/data/mysql_data/c.sql_`date '+%s'`