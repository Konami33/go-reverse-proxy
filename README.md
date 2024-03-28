### Run locally:
1. Run each server individually
```bash
go run main.go
```
<\br>
### Run using Docker file:
1. Docker build:
```bash
docker build -t origin:v1
docker build -t origin2:v1
docker build -t reverse-proxy:v1
```
2. Run docker images:
```bash
docker run -d -p 8081:8081 origin:v1
docker run -d -p 8082:8082 origin2:v2
```