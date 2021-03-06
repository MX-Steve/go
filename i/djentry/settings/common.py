"""
Django settings for this project.

Generated by 'django-admin startproject' using Django 3.2.7.

For more information on this file, see
https://docs.djangoproject.com/en/3.2/topics/settings/

For the full list of settings and their values, see
https://docs.djangoproject.com/en/3.2/ref/settings/
"""

import os
from pathlib import Path
from os import environ
import datetime

# Build paths inside the project like this: BASE_DIR / 'subdir'.
BASE_DIR = Path(__file__).resolve().parent.parent.parent

# Application definition

INSTALLED_APPS = [
    'django.contrib.admin', 'django.contrib.auth',
    'django.contrib.contenttypes', 'django.contrib.sessions',
    'django.contrib.messages', 'django.contrib.staticfiles', 'rest_framework',
    'rest_framework.authtoken', 'approval'
]

MIDDLEWARE = [
    'django.middleware.security.SecurityMiddleware',
    'django.contrib.sessions.middleware.SessionMiddleware',
    'django.middleware.common.CommonMiddleware',
    'django.middleware.csrf.CsrfViewMiddleware',
    'django.contrib.auth.middleware.AuthenticationMiddleware',
    'django.contrib.messages.middleware.MessageMiddleware',
    'django.middleware.clickjacking.XFrameOptionsMiddleware',
]

ROOT_URLCONF = 'djentry.urls'

TEMPLATES = [
    {
        'BACKEND': 'django.template.backends.django.DjangoTemplates',
        'DIRS': [],
        'APP_DIRS': True,
        'OPTIONS': {
            'context_processors': [
                'django.template.context_processors.debug',
                'django.template.context_processors.request',
                'django.contrib.auth.context_processors.auth',
                'django.contrib.messages.context_processors.messages',
            ],
        },
    },
]

WSGI_APPLICATION = 'djentry.wsgi.application'

# Password validation
# https://docs.djangoproject.com/en/3.2/ref/settings/#auth-password-validators

AUTH_PASSWORD_VALIDATORS = [
    {
        'NAME':
        'django.contrib.auth.password_validation.UserAttributeSimilarityValidator',
    },
    {
        'NAME':
        'django.contrib.auth.password_validation.MinimumLengthValidator',
    },
    {
        'NAME':
        'django.contrib.auth.password_validation.CommonPasswordValidator',
    },
    {
        'NAME':
        'django.contrib.auth.password_validation.NumericPasswordValidator',
    },
]

# Internationalization
# https://docs.djangoproject.com/en/3.2/topics/i18n/

LANGUAGE_CODE = 'en-us'

TIME_ZONE = 'Asia/Shanghai'

USE_I18N = True

USE_L10N = True

USE_TZ = True

# Static files (CSS, JavaScript, Images)
# https://docs.djangoproject.com/en/3.2/howto/static-files/

STATIC_URL = '/static/'

STATICFILES_DIRS = [
    os.path.join(BASE_DIR, "statics"),
]

# Default primary key field type
# https://docs.djangoproject.com/en/3.2/ref/settings/#default-auto-field

DEFAULT_AUTO_FIELD = 'django.db.models.BigAutoField'
AUTH_USER_MODEL = "approval.User"

CORS_ALLOW_CREDENTIALS = True
CORS_ORIGIN_ALLOW_ALL = True
CORS_ORIGIN_WHITELIST = ()

CORS_ALLOW_METHODS = (
    'DELETE',
    'GET',
    'OPTIONS',
    'PATCH',
    'POST',
    'PUT',
    'VIEW',
)

CORS_ALLOW_HEADERS = (
    'accept',
    'accept-encoding',
    'authorization',
    'content-type',
    'dnt',
    'origin',
    'user-agent',
    'x-csrftoken',
    'X-Token',
    'x-requested-with',
)
REST_FRAMEWORK = {
    'DEFAULT_PERMISSION_CLASSES':
    ('rest_framework.permissions.IsAuthenticated', ),
    'DEFAULT_AUTHENTICATION_CLASSES':
    ('rest_framework_jwt.authentication.JSONWebTokenAuthentication', ),
}

JWT_AUTH = {
    'JWT_EXPIRATION_DELTA': datetime.timedelta(days=1),
    'JWT_AUTH_HEADER_PREFIX': 'JWT'
}
DJENTRY_STATISTIC_BACKEND_ADDR = environ.get(
    "DJENTRY_STATISTIC_BACKEND_ADDR") or '121.43.41.139:16221'
KAFKA_ADDR = environ.get("KAFKA_ADDR") or "121.43.41.139:8090"
ADMINS = (('admin_name', 'mx_steve@163.com'), )
MANAGERS = ADMINS
ENABLE_LDAP = environ.get("ENABLE")
if ENABLE_LDAP == "True":
    import ldap
    from django_auth_ldap.config import LDAPSearch, GroupOfNamesType
    AUTHENTICATION_BACKENDS = (
        'django_auth_ldap.backend.LDAPBackend',  # ??????????????????LDAP????????????????????????????????????????????????????????????
        'django.contrib.auth.backends.ModelBackend',  # django????????????????????????????????????????????????????????????????????????2????????????
    )
    AUTH_LDAP_SERVER_URI = environ.get("AUTH_LDAP_SERVER_URI")
    AUTH_LDAP_BIND_DN = environ.get("AUTH_LDAP_BIND_DN")
    AUTH_LDAP_BIND_PASSWORD = environ.get("AUTH_LDAP_BIND_PASSWORD")
    AUTH_LDAP_USER_SEARCH = LDAPSearch(environ.get("USER_SEARCH"),
                                       ldap.SCOPE_SUBTREE, "(sn=%(user)s)")
    # AUTH_LDAP_CONNECTION_OPTIONS = {
    #     ldap.OPT_DEBUG_LEVEL: 1,
    #     ldap.OPT_REFERRALS: 0,
    # }
    AUTH_LDAP_FIND_GROUP_PERMS = True
    AUTH_LDAP_ALWAYS_UPDATE_USER = True  # ???????????????ldap??????????????????
    AUTH_LDAP_USER_ATTR_MAP = {  # key???archery.sql_users????????????value???ldap?????????????????????????????????
        "username": "sAMAccountName",
        # "first_name": "displayname",
        "first_name": "givenName",
        "last_name": "sn",
        "email": "mail"
    }
