# Go API Server for svcmngserver

NRF NFManagement Service

## Overview
This server was generated by the [openapi-generator]
(https://openapi-generator.tech) project.
By using the [OpenAPI-Spec](https://github.com/OAI/OpenAPI-Specification) from a remote server, you can easily generate a server stub.  
-

To see how to make this your own, look here:

[README](https://openapi-generator.tech)

- API version: 1.0.0
- Build date: 2019-03-28T23:13:54.578+08:00[Asia/Shanghai]


### Running the server
To run the server, follow these simple steps:

```
go run main.go
```

To run the server in a docker container
```
docker build --network=host -t svcmngserver .
```

Once image is built use
```
docker run --rm -it svcmngserver 
```

