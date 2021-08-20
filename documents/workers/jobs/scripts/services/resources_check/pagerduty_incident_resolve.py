import requests
import json
import sys

dedup_key = sys.argv[1]
req_url = "https://events.pagerduty.com/v2/enqueue"
headers = {'Content-Type': 'application/json'}
payload = {
  "routing_key": "7d0ce1ffe04a4907c013fab14475af75",
  "dedup_key": dedup_key,
  "event_action": "resolve"
}
res = requests.post(url=req_url, headers=headers, data=json.dumps(payload))
result = res.json()
print(result)
if result['status'] == "success":
    print("send incident to Pagerduty successfull.")
else:
    print("send incident error,please check by manual!")
