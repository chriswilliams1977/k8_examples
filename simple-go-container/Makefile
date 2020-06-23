#To add to GSR
#git config --global credential.'https://source.developers.google.com'.helper gcloud.sh
#gcloud source repos create udemy
#git remote add google \
#https://source.developers.google.com/p/williamscj-gke-demos/r/udemy
#Add files to master doing normal git add and commit
#git push google master

#Steps to deploy image to GCR
#To build image
#make build

#To very built images
#docker images

#To run container
#docker run (check localhost)
#To view running containers
#docker ps -a

#Useful K8 commands
#kubectl port-forward pod-name 8081:3000 - maps localhost:8081 to container port 3000
#kubectl expose pod pod-name --type=NodePort --name service-name -Expose service
#kubectl expose deployment name --type=LoadBalancer --port=80 --target-port=8080 --dry-run -o yaml
#Remember NodePort open a port on every node in cluster and forwards to app

#Get all containers in pods
#kubectl get pods --all-namespaces -o=custom-columns=NameSpace:.metadata.namespace,NAME:.metadata.name,CONTAINERS:.spec.containers[*].name

#To attach to running process in containers
#kubectl attach go-simple-deploy-9cb6c88d9-lgftm -i
#see output

#To execute command line in container
#kubectl exec go-simple-deploy-9cb6c88d9-cpp8d --stdin --tty -c go-simple  /bin/sh
#Remember scratch images does not have shell or bash so to debug

#To get one pod to call another you use the endpoint
#kubectl describe service go-simple-svc
#Get endpoint 10.44.2.5:8080
#Spin up new pod with busybox
#kubectl run -i --tty busybox --image=busybox --restart=Never
#telnet 10.44.2.5:8080

#To get logs of pod
#kubectl logs go-simple-deploy-9cb6c88d9-lgftm

#Architecture
#pods on same node can talk using localhost and port number
#pods within cluster can communicate via service discovery
#kubelet responsible for launching the pods, connects to master to get info for this
#kube-proxy feeds info of pods on nodes to iptables - kube-proxy changes rules to make sure pod is routable
#iptables is firewall in Linux and can route traffic
#service is like a load balancer - has all nodes listed and traffic is forwarded to iptables - iptables forwards to nodes and pod
#usually service has a port configured, all nodes listen to that port, iptables routes traffic to the pod

#scaling in K8 done via replication controller = number of pod replicas
#kubectl get rc
#To scale kubectl scale --replicas=4  -f replication-controller.yaml - or do via rc/name
#if app is stateless cannot scale horizontally - stateless = does not write local files or store local sessions
#dbs are stateful - they write database files
#web apps can be stateless doing session management outside containers in redis or memcache
#file cannot be written locally - shared storage like GCS or volumes

#Deployment uses 2nd generation rc called replica set - supports more complex queries
#Deployment allows you to create, update, do rolling updates and rollback an app
#Show labels kubectl get pods --show-labels  - to see metadata associated to pods
#Get deployment status kubectl rollout status deployment/go-simple-deploy
#To change image kubectl set image deployment/go-simple-deploy go-simple=gcr.io/williamscj-gke-demos/go-simple:v2
#Get rollout history kubectl rollout history deployment/go-simple-deploy
#Rollback kubectl rollout undo deployment/go-simple-deploy
#or kubectl rollout undo deployment/go-simple-deploy --to-revision=3

#Service
#ClusterIP is a virtual IP only accessible internally within the cluster
#NodePort opens a port on each node that can be reached externally
#LoadBalancer is created by cloud vendor which routes external traffic to every node with NodePort open

#Labels are key/value pairs attached to objects

#healthcheck checks status of container in a pod and spins up new one if code is not 200 for example returns 500
#livenessProbe = is container running
#readinessProbe = container is ready to serve requests - if fails pod IP removed from service so swerves no requests
#to edit a deployment kubectl edit deployment/go-simple-deploy

#pod statuses
#pending - pod allocated but image downloading
#succeded - containers are terminated successfully
#failed - continers terminated and at least one has returned failure code
#unknown - state of pod unknown, maybe networking

#K8 uses scerets to access internal API
#sercret used for credentials, keys, pwds or secret data or ssh keys
#volume is a dir your app can access
#you can store secrets in file in a volume or in env vars
#to use file create txt file
#kubectl create secret generic [name] --from-file=[path to file]
#generate secrets pwd echo -n "root" | base64
#can reference secrets from env vars or a file in pod def

docker_build:
	#build and tag image
	docker build -t go-simple .
	#tag image with GCR repo details
	docker tag go-simple gcr.io/williamscj-gke-demos/go-simple
	docker push  gcr.io/williamscj-gke-demos/go-simple
docker_run:
	docker run -p 80:8080 go-simple
gcr_build:
	#Build using Cloud Build
	#This can be automated by adding code to GSR and having a GCR trigger fired on commit
	gcloud builds submit --tag gcr.io/williamscj-gke-demos/go-simple .
gke_pod:
	kubectl apply -f pod.yaml
gke_rc:
	kubectl apply -f replication-controller.yaml
gke_deploy:
	#To view get svc LB external IP
	kubectl apply -f deployment.yaml
gke_svc
	kubectl apply -f service.yaml