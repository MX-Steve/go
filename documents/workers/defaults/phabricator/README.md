1. Deploy Maria DB if it's not present, refer to Grafana deployment instructions
1. Generate base64 encoded password for Kubernetes Secrets configuration
```
$ echo -n 'fakepasswd123' | base64
ZmFrZXBhc3N3ZDEyMw==
```
1. Create Kubernetes Secret yaml file phabricator-secret.yaml, with object name as "**phabricator-secret**"
```
apiVersion: v1
kind: Secret
metadata:
  name: phabricator-secret
type: Opaque
data:
  MYSQL_PASS: ZmFrZXBhc3N3ZDEyMw==
```
1. Create Kubernetes Secret
```
$ kubectl apply -f ./phabricator-secret.yaml
secret/phabricator-secret created
```
1. Deploy Phabricator service with port 80 and 22 exposed through 30001 and 30002 on Kubernetes NodePort
```
kubectl apply -f phabricator.yaml
```
1. Manually create a Classic Load Balancer on AWS, forwarding port 22 to 30002 (TCP) and port 443 to 30001 (TCP), with Security Group to allow access to Kubernetes worker nodes, and allow port 443 and 22 TCP access from anywhere. Load Balancer's Health Check should be configured as using TCP Protocol to ping port 30001, since Phabricator's Nginx set up returns 500 for HTTP request without hostname.

**TODO**
Hide environment variables through Kubernetes Secrets instead of having them in plain text in configuration