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