build:
	#build and tag image
	docker build -t go-mysql-interface .
	#tag image with GCR repo details
	docker tag go-mysql gcr.io/williamscj-gke-demos/go-mysql-interface:v2
	docker push  gcr.io/williamscj-gke-demos/go-mysql-interface:v2
deploy_db:
	#To view get svc LB external IP
	kubectl apply -f mysql-pod.yaml
	kubectl apply -f mysql-service.yaml
deploy_app:
	#To view get svc LB external IP
	kubectl apply -f go-deploy.yaml
	kubectl apply -f go-service.yaml
	watch kubectl get svc
teardown_app:
	kubectl delete deploy go-mysql-interface-deployment
	kubectl delete svc go-mysql-interface-svc
teardown_db:
	kubectl delete pod mysql-interface-database
	kubectl delete svc mysql-interface-svc
