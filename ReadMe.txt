1. Installation von Knative wenn nicht schon geschehen: (Beinhaltet ggf. auch das installieren von kubectl und minikube)
https://github.com/knative/docs/blob/master/install/Knative-with-Minikube.md

2. Installation des des BlogWebServices:
   Aus dem Workspace das deployment.yaml importieren

kubectl apply --filename deployment.yaml

Der Service ist nun angelegt.

3. Test des Service:
3.1 "If your cluster is running outside a cloud provider (for example on Minikube), your services will never get an external IP address. In that case, use the istio hostIP and nodePort as the service IP:"

export SERVICE_IP=$(kubectl get po --selector $INGRESSGATEWAY_LABEL=ingressgateway --namespace istio-system \
--output 'jsonpath={.items[0].status.hostIP}'):$(kubectl get svc $INGRESSGATEWAY --namespace istio-system \
--output 'jsonpath={.spec.ports[?(@.port==80)].nodePort}')


3.2 Aufruf des Services mittels der istio hostIP und des nodePorts:

curl --header "Host:$SERVICE_HOST" http://${SERVICE_IP}/blogs