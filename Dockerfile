FROM fedora:latest

RUN dnf install -y golang python3-pip

WORKDIR /usr/src/stackoverflow-go

COPY go.mod go.sum requirements.txt ./
RUN go mod download && go mod verify && pip install -r requirements.txt

COPY . .
RUN go test -v -c . -o /usr/local/bin/stackoverflow-go

CMD ["./script.sh"]