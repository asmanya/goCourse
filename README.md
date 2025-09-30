# gRPC, ProtoBuf, REST, Serialization & Advanced Golang Concepts

A collection of **hands-on Go exercises** exploring:

- gRPC services and streaming  
- REST API development and benchmarking  
- Serialization & deserialization in Go  
- Advanced Go concepts: concurrency, channels, mutexes, goroutines  
- Protobuf practice and code generation  

---

## **About this Project**
This repository demonstrates my journey from Go basics to advanced topics, including building client-server applications, experimenting with HTTP/REST/gRPC, and mastering concurrency and structured data handling.

---

## **Repository Structure**

| Folder | Description |
|--------|-------------|
| **BenchmarkingHTTP1vsHTTP2** | Compare HTTP/1.1 and HTTP/2 performance with Go servers and clients. <br> - `http1_api/` – HTTP/1.1 server & benchmarking.txt <br> - `http2_api/` – HTTP/2 server with TLS & benchmarking.txt |
| **REST Practice** | REST API exercises. <br> - `httpClient/` – HTTP client example <br> - `httpServer/` – Server & benchmarking <br> - `simple_api/` – REST API with TLS and openssl config |
| **Serialize-Deserialize** | JSON serialization and deserialization examples (`serialize_deserializeJSON.go`) |
| **gRPC_practice** | gRPC hands-on projects. <br> - `gRPCserver/` – gRPC server with multiple proto services & TLS <br> - `gRPCclient/` – gRPC clients <br> - `gRPC_stream_server/` & `gRPC_stream_client/` – Streaming RPC examples <br> - `gRPC_Gateway_RESTcomboapi/` – gRPC with REST gateway |
| **go Programming** | Learning notes & exercises: <br> - `basics/` – Variables, loops, functions, etc. <br> - `intermediate/` – File handling, interfaces, JSON, logging <br> - `advanced/` – Concurrency, goroutines, channels, rate limiting <br> - `moreAboutConcurrency/` – Deadlocks, RWMutex, race conditions |
| **modernRoute** | `main.go` – Modern routing and Go HTTP server practice |
| **proto-buf-practice** | Protobuf exercises & generated Go code (`proto/` folder with `.proto` and `.pb.go` files) |

---

Udemy Course - NEW-Comprehensive Go Bootcamp with gRPC and Protocol Buffers by Ashish Sharma