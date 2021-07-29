# grpc-with-rest

This is a gRPC app build with Go, gRPC-Gateway and ProtoBuff using HTTP as transport layer.
It was build with the purpose of having both gRPC and RESTful style integrated.

grpc-with-rest performs a transcoding of HTTP calls to gRPC using a proxy server.

The server and client side is all defined in _/proto/user.proto._

## Pre-requisites 

### 1. Go
~~~~
[https://golang.org/], any one of the three latest major releases of Go.
For installation instructions, see Go's getting started guide: https://golang.org/doc/install
~~~~

### 2. PostgreSQL
~~~~
[Version 10+]
For installation instructions, please refer to this link: https://www.postgresql.org/download/
~~~~


## How to Run it?

**1. Clone the repo in your local environment:**

~~~~
git clone https://github.com/flakrimjusufi/grpc-with-rest.git
~~~~

**2. Populate environment variables in .env file:**

~~~~
db_name = (PostgreSQL database name, for example: testdb)
db_pass = (PostgreSQL database password, for example: 123456)
db_user = (PostgreSQL database user, for example: testuser)
db_type = (PostgreSQL database type, for example: postgres)
db_host = (PostgreSQL database host, for example: localhost)
db_port = (PostgreSQL database port, for example: 5434)
server_host = (The server in which you will run the app, for example: 0.0.0.0)
server_port = (The port in which you will run the server, for example: :8090) 
~~~~

**3. Run the server first:**

`go run server/main.go`

~~~~
You should recieve a response:
Serving gRPC on 0.0.0.0:8080
Serving gRPC-Gateway on 0.0.0.0:8090
~~~~

**4. Send a POST request using cURL:**

`curl -X POST -k http://localhost:8090/v1/example/echo -d '{"name": "Flakrim"}'`

~~~~
You should have a response from server: 

{
  "message": "Hello Flakrim"
}
~~~~