# 前端启动
source /etc/profile
export LD_LIBRARY_PATH="/usr/local/lib"
/usr/local/nginx/sbin/nginx
# 后端启动
python3 /app/vssh/backend/manage.py migrate
python3 /app/vssh/backend/manage.py runserver 0.0.0.0:8000 &
while true
do
  sleep 60
done
