from enum import unique
from django.db import models
import uuid

from django.db.models.expressions import F

"""
以下表属于 mysql
"""
class User(models.Model):
    """
    User table:
    uuid:       User uuid
    email:      User email
    password:   User password
    phone:      User phone
    """
    uuid = models.UUIDField(primary_key = True , auto_created = True , 
        default = uuid.uuid4, editable = False, verbose_name = "User uuid")
    email = models.CharField(max_length=100, null=True, verbose_name="User email")
    password = models.TextField(max_length=65535, null=False, verbose_name="User password")
    phone = models.CharField(max_length=100, null=True, verbose_name="User phone")
    
    class Meta:
        verbose_name = 'User'
        db_table = 'user'
        verbose_name_plural = verbose_name


class AddressWhiteList(models.Model):
    """
    AddressWhiteList table:
    uuid:           AddressWhiteList uuid
    user_uuid:      User.uuid
    address:        User pocket address
    currency:       User currency type
    ext_data:       extral information
    gmt_create:     The time When this address was created 
    gmt_update:     The time When this address was updated   
    """
    uuid = models.UUIDField(primary_key = True , auto_created = True , 
        default = uuid.uuid4, editable = False, verbose_name = "AddressWhiteList uuid")
    user_uuid = models.ForeignKey(User, to_field="uuid",on_delete=models.CASCADE)
    address = models.CharField(max_length=100, null=False, verbose_name="pocket address")
    currency = models.CharField(max_length=100, null=False, verbose_name="currency type")
    ext_data = models.CharField(max_length=100, null=True, verbose_name="extral information")
    gmt_create = models.IntegerField(null=True, verbose_name="create time")
    gmt_update = models.IntegerField(null=True, verbose_name="update time")

    class Meta:
        verbose_name = 'AddressWhiteList'
        db_table = 'address_white_list'
        verbose_name_plural = verbose_name

class FollowWallet(models.Model):
    """
    FollowWallet table:
    uuid:               FollowWallet uuid
    user_uuid:          User.uuid
    wallet:             user pocket address
    currency:           currency type
    name:               customize name
    default_wallet:     TODO, default 0
    """ 
    uuid = models.UUIDField(primary_key = True , auto_created = True , 
        default = uuid.uuid4, editable = False, verbose_name = "FollowWallet uuid")
    user_uuid = models.ForeignKey(User, to_field="uuid",on_delete=models.CASCADE)
    wallet = models.CharField(max_length=100, null=False, verbose_name="pocket address")
    currency = models.CharField(max_length=8, null=False, verbose_name="currency type")
    name = models.CharField(max_length=8, null=False, verbose_name="customize name")
    default_wallet = models.SmallIntegerField(verbose_name="todo")

    class Meta:
        verbose_name = 'FollowWallet'
        db_table = 'follow_wallet'
        verbose_name_plural = verbose_name


class Subuser(models.Model):
    """
    Subuser table:
    uuid:                   Subuser uuid
    user_uuid:              User.uuid
    name:                   subuser name
    memo:                   subuser information
    img_url:                subuser img
    is_open_account:        is open account
    account_stat:           TODO, default 0
    is_del:                 is del
    """
    uuid = models.UUIDField(primary_key = True , auto_created = True , 
        default = uuid.uuid4, editable = False, verbose_name = "FollowWallet uuid")
    user_uuid = models.ForeignKey(User, to_field="uuid",on_delete=models.CASCADE)
    name = models.CharField(max_length=100, null=False, unique=True, verbose_name="subuser name")
    memo = models.CharField(max_length=100, null=False, verbose_name="subuser information")
    img_url = models.CharField(max_length=100, null=False, verbose_name="subuser img")
    is_open_account = models.SmallIntegerField(verbose_name="is open account")
    account_stat = models.IntegerField(null=True, verbose_name="account_stat", default=0)
    is_del = models.SmallIntegerField(null=True, verbose_name="is_del")

    class Meta:
        verbose_name = 'Subuser'
        db_table = 'subuser'
        verbose_name_plural = verbose_name    

class SubuserPaymentAddress(models.Model):
    """
    SubuserPaymentAddress table:
    white_list_uuid: address_white_list.uuid
    subuser_uuid: subuser.uuid
    percent: coin percent
    """
    white_list_uuid = models.ForeignKey(AddressWhiteList, to_field="uuid",on_delete=models.CASCADE)
    subuser_uuid = models.ForeignKey(Subuser, to_field="uuid",on_delete=models.CASCADE)
    percent = models.IntegerField(null=False, verbose_name="coin percent", default=100)

    class Meta:
        verbose_name = 'SubuserPaymentAddress'
        db_table = 'subuser_payment_address'
        verbose_name_plural = verbose_name  
        unique_together = ('white_list_uuid', 'subuser_uuid')  
        

