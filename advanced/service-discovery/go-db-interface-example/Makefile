build:
	#build and tag image
	docker build -t go-mysql-interface .
	#tag image with GCR repo details
	docker tag go-mysql gcr.io/williamscj-gke-demos/go-mysql-interface:v1
	docker push  gcr.io/williamscj-gke-demos/go-mysql-interface:v1
deploy:
	#To view get svc LB external IP
	kubectl apply -f go-deploy.yaml
svc:
	kubectl apply -f go-service.yaml