#initcontainer let you execute code before app container runs
#initcontainer runs to completion
#example set up a file your app requires to run
#fetch secrets from vault and write to filesystem
#block app start until another system is ready

build:
	#build and tag image
	docker build -t go-initc .
	#tag image with GCR repo details
	docker tag go-initc gcr.io/williamscj-gke-demos/go-initc:v3
	docker push  gcr.io/williamscj-gke-demos/go-initc:v3
create_pvc:
	kubectl apply -f pvc.yaml
deploy_app:
	kubectl apply -f deployment.yaml
	kubectl apply -f service.yaml
	watch kubectl get svc
teardown_app:
	kubectl delete deploy test-app
	kubectl delete svc go-initc-svc
