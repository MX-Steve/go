from django.conf.urls import url

from ktz.api.test import TestView

app_name = 'ktz'

urlpatterns = [
    url(r'^test/$', TestView.as_view(), name='test'),
]