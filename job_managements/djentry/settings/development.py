# pylint: disable=wildcard-import,unused-wildcard-import
from os import environ
from pathlib import Path
from djentry.settings.common import *

# Build paths inside the project like this: BASE_DIR / 'subdir'.
BASE_DIR = Path(__file__).resolve().parent.parent

# SECURITY WARNING: keep the secret key used in production secret!
SECRET_KEY = environ.get('DJANGO_SECRET_KEY')

# SECURITY WARNING: don't run with debug turned on in production!
DEBUG = True

ALLOWED_HOSTS = ['*']

DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.mysql',
        'NAME': 'test',
        'HOST': "mysql8",
        'PORT': 3306,
        'USER': "root",
        'PASSWORD': "123456",
    }
}

REDIS_PORT = 6379
REDIS_PASS = "123456"

CELERY_BROKER_URL = "redis://:%s@127.0.0.1:%s/0" % (REDIS_PASS, REDIS_PORT)
ASGI_APPLICATION = 'djentry.asgi.application'
CHANNEL_LAYERS = {
    'default': {
        'BACKEND': 'channels_redis.core.RedisChannelLayer',
        'CONFIG': {
            "hosts": ["redis://:%s@127.0.0.1:%s/1" % (REDIS_PASS, REDIS_PORT)],
            "symmetric_encryption_keys": [SECRET_KEY],
        },
    },
}
