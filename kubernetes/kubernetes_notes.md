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

- **kubectl get \<component\>** = get information about the cluster components
  - **-o wide** = shows more info about the components
- **kubectl create \<component\> \<name\> \<params\>** = create a component in the cluster
  - **--image=\<image\>** = tells which image to start the deployment with
- **kubectl edit \<component\> \<name\> \<params\>** = show and edit config files
- **kubectl describe \<component\> \<name\>** = shows info about the specific component and about it's startup process
- **kubectl logs \<name\>** = shows the logs of a POD
- **kubectl delete \<component\> \<name\>** = deletes a componenet by its name
- **kubectl exec \<pod_name\> -- \<command\>** = execute commands inside of PODs
  - **-it /bin/\<shell\>** = runs an interactive terminal inside of the POD
- **kubectl apply** = executes a kubernetes commands and configurations
  - **-f \<file_path\>** = applies a config file

## YAML CONFIG FILES

### CONFIG FILE STRUCTURE

- You can use the **kubectl apply** command to apply these config files
- You should store these files where you code is stored
- A config file is made up of 3 parts
  1. Metadata
  2. Specification
  3. Status (this is automatically handled by kubernetes)

``` yaml
# 1. Metadata
apiVersion: v1 # each component has a different api version
kind: <kubernetes_componenet>
metadata:
  name: <name>
  labels: ...
# 2. Specification
# The options here depend on the kind of component
spec: ...
  # Template is for configuring the PODs for our deployments
  template:
    metadata:
      labels:
        app: <app_name>
    spec:
      # The containers we want running in a POD
      containers:
      - name: <container_name>
        image: <image_name>
        # These are the list of exposed ports
        ports:
        - containerPort: <port_number>
```

### SELECTORS & LABELS

- Deployments are linked to their PODs via labels and selectors
``` yaml
kind: Deployment
spec:
  # Using matchLabels it links all the matched labels to itself
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx # Name of the POD
```

- Selectors for services
``` yaml
kind: Service
metadata:
  name: nginx-service
spec:
  # Here the nginx Deployment and PODs are linked to the Service
  selector:
    app: nginx
```

### SERVICE PORTS

- Services redirect external connections to one of their associated PODs
``` yaml
# Service config
ports:
- protocol: TCP
  port: 80 # External cnnection will be received on this port and redirected to the targetPort
  targetPort: 8080 # This needs to match the POD's containerPort

# POD config
spec:
  containers:
  - name: <container_name>
    image: <image>
    ports:
    - containerPort: 8080
```

### EXAMPLE

``` yaml
# apiVersion specifies the Kubernetes API version this object belongs to.
# For Deployments, it's typically 'apps/v1'.
apiVersion: apps/v1
# kind specifies the type of Kubernetes object you are creating.
# In this case, it's a 'Deployment'.
kind: Deployment
metadata:
  # name is a unique identifier for your Deployment within its namespace.
  name: my-web-app-deployment
  # labels are key-value pairs used to organize and select Kubernetes objects.
  # These labels apply to the Deployment itself.
  labels:
    app: my-web-app
    environment: development

spec:
  # replicas specifies the desired number of Pod replicas (instances) for your application.
  # The Deployment controller will ensure this many Pods are running.
  replicas: 3
  # selector defines how the Deployment finds which Pods it manages.
  # It must match the labels defined in the Pod template.
  selector:
    matchLabels:
      app: my-web-app # This must match the Pod's labels below
  # template describes the Pods that the Deployment will create and manage.
  # This is essentially a Pod template.
  template:
    metadata:
      # labels here apply to the Pods created by this Deployment.
      # These labels are crucial for the Deployment's selector to work.
      labels:
        app: my-web-app # This label is matched by the Deployment's selector
    spec:
      # containers defines the container(s) that will run inside each Pod.
      containers:
      - name: my-web-app-container # Name of the container (unique within the Pod)
        image: nginx:latest       # Docker image to use for the container
        ports:
        - containerPort: 80       # The port the container listens on
        resources:
          # Define resource requests and limits for the container.
          # Requests are guaranteed, limits are the maximum allowed.
          requests:
            memory: "64Mi"
            cpu: "250m" # 250 milliCPU = 0.25 CPU core
          limits:
            memory: "128Mi"
            cpu: "500m" # 500 milliCPU = 0.5 CPU core
```