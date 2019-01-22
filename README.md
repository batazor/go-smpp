# go-smpp

Smpp client (UI + OpenTracing + support K8S)

### Feature

+ gRPC API

### Getting start

<details><summary>CLICK ME</summary>
<p>
# Run in Kubernetes


```
helm install --name go-smpp ops/Helm/go-smpp
```
</p>

<p>

```
# Generate gRPC code
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
protoc --go_out=plugins=grpc:. pb/*.proto

# Run services
$ docker-compose up
```

</p>
</details>

### ENV

| Name                             | Default value                              |
|----------------------------------|--------------------------------------------|
| GRPC_ENABLE                      | true                                       |
| GRPC_PORT                        | "50051"                                    |
| SMPP_ENABLE                      | true                                       |
| MONGO_ENABLE                     | true                                       |
| MONGO_URL                        | "telemetry"                                |
| MONGO_DATABASE_SMPP              | "go-smpp-message"                          |
| MONGO_DATABASE_WORKER            | "go-smpp-worker"                           |
| GRPC_ENABLE                      | true                                       |
| GRPC_PORT                        | "50051"                                    |
