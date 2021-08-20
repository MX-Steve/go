import requests
import json
    

def trigger(pay):
  req_url = "https://events.pagerduty.com/v2/enqueue"
  headers = {'Content-Type': 'application/json'}
  payload = {
      "payload": pay,
      "routing_key": "bc6758dc0fa8400ac064449d088a4503",
      "event_action": "trigger"
  }
  res = requests.post(url=req_url, headers=headers, data=json.dumps(payload))
  result = res.json()
  with open("/var/log/incident_trigger_ids.txt", "a") as fobj:
      fobj.write(result['dedup_key'] + "\n")
  print(result)
  if result['status'] == "success":
      print("send incident to Pagerduty successfull.")
  else:
      print("send incident error,please check by manual!")