# Run this with docker build --build-arg builder_image=<golang:x.y.z-alpine>
ARG builder_image
FROM ${builder_image} as builder
WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY ./ ./
RUN CGO_ENABLED=0 go build -o manager main.go

# runtime image
FROM alpine:3
WORKDIR /
COPY --from=builder /workspace/manager .

ENTRYPOINT ["/manager"]
