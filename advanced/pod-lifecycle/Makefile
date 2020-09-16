#Container lifecycle hooks let you execute commands at different stages of the container lifecycle
#K8 use SIGTERM to state it is terminating pods
#for example all data saved, close network connection, finish outstanding tasks
#Once K8 decides to terminate the pod the following lifecycle events occur
#Pod set to terminating = stop receiving new traffic
#pre-hook command is sent to containers
#preStop hook is a great way to trigger a graceful shutdown without modifying the application
#SIGTERM signal is sent to the container in the pod
#lets the containers know that they are going to be shut down soon
#Your code should listen for this event and start shutting down cleanly at this point.
#This may include stopping any long-lived connections (like a database connection or WebSocket stream),
#saving the current state, or anything like that.
#Kubernetes waits for a grace period - 30 secs
#parallel to the preStop hook and the SIGTERM signal. Kubernetes does not wait for the preStop hook to finish.
#if app takes longer than 30 increase using terminationGracePeriodSeconds
#SIGKILL signal is sent to pod, and the pod is removed

build:
	#build and tag image
	docker build -t go-lifecycle .
	#tag image with GCR repo details
	docker tag go-lifecycle gcr.io/williamscj-gke-demos/go-lifecycle:v3
	docker push  gcr.io/williamscj-gke-demos/go-lifecycle:v3
create_pvc:
	kubectl apply -f pvc.yaml
deploy_app:
	kubectl apply -f deployment.yaml
	kubectl apply -f service.yaml
	watch kubectl get svc
teardown_app:
	kubectl delete deploy lifecycle-app
	kubectl delete svc go-lifecycle-svc