1. 账户
    - dbserver:p9bzW5KGIYh9c5IB
2. users.xml
```
<dbserver>
    <password>p9bzW5KGIYh9c5IB</password>
    <networks incl="networks" replace="replace">
        <ip>::/0</ip>
    </networks>
    <profile>readonly</profile>
    <quota>default</quota>
    <allow_databases>
        <database>default</database>
        <database>miner_stats</database>
        <database>nicehash</database>
        <database>altchain_pool</database>
        <database>conflux_pool</database>
    </allow_databases>
</dbserver>
```
