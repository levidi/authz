# authz
Server for externalized authorization management

```bash
export PATH_CONFIG_MAP=./configMap
```

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

## test report
```bash
go test -coverprofile=report-test.out ./... 
```

```bash
go tool cover -func=report-test.out 
```

```bash
go tool cover -html=report-test.out 
```
