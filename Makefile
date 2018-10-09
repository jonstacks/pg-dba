
compose-up:
	docker-compose up -d

integration:
	SSL_MODE=disable go test -v -tags=integration ./pkg/dba/...
