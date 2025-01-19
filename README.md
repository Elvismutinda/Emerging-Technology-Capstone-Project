# Capstone Project - Budget Tracker

## Project Objectives

- **Build a basic application**: Develop a simple application that consists of multiple interconnected services.
- **Containerize each service**: Use Docker to containerize each part of the application, ensuring each service has its own Docker container.
- **Network the containers**: Set up Docker networks to allow communication between the containers.
- **Manage persistent data**: Use Docker volumes for data persistence where necessary (e.g., for a database).
- **Run and monitor the application**: Deploy and run the application locally using Docker, and monitor the resource usage.

## How to run

1. Clone the repository

2. Spin up the docker containers

```
docker compose build

docker compose up -d
```

3. To check logs

```
docker compose logs -f
```

Open [http://localhost:8080](http://localhost:8080) with your browser to access the web application.

> Note that once you spin up the containers with the command specified above,

> To stop the application, run:

```
docker compose down
```

> To start the application again, run:

```
docker compose up 
```
