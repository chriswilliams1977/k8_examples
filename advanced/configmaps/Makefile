#Config params that are not secret
#Key-value pairs
#can be read using env var, cmd line, volumes
#If you want to include these as env vars direct in code use literals
# kubectl create configmap language --from-literal=LANGUAGE=English
#If you want to read in a config file for app use file but you need to read this in code
#file is mounted into a volume (dir where app expects config)
#example create config file
#cat <<EOF > app.properties
#database=postgres
#EOF
#Then create config map kubectl create configmap <name-config> --from-file=<my-key-name>=<path-to-file>
#then access this in a pod via volumes
#to use a file set the mountpath and subpath to where your app will find the values
#kubectl create configmap db-config --from-file=config/app.properties --dry-run -o yaml - to see configmap
configmap:
	kubectl apply -f configmap.yaml
build:
	#build and tag image
	docker build -t go-configmap .
	#tag image with GCR repo details
	docker tag go-configmap gcr.io/williamscj-gke-demos/go-configmap:v5
	docker push  gcr.io/williamscj-gke-demos/go-configmap:v5
deploy:
	#To view get svc LB external IP
	kubectl apply -f deployment.yaml
svc:
	kubectl apply -f service.yaml