SELECT user.name AS user_name, email,site.name,c.card,p.coin,p.price,o.card_number* p.reward AS reward,o.card_number*p.reward_in_usd AS reward_in_usd
    FROM user
    JOIN user_site
        ON user.uuid=user_site.user_uuid
    JOIN
    (
        SELECT  site_key,
            sum(live_cards)/288 AS card_number
        FROM miner_liveness
        WHERE toYYYYMMDD(toDateTime(miner_liveness.time,'Asia/Shanghai'))=toYYYYMMDD(yesterday())
        GROUP BY site_key 
    ) AS o ON user_site.site_key=o.site_key
    JOIN
    (
        SELECT site_key, any(card_type) AS card 
        FROM machine_card 
        GROUP BY site_key
    ) AS c ON o.site_key=c.site_key
    JOIN 
    (
        SELECT card,
            coin,
            price,
            speed*reward_per_unit_hashrate as reward,
            reward * price as reward_in_usd
        FROM bminer_performance
        JOIN (
            select name as coin, avg(price) as price, avg(reward_per_unit_hashrate) as reward_per_unit_hashrate
            from algorithm_reward
            WHERE toYYYYMMDD(toDateTime(time,'Asia/Shanghai'))=toYYYYMMDD(yesterday())
            group by name
        )as t on bminer_performance.coin=t.coin
    ) AS p ON c.card=p.card
    JOIN site ON o.site_key=site.site_key
    ORDER BY user_name,site.name,reward_in_usd DESC