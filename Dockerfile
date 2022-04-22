ARG REG_HOSTNAME
ARG REG_FOLDER
FROM ${REG_HOSTNAME}/${REG_FOLDER}/golang:1.18-alpine AS builder

RUN apk --no-cache add \
    git

WORKDIR /app

COPY go.mod.dist go.mod
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go get nginx-error-server
RUN go build -o nginx-error-server

RUN cp -R $GOPATH/pkg/mod/github.com/\!gerald\!wodni/kern*/default .
COPY content content
COPY codes.txt codes.txt


FROM ${REG_HOSTNAME}/${REG_FOLDER}/golang:1.18-alpine

WORKDIR /app

COPY --from=builder /app .

EXPOSE 5000

CMD [ "/app/nginx-error-server" ]
