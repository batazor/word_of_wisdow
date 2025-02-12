# syntax=docker/dockerfile:1.5

# Link: https://github.com/moby/buildkit/blob/master/docs/attestations/sbom.md
# enable scanning for the intermediate build stage
ARG BUILDKIT_SBOM_SCAN_STAGE=true
# scan the build context only if the build is run to completion
ARG BUILDKIT_SBOM_SCAN_CONTEXT=true

FROM --platform=$BUILDPLATFORM golang:1.24-alpine AS builder

ARG CI_COMMIT_TAG

ENV GOEXPERIMENT=arenas,cgocheck2,loopvar
ENV PGO_PATH=auto

WORKDIR /go/github.com/batazor/word_of_wisdom

# Load dependencies
COPY go.mod go.sum ./
RUN go mod download

# COPY the source code AS the last step
COPY . .

# Build project
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
  go build \
  -a \
  -pgo=${PGO_PATH} \
  -ldflags "-s -w -X main.CI_COMMIT_TAG=$CI_COMMIT_TAG" \
  -installsuffix cgo \
  -trimpath \
  -o app cmd/client/client.go

FROM alpine:3.18

LABEL maintainer=batazor111@gmail.com
LABEL org.opencontainers.image.title="Word of wisdom - client"
LABEL org.opencontainers.image.description="Word of wisdom - client"
LABEL org.opencontainers.image.authors="@batazor"
LABEL org.opencontainers.image.vendor="@batazor"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.source="https://github.com/batazor/word_of_wisdom"
LABEL org.opencontainers.image.revision=$CI_COMMIT_SHA

# Install dependencies
RUN \
  apk update && \
  apk add --no-cache curl tini

RUN addgroup -S golang && adduser -S -g golang golang
USER golang

ENTRYPOINT ["/sbin/tini", "--"]

WORKDIR /app/
CMD ["./app"]
COPY --from=builder /go/github.com/batazor/word_of_wisdom/app /app
