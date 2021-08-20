# -*- coding:utf-8 -*-
import time, datetime
from dateutil.parser import parse
import logging
import os, sys
import traceback
from logging.handlers import TimedRotatingFileHandler
from django.conf import settings

BASE_DIR=settings.BASE_DIR
# 初始化logger
logger = logging.getLogger("DemoIn.app")

from django.http import HttpResponse


def loggerInFile():
    def decorator(func):
        def inner(*args, **kwargs):  # 1
            logFilePath = BASE_DIR+"/logs/err.log"  # 日志按日期滚动，保留5天
            logger = logging.getLogger()
            logger.setLevel(logging.INFO)
            handler = TimedRotatingFileHandler(logFilePath,
                                               when="d",
                                               interval=1,
                                               backupCount=5)
            formatter = logging.Formatter('%(asctime)s  - %(levelname)s - %(message)s')
            handler.setFormatter(formatter)
            logger.addHandler(handler)
            try:
                return func(*args, **kwargs)
            except:
                logger.error(traceback.format_exc())
                return HttpResponse(status=500)
        return inner
    return decorator


def now():
    '''
    datetime
    '''
    now = time.strftime('%Y-%m-%d %H:%M:%S', time.localtime(time.time()))
    return now


def utcTolocal(utcdate):
    try:
        date_tmp = parse(utcdate)
        local_date = (date_tmp + datetime.timedelta(hours=8)).strftime('%Y-%m-%d')
        return local_date
    except:
        err = '%s [%s] happend on %s line at %s' % (
        sys.exc_info()[0], sys.exc_info()[1], sys.exc_info()[2].tb_lineno, os.path.basename(__file__))
        logger.error(err)


def ser(_obj):
    '''
    orm.raw 序列化
    '''
    _list = []
    _get = []
    for i in _obj:
        _list.append(i.__dict__)

    for i in _list:
        del i['_state']
        _get.append(i)
    return _get
