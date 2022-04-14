# pylint: disable=wildcard-import,unused-wildcard-import
from os import environ
import pymysql
from djentry.settings.common import *

pymysql.install_as_MySQLdb()

DEBUG = True
# SECURITY WARNING: keep the secret key used in production secret!
SECRET_KEY = environ.get('DJANGO_SECRET_KEY')

ALLOWED_HOSTS = ["*"]

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
HELLO_APP_ID = environ.get("HELLO_APP_ID")
HELLO_APP_SECRET = environ.get("HELLO_APP_SECRET")
ITSM_BACK_PORT = 8082
JENKINS_HOST = environ.get("JENKINS_HOST")
JENKINS_USER = environ.get("JENKINS_USER")
JENKINS_PASSWD = environ.get("JENKINS_PASSWD")
DEPLOY_CHAT_ID = environ.get("DEPLOY_CHAT_ID")
APPROVAL_INSTANCES_DETAIL = environ.get("APPROVAL_INSTANCES_DETAIL") 
