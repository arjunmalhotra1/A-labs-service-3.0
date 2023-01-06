package main

/*
	Kubernetes is all abut abstracting compute power. Abstracting, machines.

	Cluster is a compute power. It is like a higher level containment inside which
	everything runs.
	It is like a higher level containment for everything that we do.

	Nodes - Can be thought of as a machine, physical or virtual.
	Node represents the actual compute power.
	A cluster could have one node and a cluster could have many nodes, depending on
	how much money we want to spend and how much compute power we need.
	When we run our cluster at our laptop/desk we only run one node, because we only got
	one machine running on.
	So, we will have single node cluster when we are running locally. But in our staging
	or production environment we will have 3 node cluster so that we can do things
	in a more dynamic way.
	The node represents the compute power that we have.
	See 1.png

	Key is to run the services in our cluster against the compute power.
	We have to abstract ourselves from if we have 2,3,4 nodes in our cluster and more about
	what's the total compute power we have and how can we distribute the applications that
	we need to run inside the cluster across that compute power.
	This is achieved with another abstraction layer called "pod".

	"pod" is a containment of one or many applications that we want to manage as a unit of
	compute.
	We can get down very granular and say "For every application that runs in this cluster
	we are going to have one pod".
	"We'll have one pod to one application".

	Sometimes we have multiple applications running and some are completely dependent
	on the others. So running them separately doesn't make sense.
	It means we will want to run multiple applications in a pod.

	In our case we will build a service, called "sales-api",
	In order to run "sales-api" in our cluster we have to configure it run in a pod.
	We could start/stop the pod, listed to the pod, have a load balancer to the pod,
	all of these things could be managed at the pod level.

	So, if we want to have multiple instances of the "sales-api" service,
	we can put that in a pod and have multiple instances of the pod running.

	For now we need just one pod of the one "sales-api" see 2.png
	Eventually we will need another service called, "metrics".

	"Metrics" is like the side car. It's providing a microservice of support,
	for the "sales-api", specifically consuming metrics and deciding where those
	metrics go, outside of the cluster. See 3.png

	Running metrics in it's own pod, we could do that but, if the "sales-api"
	isn't running then the metrics service doesn't have to run.
	This is about grouping services together that are working together
	in the one pod.

	We will be running in the original project called, "zipkin"
	Which will give us dashboard for metrics support, we should run it in it's own pod.
	Because that's like almost an entire service that could run on it's own.

	But in the original project we also run "zipkin", in the same pod.
	See 4.png because if sales-api is not running locally then we wouldn't have any
	metrics to look at either or at least tracing to look at.

	While idea of Kubernetes is that the pods, are allowed to come up and down at any point.
	Since we will not be running anything stateful.
	Talking about stateful our database needs to be stateful.

	In an environment like GCP we will have the database out of the cluster.
	See 5.png
	It would be in some other compute space where it's running.

	We might have a sidecar, DB, to provide us with special access to the DB in another cluster.
	See 6.png

	When it comes to local environment we don't need persistence so much on the
	database, while we are running we can have a local volume and we are writing to it \
	and when done, it can go away.

	In that case, what we could do and what we will do eventually is run the database,
	in it's own pod 7.png
	So our database will be running postgres which will be running in it's own pod.

	Sales.api will be able to talk to the DB using the localhost.
	Nice things among the pods that we in the cluster is that we get these
	localhost networking.

	There are databases like, DGraph which are designed to run in a kubernetes cluster.
	For production environment most likely we will not be running the pod in the cluster
	for the DB, but will run the Db outside of the cluster.
	Then we would have to make sure that we can configure the cluster and our pods will be able
	to access the database on the different network.

	For now we will have a cluster with one node, we will define the pod, we will
	get the pod to run the application "sales.api", and when we have that in place, we
	will really be able to start focussing on the production aspects of building a service
	that can run in Kubernetes.



*/
