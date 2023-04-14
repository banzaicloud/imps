ARG FROM_IMAGE=scratch

ARG GID=1000
ARG UID=1000

FROM --platform=$BUILDPLATFORM tonistiigi/xx:1.2.1@sha256:8879a398dedf0aadaacfbd332b29ff2f84bc39ae6d4e9c0a1109db27ac5ba012 AS xx

FROM --platform=$BUILDPLATFORM golang:1.20.3-alpine3.16@sha256:29c4e6e307eac79e5db29a261b243f27ffe0563fa1767e8d9a6407657c9a5f08 AS builder

COPY --from=xx / /

RUN apk add --update --no-cache ca-certificates make git curl tzdata clang lld

ARG TARGETPLATFORM

RUN xx-apk --update --no-cache add musl-dev gcc

RUN xx-go --wrap

WORKDIR /usr/local/src/imps

ARG GOPROXY

ENV CGO_ENABLED=0

COPY go.mod go.sum ./
COPY api/go.mod api/go.sum ./api/
RUN go mod download

COPY . .

RUN go build -o /usr/local/bin/manager ./cmd/controller/
RUN xx-verify /usr/local/bin/manager


FROM alpine:3.17.3@sha256:124c7d2707904eea7431fffe91522a01e5a861a624ee31d03372cc1d138a3126 AS user

ARG GID
ARG UID

RUN addgroup -g ${GID} -S appgroup && adduser -u ${UID} -S appuser -G appgroup


FROM ${FROM_IMAGE}

ARG GID
ARG UID

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /usr/local/bin/manager /usr/local/bin/manager

COPY --from=user /etc/passwd /etc/passwd
COPY --from=user /etc/group /etc/group
USER ${UID}:${GID}

ENTRYPOINT ["manager"]
