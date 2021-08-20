import incident_trigger
from subprocess import getstatusoutput
import logging
# LOG_FORMAT = "[%(asctime)s][%(name)s:%(pathname)s:%(levelname)s] %(message)s "
LOG_FORMAT = "[%(asctime)s][%(name)s:cronjob_status:%(levelname)s] %(message)s "
DATE_FORMAT = '%Y-%m-%d %H:%M:%S %a'
logging.basicConfig(level=logging.DEBUG,
                    format=LOG_FORMAT,
                    datefmt = DATE_FORMAT ,
                    filename="/var/log/warning.log" 
                    )
# logging.debug("msg1")
# logging.info("msg2")
# logging.warning("msg3")
# logging.error("msg4")
# logging.critical("msg5")

errList=[]
warningList=[]
services=[
  "athena-to-clickhouse-algorithm-reward-job",
  "athena-to-clickhouse-bminer-card-max-reward-job",
  "athena-to-clickhouse-bminer-card-reward-job",
  "athena-to-clickhouse-coin-reward-info-job",
  "athena-to-clickhouse-miner-pool-reward-job",
  "coin-stats-reporter",
  "kafka-ingestion-fast-new-prod",
  "kafka-ingestion-new-prod",
  "market-stats-reporter",
  "mining-dc-worker-reporter-job",
  "mining-hash-rate-reporter",
  "mining-pool-reward-reporter",
  "mining-pool-worker-reporter",
  "mining-site-revenue-reporter-job",
  "pagerduty-reporter",
  "pool-account-reporter-job",
  # "pool-network-reporter-job",
  "pool-payment-job-prod",
  "pool-stats-reporter-job",
  "site-estimation-reward-job",
  "site-revenue-record-new-job",
  "weatherdaemon",
  "etherscan-wallet-transaction-reporter",
  "madminer-machine-update-job",
  "clickhouse-to-s3",
  "bee-pool-reporter-job",
  "etl-fact-bminer-device-info-prod",
  "etl-fact-miner-profitability-prod",
  "etl-conflux-pool-credit-calculation",
  "etl-fact-bminer-revenue-prod",
  ####################  以上已梳理  #########################
  "influxdb-ingestion-new-prod",
  "conflux-block-validate-job",
  "conflux-mined-block-reporter",
  "conflux-pool-hashrate-15m-reporter-prod",
  "conflux-pool-hashrate-24h-reporter-prod",
  "conflux-pool-hashrate-reporter-prod",
  "conflux-pool-hashrate15m-reporter-job",
  "conflux-pool-hashrate24h-reporter-job",
  "conflux-pool-stats-reporter-prod",
  "conflux-user-hashrate-15m-reporter-prod",
  "conflux-user-hashrate-24h-reporter-prod",
  "conflux-user-worker-hashrate-15m-reporter-prod",
  "conflux-user-worker-hashrate-24h-reporter-prod",
  "etl-fact-algorithm-reward-prod",
  "etl-fact-bminer-active-gpu-count-prod",
  "etl-fact-bminer-algorithm-stats-prod",
  "etl-fact-bminer-farm-summary-stats-prod",
  # "etl-fact-bminer-market-share-prod",
  "etl-fact-site-hashrate-prod",
  # "pool-share-v2-rawdata-rewriter",
]
payload={
        "summary": "The following cronjobs failed to run!",
        "source": "",
        "severity": "critical",
        "component": "data",
        "group": "DataPipeline",
        "class": "Cronjobs",
        "custom_details": {
          "errList": "",
          "warningList": ""
        }
    }


def findErr(svc):
  errorCmd="kubectl get pods -n data | grep %s | grep Error | wc -l"%(svc)
  status,err=getstatusoutput(errorCmd)
  errCount = int(err)
  completedCmd="kubectl get pods -n data | grep %s | grep Completed | wc -l"%(svc)
  status,com=getstatusoutput(completedCmd)
  comCount = int(com)
  if errCount > 0 and comCount == 0:
    logging.error("Cronjob %s run error!"%(svc))
    errList.append(svc)
  elif errCount > 0 and comCount > 0:
    logging.warning("Cronjob %s run warning!"%(svc))
    warningList.append(svc)


def main():
  logging.info("Total Cronjobs: %d starting"%(len(services)))
  for svc in services:
    findErr(svc)
  if len(errList) > 0:
    payload["source"] = ",".join(errList)
    payload["custom_details"]["errList"] = errList
    payload["custom_details"]["warningList"] = warningList
    logging.info(payload)
    incident_trigger.trigger(payload)
  logging.info("Total Cronjobs: %d finished"%(len(services)))


if __name__ == '__main__':
  main()
      