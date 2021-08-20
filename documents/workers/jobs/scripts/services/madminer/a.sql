CREATE TABLE miner_stats.bminer_card_stats
(
    `normalized_device_name` String,
    `time` Date,
    `cards_num` Int,
    `algorithm` Int32 CODEC(DoubleDelta,
 ZSTD(1)),
    `hashrate` Int64 CODEC(DoubleDelta,
 ZSTD(1)),
    `coin` LowCardinality(String)
)
ENGINE = MergeTree
PARTITION BY toYYYYMMDD(time,
 'UTC')
ORDER BY (time,
 site_key)
TTL time + toIntervalMonth(3)
SETTINGS index_granularity = 8192