Anwendung mit istio laufen lassen:

1. Kubernetes einrichten
https://kubernetes.io/docs/setup/minikube/

2. Istio installation:

2.1 Istio download:
https://istio.io/docs/setup/kubernetes/download-release/

2.2 Starte Minikube mit genügend Ram und CPUs
minikube start --memory=8192 --cpus=4 --kubernetes-version=v1.10.0 \
    --vm-driver=hyperkit

2.3 Installiere Istio
https://istio.io/docs/setup/kubernetes/quick-start/

Wenn alle Istio Pods laufen, spiele Blog App ein.
kubectl get pods -n istio-system

2.4 Aktiviere automatische istio-Sidecar-injection für den default namespace
kubectl label namespace default istio-injection=enabled

3. Da automatisch Istio sidecar proxies in die neuen pods injeziert werden können wir unsere services regulär anlegen:
kubectl apply -f service.yaml

4. Lege gateway und virtualservice an um die kommunikation zu unseren services zu ermöglichen:
kubectl apply -f gateway.yaml

5. Teste die Schnittstelle (Minikube)
5.1 Herrausfinden des Ports für die Kommunikation mit dem Gateway: (Für die Kommunikation verwenden wir den istio-ingressgateway, unsecure http2)
export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].nodePort}')

5.2 Herrausfinden der IP:
minikube ip
export GATEWAY_URL=<minicubeIp>:$INGRESS_PORT

5.3 Test
curl http://${GATEWAY_URL}/blogs