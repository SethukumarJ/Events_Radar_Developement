FROM golang:alpine AS builder
#maintainer info
LABEL maintainer="sethukumarj <sethukumarj.76@gmail.com>"
#installing git
RUN apk update && apk add --no-cache git

WORKDIR /Events_Radar_Developement
# installing air
# RUN go get github.com/cosmtrek/air@latest

COPY . .

RUN apk add --no-cache make

RUN make deps
RUN make build
RUN go mod vendor

CMD [ "make", "run"] 