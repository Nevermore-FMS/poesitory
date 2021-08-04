FROM golang:1.16
WORKDIR /go/src/app

# Fetch dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY . .
RUN go install

EXPOSE 8080

ENTRYPOINT ["backend"]
