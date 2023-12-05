FROM golang:alpine AS buld-env

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64

WORKDIR /

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./.bin/app ./cmd/test1/main.go

# Run
FROM scratch

COPY --from=buld-env /.bin/app /.bin/app

ENTRYPOINT [".bin/app"]