# authz
Server for externalized authorization management

## Run with Docker

```bash
docker build -t {user}/authz:v1 .
```

```bash
docker run -p 50051:50051 --name authz -d {user}/authz:v1
```

```bash
docker push {user}/authz:v1
```