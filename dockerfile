############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

WORKDIR /go/src/ittech24/rest/apidemo

COPY . .

# Using go get.
RUN go get -d -v

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/apidemo
COPY demo.json /go/bin

############################
# STEP 2 build a small image
############################
FROM scratch

# Copy our static executable.
COPY --from=builder /go/bin/apidemo /go/bin/apidemo
COPY --from=builder /go/bin/demo.json /go/bin
# RUN chmod +x /go/bin/DeploymentTools
# RUN ls -la /go/bin

# Run the hello binary.
EXPOSE 80

ENTRYPOINT ["/go/bin/restapi-testapp-go"]
