# grpc-with-rest

This is a gRPC app build with Go, [gRPC-Gateway](https://github.com/grpc-ecosystem/grpc-gateway)
and [ProtoBuff](https://developers.google.com/protocol-buffers) using HTTP as transport layer. It was build with the
purpose of having both gRPC and RESTful style integrated.

grpc-with-rest performs a [transcoding of HTTP calls to gRPC](https://cloud.google.com/endpoints/docs/grpc/transcoding)
using a proxy server.

The server and client side is all defined in _/proto/user.proto._

Basically, what this app does is described in the picture below:

![](https://github.com/flakrimjusufi/grpc-with-rest/blob/develop/images/architecture_introduction_diagram.jpg)

## Pre-requisites

### 1. Go

~~~~
[https://golang.org/], any one of the three latest major releases of Go.
For installation instructions, see Go's getting started guide: https://golang.org/doc/install
~~~~

### 2. PostgreSQL (or any SQL database)

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

| HTTP call        | Endpoint           | Description  |
| :-------------: |:-------------:| :-----:|
| POST     | http://localhost:8090/user/create | Will create a user in database |
| GET      | http://localhost:8090/user/findByName/Flakrim      |  Will find a user by name in database |
| GET | http://localhost:8090/user/findById/120023      |   Will find a user by Id in database |
| GET | http://localhost:8090/user/list     |   Will find all the user in database |
| POST | http://localhost:8090/user/updateById/12003     |   Will update a user by Id in database |
| POST | http://localhost:8090/user/delete     |   Will delete a user by name in database |
| GET | http://localhost:8090/card/listCreditCards     |   Will list all credit cards in database |
| GET | http://localhost:8090/card/findByUserName/Flakrim     |   Will find a credit card by user name in database |
| POST | http://localhost:8090/card/createCreditCardApplication | Will create a credit card application |
| GET | http://localhost:8090/card/getCreditCardApplication/Flakrim | Will find a credit card application by user first name |

## User Interface

grpc-with-rest provides a User-interface to create a credit card application.

User-interface is a multi-step react integrated with gRPC services. We are performing a transcoding of HTTP over gRPC
and with this form we are sending a POST request [http://localhost:8090/card/createCreditCardApplication] to the server.

### To run the multi-step form, cd to ui directory and execute the following command:

#### `yarn start` or `npm install` 

~~~
Installs all the dependencies and runs the app in the development mode.
~~~

Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

Note that in order to interact with the server, you should also start up the gRPC server once the UI is up and running:

`go run server/main.go`

### Multi-Step form:

![](https://github.com/flakrimjusufi/grpc-with-rest/blob/develop/images/multi-step-form.png)

