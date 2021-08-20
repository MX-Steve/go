from django.conf.urls import url

from demo.api.test import TestView

app_name = 'demo'

urlpatterns = [
    url(r'^test/$', TestView.as_view(), name='test'),
]