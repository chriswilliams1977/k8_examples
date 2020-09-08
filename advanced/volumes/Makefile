#Volumes let your pod write to a filesystem that exists as long as the pod exists. ...
#ConfigMap and Secrets use filesystem by default
#Persistent volumes are long-term storage in your Kubernetes cluster.
#Persistent volumes exist beyond containers, pods, and nodes. A pod uses a persistent volume claim to to
#get read and write access to the persistent volume.
#persistent volume (PV) is the "physical" volume on the host machine that stores your persistent data.
#A persistent volume claim (PVC) is a request for the platform to create a PV for you, and you attach PVs to your pods via a PVC

##container in a pod has its own isolated filesystem
#new container will not see anything that was written to the filesystem by the previous container
#You may not need (or want) the whole filesystem to be persisted, but you do want to preserve the directories that hold actual data.
#StorageClass defined as a part of a pod and share the same lifecycle as the pod
#StorageClass is a way of describing available types of storage that can be provisioned
#A volume has a lifetime of the pod, pod dies = storage dies
#ConfigMap and Secrets are volume types
#This means a volume is created when the pod is started and is destroyed when the pod is deleted
#Because of this, a volume’s contents will persist across container restarts.
#A volume is available to all containers in the pod, but it must be mounted in each container that needs to access it.

#If you need to read and write to volume you can use persistent volumes
#Persistent volume in K8 allows you to attach volume that exists even if volume stopped - stateful apps
#volumes are attached using plugins in GKE this is Google persistent disk
#PV use storage class
#PV are not manage by K8 and need to be set up by admin or dynamic provisioning like on GKE
#PersistentVolumneClaim to use that storage
#Volume which uses the PVC

#a volume is a directory which is accessible to all of the containers in a Pod
#Ephemeral Volume types have the same lifetimes as their enclosing Pods.
#These Volumes are created when the Pod is created, and they persist through container restarts.
#PersistentVolumes however are cluster resources that exist independently of Pods
#In GKE, PersistentVolumes are typically backed by Compute Engine persistent disks
#PersistentVolumes can also be used with other storage types like NFS. Filestore is an NFS solution on Google Cloud.
#PersistentVolumes lifecycle is managed by Kubernetes. PersistentVolumes can be dynamically provisioned;
#you do not have to manually create and delete the backing storage.
#PersistentVolume resources can be provisioned dynamically through PersistentVolumeClaims
#A PersistentVolumeClaim is a request for and claim to a PersistentVolume resource.
#PersistentVolumeClaim objects request a specific size, access mode, and StorageClass for the PersistentVolume.
#If a PersistentVolume that satisfies the request exists or can be provisioned, the PersistentVolumeClaim is bound to that PersistentVolume.
#Volume implementations such as gcePersistentDisk are configured through StorageClass resources.
#GKE creates a default StorageClass for you which uses the standard persistent disk type

build:
	#build and tag image
	docker build -t go-volumes .
	#tag image with GCR repo details
	docker tag go-volumes gcr.io/williamscj-gke-demos/go-volumes:v2
	docker push  gcr.io/williamscj-gke-demos/go-volumes:v2
deploy_secrets:
	kubectl apply -f secrets.yaml
deploy_db:
	#To view get svc LB external IP
	kubectl apply -f pvc.yaml
	kubectl apply -f mysql-pod.yaml
	kubectl apply -f mysql-service.yaml
deploy_app:
	#To view get svc LB external IP
	kubectl apply -f deployment.yaml
	kubectl apply -f service.yaml
teardown_app:
	kubectl delete deploy go-mysql-vol-deploy
	kubectl delete svc go-mysql-vol-svc
teardown_db:
	kubectl delete pod mysql-vol-database
	kubectl delete svc mysql-vol-svc
