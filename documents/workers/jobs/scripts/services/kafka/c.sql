SELECT
    worker,
    max(temperature) as temperature
FROM
(
    SELECT
        worker,
        pci_id,
        temperature
    FROM miner_card_stats
    WHERE
        site_key = unhex('00DE6B6C73376743452CCE66EFAA1CB94F4BBC107B')
        and temperature > 78
        and isNotNull(temperature)
        and $timeFilter
    GROUP BY
        worker,
        pci_id,
        temperature
)
GROUP BY worker
ORDER BY temperature desc




SELECT
    worker,
    max(temperature) as temperature
FROM
(
    SELECT
        worker,
        pci_id,
        temperature
    FROM miner_card_stats
    WHERE
        site_key = unhex('00DE6B6C73376743452CCE66EFAA1CB94F4BBC107B')
        and temperature > 78
        and isNotNull(temperature)
        and $timeFilter
    GROUP BY
        worker,
        pci_id,
        temperature
)
GROUP BY worker
ORDER BY temperature desc