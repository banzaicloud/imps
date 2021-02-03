# Copyright (c) 2019 Banzai Cloud Zrt. All Rights Reserved.

# Build the manager binary
FROM golang:1.13 as builder

ARG GITHUB_ACCESS_TOKEN

ARG GOPROXY
ARG GOPRIVATE=github.com/banzaicloud

ENV GOFLAGS="-mod=readonly"

WORKDIR /workspace/services/imps
# Copy the Go Modules manifests

COPY ./services/imps/go.mod .
COPY ./services/imps/go.sum .
COPY ./services/sre/go.mod .
COPY ./services/sre/go.sum .
COPY ./go.mod /workspace/
COPY ./go.sum /workspace/
COPY ./pkg/platform/log/go.mod /workspace/pkg/platform/log/
COPY ./pkg/platform/log/go.sum /workspace/pkg/platform/log/
COPY ./pkg/common/go.mod /workspace/pkg/common/
COPY ./pkg/common/go.sum /workspace/pkg/common/
COPY ./pkg/prometheus/go.mod /workspace/pkg/prometheus/
COPY ./pkg/prometheus/go.sum /workspace/pkg/prometheus/
COPY ./pkg/template/go.mod /workspace/pkg/template/
COPY ./pkg/template/go.sum /workspace/pkg/template/
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
