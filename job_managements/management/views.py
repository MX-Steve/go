from django.http import JsonResponse
from .tasks import add


def test_celery(request):
    add.delay(3, 9)
    return JsonResponse({
        "code": 200,
        "data": {},
        "msg": "测试 celery 功能"
    })