VERSION=`cat VERSION.txt`

compose-up:
	docker-compose up -d

integration:
	SSL_MODE=disable go test -v -tags=integration ./pkg/dba/...

docker-image:
	docker build --pull -t jonstacks/pg-dba:$(VERSION) \
			    -t jonstacks/pg-dba:latest \
			    --build-arg PG_DBA_VERSION=$(VERSION) .
