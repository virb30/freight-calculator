# Freight Calculator Service

Production ready freight calculator service that suports multiple distance formulas.

Supports, by default:

- Haversine
- Cosins
- Constant distance (testing purposes)

## How To execute

After cloning the project, navigate to folder

```console
cd freight-calculator
```

### Go Version

- Install dependencies
- Start server

```console
go mod tidy
go run cmd/main.go
```

### Docker version

If you don't have or don't want to install go in your system, follow the steps below:

- Start docker compose
- Access app container
- Install dependencies
- Start server

```console
docker compose up -d
docker compose exec -it app bash
go mod tidy
go run cmd/main.go
```

## Production Ready

- Build Dockerfile.prod image
- Start builded image container

```console
docker build -f Dockerfile.prod -t your/image:verion .
docker run -p <hostPort>:<containerPort> -e "WEB_SERVER_PORT=<containerPort>" your/image:version
```
