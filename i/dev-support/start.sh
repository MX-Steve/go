#!/bin/bash
cd /app/itsm_backend
/app/itsm_backend/build/python/bin/python /app/itsm_backend/manage.py runserver 0.0.0.0:8000 &
while true
do
    sleep 60
done