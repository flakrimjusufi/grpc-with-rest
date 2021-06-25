# grpc-with-rest

This is just a simple gRPC app build with Go, gRPC-Gateway and ProtoBuff using HTTP as transport layer.

The route and client side is all defined in _/proto/helloworld.proto._

Please make sure you have Go installed if you want to give it a try:

~~~~
Go [https://golang.org/], any one of the three latest major releases of Go.
For installation instructions, see Go's getting started guide: https://golang.org/doc/install
~~~~

## How to Run it?

**1. Clone the repo in your local environment:**

~~~~
git clone https://github.com/flakrimjusufi/grpc-with-rest.git
~~~~


**2. Run the server first:**

`go run server/main.go`

~~~~
You should get a response:
Serving gRPC on 0.0.0.0:8080
Serving gRPC-Gateway on http://0.0.0.0:8090
~~~~

**3. Send a POST request using cURL:**

`curl -X POST -k http://localhost:8090/v1/example/echo -d '{"name": "Flakrim"}'`

~~~~
You should have a response from server: 

{
  "message": "Hello Flakrim"
}
~~~~