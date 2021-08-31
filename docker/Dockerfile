FROM golang:alpine

WORKDIR /app

COPY ../cli/poesitory/go.mod ../cli/poesitory/go.sum ./

RUN go mod download

COPY ../cli/poesitory/* ./

RUN go install

CMD [ "poesitory" ]