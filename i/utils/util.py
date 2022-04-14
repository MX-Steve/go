# -*- coding:utf-8 -*-
# pylint: disable=duplicate-code
'''

Some tool sets

'''
import time
import datetime
import shortuuid


def now():
    '''
    datetime
    '''
    return time.strftime('%Y-%m-%d %H:%M:%S', time.localtime(time.time()))


def time_trans(utc_time: str):  # UTC时间转化本地时间
    if utc_time:
        utc_len = len(utc_time.split(':'))
        if utc_len == 2:
            UTC_FORMAT = "%Y-%m-%dT%H:%MZ"
        elif utc_len == 3:
            if len(utc_time.split('.')) == 2:
                UTC_FORMAT = "%Y-%m-%dT%H:%M:%S.%fZ"
            else:
                UTC_FORMAT = "%Y-%m-%dT%H:%M:%SZ"
        utcTime = datetime.datetime.strptime(utc_time, UTC_FORMAT)
        trans = utcTime + datetime.timedelta(hours=8)
        return str(trans)


def gen_shortuuid():
    uuid = shortuuid.ShortUUID().random(length=20)
    return uuid
