#node affinity = node affinity attracts a pod to certain nodes
#pod affinity = pod affinity attracts a pod to certain pods
#taints = allows nodes to repel certain pods - opposite of node affinity
#Taints mark the node
#toleration applied to pod and influence the scheduling
#example of taint = make sure new pods not scheduled on the master
#or node reserved for specific teams or user
#or nodes with specific hardware
#to apply a taint
#kubectl taint nodes [nodename] key=value:NoSchedule - set your own value
#adding a toleration to nodename will allow pod to be schedule

#taints also have a preference like affinity
#NoSchedule = hard requirement for pod to not be scheduled - unless matching toleration
#PreferNoSchedule = K8 will try to avild scheduling pod on node - but not a hard requirement
#NoExecute means running pods can be evicted - you can state how long pod can run on tainted node