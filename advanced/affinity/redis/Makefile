#The general theory of pod scheduling in kubernetes is to let the scheduler handle it.
#You tell the cluster to start a pod, the cluster looks at all the available nodes and decides where to put the new thing,
#based on comparing available resources with what the pod declares it needs.
#You can often nudge the scheduler in the right direction simply by setting resource requests appropriately.
#If your new pod needs 5 GB of ram and the only node big enough is the one you added for it to run on,
#then setting the memory request for that pod to 5 GB will force the scheduler to put it there.
#This is a fairly fragile approach, however, and while it will get your pod onto a node with sufficient resources
#it won’t keep the scheduler from putting other things there as well, as long as they will fit.


#in a deployment you can use nodeSelector to state which node pods should run
#for example
#nodeSelector:
#	hardware:high-spec - only launch pods on nodes with the label hardware:high-spec
#affinity and anti-affinity are similar to this but more complex scheduling plud works on pods
#can create preferered rules (preference) = pods can still be scheduled even if rules are not met
#nodeSelector does not schedule if no labels exist
#you can create rules to take others pods into account = never schedule these two pods on same node
#node affinity = similar to nodeSelector
#pod affinity/anti-affinity = rules on how pods are scheduled
#only works during scheduling - need to recreate if already scheduled if rule changes

#Node Affinity
#Two types
#requiredDuringScheduledIgnoredDuringExecution - hard requirement (must be met to be scheduled)
#preferredDuringScheduledIgnoredDuringExecution - try to enforce but not guaranteed
#scoring done by matching rules and the  weight of the rule
#node with highest score is where pod get scheduled
#you can your own lables to nodes or use pre-populated ones like kubernetes.io/hostname

#to label node
#kubectl label node [node-id] env=dev

#Affinity comes in two flavors: node affinity and pod affinity. It’s important to read these as properties of a pod.
#So node affinity attracts a pod to certain nodes, and pod affinity attracts a pod to certain pods.
#Pod affinity furthermore has its opposite in pod anti-affinity, which as you would expect repels a pod from other pods.

#interpod affinity and anit affinity
#allows you to influence scheduling based on labels of pds already running
#applies to specific namespace, use default if not specified
#same as node affinity/anit affinity it has two types
#requiredDuringScheduledIgnoredDuringExecution and preferredDuringScheduledIgnoredDuringExecution
#Good to use to colocate pods on same node or colocate pods in same zone
#you need to specify topology domain called topologyKey in the rules
#this refers to a node label - if node label matches topologyKey pod is scheduled on that node
#so K8 checks to match pods based on label selector
#use topologyKey to decide which node to schedule on
#pod anti affinity tells k8 not to schedule pods on that node if label matches
#both use operators in/notin and exists/doesnotexist

#EXAMPLE
#4 nodes
#MySQL on node with SSD enabled
#webapp and redis cache to be scheduled together
#no node with more than one redis pod

docker_build:
	#build and tag image
	docker build -t affinityredis .
	#tag image with GCR repo details
	docker tag affinityredis gcr.io/williamscj-demos/affinityredis:v1
	docker push  gcr.io/williamscj-demos/affinityredis:v1
config_redis:
	kubectl apply -f redis-configmap.yaml
deploy_redis:
	kubectl apply -f redis-cluster.yaml
svc_redis:
	kubectl apply -f redis-svc.yaml
deploy_app:
	kubectl apply -f app-deploy.yaml
svc_app:
	kubectl apply -f app-svc.yaml

#TO TEST
#create a busybox pod or connect to running busybox pod
#kubectl run -i --tty load-generator --image=busybox /bin/sh
#kubectl exec -ti load-generator-5fb4fb465b-p9rgr /bin/sh - already running
#use wget to connect to the service
#wget http://go-affinity-svc.default.svc.cluster.local:31001
#cat index.html to see response
#rm index.html
