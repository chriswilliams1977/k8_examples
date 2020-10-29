/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

//import the meta/v1 API group
//contains metadata common to all Kubernetes Kinds
import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

//define types for the Spec and Status of our Kind
//Kubernetes functions by reconciling desired state (Spec)
////with actual cluster state (other objects’ Status) and external state,
////and then recording what it observed (Status)
//every functional object includes spec and status
// MemcachedSpec defines the desired state of Memcached
type MemcachedSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Minimum=0
	// Size is the size of the memcached deployment
	Size int32 `json:"size"`
}

// MemcachedStatus defines the observed state of Memcached
type MemcachedStatus struct {
	// Nodes are the names of the memcached pods
	Nodes []string `json:"nodes"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

//define the types corresponding to actual Kinds, Memcached and MemcachedList
//Memcached is our root type, and describes the Memcached kind
//Like all Kubernetes objects, it contains TypeMeta (which describes API version and Kind)
//ObjectMeta, which holds things like name, namespace, and labels
// Memcached is the Schema for the memcacheds API
//adds a status subresource to the CRD manifest so that the controller can update the CR status without changing the rest of the CR object
// +kubebuilder:subresource:status
type Memcached struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MemcachedSpec   `json:"spec,omitempty"`
	Status MemcachedStatus `json:"status,omitempty"`
}

//marker
//extra metadata, telling controller-tools (our code and YAML generator) extra information
//This particular one tells the object generator that this type represents a Kind.
//Then, the object generator generates an implementation of the runtime.Object interface for us,
//which is the standard interface that all types representing Kinds must implement.
// +kubebuilder:object:root=true

//container for multiple MemcachedList. It’s the Kind used in bulk operations, like LIST.
// MemcachedList contains a list of Memcached
type MemcachedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Memcached `json:"items"`
}

//Go types to the API group. This allows us to add the types in this API group to any Scheme.
func init() {
	SchemeBuilder.Register(&Memcached{}, &MemcachedList{})
}
