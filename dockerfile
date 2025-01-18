FROM golang:1.24rc1-bullseye

ENV ROOT=/api

WORKDIR ${ROOT}

COPY /go.mod go.sum ./

RUN apt-get update

RUN apt-get install -y git curl tree

COPY . .

RUN go mod download
RUN go install github.com/air-verse/air@latest



CMD ["air", "-c", ".air.toml"]
