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

### NAMESPACES

- Used for separating applications into logical parts
- Helps when multiple teams work on the same cluster
- Accessing limits for teams
- You can add which namespace this belongs to in a config file (in metadata) or after the apply command
``` yaml
metadata:
  namespace: <namespace_name>
```

### INGRESS

- Works like a DNS
- Redirects external calls to internal services

# MINIKUBE

- A local kubernetes cluster for testing and development

## MINIKUBE COMMANDS

- **minikube start** = starts a minikube cluster
  - **--vm-driver=\<vm_driver\>** = sets which VM to start on (use docker for it)
- **minikube stop** = stops the minkube cluster but saves the data to a docker volume
- **minikube delete** = removes all of the stored data about the cluster and its volume
- **minikube service \<external_service_name\>** = assigns a public IP to the external service
- **minikube tunnel** = makes extenral (LoadBalancer) services accessible from the host machine

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
  - **--namespace=\<namespace_name\>** = apply the config in a specific namespace
- **kubectl api-resources** = shows info about resources
  - **--namespaced=\<bool\>** = only shows the namespaced or not namespaced resources

# YAML CONFIG FILES

## CONFIG FILE STRUCTURE

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

### SERVICE CONFIG

- Services redirect connections to one of their associated PODs
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

- External service config needs a **type** key in its **spec** section and a **nodePort** for external access
- **LoadBalancer type** = the cloud environment does the forwarding and you can access the service on the node's public IP and the specified port number
``` yaml
spec:
  type: LoadBalancer # LoadBalancer = external service
ports:
  - protocol: TCP
    port: 80
    targetPort: 8080
```
- **NodePort type** = with this you set a specific static node port for accessing the service
``` yaml
spec:
  type: NodePort
ports:
- protocol: TCP
  port: 80
  targetPort: 8080
  nodePort: <number between 30K-32K>
```

### SECRET CONFIG

- Secrets should already be applied to the cluster if we want to reference it somewhere
- These should never be pushed to a repository

``` yaml
apiVersion: v1
kind: Secret
metadata:
  name: <secret_name>
type: Opaque
# Values should be base64 encoded
data:
  <key>: <value>
```

- Below is how to reference a secret in a config yaml file
``` yaml
valueFrom:
  secretKeyRef:
    name: <secret_name>
    key: <key_from_secret_data_section>
```

### CONFIGMAP

- Very similar to a secret just not encoded
- External configuration
- More components can use it

``` yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: <name>
data:
  <key>: <value>
```

- Below is how to reference a ConfigMap in a config yaml file
``` yaml
valueFrom:
  configMapKeyRef:
    name: <configMap_name>
    key: <key_from_configMap_data_section>
```

### INGRESS

- Everything is defined in the rules section
- When a user enters the host address into the browser or any accessing tool, 
  Ingress will redirect the connection to the internal service

``` yaml
apiVersion: networking.k8s.io/v1beta1
kind: Ingress
metadata:
  name: <ingress_name>
spec:
  rules:
  - host: <host_name.com>
    http:
      paths:
      - backend:
          serviceName: <internal_service_name>
          servicePort: 8080
```

# HELM

- Package manager for K8s

## HELM CHART

- A bundle of yaml configuration files
- You can make your own Helm Charts with Helm
- You can bundle the helm charts into *umbrella chart*

## TEMPLATING ENGINE

- Create template files for you config files
- There are 2 parts
  1. Fix values
  2. Placeholder values

``` yaml
apiVersion: v1
kind: Pod
metadata:
  name: {{ .Values.name }}
spec:
  containers:
  - name: {{ .Values.container.name }}
    image: {{ .Values.container.image }}
    port: {{ .Values.container.port }}
```

- The placeholder values come from an additional yaml file called values.yaml
``` yaml
name: <name>
container:
  name: <containter_name>
  image: <image_name>
  port: <port_num>
```

## HELM CHART STRUCTURE

``` txt
mychart/        -> name of the chart
  Chart.yaml    -> chart metadata
  values.yaml   -> default values config file
  charts/       -> chart dependencies
  templates/    -> the actual template files
```

### Chart.yaml

``` yaml
apiVersion: v2
name: mysql-chart
version: "1.0"
description: A basic mysql db chart
```

## UMBRELLA CHART / INTEGRATION CHART

- Needs a requirements.yaml config file in which you specify other helm charts as dependencies
- This enables better and more streamlined deployments of multiple microservices which each have their own helm chart

``` yaml
dependencies:
- name: <helm_chart_name>
  repository: <url_to_helm_chart_repo>
  version: <helm_chart_version_number>
```

- You can also set values in the umbrella charts values.yaml for the dependencies

``` yaml
<helm_chart_name>:
  <key>: <value>
```

## HELM COMMANDS

- **helm install \<args\> \<chart_name\>** = applies the yaml files from the chart onto kubernetes
  - **--values=\<values_config_file_path\>** = applies a different set of values for the templates
  - **--set \<\<object_name\>.value_name\>=\<\value\>** = overwrites default value.yaml settings
- **helm uninstall \<args\> \<chart_name\>** = uninstalls and deletes all of the K8s components that are inside the chart
- **helm status \<chart_name\>** = shows info about the status of the chart
- **helm get all <\chart_name\>** = shows all of the instatiated templates from the deployed chart
- **helm upgrade \<args\> <\chart_name\>** = upgrades and already deployed chart
  - **--dry-run** = this argument will print all of the modified YAML files to stdout without applying them
- **helm history \<chart_name\>** = outputs different revisions of the chart
- **helm rollback \<chart_name\> \<revision_number\>** = rolls the deployed chart to the specified revision
- **helm dependency \<args\> \<integration_chart_name\>** = command enables working with dependencies with integration charts
  - **update** = updates the dependencies in requirements.yaml
  - **list** = lists the dependencies in the integration chart