from django.http import JsonResponse
from celery import shared_task
import time
import requests
from management.models import *

@shared_task
def add(x, y):
    pts = PeriodicTask.objects.all()
    print(pts.count())
    time.sleep(2)
    return x + y

@shared_task
def crontab_job_add():
    return JsonResponse({"code": 200, "data": {}, "msg": "success"})

@shared_task
def one_job_add():
    return JsonResponse({"code": 200, "data": {}, "msg": "success"})

@shared_task
def interval_job_add():
    return JsonResponse({"code": 200, "data": {}, "msg": "success"})

@shared_task
def clock_job_add():
    return JsonResponse({"code": 200, "data": {}, "msg": "success"})
