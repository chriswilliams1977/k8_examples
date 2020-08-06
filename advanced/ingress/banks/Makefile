build:
	#build and tag image
	docker build -t ingress-banks .
	#tag image with GCR repo details
	docker tag ingress-banks gcr.io/williamscj-gke-demos/ingress-example-banks-svc:v1
	docker push  gcr.io/williamscj-gke-demos/ingress-example-banks-svc:v1
deploy:
	#To view get svc LB external IP
	kubectl apply -f deployment.yaml
svc:
	kubectl apply -f service.yaml