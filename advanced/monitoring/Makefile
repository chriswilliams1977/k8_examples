#Prometheus came from Borgmon - monitoring for Borg
#collects metrics from monitored targets by scraping (querying) metrics HTTP endpoints on these targets
#configure scrape targets = controls what resources Prometheus monitors
#Scrape targets are generally HTTP endpoints exposed from an app
#Scraped at configurable intervals to discover and pull in metrics from fleets of machines
#Metrics collected from a scrape target are stored in the Prometheus time-series database
#In production environments, you export the metrics from the local Prometheus database to a more robust monitoring solution like Cloud Monitoring
#loud Monitoring provides a robust storage solution for metrics, logs, traces, and events.
#It also provides a suite of observability tooling that includes dashboards, reporting, alerting, and many other features.

#On GKE
#use a prometheus-stackdriver-sidecar container, running alongside a prometheus container
#Used to authenticate and send metrics to Monitoring, where they are stored and used for analysis and alerting
#best practice because it isolates scraping for metrics from the workload to be monitored.
#This approach ensures that the scraping and monitoring processes don't interfere with each other and that resources can be allocated as needed for each.

#Set wkdir
set_wkdir:
	cd $HOME
    git clone https://github.com/GoogleCloudPlatform/prometheus-stackdriver-gke
    cd $HOME/prometheus-stackdriver-gke
    WORKDIR=$(pwd)
#install Helm
install_helm:
	HELM_VERSION=v2.13.0
    HELM_PATH="$WORKDIR"/helm-"$HELM_VERSION"
    wget https://storage.googleapis.com/kubernetes-helm/helm-"$HELM_VERSION"-linux-amd64.tar.gz
    tar -xvzf helm-"$HELM_VERSION"-linux-amd64.tar.gz
    mv linux-amd64 "$HELM_PATH"
    rm $WORKDIR/helm-"$HELM_VERSION"-linux-amd64.tar.gz
#install pgbench to simulate traffic to PostgreSQL
install_pgbench:
	sudo apt-get install postgresql-contrib
#Configure IAM
configure_iam:
	#create service account to be used by prometheus monitoring sidecar (used with Cloud Monitoring)
	gcloud iam service-accounts create prometheus --display-name prometheus-service-account
	#Store the service account email address and your current project ID in environment variables for use in later commands
	export PROJECT_ID=$(gcloud info --format='value(config.project)')
    PROMETHEUS_SA_EMAIL=$(gcloud iam service-accounts list \
        --filter="displayName:prometheus-service-account" \
        --format='value(email)')
    #monitoring.metricWriter role to the Prometheus service account
    #The monitoring.metricWriter role lets the Prometheus sidecar container write metrics data to Monitoring.
    #Monitoring stores and uses these metrics for dashboards, alerts, and more
    gcloud projects add-iam-policy-binding ${PROJECT_ID} --role roles/monitoring.metricWriter --member serviceAccount:${PROMETHEUS_SA_EMAIL}
    #Download the service account key to your working directory
    #The Prometheus sidecar container uses the service account key as a secret to authenticate to the Cloud Monitoring API
    #GKE clusters have access to Google Cloud service account  thus  the Prometheus sidecar container can authenticate
    #to the Cloud Monitoring API by using the service account information from the instance metadata service
    #Non GKE clusters (on-prem) don't have access to Google Cloud service accounts thus need to use a manually configured service account key to authenticate themselves
    #n this tutorial, the GKE cluster called gke uses the Compute Engine instance metadata to authenticate to the Cloud Monitoring API to write metrics.
    #The second GKE cluster, called onprem, simulates a non-GKE cluster that doesn't have access to the Monitoring API to write metrics,
    #thus requiring you to use the service account key.
    gcloud iam service-accounts keys create $WORKDIR/prometheus-service-account.json --iam-account ${PROMETHEUS_SA_EMAIL}
    #Each GCE instance has an identity (Google Cloud service account associated on creation)
    #IAM roles and permissions associated with that service account also determine the permissions associated with that instance.
    #IF YOU DONT DEFINE SERVICE ACCOUNT WHEN CREATING CLUSTER IT USE DEFAULT SERVICE ACCOUNT = DEFAULT PERMISSIONS
