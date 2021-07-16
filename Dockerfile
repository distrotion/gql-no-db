FROM golang:1.16-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the source code
COPY *.go .
RUN go get github.com/distrotion/gql-no-db/graph
RUN go get github.com/distrotion/gql-no-db/graph/generated
RUN go get github.com/distrotion/gql-no-db/internal/auth
# Build
RUN go build -o /gql-no-db

# This is for documentation purposes only.
# To actually open the port, runtime parameters
# must be supplied to the docker command.
EXPOSE 9100

# (Optional) environment variable that our dockerised
# application can make use of. The value of environment
# variables can also be set via parameters supplied
# to the docker command on the command line.
#ENV HTTP_PORT=8081

# Run
CMD [ "/gql-no-db" ]
