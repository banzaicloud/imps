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


FROM redhat/ubi8-micro:8.7@sha256:6a56010de933f172b195a1a575855d37b70a4968be8edb35157f6ca193969ad2 AS ubi8

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/local/src/imps/LICENSE.md /usr/local/src/imps/LICENSE.md

COPY --from=builder /usr/local/bin/manager /usr/local/bin/manager

USER nobody:nobody

ENTRYPOINT ["manager"]


FROM gcr.io/distroless/base-debian11:latest@sha256:e711a716d8b7fe9c4f7bbf1477e8e6b451619fcae0bc94fdf6109d490bf6cea0 AS distroless

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/local/src/imps/LICENSE.md /usr/local/src/imps/LICENSE.md

COPY --from=builder /usr/local/bin/manager /usr/local/bin/manager

USER nobody:nobody

ENTRYPOINT ["manager"]