create_gke_clusters:
	#Create the first cluster, gke, in the us-west2-a zone, and enable Cloud Monitoring for GKE on the cluster
	gcloud beta container clusters create gke \
        --zone us-west2-a \
        --num-nodes 3 \
        --machine-type n1-standard-2 \
        --enable-stackdriver-kubernetes \
        --verbosity=none --async
create_onprem_clusters:
    #Create the second cluster, onprem, in the us-east4-a zone:
    gcloud container clusters create onprem \
        --zone us-east4-a \
        --num-nodes 3 \
        --machine-type n1-standard-2 \
        --verbosity=none
    watch gcloud container clusters list
#Configure cluster connections and access
configure_cluster_access:
	#Connect to both clusters to generate entries in the kubeconfig file
	#The kubeconfig file is used to create authentication to clusters by creating users and contexts for each cluster.
	gcloud container clusters get-credentials gke --zone us-west2-a --project ${PROJECT_ID}
    gcloud container clusters get-credentials onprem --zone us-east4-a --project ${PROJECT_ID}
    #Use kubectx to rename the context names for convenience:
    kubectx gke=gke_${PROJECT_ID}_us-west2-a_gke
    kubectx onprem=gke_${PROJECT_ID}_us-east4-a_onprem
    #Give your Google identity cluster-admin permission for both clusters so that you can perform administrator-level tasks on the clusters:
    kubectl create clusterrolebinding user-admin-binding \
        --clusterrole=cluster-admin \
        --user=$(gcloud config get-value account) \
              --context gke
    kubectl create clusterrolebinding user-admin-binding \
        --clusterrole=cluster-admin \
        --user=$(gcloud config get-value account) \
        --context onprem
#Install Prometheus on GKE cluster
install_prom_gke:
	#create a dedicated namespace in the gke cluster for Prometheus
	kubectl create namespace prometheus --context gke
	#Create a Prometheus Kubernetes service account, a ClusterRole role, and a cluster role binding
	#The cluster role gives permissions to retrieve metrics from deployments running in all namespaces.
    #The cluster role binding assigns this role to the Prometheus service account.
	cd $WORKDIR
    kubectl apply -f prometheus-service-account.yaml --context gke
    #Create the Prometheus configmap to scrape metrics from apps running in the gke cluster
    kubectl apply -f prometheus-configmap.yaml --context gke
    #create the Prometheus server deployment
    #creates the Prometheus deployment with a single pod
    #pod is composed of two containers: the Prometheus server container and the Monitoring sidecar
    #Prometheus server container collects metrics from pods in the GKE cluster that are exporting Prometheus metrics
    #The server uses the Monitoring sidecar container to push metrics to Monitoring.
    #Define environment variables used in the Prometheus deployment manifest
    export KUBE_NAMESPACE=prometheus
    export KUBE_CLUSTER=gke
    export GCP_REGION=us-west2-a
    export GCP_PROJECT=$(gcloud info --format='value(config.project)')
    export DATA_DIR=/prometheus
    export DATA_VOLUME=prometheus-storage-volume
    export SIDECAR_IMAGE_TAG=release-0.3.2
    #Apply the Prometheus deployment manifest by using the environment variables you just defined
    envsubst < gke-prometheus-deployment.yaml | kubectl --context gke apply -f -
    #Wait a few moments and confirm that the Prometheus pod is running
    kubectl get pods -n prometheus --context gke
    #Verify the container images
    kubectl --context gke get pods -n prometheus -o json | jq '.items[].spec.containers[].image'
    #Set up port forwarding to the Prometheus server UI that's running in the gke cluster
    export PROMETHEUS_POD_GKE=$(kubectl get pods --namespace prometheus -l "app=prometheus-server" \
        -o jsonpath="{.items[0].metadata.name}" \
        --context gke)
    kubectl --context gke port-forward --namespace prometheus $PROMETHEUS_POD_GKE 9090:9090 >> /dev/null &
