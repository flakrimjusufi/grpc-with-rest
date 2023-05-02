# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.18-alpine AS build-stage

ADD . /app
WORKDIR /app

RUN go mod download

# Compile the application with the optimizations turned off
# This is important for the debugger to correctly work with the binary
RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -o /main

##
## Deploy
##
FROM gcr.io/distroless/base-debian10 AS build-release-stage

WORKDIR /
COPY --from=build-stage /main /

CMD go build -o ./seeds/migrate_models_in_database.go

EXPOSE 8090

USER nonroot:nonroot

ENTRYPOINT ["/main"]