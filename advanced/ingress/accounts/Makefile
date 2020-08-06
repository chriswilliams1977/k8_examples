build:
	#build and tag image
	docker build -t ingress-accounts .
	#tag image with GCR repo details
	docker tag ingress-accounts gcr.io/williamscj-gke-demos/ingress-example-accounts-svc:v1
	docker push  gcr.io/williamscj-gke-demos/ingress-example-accounts-svc:v1
deploy:
	#To view get svc LB external IP
	kubectl apply -f deployment.yaml
svc:
	kubectl apply -f service.yaml