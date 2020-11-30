start-ad:
	docker stop $$(docker ps -q --filter ancestor=dwimberger/ldap-ad-it) || true
	docker run -it --rm -p 10389:10389 dwimberger/ldap-ad-it

test:
	go test -v ./pkg/...

test-cover:
	go test -coverprofile=cover.out ./pkg/... && go tool cover -func=cover.out

start:
	go mod vendor
	gow -c run cmd/rest-server/main.go

.PHONY: start-ad
