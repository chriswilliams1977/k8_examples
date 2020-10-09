#K8 can scale pods based on metrics
#metrics are gathered from the metrics server
#pods are scaled based on
#Actual resource usage: when a given Pod's CPU or memory usage exceeds a threshold. - per-pod resource metrics based on utilization
#Custom metrics: based on any metric reported by a Kubernetes object in a cluster
#External metrics: based on a metric from an application or service external to your cluster.

#this can be applied to deployment, statefulset, rs and rc
#Each configured Horizontal Pod Autoscaler object operates using a control loop
#pods are queried every 15 secs to check CPU utilization by the pod (controller loop)
#this can be overriden using the --horizontal-pod-autoscaler-sync-period flag
#Horizontal Pod Autoscaler (HPA) is a K8 resource

#Service resolution
#Kubernetes DNS schedules a DNS Pod and Service on the cluster, and configures the kubelets to tell
#individual containers to use the DNS Service’s IP to resolve DNS names.
#Every Service defined in the cluster (including the DNS server itself) is assigned a DNS name
#By default, a client Pod’s DNS search list will include the Pod’s own namespace and the cluster’s default domain.
#To make these names resolvable you need a tool that watches the Kubernetes API so as to allow you to know
#your container IPs, and then answer your DNS requests with the correct IP. At present the most common tool to achieve this is Kube-DNS.
#ClusterIPs are governed by IPTables rules created by kube-proxy on Workers that NAT your request to the final container’s IP.
#The Kube-DNS naming convention is service.namespace.svc.cluster-domain.tld
#the default cluster domain is cluster.local.
#if you want to contact a service called mysql in the db namespace from any namespace, you can simply speak to mysql.db.svc.cluster.local.


#In order to put load on the service to see the hpa scale you need to
#create a busybox pod or connect to running busybox pod
#kubectl run -i --tty load-generator --image=busybox /bin/sh
#kubectl exec -ti load-generator-5fb4fb465b-p9rgr /bin/sh - already running
#use wget to cponnect to the service
#wget http://go-autoscaling-backend.default.svc.cluster.local:31001
#cat index.html to see response
#rm index.html
#while true; do wget -q -O- http://go-autoscaling-svc.default.svc.cluster.local:31001; done
#watch kubectl get hpa

#potential issues
#to find svc dns do
#kubectl exec -ti go-autoscaling-656544749c-crl8z /bin/sh - into a pod
#nslookup go-autoscaling-svc
#if does not work check endpoints
#kubectl get ep -o wide
#also check is svc selector matches any pods
#kubectl get pods --show-labels |egrep 'app=go-autoscaling'
#check the container port in deployment and then check target port in svc - they should match
#Endpoints are only created if your deployment is considered healthy. If you have defined your readinessProbe
#incorrectly or the deployment does not react to it correctly, an endpoint will not be created.


docker_build:
	#build and tag image
	docker build -t go-autoscaling .
	#tag image with GCR repo details
	docker tag go-autoscaling gcr.io/williamscjdeploy-demos/go-autoscaling:v1
	docker push  gcr.io/williamscj-demos/go-autoscaling:v1
create_cluster:
	gcloud container clusters create udemy-autoscaling \
            --zone europe-west4-a \
            --num-nodes 3 \
            --machine-type n1-standard-2 \
            --enable-stackdriver-kubernetes
deploy:
	kubectl apply -f deployment.yaml
hpa:
	kubectl apply -f hpa.yaml
static_ip:
	gcloud compute addresses create go-autoscaling-ip --global
	gcloud compute addresses describe go-autoscaling-ip --global
ingress:
	kubectl apply -f ingress.yaml
generate_load:
	 kubectl exec -it deployments/go-autoscaling -- /bin/sh -c \
         "for i in $(seq -s' ' 1 10000); do wget -q -O- localhost:8080; done"
teardown:
	kubectl delete ingress go-autoscaling
	kubectl delete svc go-autoscaling-backend
	kubectl delete hpa go-autoscaling
	kubectl delete deploy go-autoscaling
	gcloud compute addresses delete go-autoscaling-ip
	gcloud container cluster delete udemy-autoscaling

