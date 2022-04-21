FROM golang:1.18 AS builder

WORKDIR /app

COPY go.mod.dist go.mod
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go get nginx-error-server
RUN go build -o nginx-error-server

RUN cp -R $GOPATH/pkg/mod/github.com/\!gerald\!wodni/kern*/default .
COPY content ./


FROM golang:1.18-alpine

WORKDIR /app

COPY --from=builder /app/* .

EXPOSE 5000

CMD [ nginx-error-server ]
