from django.conf.urls import url
from django.contrib import admin
from django.urls import include
from demo.check.check import check

urlpatterns = [
    url(r'^demo/',include('demo.urls')),
    url(r'isAlive',check.as_view())
]
