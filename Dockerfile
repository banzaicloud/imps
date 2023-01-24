# Copyright (c) 2019 Banzai Cloud Zrt. All Rights Reserved.
ARG FROM_IMAGE=scratch
ARG GO_VERSION=1.17
ARG GID=1000
ARG UID=1000

# Build the manager binary
FROM golang:${GO_VERSION}-alpine3.14 as builder
ARG GID
ARG UID

# Create user and group
RUN addgroup -g ${GID} -S appgroup && \
    adduser -u ${UID} -S appuser -G appgroup

# set up nsswitch.conf for Go's "netgo" implementation
# https://github.com/gliderlabs/docker-alpine/issues/367#issuecomment-424546457
RUN echo 'hosts: files dns' > /etc/nsswitch.conf.build

RUN apk add --update --no-cache make bash curl ca-certificates git tzdata

ARG GOPROXY
ENV GOFLAGS="-mod=readonly"

WORKDIR /workspace/
# Copy the Go Modules manifests

COPY ./go.mod /workspace/
COPY ./go.sum /workspace/
COPY ./api/ /workspace/api/
RUN go mod download

COPY ./ /workspace/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o manager ./cmd/controller/

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM ${FROM_IMAGE}
ARG GID
ARG UID

WORKDIR /

COPY --from=builder /etc/nsswitch.conf.build /etc/nsswitch.conf
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /workspace/manager .

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
USER ${UID}:${GID}

ENTRYPOINT ["/manager"]
