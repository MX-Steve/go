# 1. site_key的创建：
  (1). 目前用的site_key实际上是比特币地址格式，可以到网站上生成一个地址，http://app.hubwiz.com/btc-address-generator/ 选择legecy，生成，即可
    - 1HmXau9Vc8KEk8Dgy1AGxcL8dAyKTkpRY
  (2). 入库：
    local ]# cd /data/go/src/github.com/MX-Steve/tools/genSiteKey
    local ]# ./genSiteKey 1HmXau9Vc8KEk8Dgy1AGxcL8dAyKTkpRY
    00032bd32e1993b53da3f38019af41768a95ab0e58
    controler ]#  kubectl exec -it mariadb-5ccc59b477-jc5kc -- /bin/bash
    mariadb ]# mysql -uroot -p'wl4Y5FJMLd'
    > use bminer;
    > INSERT INTO bminer VALUES(UNHEX("00032bd32e1993b53da3f38019af41768a95ab0e58"),"lihan_test","","1HmXau9Vc8KEk8Dgy1AGxcL8dAyKTkpRY")
    > SELECT HEX(uuid),name from user where name='lihan';
    13251576C30F483099CB9AF96B216F89
    > INSERT INTO user_site VALUE(NULL,UNHEX("00032bd32e1993b53da3f38019af41768a95ab0e58"),UNHEX("13251576C30F483099CB9AF96B216F89"))

# 2. database新用户创建:
  a. 生成一个密码： dd if=/dev/urandom bs=1 count=32 2> /dev/null | base64 | tr -dc "A-Za-z0-9" | head -c${1:-16}
  b. 创建用户 GRANT USAGE ON *.* TO `xxxx`@`%` IDENTIFIED BY PASSWORD 'xxxx'
  c. 授权 GRANT ALL PRIVILEGES ON `ddxxxx`.* TO `xxxx`@`%`

# 3. 2060 cfx挖矿收益与devfee，此项可与yating协同
  a. 挖矿收益: 登陆conflux浏览器https://www.confluxscan.io/，查询地址cfx:aatxetsp0kdarpdb5stdyex11dr3x6sb0jw2gykec0，
    查询当月cfx:aamwwx800rcw63n42kbehesuukjdjcnu4ueu84nhp5转给cfx:aatxetsp0kdarpdb5stdyex11dr3x6sb0jw2gykec0的数据，为挖矿收益
  b. cfx devfee:cfx:aajwd6fjbscapyrp50w976x76f08xdb8v6uv3k4zud
    查询当月cfx:aamwwx800rcw63n42kbehesuukjdjcnu4ueu84nhp5转给cfx:aajwd6fjbscapyrp50w976x76f08xdb8v6uv3k4zud的数据，为devfee
    其中cfx:aajwd6fjbscapyrp50w976x76f08xdb8v6uv3k4zud的私钥记录在configmap pool-payment-job-configs下。
  c. 2060 收益,由于我们运维2060不收devfee，因此在计算2060收益时，把它的devfee退还给它(注：devfee是miner所有的收入，但是当前除了我们自己维护的2060，几乎没有算力接到devfee)
    由于6月对cfx devfee进行过一次切换，之前的devfee收益已经打给了2060用户，之前的devfee如下：
    4112,查询方法select sum(amount) from ledger_event where timestamp>=1622476800 and timestamp<1625068800 and miner='bminer' and amount>0;
    后续的devfee可以暂存在这里cfx:aajwd6fjbscapyrp50w976x76f08xdb8v6uv3k4zud，后续运维方式改变后，统一一次打给用户
    - https://www.beepool.com/observer?url_hash=ee2177780c07f75238ef85a9a3e0b947#machines 查 devfee 地址蜜蜂矿场

# 4. 新疆电费统计，每月需要查询在线卡数，查询的办法有多种，我直接列出我的查询：
  以半个小时为单位，查询每个site有多少张卡在线
  先查询一个时间段，有多长时间
  select dateDiff('hour', toDateTime('2021-06-06 16:00:00'), toDateTime('2021-06-25 12:00:00'))*2
  例如得到904
  再查询这段时间在线情况
  select site.name,online_card from site 
    join (
    select site_key,sum(online)/904 as online_card from (
    select toStartOfInterval(time, INTERVAL 30 minute) as time,
    site_key,
    count(distinct worker,pci_id) as online 
    from miner_card_stats
    where time>toDateTime('2021-06-06 16:00:00')
    and time<=toDateTime('2021-06-25 12:00:00')
    and site_key in (unhex('00f7a298416acba219c449d51878fdc20423ca8396'),
    unhex('00c60198e81cbba05916463195d69ad7d376e415e3'),
    unhex('004ee205e0a0a6bdc3b08d2fc5e3d557908e55fd75'),
    unhex('005a3eaca372534660bb7ddcf0cecce6466b45744c'))
    group by time,site_key
    )group by site_key
    ) as o on site.site_key=o.site_key    











密语助记词
water spike muscle flame thought brave security field swap scan weasel love