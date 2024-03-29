### Run locally:
1. Run each server individually

path: go-reverse-proxy/origin-server
```bash
go run main.go
```
path: go-reverse-proxy/origin-server-2
```bash
go run main.go
```
path: go-reverse-proxy/proxy-server
```bash
go run main.go
```

### Run using Docker file:
1. Docker build and run:
path: go-reverse-proxy/origin-server
```bash
docker build -t origin:v1
docker run -d -p 8081:8081 origin:v1
```
path: go-reverse-proxy/origin-server-2
```bash
docker build -t origin2:v1
docker run -d -p 8082:8082 origin2:v1
```
path: go-reverse-proxy/proxy-server
```bash
docker build -t reverse-proxy:v1
docker run -d -p 8080:8080 reverse-proxy:v1
```