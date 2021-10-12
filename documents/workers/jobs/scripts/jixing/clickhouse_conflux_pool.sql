drop table credit ;
CREATE TABLE conflux_pool.credit
(
    `timestamp` Int32,
    `username` String,
    `coin` String,
    `amount` Float64
)
ENGINE = MySQL('100.70.102.59:3306', 'conflux_pool', 'credit', 'pool', '3IxFZUpIrwXuLyzh')

drop table ledger_event
CREATE TABLE conflux_pool.ledger_event
(
    `id` Int32,
    `timestamp` Int32,
    `miner` String,
    `coin` String,
    `amount` Float64,
    `transaction_id` String,
    `transaction_status` Int32
)
ENGINE = MySQL('100.70.102.59:3306', 'conflux_pool', 'ledger_event', 'pool', '3IxFZUpIrwXuLyzh')