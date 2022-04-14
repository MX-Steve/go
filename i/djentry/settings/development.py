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
        'HOST': "121.43.41.139",
        'PORT': 3306,
        'USER': "root",
        'PASSWORD': "YHhg9-5.*",
    }
}
# HELLO_APP_ID = "cli_a1543517eaf9100b"
# HELLO_APP_SECRET = "hjaC6YU7LF25DQNY3aNkkcjfu6PKD7xr"
HELLO_APP_ID = "cli_a15488ef3829d00e"
HELLO_APP_SECRET = "8ciJX1U4foZzN9wu9UP5ucVO0j1cG6Qf"
ITSM_BACK_PORT = 8000
JENKINS_HOST = "http://121.43.41.139:8080/"
JENKINS_USER = "admin"
JENKINS_PASSWD = "root123"
DEPLOY_CHAT_ID = "oc_a9c92c257c0811caf8429df6a8fb596d"
APPROVAL_INSTANCES_DETAIL = "http://localhost:9527/approvals/use/detail?instance_id="
