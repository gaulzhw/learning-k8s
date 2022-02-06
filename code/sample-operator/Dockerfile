# Build the manager binary
# Run this with docker build --build-arg builder_image=<golang:x.y.z-alpine>
ARG builder_image
FROM ${builder_image} as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY ./ ./

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o manager main.go

FROM alpine:3
WORKDIR /
COPY --from=builder /workspace/manager .

ENTRYPOINT ["/manager"]
