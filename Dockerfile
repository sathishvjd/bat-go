FROM golang:1.18-alpine as builder


# put certs in builder image
RUN apk update
RUN apk add -U --no-cache ca-certificates && update-ca-certificates
RUN apk add make
RUN apk add build-base
RUN apk add git
RUN apk add bash

ARG VERSION
ARG BUILD_TIME
ARG COMMIT

WORKDIR /src
COPY . ./
RUN chown -R nobody:nobody /src/
RUN mkdir /.cache
RUN chown -R nobody:nobody /.cache

USER nobody
RUN cd main && go mod download

RUN cd main && CGO_ENABLED=0 GOOS=linux go build \
    -ldflags "-w -s -X main.version=${VERSION} -X main.buildTime=${BUILD_TIME} -X main.commit=${COMMIT}" \
    -o bat-go main.go

FROM alpine:3.16 as base
# put certs in artifact from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /src/main/bat-go /bin/

FROM base as payments
CMD ["bat-go", "serve", "nitro", "inside-enclave", "--log-address", "vm(3):2345", "--egress-address", "vm(3):1234", "--upstream-url", "http://0.0.0.0:8080", "--address", ":8080"]

FROM base as artifact
COPY --from=builder /src/migrations/ /migrations/
EXPOSE 3333
CMD ["bat-go", "serve", "grant", "--enable-job-workers", "true"]

