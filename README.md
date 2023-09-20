
# Movie Title Hashing

Solution to the Redbrain team coding challenge.

## Task:

```
We would like you to build two services which together will allow the User to submit the title of any movie and receive back the sha256 hash of the title they have provided.

For security, the User must make the request to an api-service which connects to a separate crypo-service to generate the hash for the movie title.

Communication between the User and the api-service should be RESTful.Communication between the api-service and the crypto-service should be done using Protocol Buffers.
```

## Solution

Two services are implemented as part of the solution:

* `gateway-service` - This is a REST service that handles users requests for hashing movie titles. It is accessible on port `8080` when run locally
* `encryption-service` - This is a gRPC service that provides api for hashing input data. Proto definition can be found: `api/proto/encryptor.proto`. This service is accessible on port `9001` when run locally

_NOTE: For simplicity I kept both services in the same repo. Two different entry points are:_
* `cmd/gatewayservice/main.go` - Sets up the gateway service
* `cmd/encryptionservice/main.go` - Sets up the gRPC encryption service

## How to run services:

There is a Makefile targe that can be executed to run both services locally:

`make start`

This will build `Docker` images for both services and run them locally with `docker-compose`

If you want to stop both services:

`make stop`


## How to test:

Once services are running locally you can invoke the API like:

```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"movie":"Titanic"}' \
  http://localhost:8080/hash-movie-name
```