install_postgress:
	#install PostgreSQL by using the stable Helm chart
	#Switch to the gke cluster context
	kubectx gke
	#Grant Tiller, the server side of Helm, the cluster-admin role in the gke cluster
	kubectl create serviceaccount tiller --namespace kube-system
    kubectl create clusterrolebinding tiller-admin-binding \
        --clusterrole=cluster-admin --serviceaccount=kube-system:tiller
    #Initialize Helm and install Tiller in the gke cluster
    ${HELM_PATH}/helm init --service-account=tiller
    ${HELM_PATH}/helm repo update
    #Ensure that Helm is properly installed by running the following command:
    ${HELM_PATH}/helm version
    #Install PostgreSQL using the stable Helm release
    ${HELM_PATH}/helm install --name gke --set postgresUser=user,postgresPassword=password,postgresDatabase=postgres \
        stable/postgresql --set metrics.enabled=true \
        --set postgresqlDatabase=prometheusdb
    #Ensure that PostgreSQL is running before you proceed:
    kubectl get pods
    #Inspect the two containers running inside the PostgreSQL pod
    #The first container is the PostgreSQL container. The second container is the Prometheus metrics exporter for Postgres.
    kubectl --context gke get pods gke-postgresql-0 -ojson | jq '.spec.containers[].image'
    #Inspect the annotations on the gke-postgresql-metrics service
    #Prometheus discovers Kubernetes services by using the Kubernetes service registry.
    #It uses Kubernetes annotations on the services to determine the Prometheus target configuration
    #These annotations describe Prometheus configurations such as port, target endpoints, which services to scrape, and other settings.
    kubectl --context gke get service gke-postgresql-metrics -ojson | jq '.metadata.annotations'
install_prometheus_onprem
	#Create a dedicated namespace in the onprem cluster for Prometheus
	kubectl create namespace prometheus --context onprem
	#Create a Prometheus Kubernetes service account and the ClusterRole role
	#The ClusterRole role gives the Prometheus server permissions to retrieve metrics from deployments running in all namespaces.
	kubectl apply -f prometheus-service-account.yaml --context onprem
	#Create the Prometheus configmap to scrape metrics from apps running in the onprem cluster
	kubectl apply -f prometheus-configmap.yaml --context onprem
	#create the Prometheus deployment
	#The Prometheus server uses the Kubernetes service account called prometheus (created in step 2) to scrape metrics from apps running in the Kubernetes cluster
	#The Prometheus server uses the Monitoring sidecar container to send these metrics to Monitoring
	#The sidecar requires a Google Cloud service account to authenticate to the Cloud Monitoring API and to push metrics
	#The onprem cluster can't access Google Cloud service accounts using the Compute Engine instance metadata service.
	#To configure the sidecar, you use the Google Cloud service account's JSON key that you downloaded earlier
	#Create a secret using the Google Cloud service account's JSON key
	kubectl create secret -n prometheus generic prometheus-key --from-file=$WORKDIR/prometheus-service-account.json --context onprem
	#Define the environment variables needed to create the Prometheus deployment
	export KUBE_NAMESPACE=prometheus
    export KUBE_CLUSTER=onprem
    export GCP_REGION=us-east4-a
    export GCP_PROJECT=$(gcloud info --format='value(config.project)')
    export DATA_DIR=/prometheus
    export DATA_VOLUME=prometheus-storage-volume
    export SIDECAR_IMAGE_TAG=release-0.3.2
    #Create the Prometheus deployment
    envsubst < onprem-prometheus-deployment.yaml | kubectl --context onprem apply -f -
    #Wait a few moments and confirm that the Prometheus pod is running
    kubectl get pods -n prometheus --context onprem
    #Verify the container images by running the following command
    #You see the Prometheus server and the stackdriver-prometheus-sidecar container running in the pod.
    #The sidecar pod uses the secret prometheus-key for authentication to the Cloud Monitoring API.
    #Inspect the volumes, volumeMounts, and env for the Prometheus deployment.
    kubectl --context onprem get pods -n prometheus -ojson | jq '.items[].spec.containers[].image'
    #Inspect the volumes in Prometheus deployment
    kubectl --context onprem -n prometheus get deploy prometheus-deployment -ojson | jq '.spec.template.spec.volumes'
    #Inspect volumeMounts on the sidecar container
    kubectl --context onprem -n prometheus get deploy prometheus-deployment -ojson | jq '.spec.template.spec.containers[1].volumeMounts'
    #The volume prometheus-key is mounted in the mount path /var/secrets/google in the stackdriver-sidecar container. This path is where the service account key resides.
	#Inspect the environment variables defined for the Monitoring sidecar container.
	#The environment variable [GOOGLE_APPLICATION_CREDENTIALS](/docs/authentication/getting-started#setting_the_environment_variable)
	#is a special environment variable used by Google client libraries for authentication.
	#This variable is assigned the value that points to the path of the prometheus-service-account.json key.
	#The Monitoring sidecar uses Google client libraries and therefore uses this variable to authenticate to the Cloud Monitoring API.
	kubectl --context onprem -n prometheus get deploy prometheus-deployment -ojson | jq '.spec.template.spec.containers[1].env'
	#Set up port forwarding to the Prometheus server UI running in the onprem cluster
	export PROMETHEUS_POD_ONPREM=$(kubectl get pods --namespace prometheus -l "app=prometheus-server" \
        -o jsonpath="{.items[0].metadata.name}" --context onprem)
    kubectl --context onprem port-forward \
        --namespace prometheus $PROMETHEUS_POD_ONPREM 9091:9090 >> /dev/null &
