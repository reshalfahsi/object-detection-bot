FROM golang:1.17.0-alpine3.13

WORKDIR /work

ADD . .

RUN go mod init main
RUN go mod tidy
RUN go build -o /bin/jwt-auth-api .

WORKDIR /

RUN rm -r /work

ENTRYPOINT ["/bin/jwt-auth-api"]
