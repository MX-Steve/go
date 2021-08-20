import os
import sys
import datetime
import configparser as ConfigParser


BASE_DIR = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))


SECRET_KEY = 'qe187p7jy6w%hrxdly6t=rq%2izaq@owxtefa=sc(gpnialsa!'
cf = ConfigParser.ConfigParser()
cf.read(BASE_DIR + "/ktzIn/config.ini")

DEBUG = False 

ALLOWED_HOSTS = ["*"]

INSTALLED_APPS = [
    'django.contrib.admin',
    'django.contrib.auth',
    'django.contrib.contenttypes',
    'django.contrib.sessions',
    'django.contrib.messages',
    'django.contrib.staticfiles',
    'rest_framework',
    'rest_framework.authtoken',
    'drf_multiple_model',
    'corsheaders',
    'ktz',
]
AUTH_USER_MODEL = 'ktz.Accounts'

MIDDLEWARE = [
    'corsheaders.middleware.CorsMiddleware',
    'django.middleware.security.SecurityMiddleware',
    'django.contrib.sessions.middleware.SessionMiddleware',
    'django.middleware.common.CommonMiddleware',
    'django.middleware.csrf.CsrfViewMiddleware',
    'django.contrib.auth.middleware.AuthenticationMiddleware',
    'django.contrib.messages.middleware.MessageMiddleware',
    'django.middleware.clickjacking.XFrameOptionsMiddleware',
    'ktz.utils.md1.MD1'
]

ROOT_URLCONF = 'ktzIn.urls'

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
    'DEFAULT_PERMISSION_CLASSES': (
        'rest_framework.permissions.IsAuthenticated',
    ),
    'DEFAULT_AUTHENTICATION_CLASSES': (
        'rest_framework_jwt.authentication.JSONWebTokenAuthentication',
    ),
}

JWT_AUTH = {
    'JWT_EXPIRATION_DELTA': datetime.timedelta(days=5),
    'JWT_AUTH_HEADER_PREFIX': 'JWT'
}

WSGI_APPLICATION = 'ktzIn.wsgi.application'


DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.mysql',
        'NAME': cf.get("db", "name"),
        'USER': cf.get("db", "user"),
        'PASSWORD': cf.get("db", "pass"),
        'HOST': cf.get("db", "host"),
        'PORT': cf.get("db", "port"),
        'OPTIONS': {
            "init_command": "SET foreign_key_checks = 0;",
        }
    }
}


AUTH_PASSWORD_VALIDATORS = [
    {
        'NAME': 'django.contrib.auth.password_validation.UserAttributeSimilarityValidator',
    },
    {
        'NAME': 'django.contrib.auth.password_validation.MinimumLengthValidator',
    },
    {
        'NAME': 'django.contrib.auth.password_validation.CommonPasswordValidator',
    },
    {
        'NAME': 'django.contrib.auth.password_validation.NumericPasswordValidator',
    },
]

LANGUAGE_CODE = 'zh-hans'

TIME_ZONE = 'Asia/Shanghai'

USE_I18N = True

USE_L10N = True

USE_TZ = False

STATIC_URL = '/static/'
STATIC_ROOT = os.path.join(BASE_DIR, 'static')

#BASIC_URL = cf.get("apis", "BASIC_URL")
AUTH_URL = cf.get('apis', "AUTH_URL")


LOGGING = {
    'version': 1,
    'disable_existing_loggers': True,
    'formatters': {
        'standard': {
            'format': '%(asctime)s [%(threadName)s:%(thread)d] [%(name)s:%(lineno)d] [%(levelname)s]- %(message)s'
        },
    },
    'filters': {
    },
    'handlers': {
        'mail_admins': {
            'level': 'ERROR',
            'class': 'django.utils.log.AdminEmailHandler',
            'include_html': True,
        },
        'default': {
            'level':'DEBUG',
            'class':'logging.handlers.RotatingFileHandler',
            'filename': os.path.join(BASE_DIR+'/logs/','all.log'), 
            'maxBytes': 1024*1024*5, 
            'backupCount': 5,
            'formatter':'standard',
        },
        'console':{
            'level': 'DEBUG',
            'class': 'logging.StreamHandler',
            'formatter': 'standard'
        },
        'request_handler': {
            'level':'DEBUG',
            'class':'logging.handlers.RotatingFileHandler',
            'filename': os.path.join(BASE_DIR+'/logs/','script.log'), 
            'maxBytes': 1024*1024*5, 
            'backupCount': 5,
            'formatter':'standard',
        },
        'scprits_handler': {
            'level':'DEBUG',
            'class':'logging.handlers.RotatingFileHandler',
            'filename': os.path.join(BASE_DIR+'/logs/','script.log'), 
            'maxBytes': 1024*1024*5, 
            'backupCount': 5,
            'formatter':'standard',
        },
        'out_handler': {
                    'level':'DEBUG',
                    'class':'logging.handlers.RotatingFileHandler',
                    'filename': os.path.join(BASE_DIR+'/logs/','out.log'), 
                    'maxBytes': 1024*1024*5, 
                    'backupCount': 5,
                    'formatter':'standard',
        },
        'info': {
                    'level':'DEBUG',
                    'class':'logging.handlers.RotatingFileHandler',
                    'filename': os.path.join(BASE_DIR+'/logs/','out.log'), 
                    'maxBytes': 1024*1024*5,
                    'backupCount': 5,
                    'formatter':'standard',
        },
        'error': {
                    'level':'DEBUG',
                    'class':'logging.handlers.RotatingFileHandler',
                    'filename': os.path.join(BASE_DIR+'/logs/','script.log'), 
                    'maxBytes': 1024*1024*5, 
                    'backupCount': 5,
                    'formatter':'standard',
        }
    },
    'loggers': {
        'django': {
            'handlers': ['default','console'],
            'level': 'DEBUG',
            'propagate': False
        },
        'ttool.app':{
            'handlers': ['error', 'info', 'console', 'default'],
            'level': 'INFO',
            'propagate': True
        },
        'django.request': {
            'handlers': ['request_handler'],
            'level': 'DEBUG',
            'propagate': False
        },
        'scripts': { 
            'handlers': ['scprits_handler'],
            'level': 'INFO',
            'propagate': False
        },
        'Yearning.core.views': {
            'handlers': ['out_handler'],
            'level': 'DEBUG',
            'propagate': True
        },
        'out.app':{
            'handlers': ['out_handler','console'],
            'level': 'DEBUG',
            'propagate': True
        }
    }
}
