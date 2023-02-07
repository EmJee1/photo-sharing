# Dockerfile based on https://docs.docker.com/language/golang/build-images/

FROM golang:1.19-alpine

WORKDIR /ps
#ENV GOPATH /photo-sharing

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go .
COPY *.html .
COPY *.css .
COPY *.js .
COPY .env ./.env

RUN go build -o /dist

EXPOSE 8080

CMD [ "/dist" ]
