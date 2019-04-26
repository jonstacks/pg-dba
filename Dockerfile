FROM golang:1.12-alpine as builder
RUN apk add --no-cache git
ENV GO111MODULE=on
WORKDIR /pg-dba
COPY . .
RUN go get ./...
RUN CGO_ENABLED=0 go build -v -o /usr/local/bin/pg-dba ./cmd/pg-dba/...

FROM alpine
ARG PG_DBA_VERSION
ENV PG_DBA_VERSION=$PG_DBA_VERSION
COPY --from=builder /usr/local/bin/pg-dba /usr/local/bin/pg-dba
CMD ["pg-dba"]
