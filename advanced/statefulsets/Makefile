#StatefulSets are suitable for deploying Kafka, MySQL, Redis, ZooKeeper, and other applications needing unique,
#persistent identities and stable hostnames - pods need to be aware of each other and have stable identity
#Not recommended to run stateful datastore in K8 - can be complicated to manage
#If you can run a database on a managed service, do it. Let someone else wake up at 2 am to fix your database.

#For stateless applications, use Deployments.
#with Deployment pods are interchangeable - can be restarted with new identity
#use a service to route request to any backend - does not care where request comes from
#A stateless application is one that does not care which network it is using, and it does not need permanent storage.
#Examples of stateless apps may include web servers (Apache, Nginx, or Tomcat).

#StatefulSets use headless service (does not contain a ClusterIP)
#Pods in a StatefulSet will have an IP address mapped directly to them instead of a virtual IP address. This lets other Pods in the cluster directly communicate to an application, as in database connections.
#Instead, it creates several Endpoints that are used to produce DNS records. Each DNS record is bound to a pod
#Provide a stable hostname so not pod-randomstring
#pod name has sticky identity like podname-0
#when pod gets rescheduled it keep that identity (hostname)
#set of Pods with unique, persistent identities and stable hostnames that GKE maintains regardless of where they are scheduled
#can reference pod name in config and know it does not change
#state information and other resilient data for any given StatefulSet Pod is maintained in persistent disk storage
#associated with the StatefulSet

#also provide stateful apps with stable storage = volume mapped to origianl pod
#does not remove volumes if you scale down pods - you need to clean up
#stateful apps use DNS to find other pods

#stateful apps can also order startup and teardown - zero based
#StatefulSet Pods are deployed in sequential order and are terminated in reverse ordinal order
#startup 0 to n-1 (3 pods 0,1,2)
#teardown n-1 to 0 (3 pods 2,1,0)
#useful if you need to drain data before shutdown

deploy_storageclass:
	kubectl apply -f storageclass.yaml
deploy_headlesssvc:
	kubectl apply -f headlesssvc.yaml
deploy_statefulset:
	kubectl apply -f statefulset.yaml



