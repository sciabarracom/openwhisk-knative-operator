#!/bin/bash
CONF="$(dirname $0)/conf/knative-install.yaml"
if ! which kubectl
then echo "install kubectl please" ; exit 1
fi
if ! kubectl get nodes
then echo "configure access to a kubernetes cluster please" ; exit 1
fi 
kubectl apply -f "$CONF"
while true 
do POD="$(kubectl get po -l app=knative-install -o jsonpath="{.items[0].metadata.name}")"
   test -z "$POD" || break
   sleep 1
done
while ! kubectl logs $POD 2>/dev/null
do sleep 1
done
kubectl logs $POD -f &
kubectl wait --timeout=10m --for=condition=complete job/knative-install
kubectl create ns knative-whisk
kubectl get ns
echo "*** Check you have istio-system, knative-serving and tekton-pipelines"
