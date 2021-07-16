VERSION=`cat VERSION.txt`

clean:
	rm -rf coverage.txt
	rm -rf integration-coverage.txt

compose-up:
	docker-compose up -d

integration:
	SSL_MODE=disable go test -v -tags=integration ./pkg/dba/...

docker-image:
	docker build --pull -t jonstacks/pg-dba:$(VERSION) \
			    -t jonstacks/pg-dba:latest \
			    --build-arg PG_DBA_VERSION=$(VERSION) .

# Run the unit and integration tests and combine them.
test: unit-tests integration-tests
	gocovmerge coverage.txt integration-coverage.txt > coverage.txt

unit-tests:
	go test -v -coverprofile=coverage.txt -covermode=atomic ./...

integration-tests:
	SSL_MODE=disable go test -v -tags=integration -coverprofile=integration-coverage.txt -covermode=atomic ./...

doc-server:
	$(MAKE) -C docs doc-server

kill-doc-server:
	$(MAKE) -C docs kill-doc-server
