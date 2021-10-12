SELECT user.name AS user_name,
	       email,
		   site_name,
		   coin,
		   reward,
		   reward * price AS reward_in_usd,
		   income_amount,
		   income_amount * price AS income_in_usd
	FROM user
	JOIN user_site ON user.uuid=user_site.user_uuid
	JOIN (
		SELECT site_key,
		       site_name,
		       coin,
		       estimation_reward AS reward,
			   pool_income AS income_amount
		FROM site_revenue_new
		WHERE toYYYYMMDD(date) = '2021-10-07'
	) AS o on user_site.site_key = o.site_key
	JOIN (
		SELECT name,
		avg(price) AS price
		FROM algorithm_reward
		WHERE toYYYYMMDD(toDateTime(time,'Asia/Shanghai'))='2021-10-07'
		GROUP BY  name
	) AS p on o.coin=p.name
	WHERE reward > 0.02
	ORDER BY  user_name,site_name