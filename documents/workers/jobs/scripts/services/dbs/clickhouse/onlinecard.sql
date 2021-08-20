use miner_stats;
select site.name,online_card from site
join (
select site_key,sum(online)/48 as online_card from (
select toStartOfInterval(time, INTERVAL 30 minute) as time,
site_key,
count(distinct worker,pci_id) as online 
from miner_card_stats
where time>toDateTime('2021-07-01 00:00:00')
and time<=toDateTime('2021-07-02 00:00:00')
and site_key in (unhex('00F7A298416ACBA219C449D51878FDC20423CA8396'),
unhex('00C60198E81CBBA05916463195D69AD7D376E415E3'),
unhex('005A3EACA372534660BB7DDCF0CECCE6466B45744C'),
unhex('004EE205E0A0A6BDC3B08D2FC5E3D557908E55FD75'))
group by time,site_key
) group by site_key
) as o on site.site_key=o.site_key;