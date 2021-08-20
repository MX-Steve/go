from django.conf.urls import url
from django.contrib import admin
from django.urls import include
from ktz.check.check import check

urlpatterns = [
    url(r'^ktz/',include('ktz.urls')),
    url(r'isAlive',check.as_view())
]
