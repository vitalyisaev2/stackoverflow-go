FROM fedora:latest

ADD stackoverflow-go /usr/local/bin

WORKDIR /usr/local/bin

RUN ls -l .

CMD ["./stackoverflow-go"]