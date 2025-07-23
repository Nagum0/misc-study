# CLOUD ARCHITECTURE

## Adoption of Cloud-Native Architecture, Part 1: Architecture Evolution and Maturity

- [Link to article](https://www.infoq.com/articles/cloud-native-architecture-adoption-part1/)

### Architecture evolution

- **Centralized**
  - One big centralized computer that held the data and handled user interaction
- **Distributed**
  - Client/Server model
  - Applications were split into 3 parts
    1. UI
    2. Web/Application server (API)
    3. Database server
- **Cloud hosted**
  - Servers were sold and developers had to only focus on their applications because the server infrastructure was handled by the cloud
  - **Infrastructure as a service (IaaS):** Developers choose the server specifications to host their applications and the cloud provides the hardware, OS, and networking
  - **Platform as a service (PaaS):** Developers only need to worry about their application and configuration. The cloud provider takes care of all the server infrastructure, network, and monitoring tasks.
  - **Software as a service (SaaS):** The cloud provider offers the actual applications hosted on the cloud so the client organizations can consume the application as a whole without responsibility even for the application code. This option provides the software services out of the box but it’s inflexible if the client needs to have any custom business functions outside of what’s offered by the provider.
- **Microservices**
  - Applications were split into multiple smaller independent units called microservices
  - Most of the time relies on cloud and containerization
- **Serverless**
  - This completely abstracts away the server from the developer and relies on serverless functions defined by the host

- The [article](https://www.infoq.com/articles/cloud-native-architecture-adoption-part1/) goes into much more detail about the specific pros and cons of all of the things mentioned above

## Adoption of Cloud Native Architecture, Part 2: Stabilization Gaps and Anti-Patterns

- [Link to article](https://www.infoq.com/articles/cloud-native-architecture-adoption-part2/)

### Microservice architecture anti-patterns

- **Monolithic Hell**
- **Death Star**
- **Jenga Tower**
- **Square Wheel**

#### Monolithic Hell

- Different versions
  - client/server (monolith) apps migrated to a cloud platform as is
  - a monolithic application with bundled services
  - a monolithic application artificially broken down into microservices

#### Death Star

- Evolves when applications start relying on more and more microservices and container technologies
- This anti-pattern usually occurs after a long period of time
- It gets hard to monitor the calls between the microservices (api calls call other api calls)

#### Jenga Tower and Square Wheel

- Jenga Tower/Logo Slide/Frankenstein
  - Using too many independent new dependencies
  - Stitching together the tech stack (Frankenstein)
  - Using dependencies for just specific things and utilitizing only 20% of their functionality (Jenga)
  - A lot of dependency management of versions (leads to instability)
- Square Wheel
  - Architects and teams have a tool of preference in which they are well versed and know how to use the tool to produce the necessary results, but at the cost of an architectural compromise
  - But the most common appearance of Square Wheel comes when the tool does part of the work needed and it does it well but at a cost, since the tool cannot be decoupled from unintended behavior and work needs to be done to eliminate or mask the unintended behavior

- The [article](https://www.infoq.com/articles/cloud-native-architecture-adoption-part2/) has a very good summary table about the anti-patterns and how to avoid them

### Architecture adoption pendulum

- One one side we have the DIY approach where teams implement everything from scratch (this leads to a long dev time); On the other side we have teams that overuse external dependencies which leads to a Frankenstein or Square Wheel anti-patterns.

### Architecture Goldilocks zone: Best of both worlds

- Use and reuse of best-of-breed technologies and frameworks
- Single and centralized implementation for each business and technical capability
- Service-layer abstraction for commonly used business and technical functions like authentication, user authorization, data caching, customer notifications, and so on

## Adoption of Cloud Native Architecture, Part 3: Service Orchestration and Service Mesh

- [Link to article](https://www.infoq.com/articles/cloud-native-architecture-adoption-part3/)

- Before microservices applications and services had a lot of redundant and non-functional code inside of them which was common among all of them

### Common services

- The common functionality are split into services which are running indenependelty and the client application just calls to these when it needs those functionalities
- These common services could be hosted on cloud platforms
- Common services come with their own issues and challanges

### Challenge 2: App/service and service/service communication

## 5 principles for cloud-native architecture—what it is and how to master it

- [Link to article](https://cloud.google.com/blog/products/application-development/5-principles-for-cloud-native-architecture-what-it-is-and-how-to-master-it)