class SubuserObserver(models.Model):
    """
    SubuserObserver table:
    subuser_uuid:               Subuser.uuid
    observer_user_uuid:         User.uuid
    """
    subuser_uuid = models.ForeignKey(Subuser, to_field="uuid",on_delete=models.CASCADE)
    observer_user_uuid = models.ForeignKey(User, to_field="uuid",on_delete=models.CASCADE)

    class Meta:
        verbose_name = 'SubuserObserver'
        db_table = 'subuser_observer'
        verbose_name_plural = verbose_name    
        unique_together = ('subuser_uuid', 'observer_user_uuid')


class DefaultAccount(models.Model):
    """
    DefaultAccount table:
    user_uuid:               User.uuid
    subuser_uuid:         Subuser.uuid
    """
    user_uuid = models.ForeignKey(User, to_field="uuid",on_delete=models.CASCADE)
    subuser_uuid = models.ForeignKey(Subuser, to_field="uuid",on_delete=models.CASCADE)

    class Meta:
        verbose_name = 'DefaultAccount'
        db_table = 'default_account'
        verbose_name_plural = verbose_name  
        unique_together = ('user_uuid', 'subuser_uuid')  


class DefaultMiner(models.Model):
    """
    DefaultMiner table:
    user_uuid:               User.uuid
    subuser_uuid:            Subuser.uuid
    currency:                currency type
    """
    user_uuid = models.ForeignKey(User, to_field="uuid",on_delete=models.CASCADE)
    subuser_uuid = models.ForeignKey(Subuser, to_field="uuid",on_delete=models.CASCADE)
    currency = models.CharField(max_length=8, null=False, verbose_name="currency type")

    class Meta:
        verbose_name = 'DefaultMiner'
        db_table = 'default_miner'
        verbose_name_plural = verbose_name    
        unique_together = ('user_uuid', 'subuser_uuid')


class EmailCode(models.Model):
    """
    EmailCode table:
    email:      User email
    code:       Email code
    status:     Email status
    create:     create time 
    """
    email = models.CharField(primary_key=True, max_length=100, null=False, 
        verbose_name="email")
    code = models.CharField(max_length=100, null=True, verbose_name="code")
    status = models.IntegerField(null=False, verbose_name="status", default=100)
    create = models.IntegerField(null=False, verbose_name="create", default=100)

    class Meta:
        verbose_name = 'EmailCode'
        db_table = 'email code'
        verbose_name_plural = verbose_name    


class Coin(models.Model):
    """
    currency:       currency type
    algorithm:      algorithm name
    """
    currency = models.CharField(max_length=8, null=False, verbose_name="currency type")
    algorithm = models.CharField(max_length=100, null=False, verbose_name="algorithm name")

    class Meta:
        verbose_name = 'Coin'
        db_table = 'coin'
        verbose_name_plural = verbose_name    


class CurrencyPoolStatus(models.Model):
    """
    currency:  
    blocks:
    hashrate:
    miners:
    workers:
    """
    currency = models.CharField(max_length=8, null=False,unique=True, verbose_name="currency type")
    blocks = models.IntegerField(verbose_name="blocks", default=0)
    hashrate = models.FloatField(verbose_name="hashrate", default=0)
    miners = models.IntegerField(verbose_name="miners", default=0)
    workers = models.IntegerField(verbose_name="workers", default=0)

    class Meta:
        verbose_name = 'CurrencyPoolStatus'
        db_table = 'currency_pool_status'
        verbose_name_plural = verbose_name    


class CurrencyStatus(models.Model):
    """
    CurrencyStatus table:
    currency:           currency table
    income:
    mean_income_24h:
    income_hashrate:
    usd:
    cny:
    network_hashrate:
    """
    currency = models.CharField(max_length=8, null=False,unique=True, verbose_name="currency type")
    income = models.FloatField(verbose_name="income", default=0)
    mean_income_24h = models.FloatField(verbose_name="mean_income_24h", default=0)
    income_hashrate = models.IntegerField(verbose_name="income_hashrate", default=0)
    usd = models.FloatField(verbose_name="usd", default=0)
    cny = models.FloatField(verbose_name="cny", default=0)
    network_hashrate = models.IntegerField(verbose_name="network_hashrate", default=0)

    class Meta:
        verbose_name = 'CurrencyStatus'
        db_table = 'currency_status'
        verbose_name_plural = verbose_name    

class LedgerCheckpoint(models.Model):
    """
    LedgerCheckpoint table:
    coin:               coin id
    timestamp:          timestamp
    """
    coin = models.IntegerField(verbose_name="coin",unique=True)
    timestamp = models.IntegerField(verbose_name="timestamp", default=0)

    class Meta:
        verbose_name = 'LedgerCheckpoint'
        db_table = 'ledger_checkpoint'
        verbose_name_plural = verbose_name   

