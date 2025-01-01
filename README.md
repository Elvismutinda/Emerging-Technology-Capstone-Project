# Capstone Project - Budget Tracker

This is a group project for building a multi-container application using Docker.

## How to run

1. Clone the repository

2. Spin up the docker containers

```
docker compose --env-file ./backend/.env up --build -d
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
docker compose --env-file backend/.env up -d
```
