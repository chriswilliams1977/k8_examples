#An operator is a way of packaging, running, and managing Kubernetes applications
#Based on control theory
#A controller is a special software that, in the loop, responds to changes and performs adaptation actions in the cluster.
#The first Kubernetes controller was a kube-controller-manager. It is regarded as an ancestor for all operators, which were built later.
#To put it simply, controller loops are the base of controller actions.
#there is a non-terminating process (called a reconciliation loop, in Kubernetes) happening over and over again
#This process observes at least one Kubernetes object, which contains information about the desired state.
#defined by configuration files made up of manifests in JSON or YAML
#the controller(s) makes a continuous adjustment via Kubernetes API to imitate the desired state until the current
#state becomes the desired state, according to the built-in logic.
#In that way, Kubernetes deals with the dynamic nature of Cloud Native systems, by handling a constant change

#An operator is an application-specific controller.
#It extends a Kubernetes API to create, configure, and manage complex applications on behalf of humans (operation engineers or site reliability engineers)
#Operators are software extensions to Kubernetes that make use of custom resources to manage applications and their components
#OPERATORS trace custom objects, often called custom resources (CR)
#Operators are controllers working in association with custom resources to perform tasks that well… “human operators” have to take care of
#CR is an extension of the Kubernetes API that provides a place where you can store and retrieve structured data—the desired state of your application
#When the operator receives any information, it will take action to adjust the Kubernetes cluster or external system to the desired state as part of its reconciliation loop in the custom controller.
#A custom resource extends Kubernetes capabilities by adding new kinds of objects that can be helpful for your application
#Custom Resource Definitions (CRD), a simple way that can be created without any programming knowledge, as an extension to the original Kubernetes API server.

#In its simplest form, an Operator adds an endpoint to the Kubernetes API, called a custom resource (CR),
#along with a control plane component that monitors and maintains resources of the new type.
#Operators comprises of a set of controllers which watches Kubernetes resources.
#Each controller implements a reconciliation loop which is triggered when a watched resource is created, updated or deleted.
#Each controller is responsible for watching one specific resource.
#Controller is a non-terminating loop that regulates the state of a system. It tries to make the current state close as possible to the desired state

#Example
#Think about deploying a database cluster with a manageable number of nodes, managing upgrades and even performing backups.
#A custom resource would specify the version of the database, the number of nodes to deploy and
#the frequency of the backups as well as their target storage,
#the controller would implement all the business logic needed to perform these operations.

#Often, there’s a one-to-one mapping between Kinds and resources.
#For instance, the pods resource corresponds to the Pod Kind.

#STEPS TO SET UP
#to check makefile tabs - cat -e -t -v makefile_name
#1)Download and install operator SDK
#Operator-SDK is a component of the Operator Framework  for scaffolding the code needed to create Kubernetes native applications.
#has 3 types of operator - Helm, Ansbile and Go (have to write  custom logic for which resources to watch)
#2)Initialize project
#3)Create api types and controller
#4)Define the API for the Memcached Custom Resource(CR) by modifying the Go type definitions
#5)Add the +kubebuilder:subresource:status marker
#After modifying the *_types.go file always run make generate
#The above makefile target will invoke the controller-gen utility to update the api/v1alpha1/zz_generated.deepcopy.go file
#to ensure our API’s Go type definitons implement the runtime.Object interface that all Kind types must implement.
#6)generate manifests - will invoke controller-gen to generate the CRD manifests at config/crd/bases/cache.example.com_memcacheds.yaml
#OpenAPIv3 schemas are added to CRD manifests in the spec.validation block when the manifests are generated.
#This validation block allows Kubernetes to validate the properties in a Memcached Custom Resource when it is created or updated.
#7)Implement the Controller
#The SetupWithManager() function in controllers/memcached_controller.go specifies how the controller is built to watch a CR and other resources that are owned and managed by that controller.
#8)Build and run the operator - build image and push
#9)deploy operator
#set namespace cd config/default/ && kustomize edit set namespace "default" && cd ../..
#10)Create a example Memcached custom resource
#kubectl apply -f config/samples/cache_v1alpha1_memcached.yaml

#for go make sure you set the right RBAC permission above reconcile func in controller
#check here https://sdk.operatorframework.io/docs/building-operators/golang/tutorial/#build-and-run-the-operator

create_cluster:
	#Create the first cluster, gke, in the us-west2-a zone, and enable Cloud Monitoring for GKE on the cluster
	gcloud beta container clusters create memcacheop \
        --zone europe-west4-a \
        --num-nodes 3 \
        --machine-type n1-standard-2 \
        --enable-stackdriver-kubernetes \
        --verbosity=none --async
install_sdk:
	brew install operator-sdk
init_operator:
	#generates a go.mod file to be used with go modules for operators
	#--repo used to specify mod path
	operator-sdk init --domain=example.com --repo=github.com/chriswilliams1977/memcached-operator
create_api_controller:
	#run this from within generated folder above
	#v1beta1 is the group version
	#
	#This will scaffold the Memcached resource API at api/[version]/memcached_types.go and the controller at controllers/memcached_controller.go
	operator-sdk create api --group=cache --version=v1beta1 --kind=Memcached
update_types:
	make generate
generate_crd_manifest:
	make manifests
register_crd:
	#Before running the operator, the CRD must be registered with the Kubernetes apiserver:
	make install
build_operator:
	make docker-build IMG=gcr.io/williamscj-demos/memcached-operator:v0.0.1
	make docker-push IMG=gcr.io/williamscj-demos/memcached-operator:v0.0.1
deploy_operator:
	make deploy IMG=gcr.io/williamscj-demos/memcached-operator:v0.0.1
create_custom_resource:
	kubectl apply -f config/samples/cache_v1beta1_memcached.yaml
get_logs:
	kubectl logs deployment.apps/memcached-operator-controller-manager -n memcached-operator-system -c manager
update_memcache_size:
	kubectl patch memcached memcached-sample -p '{"spec":{"size": 5}}' --type=merge
check_setup:
	kubectl get deployment
	kubectl get pods
	kubectl get memcached/memcached-sample -o yaml
clean_up:
	kubectl delete -f ./memcached-operator/config/samples/cache_v1beta1_memcached.yaml
	kubectl delete deployments,service -l control-plane=controller-manager
	kubectl delete role,rolebinding --all