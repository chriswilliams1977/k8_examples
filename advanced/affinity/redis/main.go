package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/go-redis/redis"
)

//Thread: A process is a part of an operating system which is responsible for executing an application.
//Every program that executes on your system is a process and to run the code inside the application a process
//uses a term known as a thread. A thread is a lightweight process, or in other words,
//a thread is a unit which executes the code under the program.
//So every program has logic and a thread is responsible for executing this logic.

//Many programs need to make use of resources that are expensive to create and maintain.
//Examples of these are database connections and threads. A resource pool provides a good way to manage these resources.
//
//When a client needs to use a database connection, rather than create one itself, it asks the pool to give it one.
//After the client is done with the connection it returns it to the pool.
//Often a pool will create a bunch of connections when it starts up.
//The size of the pool can usually be set through configuration controls and can be adapted depending on how much
//need there is from clients and the costs of maintaining the resources.
//
//If all the resources are in use when a client requests a new one,
//there are a couple of different responses. One is to throw an exception,
//another is to make the client wait until a resource comes available.

//when an application runs its uses memory
//CPU is used for garbage collection (GC) - cleaning up processes so there is memory available
//is there is not enough memory for app to run 100% of CPU time will be spend on garbage collection trying to clear space to find memory
//= all CPU cycles used trying to find memory
//> memory available < CPU cycles needed for GC
//We want to give GC no reason to run - does not need to free up memory
//use resource pool to prevent GC from running
//Instead of app going to system to get resources it goes to resource pool
//resource pool is responsible for assigning resources
//it know when it needs to get memory and when to release it
//we are not reallocating objects which impact memory and GC CPU cycles

//The purpose of the pool package is to show how you can use a buffered channel to pool a set of resources that can
//be shared and individually used by any number of goroutines.
//This pattern is useful when you have a static set of resources to share, such as database connections or memory buffers.
//When a goroutine needs one of these resources from the pool, it can acquire the resource, use it, and then return it to the pool.

//The HTTP server uses Redis as a cache database, counts the number of requests it receives,
//and prints out the number on the website.
//If the Redis service works, the number should keep increasing.
//If it is overloaded resource wise it will through an error

//to look at conf go up one dir its in conf folder - in data dir by default
//cli is in usr/local/bin
//./usr/local/bin/redis-server ./conf/redis.conf - to load conf
//You have to create the cluster first by running
//redis-cli --cluster create 10.28.0.8:6379 10.28.3.8:6379 10.28.2.6:6379 10.28.1.7:6379.... - pod endpoints

//Create our pool of resources
// Start fixed size of resource pool code.
const resourcePoolSize = 50

//resourcePool struct
//purpose is to create a pool of resources for redis
//it locks pool when a new resource is requested
//check if there are resources available
//if available increments a count
//then unlock resource to be used
//it throws an  error if resources are not available
type resourcePool struct {
	//A Mutex is used to provide a locking mechanism to ensure that only one Goroutine is running
	//the critical section of code at any point of time to prevent race condition from happening
	mtx       sync.Mutex
	allocated int
}

//create instance of resourcePool
var pool resourcePool

//ClusterClient is a Redis Cluster client representing a pool of zero or more underlying connections.
//It's safe for concurrent use by multiple goroutines.
var redisClusterClient *redis.ClusterClient

//resourcePool allocation
//locks resourcepool to stop further resources being allocated
//checks if resources are available
//increments resource count
//unlocks resourcepool to allocate resource
func (p *resourcePool) alloc() bool {
	//lock thread - only allows single go routine
	p.mtx.Lock()
	//unlock after writing
	defer p.mtx.Unlock()
	//check if there is space to allocate request
	if p.allocated < resourcePoolSize*0.9 {
		//if possible increment number of allocated requests
		p.allocated++
		return true
	}
	return false
}

//release lock
func (p *resourcePool) release() {
	//lock threads
	p.mtx.Lock()
	defer p.mtx.Unlock()
	//release thread an unlock
	p.allocated--
}

//check if resourcePool has available resources to serve request
func (p *resourcePool) hasResources() bool {
	//lock thread
	p.mtx.Lock()
	//unlock thread at end
	defer p.mtx.Unlock()
	//return bool if number of allocated requests if lower than 50
	return p.allocated < resourcePoolSize
}

// End of resource pool code.

func main() {
	// use PORT environment variable, or default to 8080
	port := "8080"
	//get port number from env vars if set
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	// register hello function to handle all requests
	server := http.NewServeMux()
	//health check
	server.HandleFunc("/healthz", healthz)
	//hello handler called from root
	server.HandleFunc("/", hello)

	// connect to redis cluster
	redisClusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:          []string{"redis-cluster:6379"},
		RouteByLatency: true})
	//close connection  at the end
	defer redisClusterClient.Close()

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	err := http.ListenAndServe(":"+port, server)
	log.Fatal(err)
}

// Start of healthz code.
func healthz(w http.ResponseWriter, r *http.Request) {
	// Log to make it simple to validate if healt checks are happening.
	log.Printf("Serving healthcheck: %s", r.URL.Path)
	//resourcePool has resources
	if pool.hasResources() {
		//write ok
		fmt.Fprintf(w, "Ok\n")
		//return
		return
	}

	//if resources unavailable return error
	w.WriteHeader(http.StatusServiceUnavailable)
	w.Write([]byte("503 - Error due to tight resource constraints in the pool!"))
}

// End of healthz code.

// hello responds to the request with a plain-text "Hello, world" message and a timestamp.
func hello(w http.ResponseWriter, r *http.Request) {

	log.Printf("Serving request: %s", r.URL.Path)

	//if unable to allocate resources return error
	if !pool.alloc() {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("503 - Error due to tight resource constraints in the pool!\n"))
		return
	} else {
		//otherwise release resourcePool when completed
		defer pool.release()
	}

	//read counts from redis
	count, err := redisClusterClient.Incr("hits").Result()
	//if error through err
	if err != nil {
		w.Write([]byte("500 - Error due to redis cluster broken!\n" + fmt.Sprintf("%v", err)))
		return
	}
	//if successful state how many request have been served
	fmt.Fprintf(w, "I have been hit [%v] times since deployment!", count)
}
