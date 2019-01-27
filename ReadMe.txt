Plain Kubernetes Cluster einrichten mit zwei kommunizierenden Services:

1. Kubernetes einrichten
https://kubernetes.io/docs/setup/minikube/


2. Importieren des BlogWebServices (frontend)
https://github.com/davswo/BlogWebServices
    Importiere die service.yaml
kubectl create -f service.yaml


3. Importieren des BlogServices (backend)
https://github.com/davswo/BlogServices
    Importiere die service.yaml
kubectl create -f service.yaml


4. Zum testen musst du herrausfinden auf welcher IP dein minikube läuft

minikube ip

5. Anfrage über NodePort an BlogWebServices:

curl {minikubeIp}:30001/blogs

/blogs frägt den backend Service an.