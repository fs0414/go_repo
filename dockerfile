FROM golang:1.20-bullseye

ENV ROOT=/api

WORKDIR ${ROOT}

COPY api/go.mod api/go.sum ./

RUN apt-get update

RUN apt-get install -y git curl tree

COPY . .

RUN go mod download
RUN go install github.com/cosmtrek/air@v1.29.0



CMD ["air", "-c", ".air.toml"]
