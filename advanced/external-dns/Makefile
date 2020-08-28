#External DNS allows you to automatically register service domain with CloudDNS
#It runs in a separate pod
#The external-dns pod checks the ingress (and/or services) for annotations and creates DNS records in upstream DNS provider.
#These can then be used with ingress and the GCLB to route traffic to svcs based on host (domain)
#This can also be done reserving s static IP and map to A record with ingress IP

#Steps to create
#Create project settings
gcloud config set project "williamscj-gke-demos"
gcloud config set compute/region "europe-west4"
gcloud config set compute/zone "europe-west4-a"
#Create cluster with scope to update clouddns
#Scopes inform the access level your cluster nodes will have to specific GCP services as a whole.
#All Google Compute Engine VMs come with a built in OAuth2 Service Account that can be used to automatically
#authenticate to various GCP services. You can choose which services this service account has permission to
#access by assigning different “scopes” to the account.
gcloud container clusters create "udemy-demos" \
    --num-nodes 3 \
    --scopes "https://www.googleapis.com/auth/ndev.clouddns.readwrite"
#if existing cluster
#gcloud container clusters describe udemy-demos --zone=europe-west4-a - to see existing config
#look under node pool oauthScopes
#if you need new scope you need to create a new node pool and drain
#view existing node pools
#gcloud container node-pools list --zone=europe-west4-a --cluster=udemy-demos
gcloud container node-pools create udemy-pool \
   --cluster udemy-demos --zone europe-west4-a \
   --num-nodes 3 \
   --scopes https://www.googleapis.com/auth/ndev.clouddns.readwrite
#then cordon and drain each node from old pool so they get reassigned
#kubectl get nodes
#if refused do gcloud container clusters get-credentials udemy-demos
#make sure zone is set gcloud config set compute/zone europe-west4-a
kubectl cordon gke-udemy-demos-default-pool-b2fca1b7-bstl gke-udemy-demos-default-pool-b2fca1b7-d3gb gke-udemy-demos-default-pool-b2fca1b7-z9zf
#then drain
kubectl drain gke-udemy-demos-default-pool-b2fca1b7-bstl gke-udemy-demos-default-pool-b2fca1b7-d3gb gke-udemy-demos-default-pool-b2fca1b7-z9zf --force
#then delete old pool
gcloud container node-pools delete default-pool \
   --cluster udemy-demos --zone europe-west4-a


#DNS is a hierarchical distributed database that lets you store IP addresses and other data, and look them up by name.
#Cloud DNS lets you publish your zones and records in the DNS without the burden of managing your own DNS servers and software.
#Create a DNS parent zone for your domain name
 gcloud dns managed-zones create "jackburton-io" \
    --dns-name "jackburton.io." \
    --description "Parent Zone"
#POINT GOOGLE DOMAIN NS AT CLOUD DNS PARENT MANAGED ZONE OTHERWISE THIS DOES NOT WORK
#Example ns-cloud-d1.googledomains.com


#Create DNS zone which will contain the managed DNS records.
#The managed zone holds DNS records for the same DNS name suffix e.g example.com
#In Cloud DNS, the managed zone is the resource that models a DNS zone.
#MAKE SURE dns-name IS IN VALID FORMAT AND SUB DOMAIN
 gcloud dns managed-zones create "externaldns-jackburton-io" \
    --dns-name "externaldns.jackburton.io." \
    --description "Automatically managed zone by kubernetes.io/external-dns"
#Make a note of the nameservers that were assigned to your new zone for example ns-cloud-a1.googledomains.com.
gcloud dns record-sets list \
                 --zone "externaldns-jackburton-io" \
                 --name "externaldns.jackburton.io." \
                 --type NS

#Tell the parent zone where to find the DNS records for this zone by adding the corresponding NS records there.
#Assuming the parent zone is "jackburton-io" and the domain is "jackburton.io" and that it's also hosted at Google we would do the following.
gcloud dns record-sets transaction start --zone "jackburton-io"
gcloud dns record-sets transaction add ns-cloud-c{1..4}.googledomains.com. \
    --name "externaldns.jackburton.io." --ttl 300 --type NS --zone "jackburton-io"
gcloud dns record-sets transaction execute --zone "jackburton-io"
#
#Deploy ExternalDNS
#Role-Based Access Control (RBAC)
#RBAC is enabled by default on all Container clusters which are running Kubernetes version 1.6 or higher.
#Because of the way Container Engine checks permissions when you create a Role or ClusterRole, you must first create a
#RoleBinding that grants you all of the permissions included in the role you want to create.
kubectl create clusterrolebinding your-user-cluster-admin-binding --clusterrole=cluster-admin --user=williamscj@google.com
#Connect your kubectl client to the cluster you just created.
gcloud container clusters get-credentials "udemy-demos"
#Then apply externalednswithrbac.yaml to deploy ExternalDNS.
#Check to see what would of been created kubectl apply -f externalednswithrbac.yaml --dry-run
#Verify ExternalDNS works
#NOTE: Annotations added to service
#NOTE: externalednswithrbac svc looks at service or ingress running to set A and TXT records
#Check DNS records created for service
gcloud dns record-sets list \
    --zone "externaldns-jackburton-io" \
    --name "externaldns.jackburton.io."
#Note created TXT record alongside A record. TXT record signifies that the corresponding
#A record is managed by ExternalDNS. This makes ExternalDNS safe for running in environments where
#there are other records managed via other means.
#Let's check that we can resolve this DNS name. We'll ask the nameservers assigned to your zone first.
#Make sure you use correct namespace record numbers
dig +short @ns-cloud-b1.googledomains.com. customers.externaldns.jackburton.io.
#Given you hooked up your DNS zone with its parent zone you can use curl to access your site.
curl customers.externaldns.jackburton.io.

#Clean Up
gcloud dns record-sets transaction start --zone "jackburton-io"
gcloud dns record-sets transaction remove ns-cloud-c{1..4}.googledomains.com. \
    --name "externaldns.jackburton.io." --ttl 300 --type NS --zone "jackburton-io"
gcloud dns record-sets transaction execute --zone "jackburton-io"