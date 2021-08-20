import requests
import json

req_url = "https://events.pagerduty.com/v2/enqueue"
headers = {'Content-Type': 'application/json'}
payload = {
    "payload": {
        "summary": "Warning: The space usage is over than 90%!",
        "source": "54.187.56.18/data/pod/clickhouse-miner-stats-0",
        "severity": "critical",
        "component": "ClickHouse",
        "group": "Database",
        "class": "Disk Space",
        "custom_details": {
            "dir": "/var/lib/clickhouse",
            "disk space percent": 92
        }
    },
    "routing_key": "7d0ce1ffe04a4907c013fab14475af75",
    "event_action": "trigger"
}
res = requests.post(url=req_url, headers=headers, data=json.dumps(payload))
result = res.json()
with open("check_df.log", "a") as fobj:
    fobj.write(result['dedup_key'] + "\n")
print(result)
if result['status'] == "success":
    print("send incident to Pagerduty successfull.")
else:
    print("send incident error,please check by manual!")
