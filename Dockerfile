# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build-env

ADD . /app
WORKDIR /app

# Compile the application with the optimizations turned off
# This is important for the debugger to correctly work with the binary
RUN go build -gcflags "all=-N -l" -o /main

##
## Deploy
##
FROM debian:buster

WORKDIR /
COPY --from=build-env /main /

CMD go build -o ./seeds/migrate_models_in_database.go

EXPOSE 8090

ENTRYPOINT ["/main"]