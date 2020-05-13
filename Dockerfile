FROM golang:alpine AS builder

# Move to working directory /build
WORKDIR /build

# use modules
COPY go.mod .

ENV GO111MODULE=on
RUN go mod download
RUN go mod verify

COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o swe .

FROM scratch

# Copy our static executable
COPY --from=builder /build/swe /

# Run the hello binary.
ENTRYPOINT ["/swe"]
