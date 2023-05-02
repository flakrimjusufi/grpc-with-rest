# grpc-with-rest

This is a gRPC app build with Go, [gRPC-Gateway](https://github.com/grpc-ecosystem/grpc-gateway)
and [ProtoBuff](https://developers.google.com/protocol-buffers) using HTTP as transport layer. It was build with the
purpose of having both gRPC and REST-ful style integrated and can be used as a template to start a project in gRPC 
with a SQL database and a UI integrated.


grpc-with-rest performs a [transcoding of HTTP calls to gRPC](https://cloud.google.com/endpoints/docs/grpc/transcoding)
using a proxy server.

The server and client side is all defined in _/proto/user.proto._

Basically, what this app does is described in the picture below:

![](https://github.com/flakrimjusufi/grpc-with-rest/blob/develop/images/architecture_introduction_diagram.jpg)

## How to Run it?

### Docker 

**Clone the repo in your local environment:**

~~~~
git clone https://github.com/flakrimjusufi/grpc-with-rest.git
~~~~

*Remove ".example" from both .env.example and /seeds/init.sql.example and populate them with your environment variables.*


**In case you have docker-compose installed in your machine, just execute the following:**

~~~~
docker-compose up 
~~~~

or 

~~~
make docker-up 
~~~

Docker-compose will build all the dependencies and will add a PostgresSQL image in your container alongside
with the server so that we can interact with data.

*Once the docker-compose is finished, you should see an output in terminal:*

~~~
Serving gRPC on localhost:8080
Serving gRPC-Gateway on localhost:8090
~~~

*Send a POST request using cURL:*

`curl -X POST -k http://localhost:8090/api/v1/example/echo -d '{"name": "Flakrim"}'`

~~~~
You should have a response from server: 

{
  "message": "Hello Flakrim"
}
~~~~

### Bazel 

#### To set up the project with Bazel

~~~
make bazel-setup
~~~

#### To run the project with Bazel

~~~
make bazel-run 
~~~

#### In case you don't have docker or bazel installed, you need to do the following:

## Pre-requisites

### 1. Go

~~~~
[https://golang.org/], any one of the three latest major releases of Go (1.18 was used in this project) 
For installation instructions, see Go's getting started guide: https://golang.org/doc/install
~~~~

### 2. PostgresSQL (or any SQL database)

~~~~
[Version 9+]
For installation instructions, please refer to this link: https://www.postgresql.org/download/
~~~~

## How to Run it?

**1. Clone the repo in your local environment:**

~~~~
git clone https://github.com/flakrimjusufi/grpc-with-rest.git
~~~~

**2. Populate environment variables in .env file:**

~~~~
DB_USERNAME=
DB_DATABASE=
DB_HOSTNAME=
DB_PORT=
DB_TYPE=
POSTGRES_PASSWORD=
SERVER_HOST=
GRPC_SERVER_PORT=
GRPC_GATEWAY_SERVER_PORT=
~~~~

**.env file example**
~~~~
DB_USERNAME = (PostgreSQL database user, for example: testuser)
DB_DATABASE = (PostgreSQL database name, for example: testdb)
DB_HOSTNAME = (PostgreSQL database host, for example: localhost)
DB_TYPE = (PostgreSQL database type, for example: postgres)
DB_PORT = (PostgreSQL database port, for example: 5432)
POSTGRES_PASSWORD = (PostgreSQL database password, for example: 123456)
SERVER_HOST = (The server in which you will run the app, for example: 0.0.0.0)
GRPC_SERVER_PORT = (The port in which you will run the app, for example: 8080
GRPC_GATEWAY_SERVER_PORT = (The port in which you will run the server, for example: 8090) 
~~~~

**3. Run the server first:**

`go run server/main.go`

~~~~
You should recieve a response:
Serving gRPC on 0.0.0.0:8080
Serving gRPC-Gateway on 0.0.0.0:8090
~~~~

**4. Send a POST request using cURL:**

`curl -X POST -k http://localhost:8090/api/v1/example/echo -d '{"name": "Flakrim"}'`

~~~~
You should have a response from server: 

{
  "message": "Hello Flakrim"
}
~~~~

## Populate the database

### To interact with data, we are using:

- **[gORM](https://gorm.io/)** - to interact with our database.
- **[gofakeit](https://github.com/brianvoe/gofakeit)** - to populate our database.

There are three data models which we are using in this app:

- users (which can be found
  under[ /models/data_struct.go](https://github.com/flakrimjusufi/grpc-with-rest/blob/develop/models/data_struct.go))
- credit_cards (which can also be found
  under [ /models/data_struct.go](https://github.com/flakrimjusufi/grpc-with-rest/blob/develop/models/data_struct.go))
- credit_card_application (which is also located
  under [/models/data_struct.go ](https://github.com/flakrimjusufi/grpc-with-rest/blob/develop/models/data_struct.go))

### There are three seed scripts for populating the database:

- [create_fake_users.go](https://github.com/flakrimjusufi/grpc-with-rest/blob/develop/seeds/create_fake_users.go)
- [create_fake_credit_cards.go](https://github.com/flakrimjusufi/grpc-with-rest/blob/develop/seeds/create_fake_credit_cards.go)
- [create_fake_cc_applications.go](https://github.com/flakrimjusufi/grpc-with-rest/blob/develop/seeds/create_fake_cc_applications.go)

#### To execute the script which populates the database with fake users, run this command:

`go run seeds/create_fake_users.go`

~~~
** This command will auto-migrate user data struct and will create a table in database named users, afterwards will populate the table with fake users. 
~~~

#### To execute the script which populates the database with fake credit cards, run this command:

`go run seeds/create_fake_credit_cards.go`

~~~
** This command will auto-migrate credit_card data struct and will create a table in database named creidt_cards, afterwards will populate the table with fake credit cards. 
~~~

#### To execute the script which populates the database with fake credit card applications, run this command:

`go run seeds/create_fake_cc_application.go`

~~~
** This command will auto-migrate credit_card_application data struct and will create a table in database named credit_card_applications, afterwards will populate the table with fake credit card applications. 
~~~

## Available HTTP endpoints

After populating the database, you can use the below es to hit the reverse proxy server and to get the data from the
server:

| HTTP call |                       Endpoint                        |                      Description                       |
|:---------:|:-----------------------------------------------------:|:------------------------------------------------------:|
|   POST    |           http://localhost:8090/api/v1/user           |             Will create a user in database             |
|    GET    |     http://localhost:8090/api/v1/user/name/:name      |          Will find a user by name in database          |
|    GET    |       http://localhost:8090/api/v1/user/id/:id        |           Will find a user by Id in database           |
|    GET    |           http://localhost:8090/api/v1/user           |           Will find all the user in database           |
|    PUT    |       http://localhost:8090/api/v1/user/id/:id        |          Will update a user by Id in database          |
|    PUT    |     http://localhost:8090/api/v1/user/name/:name      |         Will update a user by name in database         |
|  DELETE   |     http://localhost:8090/api/v1/user/name/:name      |         Will delete a user by name in database         |
|    GET    |           http://localhost:8090/api/v1/card           |         Will list all credit cards in database         |
|    GET    |    http://localhost:8090/api/v1/card/name/Flakrim     |    Will find a credit card by user name in database    |
|   POST    |           http://localhost:8090/api/v1/card           |         Will create a credit card application          |
|    GET    | http://localhost:8090/api/v1/card/application/Flakrim | Will find a credit card application by user first name |

## User Interface

grpc-with-rest provides a User-interface to create a credit card application.

User-interface is a multistep React integrated with gRPC services. We are performing a transcoding of HTTP over gRPC
and with this form we are sending a POST request [http://localhost:8090/card/createCreditCardApplication] to the server.

### To run the multistep form, cd to ui directory and execute the following command:

#### `yarn start` or `npm install` 

~~~
Installs all the dependencies and runs the app in the development mode.
~~~

Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

Note that in order to interact with the server, you should also start up the gRPC server once the UI is up and running:

`go run server/main.go`

### Multi-Step form:

![](https://github.com/flakrimjusufi/grpc-with-rest/blob/develop/images/multi-step-form.png)

**[Watch a Demo Video](https://youtu.be/gIiTUbvQRzw)**

[![DEMO](https://github.com/flakrimjusufi/grpc-with-rest/blob/develop/images/demo-screen-shoot.png)](https://youtu.be/gIiTUbvQRzw)


### Acknowledgments

Many thanks to **[JetBrains](https://www.jetbrains.com/)** for their help and their support on this project.

Using their fantastic tools, we were able to write code, debug, test, document, build and ship quickly.

We totally recommend all of their tools. 

Here are the ones that were used to build this project: 

- **[GoLand](https://www.jetbrains.com/go/)** 
- **[DataGrip](https://www.jetbrains.com/datagrip/)**