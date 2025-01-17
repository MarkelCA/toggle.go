# syntax=docker/dockerfile:1
FROM golang:1.21 AS builder

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

ARG APP

RUN CGO_ENABLED=0 go build -v -a -installsuffix cgo  -o app cmd/${APP}/*.go

FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /usr/src/app/testdata ./testdata
COPY --from=builder /usr/src/app/app ./
CMD ["./app"]
