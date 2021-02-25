# Copyright (c) 2019 Banzai Cloud Zrt. All Rights Reserved.

# Build the manager binary
FROM golang:1.15 as builder

ARG GOPROXY

ENV GOFLAGS="-mod=readonly"

WORKDIR /workspace/
# Copy the Go Modules manifests

COPY ./go.mod /workspace/
COPY ./go.sum /workspace/
RUN go mod download

COPY ./ /workspace/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o manager main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/manager .
USER nonroot:nonroot

ENTRYPOINT ["/manager"]
