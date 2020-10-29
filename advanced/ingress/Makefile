#Ingress is an alternative to loadbalancer and nodeport
#With an ingress controller you can potentially save costs.
#Rather than using 1 LoadBalancer per public application, you can use the ingress controller as a gateway for
#all your public apps, and only use 1 LoadBalancer in front of the ingress-controller

#exposes services that need access outside the cluster
#can use default or custom ingress controllers
#GKE Ingress Controller uses GCLB
#Can be external (LB across edge network) or internal (LB powered by Envoy proxies outside cluster but in VPC)
#external traffic comes into controller (runs in a pod)
#controller distributes to services
#ingress rules define where traffic goes (host or path based)
#ingress rules are defined in ingress obj
#to see gcp lb do
#gcloud compute url-maps list
ingress:
	kubectl apply -f ingress.yaml