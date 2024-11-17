FROM --platform=$BUILDPLATFORM golang:1.23.3 AS build

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.sum ./

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY cmd ./cmd
COPY pkg ./pkg
COPY redisdb ./redisdb
COPY server ./server
COPY plugins ./plugins
COPY gate.go ./

# Automatically provided by the buildkit
ARG TARGETOS TARGETARCH

# Build
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
    go build -ldflags="-s -w" -a -o gate gate.go

# Final image stage (debug version)
FROM --platform=$BUILDPLATFORM debian:11 AS app
COPY --from=build /workspace/gate /gate
# Add the Gate configuration file
COPY config.yml /config.yml
COPY servers.json /servers.json
RUN apt-get update && apt-get install -y bash && apt-get install -y ca-certificates && update-ca-certificates

CMD ["/gate"]
