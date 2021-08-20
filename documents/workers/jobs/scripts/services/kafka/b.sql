SELECT   multiIf( 
    t1.site_key = '1MH3hWen4ftgG7H3YrryMVR4oGzS1vquCB', 'icn01', 
    t1.site_key = '1JRDc4beMDCjjce1S6Xg8ubn5DN84fQNHw', 'kwj01', 
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