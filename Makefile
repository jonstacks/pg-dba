VERSION=`cat VERSION.txt`

compose-up:
	docker-compose up -d

integration:
	SSL_MODE=disable go test -v -tags=integration ./pkg/dba/...

docker-image:
	docker build --pull -t jonstacks/pg-dba:$(VERSION) \
			    -t jonstacks/pg-dba:latest \
			    --build-arg PG_DBA_VERSION=$(VERSION) .

unit-tests:
	go test -v -coverprofile=coverage.txt -covermode=atomic ./...

integration-tests:
	go test -v -tags=integration -coverprofile=coverage.txt -covermode=atomic ./...

doc-server:
	$(MAKE) -C docs doc-server

kill-doc-server:
	$(MAKE) -C docs kill-doc-server
