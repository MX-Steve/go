ns=$1 # namespace
rs=$2 # resource type: cm/deployment/svc/ingress/secret
mkdir $rs
for cm in `kubectl get $rs -n $ns | awk 'NR!=1{print $1}'`
do
  kubectl get $rs $cm -n $ns -o yaml > $rs/$cm.yaml
done