install_postgresql_onprem:
   	#Switch to the onprem context
   	kubectx onprem
   	#Grant Tiller, the server side of Helm, the cluster-admin role in the onprem cluster:
   	kubectl create serviceaccount tiller --namespace kube-system
    kubectl create clusterrolebinding tiller-admin-binding --clusterrole=cluster-admin --serviceaccount=kube-system:tiller
    #Initialize Helm and install Tiller in the onprem cluster
    ${HELM_PATH}/helm init --service-account=tiller
    ${HELM_PATH}/helm update
    #Ensure that Helm is properly installed by running the following command:
    ${HELM_PATH}/helm version
    #Install PostgreSQL using the stable Helm release
    ${HELM_PATH}/helm install --name onprem \
    --set postgresUser=user,postgresPassword=password,postgresDatabase=postgres \
    stable/postgresql --set metrics.enabled=true --set postgresqlDatabase=prometheusdb
    #Ensure that PostgreSQL is running
    kubectl get pods
    #Inspect the two containers running in the PostgreSQL pod:
    kubectl --context onprem get pods onprem-postgresql-0 -ojson | jq '.spec.containers[].image'
    #Inspect the services running in the onprem cluster. PostgreSQL exposes Prometheus metrics using the onprem-postgresql-metrics service.
    kubectl get services --context onprem
    #Inspect the annotations on the onprem-postgresql-metrics service
    kubectl --context onprem get service onprem-postgresql-metrics -ojson | jq '.metadata.annotations'
generate_traffic:
	#Get the PostgreSQL database password in the gke cluster
	export POSTGRES_PASSWORD_GKE=$(kubectl --context gke get secret --namespace default gke-postgresql -o jsonpath="{.data.postgresql-password}" | base64 --decode)
	#Use port-forwarding to expose the PostgreSQL database on port 5432 in the gke cluster
	kubectl --context gke port-forward --namespace default svc/gke-postgresql 5432:5432 >> /dev/null &
	#Log in to the PostgreSQL database
	PGPASSWORD="$POSTGRES_PASSWORD_GKE" psql --host 127.0.0.1 -p 5432 -U postgres -d prometheusdb
	#Create a database named gketest for benchmarking
	CREATE DATABASE gketest;
	#Exit the PostgreSQL database
	\q
	#Initialize the gketest database for benchmarking
	PGPASSWORD="$POSTGRES_PASSWORD_GKE" pgbench -i -h localhost -p 5432 -U postgres -d gketest
	#Start a benchmark on gketest database in the gke cluster. The following test runs for 10 mins as configured by the -T option (in seconds).
	PGPASSWORD="$POSTGRES_PASSWORD_GKE" pgbench -c 10 -T 600 -h localhost -p 5432 -U postgres -d gketest