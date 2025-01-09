# Oktopus: Build Image From Source Code

This project is forked from [Oktopus](https://github.com/OktopUSP/oktopus).

The main difference between this branch and the original version is that this project builds the image from the source code. The original project used pre-built images from Docker Hub.

## Introduction

Oktopus is a powerful and flexible device management platform that supports both USP (User Services Platform) and CWMP (CPE WAN Management Protocol) protocols. It provides a centralized solution for managing, monitoring, and controlling various network devices.

The platform consists of multiple components including a message broker, API REST controller, databases, message transfer protocols, adapters, and a user-friendly frontend. Oktopus is designed to be scalable, efficient, and easy to deploy using Docker containers or Kubernetes.

## Repository Structure

```
.
├── agent
│   └── Dockerfile
├── backend
│   └── services
│       ├── acs
│       ├── bulkdata
│       ├── controller
│       ├── mtp
│       └── utils
├── deploy
│   ├── oktopus
│   │   ├── docker-compose.yaml
│   │   ├── run.sh
│   │   └── stop.sh
│   └── kubernetes
│       ├── adapter.yaml
│       ├── controller.yaml
│       ├── mqtt.yaml
│       ├── ws.yaml
│       └── ...
├── frontend
│   ├── public
│   └── src
│       ├── components
│       ├── contexts
│       ├── hooks
│       ├── layouts
│       ├── pages
│       ├── sections
│       ├── theme
│       └── utils
└── README.md
```

Key Files:
- `deploy/oktopus/docker-compose.yaml`: Main configuration file for Docker Compose deployment
- `deploy/kubernetes/*.yaml`: Kubernetes deployment configurations
- `frontend/src/pages/_app.js`: Main entry point for the frontend application
- `frontend/src/contexts/auth-context.js`: Authentication context for the frontend
- `agent/Dockerfile`: Dockerfile for building the OB-USP-A agent image

## Usage Instructions

### Installation

Prerequisites:
- Docker and Docker Compose (for Docker deployment)
- Kubernetes cluster (for Kubernetes deployment)
- Node.js and npm (for frontend development)

#### Docker Compose Deployment

1. Clone the repository:
   ```
   git clone https://github.com/Wayne45/oktopus.git
   cd oktopus
   ```

2. Set up environment variables:
   Create `.env` files for each service in the `deploy/oktopus` directory (e.g., `.env.nats`, `.env.controller`, etc.)

3. Run the deployment script:
   ```
   cd deploy/oktopus
   ./run.sh
   ```

4. To stop the deployment:
   ```
   ./stop.sh
   ```

#### Kubernetes Deployment

1. Apply the Kubernetes configurations:
   ```
   kubectl apply -f deploy/kubernetes/
   ```

2. Check the status of the pods:
   ```
   kubectl get pods
   ```

### Configuration

- NATS: Configure the message broker in `deploy/oktopus/nats_config/nats.cfg`
- Controller: Set environment variables in `.env.controller`
- Frontend: Update API endpoints in `frontend/src/contexts/backend-context.js`

### Getting Started

1. Access the frontend application at `http://localhost:80`
2. Log in using the default admin credentials (if set up)
3. Start managing devices through the dashboard

## Data Flow

The Oktopus platform follows this general data flow for device management:

1. Devices connect to the platform using various Message Transfer Protocols (MQTT, WebSocket, STOMP)
2. MTP adapters convert protocol-specific messages to a common format and publish them to NATS
3. The controller subscribes to relevant NATS topics and processes incoming messages
4. The controller stores device data and state in MongoDB
5. The frontend application communicates with the controller's REST API to display and manage device information
6. Real-time updates are pushed to the frontend using Socket.IO

```
Device -> MTP -> Adapter -> NATS -> Controller <-> MongoDB
                                         ^
                                         |
                        Frontend <-> REST API
```

## Deployment

For detailed deployment instructions, refer to the `deploy/kubernetes/README.md` file for Kubernetes deployment or use the Docker Compose files in the `deploy/oktopus` directory for container-based deployment.

## Infrastructure

The Oktopus platform consists of the following key infrastructure components:

- NATS: Message broker (nats:latest)
- Controller: API REST / USP Controller (oktopusp/controller)
- MongoDB: Database for storing device and user data (mongo)
- MQTT: MQTT broker for device communication (oktopusp/mqtt)
- WebSocket: WebSocket server for device communication (oktopusp/ws)
- STOMP: STOMP server for device communication (oktopusp/stomp)
- Adapters: Protocol-specific adapters (oktopusp/mqtt-adapter, oktopusp/ws-adapter, oktopusp/stomp-adapter)
- Frontend: User interface (oktopusp/frontend)
- Nginx: Reverse proxy and load balancer (nginx:latest)
- File Server: For serving firmware and image files (oktopusp/file-server)

Each component is containerized and can be deployed using Docker Compose or Kubernetes, allowing for easy scaling and management of the platform.