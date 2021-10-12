drop table user;
CREATE TABLE miner_stats.user
(
    `uuid` String,
    `name` String,
    `github_handle` String,
    `wechat_handle` String,
    `google_handle` String,
    `oauth_detail_uuid` String,
    `authority` Int32,
    `email` String
)
ENGINE = MySQL('100.70.102.59:3306', 'bminer', 'user', 'bminer', 'aNpalSEO7BsSwhEQ');

drop table user_site;
CREATE TABLE miner_stats.user_site
(
    `id` Int32,
    `site_key` String,
    `user_uuid` String
)
ENGINE = MySQL('100.70.102.59:3306', 'bminer', 'user_site', 'bminer', 'aNpalSEO7BsSwhEQ');

drop table site_info;
CREATE TABLE miner_stats.site_info
(
    `site_key` String,
    `annotation` String,
    `pool` String,
    `coin` String,
    `normalized_uri` String
)
ENGINE = MySQL('100.70.102.59:3306', 'bminer', 'farm_stats', 'bminer', 'aNpalSEO7BsSwhEQ');

drop table site_blacklist;
CREATE TABLE miner_stats.site_blacklist
(
    `site_key` String
)
ENGINE = MySQL('100.70.102.59:3306', 'bminer', 'site_blacklist', 'bminer', 'aNpalSEO7BsSwhEQ');

drop table site;
CREATE TABLE miner_stats.site
(
    `site_key` String,
    `name` String,
    `description` String,
    `encoded_key` String
)
ENGINE = MySQL('100.70.102.59:3306', 'bminer', 'site', 'bminer', 'aNpalSEO7BsSwhEQ') ;

drop table pool_income_uri;
CREATE TABLE miner_stats.pool_income_uri
(
    `site_key` String,
    `annotation` String,
    `pool` String,
    `coin` String,
    `normalized_uri` String,
    `drift` Int32
)
ENGINE = MySQL('100.70.102.59:3306', 'bminer', 'pool_income_uri', 'bminer', 'aNpalSEO7BsSwhEQ');

drop table missing_card_machine_blacklist;
CREATE TABLE miner_stats.missing_card_machine_blacklist
(
    `site_key` String,
    `worker` String
)
ENGINE = MySQL('100.70.102.59:3306', 'bminer', 'missing_card_machine_blacklist', 'bminer', 'aNpalSEO7BsSwhEQ');

drop table miner_overview;
CREATE TABLE miner_stats.miner_overview
(
    `site_key` String,
    `worker` String,
    `algorithm` String,
    `updated_at` Int64,
    `h1_rate` Float64,
    `h24_rate` Float64
)
ENGINE = MySQL('100.70.102.59:3306', 'bminer', 'miner_overview', 'bminer', 'aNpalSEO7BsSwhEQ');

drop table machine_device_config;
CREATE TABLE miner_stats.machine_device_config
(
    `mac` String,
    `device` String
)
ENGINE = MySQL('100.70.102.59:3306', 'bminer', 'machine_device_config', 'bminer', 'aNpalSEO7BsSwhEQ');

drop table machine_card;
CREATE TABLE miner_stats.machine_card
(
    `mac` String,
    `site_key` String,
    `worker` String,
    `card_number` Int32,
    `card_type` String
)
ENGINE = MySQL('100.70.102.59:3306', 'bminer', 'machine_card', 'bminer', 'aNpalSEO7BsSwhEQ');

drop table machine_blacklist;
CREATE TABLE miner_stats.machine_blacklist
(
    `site_key` String,
    `worker` String
)
ENGINE = MySQL('100.70.102.59:3306', 'bminer', 'machine_blacklist', 'bminer', 'aNpalSEO7BsSwhEQ');

drop table machine;
CREATE TABLE miner_stats.machine
(
    `mac` String,
    `site_key` String,
    `worker` String,
    `config_group_id` String,
    `memo` String
)
ENGINE = MySQL('100.70.102.59:3306', 'bminer', 'machine', 'bminer', 'aNpalSEO7BsSwhEQ');

drop table config_group;
CREATE TABLE miner_stats.config_group
(
    `uuid` String,
    `name` String,
    `parameters` String
)
ENGINE = MySQL('100.70.102.59:3306', 'bminer', 'config_group', 'bminer', 'aNpalSEO7BsSwhEQ');

drop table card_hashrate_blacklist;
CREATE TABLE miner_stats.card_hashrate_blacklist
(
    `site_key` String,
    `worker` String,
    `pci_id` String
)
ENGINE = MySQL('100.70.102.59:3306', 'bminer', 'card_hashrate_blacklist', 'bminer', 'aNpalSEO7BsSwhEQ');

drop table bminer_performance;
CREATE TABLE miner_stats.bminer_performance
(
    `card` String,
    `coin` String,
    `speed` Float32
)
ENGINE = MySQL('100.70.102.59:3306', 'bminer', 'bminer_performance', 'bminer', 'aNpalSEO7BsSwhEQ');

drop table bminer_algorithm_info;
CREATE TABLE miner_stats.bminer_algorithm_info
(
    `id` Int32,
    `algorithm` String,
    `coin` String,
    `hashrate_factor` Float32
)
ENGINE = MySQL('100.70.102.59:3306', 'bminer', 'bminer_algorithm_info', 'bminer', 'aNpalSEO7BsSwhEQ');

drop table card_statistics;
CREATE TABLE miner_stats.card_statistics
(
    `site_key` String,
    `total_machine` Int32,
    `total_cards` Int32
)
ENGINE = MySQL('100.70.102.59:3306', 'bminer', 'card_statistics', 'bminer', 'aNpalSEO7BsSwhEQ');