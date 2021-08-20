SELECT   multiIf( 
    t1.site_key = '1AgfHaZwJ7YqhBHibukCukjkjafjCb8bax', 'yul01', 
    t1.site_key = '1CCH7ahHUatLuiYdDEbn5n9RwZxNoD6D1n', 'yul02', 
    t1.site_key = '39C4rnYN4FpSE6tcnB7iaWGU3niDJGwvwn', 'yul02_amd', 
    t1.site_key = '18C6UKAprzSbR6R2FMFxpnak991cgjjuDb', 'urc01_p104', 
    t1.site_key = '1K3xhxvZvKaj7EG4NAfw7rY65ghA4r1Qqr', 'urc01_p102',
    t1.site_key = '19EAtxTncfYRwcyVw3reNZmVdXspyfixoR', 'urc01_1070',
    t1.site_key = '1PaNgfUFPTWPCHhXiLxUv9VADWVWqD3ztt', 'urc01_2060',
    t1.site_key = '1MH3hWen4ftgG7H3YrryMVR4oGzS1vquCB', 'icn01',
    t1.site_key = '1JRDc4beMDCjjce1S6Xg8ubn5DN84fQNHw', 'kwj01',
    t1.site_key = '1GnyNrDxA8bKDBDZQhpRhWsr8LPp52s1YK', 'nkg01', \r\n    t1.site_key) as Site,\r\n    t1.Live_Nodes,\r\n    t2.Algorithm,\r\n    t1.Total_Live_Cards,\r\n    concat(toString(floor((t1.Total_Live_Cards/t1.total_cards)*100,2)),'%') as online_rate,\r\n    t2.Total_Hashrate,\r\n    t1.Max_Temperature\r\nFROM     ( \r\n   SELECT a.site_key as site_key,algorithm,Live_Nodes,Total_Live_Cards,Max_Temperature,total_cards From(\r\nSELECT site_key,algorithm, \r\n    count(distinct worker) as Live_Nodes,\r\n    count(distinct worker,pci_id) as Total_Live_Cards\r\n    FROM miner_stats.miner_speed\r\n    WHERE time>subtractMinutes(now(),5)\r\n    GROUP BY site_key,algorithm\r\n) as a Join (select site_key,total_cards from grafana_total_cards) as c on a.site_key=c.site_key\r\nJoin (\r\n    SELECT\r\n        site_key,\r\n        max(temperature)  as Max_Temperature\r\n    FROM miner_stats.miner_temperature\r\n    WHERE time>subtractMinutes(now(),5)\r\n    and temperature >0 \r\n    GROUP BY site_key\r\n) as b on a.site_key=b.site_key\r\n    ) as t1 JOIN (\r\n        SELECT \r\n            site_key,\r\n            algorithm as Algorithm, \r\n            concat(toString(floor((sum(hash_rate)/1000000),2)),'Mh/s') as Total_Hashrate\r\n        FROM (\r\n            select \r\n                site_key,\r\n                algorithm,\r\n                avg(hash_rate) as hash_rate \r\n            from miner_stats.miner_speed \r\n        WHERE time>subtractMinutes(now(),5)\r\n        \r\n        group by \r\n            site_key,\r\n            worker,pci_id,\r\n            algorithm) \r\n    group by \r\n        site_key,\r\n        algorithm) as t2\r\n    ON t1.site_key=t2.site_key and t1.algorithm=t2.Algorithm"