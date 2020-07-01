#service discovery
#containers within pod can use localhost:port to connect
#pods running on other nodes in cluster use service discovery to find each other
#because you dont know ip or port
#you need service definition for service discovery to work
#app do DNS lookup to get service IP addresses - only works with apps in same ns
#they can do this by doing lookup on service dns [svc_name].[ns].svc.cluster.local - gets ip address
#use A type resolution host -> IP address
#under the hood the container calls nameserver which is kube-dns running on another pod - runs in kube-system ns
#this does the DNS discovery

#Debug
#kubectl describe pod go-mysql-deployment-67bf789bf-fxncg
#kubectl logs go-mysql-deployment-67bf789bf-fxncg
#kubectl logs go-mysql-deployment-67bf789bf-fxncg -c go-msql
#kubectl get events --sort-by=.metadata.creationTimestamp
#Go into container
#Add debug init container

#To connect to db
#kubectl exec -it mysql-database bash
#Then use mysql connection
#mysql -u root -p databasename
#then type password
#Then execute commands
#select * from db;

build:
	#build and tag image
	docker build -t go-mysql .
	#tag image with GCR repo details
	docker tag go-mysql gcr.io/williamscj-gke-demos/go-mysql:v60
	docker push  gcr.io/williamscj-gke-demos/go-mysql:v60
deploy:
	#To view get svc LB external IP
	kubectl apply -f go-deploy.yaml
svc:
	kubectl apply -f go-service.yaml