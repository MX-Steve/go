SELECT   multiIf(      t1.site_key = '1AgfHaZwJ7YqhBHibukCukjkjafjCb8bax', 'yul01',      t1.site_key = '1CCH7ahHUatLuiYdDEbn5n9RwZxNoD6D1n', 'yul02',      t1.site_key = '17pUf1vukAzLFavDWj34VBqfUAhALqJ4v5', 'yul02_p102',     t1.site_key = '39C4rnYN4FpSE6tcnB7iaWGU3niDJGwvwn', 'yul02_amd',      t1.site_key = '18C6UKAprzSbR6R2FMFxpnak991cgjjuDb', 'urc01_p104',      t1.site_key = '1K3xhxvZvKaj7EG4NAfw7rY65ghA4r1Qqr', 'urc01_p102',      t1.site_key = '19EAtxTncfYRwcyVw3reNZmVdXspyfixoR', 'urc01_1070',      t1.site_key = '1PaNgfUFPTWPCHhXiLxUv9VADWVWqD3ztt', 'urc01_2060',      t1.site_key = '1MH3hWen4ftgG7H3YrryMVR4oGzS1vquCB', 'icn01',      t1.site_key = '1JRDc4beMDCjjce1S6Xg8ubn5DN84fQNHw', 'kwj01',      t1.site_key = '1GnyNrDxA8bKDBDZQhpRhWsr8LPp52s1YK', 'nkg01',      t1.site_key) as Site,     t1.Live_Nodes,     t2.Algorithm,     t1.Total_Live_Cards,     concat(toString(floor((t1.Total_Live_Cards/t1.total_cards)*100,2)),'%') as online_rate,     t2.Total_Hashrate,     t1.Max_Temperature FROM     (     SELECT a.site_key as site_key,algorithm,Live_Nodes,Total_Live_Cards,Max_Temperature,total_cards From( SELECT site_key,algorithm,      count(distinct worker) as Live_Nodes,     count(distinct worker,pci_id) as Total_Live_Cards     FROM miner_stats.miner_speed     WHERE time>subtractMinutes(now(),5)     GROUP BY site_key,algorithm ) as a Join (select site_key,total_cards from grafana_total_cards) as c on a.site_key=c.site_key Join (     SELECT         site_key,         max(temperature)  as Max_Temperature     FROM miner_stats.miner_temperature     WHERE time>subtractMinutes(now(),5)     and temperature >0      GROUP BY site_key ) as b on a.site_key=b.site_key     ) as t1 JOIN (         SELECT              site_key,             algorithm as Algorithm,              concat(toString(floor((sum(hash_rate)/1000000),2)),'Mh/s') as Total_Hashrate         FROM (             select                  site_key,                 algorithm,                 avg(hash_rate) as hash_rate              from miner_stats.miner_speed          WHERE time>subtractMinutes(now(),5)                  group by              site_key,             worker,pci_id,             algorithm)      group by          site_key,         algorithm) as t2     ON t1.site_key=t2.site_key and t1.algorithm=t2.Algorithm


00D36329A97C3CEF82A1BF31A96C09D462F014EFC9	madminer	1LGiNRcWXpxWiMgj7hppEL34KAkpxbQ6o9

--- SELECT hex(site_key),name,encoded_key FROM bminer.site s ;
SELECT hex(site_key),hex(user_uuid) FROM bminer.user_site WHERE site_key = unhex('00D36329A97C3CEF82A1BF31A96C09D462F014EFC9') ;
00D36329A97C3CEF82A1BF31A96C09D462F014EFC9	849F4EFF47594E54BDAEE42C3030EEB9
00D36329A97C3CEF82A1BF31A96C09D462F014EFC9	77B0FC05BC9944378E73732963158DBB
00D36329A97C3CEF82A1BF31A96C09D462F014EFC9	13251576C30F483099CB9AF96B216F89
00D36329A97C3CEF82A1BF31A96C09D462F014EFC9	4CFDE7FDC0034543A303773CFF495DCC
00D36329A97C3CEF82A1BF31A96C09D462F014EFC9	148ED1D97F9D420C8C83400E7A47471E

SELECT hex(uuid),register_email FROM user WHERE register_email = '9877cool@163.com'
148ED1D97F9D420C8C83400E7A47471E	9877cool@163.com

insert into user_site values(NULL,unhex('00D36329A97C3CEF82A1BF31A96C09D462F014EFC9'),unhex('148ED1D97F9D420C8C83400E7A47471E'))

004DD0957AA430A8CC71513E6FA1F5D7159E3B38E4	HDI-1	186SuSCtdgEybzbD275UHoz4V2PAwEMsaS
004EE205E0A0A6BDC3B08D2FC5E3D557908E55FD75	urumqi	18C6UKAprzSbR6R2FMFxpnak991cgjjuDb
005A3EACA372534660BB7DDCF0CECCE6466B45744C	urumqi_1070	19EAtxTncfYRwcyVw3reNZmVdXspyfixoR
007ACB68F8A4FFDFC7741C413E4BFC94FF4D7D85B1	yul02	1CCH7ahHUatLuiYdDEbn5n9RwZxNoD6D1n
00AD3BAA7AE10BBAFD4C64B3413862A71825C362DE	nkg01	1GnyNrDxA8bKDBDZQhpRhWsr8LPp52s1YK
00BF0EA4A4139A500817FF05FC5C2FA77A8448F017	kwj01	1JRDc4beMDCjjce1S6Xg8ubn5DN84fQNHw
00C60198E81CBBA05916463195D69AD7D376E415E3	urumqi_p102	1K3xhxvZvKaj7EG4NAfw7rY65ghA4r1Qqr
00D36329A97C3CEF82A1BF31A96C09D462F014EFC8	madminer2	1LGiNRcWXpxWiMgj7hppEL34KAkpxbQ6o8
00D36329A97C3CEF82A1BF31A96C09D462F014EFC9	madminer	1LGiNRcWXpxWiMgj7hppEL34KAkpxbQ6o9



4CFDE7FDC0034543A303773CFF495DCC