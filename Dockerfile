# prod version
FROM golang:alpine as builder
RUN apk add --no-cache openssh-client git

# configuring dirs
WORKDIR /build
ADD src .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-extldflags "-static"' -o /main ./cmd/app/main.go

# resulted image
FROM alpine:latest
COPY --from=builder main /bin/main
ENTRYPOINT ["/bin/main"]
