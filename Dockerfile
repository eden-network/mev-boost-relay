# syntax=docker/dockerfile:1
FROM golang:1.18 as builder
ARG VERSION
WORKDIR /build
ADD . /build/
RUN --mount=type=cache,target=/root/.cache/go-build make build

FROM alpine
RUN apk add --no-cache libstdc++ libc6-compat
WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/boost-relay /app/boost-relay
EXPOSE 9062
ENTRYPOINT ["/app/boost-relay"]
