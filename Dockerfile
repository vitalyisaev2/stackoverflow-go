FROM golang:1.17

WORKDIR /usr/src/stackoverflow-go

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go test -v -c . -o /usr/local/bin/stackoverflow-go

CMD ["stackoverflow-go"]