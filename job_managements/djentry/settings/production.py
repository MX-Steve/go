# pylint: disable=wildcard-import,unused-wildcard-import
from os import environ
import pymysql
from djentry.settings.common import *

pymysql.install_as_MySQLdb()

# SECURITY WARNING: keep the secret key used in production secret!
SECRET_KEY = environ.get('DJANGO_SECRET_KEY')

ALLOWED_HOSTS = ["*"]
DEBUG = True

# Database
# https://docs.djangoproject.com/en/3.2/ref/settings/#databases

DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.mysql',
        'NAME': environ.get("CLOVER_MYSQL_DB"),
        'USER': environ.get("CLOVER_MYSQL_USER"),
        'PASSWORD': environ.get("CLOVER_MYSQL_PASSWORD"),
        'HOST': environ.get("CLOVER_MYSQL_HOST"),
        'PORT': environ.get("CLOVER_MYSQL_PORT") or '3306'
    }
}
REDIS_PORT = environ.get("REDIS_PORT")
REDIS_PASS = environ.get("REDIS_PASS")
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
