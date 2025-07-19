# KUBERNETES

- Kubernetes helps you manage applications made up of a large amount of containers
- **High availability** = no downtime
- **Scalability** = good performance
- **Disaster recovery** = backups and resoration

## KUBERNETES/K8s COMPONENTS

### NODE

- A server

### POD

- Smallest unit/component in K8s
- An abstraction over a container; basically instead of working directly with a Docker container or other containers you use the POD to abstract it away and have the ability to change it later
- Each POD gets its own IP address
- When a POD dies, crashes or for any reason stops working a new one is spun up with a different IP address

### SERVICE

- A service gets their own IP address
- A POD can be connected to a Service and will use the services IP
  - This way if a POD dies and a new one is created the IP will stay the same
- Also works as a load balancer if the applications is replicated across mutliple nodes

### EXTERNAL & INTERNAL SERVICES

- An *external service* is exposed to external sources and opens communications with them
- An *internal service* is exposed to intenral sources and communicates with them
  - For example a database can be internal because only our apps should communicat with it

### INGRESS

- Works similarly to a DNS; basically routes the name of the service to its ip and port number; for example: https://my-app.com -> 128.0.1.98:8080

### CONFIGMAP

- Holds config data of the applications
- You can connect a POD to a ConfigMap so that it can access the config data
- These settings and configs could be accessed in the PODs as environment variables or as a properties file

### SECRET

- Works just like a ConfigMap but is used to store secret data
- Is stored in base64 encoded format

### VOLUMES

- Used for data persistance
- Connects a POD to a storage (local or remote)

### DEPLOYMENT

- A blueprint for a POD so you can easily start new ones on new nodes
- You can specify the number of POD replicas you want across nodes

### STATEFULSET

- Works similarly to a Deployment but is used for stateful apps (apps requiring state like a database for example)

## K8s ARCHITECTURE

### SLAVE NODES / WORKER NODES

- Multiple nodes in a K8s cluster
- These nodes require 3 processes to be running
  1. **Containerization runtime**
  2. **Kubelet** = handles the PODs and works as a load balancer between nodes
  3. **Kube proxy** = makes communication between PODs more performant

### MASTER NODE / CONTROL PLANE

- These nodes require 4 processes to be running
  1. **API Server** = gateway to a cluster (GUI app or anything that allows communication with the K8s cluster); also works as an authenticator
  2. **Scheduler** = schedules the newly created PODs (decides which slave to start the POD on and then it hands it off to Kubelet)
  3. **Controller** Manager = detects cluster state changes and tries to recover the cluster state
  4. **etcd** = every cluster change, update is logged and saved into this storage (its a key value store); this is how the scheduler and controller manager know about the state of the cluster

# MINIKUBE

- A local kubernetes cluster for testing and development

## MINIKUBE COMMANDS

- **minikube start** = starts a minikube cluster
  - **--vm-driver=\<vm_driver\>** = sets which VM to start on (use docker for it)
- **minikube stop** = stops the minkube cluster but saves the data to a docker volume
- **minikube delete** = removes all of the stored data about the cluster and its volume

# KUBECTL

- A CLI for interacting with the master node's API Server

## KUBECTL COMMANDS

- **kubectl get** = get information about the clusters; below are a couple of examples
  - **nodes** = shows info about the nodes inside the cluster 
  - **pod** = shows info about the PODs
  - **servies** = shows info about the services
- **kubectl create \<component\> \<params\>** = create a component in the cluster