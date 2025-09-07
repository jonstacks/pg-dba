FROM golang:1.24.6-alpine AS builder
RUN apk add --no-cache git
WORKDIR /pg-dba
COPY . .
RUN CGO_ENABLED=0 go build -v -o /usr/local/bin/pg-dba ./cmd/pg-dba/...

FROM alpine
ARG PG_DBA_VERSION
ENV PG_DBA_VERSION=$PG_DBA_VERSION
COPY --from=builder /usr/local/bin/pg-dba /usr/local/bin/pg-dba
CMD ["pg-dba"]
