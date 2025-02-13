## Oktopus: Build Image From Source Code

This project is forked from [Oktopus](https://github.com/OktopUSP/oktopus).

The main difference between this branch and the original version is that this project builds the image from the source code. The original project used pre-built images from Docker Hub.

### Prerequisites

Setup MQTT_USERNAMES and MQTT_PASSWORDS on .env.mqtt file for agents and mqtt-adapter services.
They are comma-separated list of MQTT usernames and passwords.

### Docker Compose Deployment

1. Clone the repository:
   ```
   git clone https://github.com/Wayne45/oktopus.git
   cd oktopus
   ```

2. Set up environment variables:
   Create `.env` files for each service in the `deploy/oktopus` directory (e.g., `.env.nats`, `.env.controller`, etc.)

3. Build and run the deployment script:
   ```
   cd deploy/oktopus
   ./build.sh
   ./run.sh
   ```

4. To stop the deployment:
   ```
   ./stop.sh
   ```

### Getting Started

1. Access the frontend application at `http://localhost:80`
2. Log in using the default admin credentials (if set up)
3. Start managing devices through the dashboard