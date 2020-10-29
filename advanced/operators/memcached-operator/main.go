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

//Manager initializes shared dependencies such as Caches and Clients, and provides them to Runnables.
//A Manager is required to create Controllers.
package main

import (
	"flag"
	"os"

	//import core dependencies for manager - controller logic
	//The core controller-runtime library - allows you to work with API objects
	//The default controller-runtime logging, Zap (more on that a bit later)
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	cachev1beta1 "github.com/chriswilliams1977/memcached-operator/api/v1beta1"
	"github.com/chriswilliams1977/memcached-operator/controllers"
	// +kubebuilder:scaffold:imports
)

//Every set of controllers needs a Scheme, which provides mappings between Kinds and their corresponding Go types
//When we talk about APIs in Kubernetes, we often use 4 terms: groups, versions, kinds, and resources.
//An API Group in Kubernetes is simply a collection of related functionality
//Each group has one or more versions, which, as the name suggests, allow us to change how an API works over time.
//Each API group-version contains one or more API types, which we call Kinds - k8 type e.g a pod kind
//You’ll also hear mention of resources on occasion.
//A resource is simply a use of a Kind in the API.
//Often, there’s a one-to-one mapping between Kinds and resources.
//For instance, the pods resource corresponds to the Pod Kind.
//When we refer to a kind in a particular group-version, we’ll call it a GroupVersionKind, or GVK for short.
//The Scheme we saw before is simply a way to keep track of what Go type corresponds to a given GVK
var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(cachev1beta1.AddToScheme(scheme))
	// +kubebuilder:scaffold:scheme
}

func main() {
	//basic flags for metrics
	var metricsAddr string
	var enableLeaderElection bool
	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	flag.Parse()

	ctrl.SetLogger(zap.New(zap.UseDevMode(true)))

	//how the manager registers the Scheme for the custom resource API defintions, and sets up and runs controllers and webhooks.
	//We instantiate a manager, which keeps track of running all of our controllers, as well as setting up
	//shared caches and clients to the API server (notice we tell the manager about our Scheme).
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme: scheme,
		//you can change scope of project to a single namespace
		//restricts the provided authorization to this namespace by replacing the default ClusterRole and
		//ClusterRoleBinding to Role and RoleBinding respectively
		//Namespace:          namespace,
		//NewCache:           cache.MultiNamespacedCacheBuilder(namespaces),
		MetricsBindAddress: metricsAddr,
		Port:               9443,
		LeaderElection:     enableLeaderElection,
		LeaderElectionID:   "467969ea.example.com",
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	//We run our manager, which in turn runs all of our controllers and webhooks.
	//The manager is set up to run until it receives a graceful shutdown signal.
	//This way, when we’re running on Kubernetes, we behave nicely with graceful pod termination.
	if err = (&controllers.MemcachedReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("Memcached"),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Memcached")
		os.Exit(1)
	}
	// +kubebuilder:scaffold:builder

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
