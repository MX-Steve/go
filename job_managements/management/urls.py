from django.urls import path
from .views import test_celery
from management.apis import test, interval

urlpatterns = [
    path('v1/test', test_celery, name="tes"),
    path('v1/intervals', interval.IntervalScheduleView.as_view(), name="intervals"),
    path('v1/tasks', interval.PeriodicTaskView.as_view(), name="tasks"),
    path('v1/results', interval.ResultsTasksView.as_view(), name="results"),
    path('v1/crontab', interval.CrontabScheduleView.as_view(), name="crontab")
]