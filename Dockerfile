FROM golang:alpine as builder
WORKDIR /go/src/github.com/jonstacks/pg-dba
RUN apk add git && \
    go get -u github.com/golang/dep/cmd/dep
COPY Gopkg.lock Gopkg.toml ./
RUN dep ensure -v --vendor-only
COPY . ./
RUN CGO_ENABLED=0 go build -v -o /usr/local/bin/pg-dba ./cmd/pg-dba/...

FROM alpine
ARG PG_DBA_VERSION
ENV PG_DBA_VERSION=$PG_DBA_VERSION
COPY --from=builder /usr/local/bin/pg-dba /usr/local/bin/pg-dba
CMD ["pg-dba"]
