FROM fedora:latest

RUN dnf install -y golang

WORKDIR /usr/src/stackoverflow-go

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go test -v -c . -o /usr/local/bin/stackoverflow-go

CMD ["./script.sh"]