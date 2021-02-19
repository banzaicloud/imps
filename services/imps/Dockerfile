# Copyright (c) 2019 Banzai Cloud Zrt. All Rights Reserved.

# Build the manager binary
FROM golang:1.13 as builder

ARG GITHUB_ACCESS_TOKEN

ARG GOPROXY
ARG GOPRIVATE=github.com/banzaicloud

ENV GOFLAGS="-mod=readonly"

WORKDIR /workspace/services/imps
# Copy the Go Modules manifests

COPY ./services/imps/go.mod /workspace/services/imps/
COPY ./services/imps/go.sum /workspace/services/imps/
COPY ./services/sre/go.mod /workspace/services/sre/
COPY ./services/sre/go.sum /workspace/services/sre/
COPY ./services/xrate/go.mod /workspace/services/xrate/
COPY ./services/xrate/go.sum /workspace/services/xrate/
COPY ./go.mod /workspace/
COPY ./go.sum /workspace/
COPY ./pkg/common/go.mod /workspace/pkg/common/
COPY ./pkg/common/go.sum /workspace/pkg/common/
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN git config --global url."https://${GITHUB_ACCESS_TOKEN}:@github.com/".insteadOf "https://github.com/"
RUN go mod download

COPY ./ /workspace/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o manager main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/services/imps/manager .
USER nonroot:nonroot

ENTRYPOINT ["/manager"]
