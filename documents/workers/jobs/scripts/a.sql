SELECT   multiIf( 
    t1.site_key = '1AgfHaZwJ7YqhBHibukCukjkjafjCb8bax', 'yul01', 
    t1.site_key = '1CCH7ahHUatLuiYdDEbn5n9RwZxNoD6D1n', 'yul02', 
    t1.site_key = '2a6H4ahHUDtLuiYdDEbn3n9RwZxNoD2D3n', 'yul02_p102',
    t1.site_key = '39C4rnYN4FpSE6tcnB7iaWGU3niDJGwvwn', 'yul02_amd', 
    t1.site_key = '18C6UKAprzSbR6R2FMFxpnak991cgjjuDb', 'urc01_p104', 
    t1.site_key = '1K3xhxvZvKaj7EG4NAfw7rY65ghA4r1Qqr', 'urc01_p102', 
    t1.site_key = '19EAtxTncfYRwcyVw3reNZmVdXspyfixoR', 'urc01_1070', 
    t1.site_key = '1PaNgfUFPTWPCHhXiLxUv9VADWVWqD3ztt', 'urc01_2060s', 
    t1.site_key = '1MH3hWen4ftgG7H3YrryMVR4oGzS1vquCB', 'icn01', 
    t1.site_key = '1JRDc4beMDCjjce1S6Xg8ubn5DN84fQNHw', 'kwj01', 
    t1.site_key = '1GnyNrDxA8bKDBDZQhpRhWsr8LPp52s1YK', 'nkg01', 
    t1.site_key) as Site,
    t1.Live_Nodes, 
    t2.Algorithm,
    t1.Total_Cards, 
    t2.Total_Hashrate, 
    t1.Max_Temperature,
    t1.Min_Temperature
FROM     ( 
    SELECT
        site_key,
        count(distinct worker) as Live_Nodes,
        count(distinct worker,pci_id) as Total_Cards,
        max(temperature)  as Max_Temperature,
        min(temperature) as Min_Temperature
    FROM miner_stats.miner_temperature
    WHERE $timeFilter
    and temperature >0 
    GROUP BY site_key
    ) as t1 JOIN (
        SELECT 
            site_key,
            algorithm as Algorithm, 
            sum(hash_rate)  as Total_Hashrate
        FROM (
            select 
                site_key,
                algorithm,
                avg(hash_rate) as hash_rate 
            from miner_stats.miner_speed 
        WHERE $timeFilter
        
        group by 
            site_key,
            worker,pci_id,
            algorithm) 
    group by 
        site_key,
        algorithm) as t2
    ON t1.site_key=t2.site_key