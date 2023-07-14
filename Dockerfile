ARG UID=1000
ARG GID=1000

FROM --platform=$BUILDPLATFORM tonistiigi/xx:1.2.1@sha256:8879a398dedf0aadaacfbd332b29ff2f84bc39ae6d4e9c0a1109db27ac5ba012 AS xx

FROM --platform=$BUILDPLATFORM golang:1.20.4-alpine3.16@sha256:6469405d7297f82d56195c90a3270b0806ef4bd897aa0628477d9959ab97a577 AS builder
ARG UID
ARG GID

# Create user and group
RUN addgroup -g ${GID} -S appgroup
RUN adduser -u ${UID} -S appuser -G appgroup

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
ARG UID
ARG GID

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# RedHat certification requires application license to be in /licenses dir
COPY --from=builder /usr/local/src/imps/LICENSE /licenses/LICENSE
COPY --from=builder /usr/local/bin/manager /manager

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
USER ${UID}:${GID}

ENTRYPOINT ["/manager"]


FROM gcr.io/distroless/base-debian11:latest@sha256:e711a716d8b7fe9c4f7bbf1477e8e6b451619fcae0bc94fdf6109d490bf6cea0 AS distroless
ARG UID
ARG GID

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# RedHat certification requires application license to be in /licenses dir
COPY --from=builder /usr/local/src/imps/LICENSE /licenses/LICENSE
COPY --from=builder /usr/local/bin/manager /manager

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
USER ${UID}:${GID}

ENTRYPOINT ["/manager"]
