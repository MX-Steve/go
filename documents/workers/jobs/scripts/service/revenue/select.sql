SELECT  yesterday(),
          o.site_key,
          site.name,
          coin,
          estimation_reward,
          income_amount
  FROM 
      (
          SELECT  site_key,
                  coin,
                  avg(estimation_reward) AS estimation_reward
          FROM 
          (
              SELECT  time,
                      site_key,
                      coin,
                      estimation_reward,
                      drift
              FROM site_estimation_reward
              LEFT JOIN pool_income_uri
                  ON site_estimation_reward.site_key=pool_income_uri.site_key
                      AND site_estimation_reward.coin =pool_income_uri.coin
              WHERE site_estimation_reward.time>subtractDays(now(),3) 
          )
          WHERE toYYYYMMDD(toDateTime(subtractHours(time,drift),'Asia/Shanghai'))=toYYYYMMDD(yesterday())
          GROUP BY  site_key,coin 
      )AS o LEFT JOIN 
      (
          SELECT  site_key,
                  name,
                  sum(income_amount) AS income_amount
          FROM 
          (
              SELECT  site_key,
                      coin AS name,
                      income_amount
              FROM pool_income_uri
              LEFT JOIN 
              (
                  SELECT  sum(amount) AS income_amount,
                          normalized_uri
                  FROM miner_pool_reward
                  WHERE toYYYYMMDD(toDateTime(time,'Asia/Shanghai'))=toYYYYMMDD(yesterday())
                  GROUP BY  normalized_uri 
              ) AS u ON pool_income_uri.normalized_uri = u.normalized_uri 
          )
          GROUP BY  site_key,name 
      ) AS q ON o.site_key = q.site_key AND coin=q.name
      JOIN site ON o.site_key=site.site_key