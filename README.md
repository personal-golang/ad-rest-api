# ad-rest-api

## Pre-requisites

1. `docker pull dwimberger/ldap-ad-it`

## Running locally

In a terminal window: 
1. `make start-ad`

In a new terminal window: 
2. `go mod vendor`
3. `gow -c run cmd/rest-server/main.go`