class LedgerEvent(models.Model):
    """
    LedgerEvent table:
    timestamp:                       timestamp
    miner:                           miner name
    coin:                            coin id
    amount:
    transaction_id:
    transaction_status:
    transaction_proposed_height:
    """
    timestamp = models.BigIntegerField(verbose_name="timestamp")
    miner = models.CharField(max_length=100, verbose_name="miner")
    coin = models.IntegerField(verbose_name="coin")
    amount = models.FloatField(verbose_name="amount")
    transaction_id = models.CharField(max_length=255, verbose_name="transaction_id")
    transaction_status = models.IntegerField(verbose_name="transaction_status")
    transaction_proposed_height = models.IntegerField(verbose_name="transaction_proposed_height")

    class Meta:
        verbose_name = 'LedgerEvent'
        db_table = 'ledger_event'
        verbose_name_plural = verbose_name           


class BillStatus(models.Model):
    """
    BillStatus table:
    miner:                  miner name
    currency:               currency type
    balance:
    init_balance:
    pendding_balance:
    total_paid:
    pay1day:
    pay1week:
    paid30days:
    """
    miner = models.CharField(max_length=100, verbose_name="miner")
    currency = models.CharField(max_length=8, null=False,unique=True, verbose_name="currency type")
    balance = models.FloatField(verbose_name="balance")
    init_balance = models.FloatField(verbose_name="init_balance")
    pendding_balance = models.FloatField(verbose_name="pendding_balance")
    total_paid = models.FloatField(verbose_name="total_paid")
    pay1day = models.FloatField(verbose_name="pay1day")
    pay1week = models.FloatField(verbose_name="pay1week")
    paid30days = models.FloatField(verbose_name="paid30days")

    class Meta:
        verbose_name = 'BillStatus'
        db_table = 'bill_status'
        verbose_name_plural = verbose_name    


class MinerStatus(models.Model):
    """
    MinerStatus table:
    miner:               miner name
    currency:            currency type
    mean_hashrate_24h:
    local_hashrate:
    mean_local_hashrate_24h:
    valid_shares_24h:
    stale_shares_24h:
    invalid_shares_24h:
    online_worker_count:
    offline_worker_count:
    hashrate:
    """
    miner = models.CharField(max_length=100, verbose_name="miner")
    currency = models.CharField(max_length=8, null=False, verbose_name="currency type")
    mean_hashrate_24h = models.FloatField(verbose_name="mean_hashrate_24h")
    local_hashrate = models.FloatField(verbose_name="local_hashrate")
    mean_local_hashrate_24h = models.FloatField(verbose_name="mean_local_hashrate_24h")
    valid_shares_24h = models.IntegerField(verbose_name="valid_shares_24h")
    stale_shares_24h = models.IntegerField(verbose_name="stale_shares_24h")
    invalid_shares_24h = models.IntegerField(verbose_name="invalid_shares_24h")
    online_worker_count = models.IntegerField(verbose_name="online_worker_count")
    offline_worker_count = models.IntegerField(verbose_name="offline_worker_count")
    hashrate = models.FloatField(verbose_name="hashrate")

    class Meta:
        verbose_name = 'MinerStatus'
        db_table = 'miner_status'
        verbose_name_plural = verbose_name   
        unique_together = ('miner', 'currency')


class WorkerStatus(models.Model):
    """
    miner
    worker
    currency
    hashrate
    mean_hashrate_24h
    local_hashrate
    mean_local_hashrate_24h
    mean_hashrate_diff
    valid_shares
    stale_shares
    invalid_shares
    valid_shares_24h
    stale_shares_24h
    invalid_shares_24h
    stale_rate
    invalid_rate
    online
    last_report_time
    group_name
    """
    miner = models.CharField(max_length=100, verbose_name="miner")
    worker = models.CharField(max_length=100, verbose_name="worker")
    currency = models.CharField(max_length=8, null=False,unique=True, verbose_name="currency type")
    hashrate = models.FloatField(verbose_name="hashrate")
    mean_hashrate_24h = models.FloatField(verbose_name="mean_hashrate_24h")
    local_hashrate = models.FloatField(verbose_name="local_hashrate")
    mean_local_hashrate_24h = models.FloatField(verbose_name="mean_local_hashrate_24h")
    mean_hashrate_diff = models.FloatField(verbose_name="mean_hashrate_diff")
    valid_shares = models.IntegerField(verbose_name="valid_shares")
    stale_shares = models.IntegerField(verbose_name="stale_shares")
    invalid_shares = models.IntegerField(verbose_name="invalid_shares")
    valid_shares_24h = models.IntegerField(verbose_name="valid_shares_24h")
    stale_shares_24h = models.IntegerField(verbose_name="stale_shares_24h")
    invalid_shares_24h = models.IntegerField(verbose_name="invalid_shares_24h")
    stale_rate = models.FloatField(verbose_name="stale_rate")
    invalid_rate = models.FloatField(verbose_name="invalid_rate")
    online = models.IntegerField(verbose_name="online")
    last_report_time = models.IntegerField(verbose_name="last_report_time")
    group_name = models.CharField(max_length=100, verbose_name="group_name")

    class Meta:
        verbose_name = 'WorkerStatus'
        db_table = 'worker_status'
        verbose_name_plural = verbose_name   

"""
以下表属于 clickhouse
"""