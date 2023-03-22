FROM golang:alpine AS builder
#maintainer info
LABEL maintainer="sethukumarj <sethukumarj.76@gmail.com>"
#installing git
RUN apk update && apk add --no-cache git

# Add docker-compose-wait tool -------------------




#current working directory
#COPY templates /.
WORKDIR /Events_Radar_Developement
#installing air
# RUN go get github.com/cosmtrek/air@latest

# # Copy go mod and sum files
# COPY go.mod .
# COPY go.sum .


COPY . .

RUN apk add --no-cache make



RUN make deps
RUN make build
RUN go mod vendor

CMD [ "make", "run"] 