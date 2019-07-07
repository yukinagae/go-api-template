# Use the offical Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.11 as builder

# Copy local code to the container image.
WORKDIR /go/src/github.com/knative/docs/go-api-template
COPY . .

# Setup env variables
ENV GOBIN $(pwd)/bin
ENV GOPATH $HOME/go
ENV PATH $PATH:$GOBIN
ENV GO111MODULE on

# Move to `app` directory to build because `main.go` is insided the directory.
WORKDIR /go/src/github.com/knative/docs/go-api-template/app

RUN CGO_ENABLED=0 GOOS=linux go build -v -o main

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine

# Copy the binary to the production image from the builder stage.
COPY --from=builder /go/src/github.com/knative/docs/go-api-template/app/main /main

EXPOSE 8080

CMD ["/main"]
