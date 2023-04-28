# build stage
FROM golang:1.19-alpine3.16 AS builder
#maintainer info
LABEL maintainer="Sethukumarj <Sethukumarj.76@gmail.com>"
#installing git
RUN apk update && apk add --no-cache git

WORKDIR /Events_Radar_Developement

COPY . .

RUN apk add --no-cache make

RUN make deps
RUN go mod vendor
RUN make build



# Run stage
FROM alpine:3.16

WORKDIR /Events_Radar_Developement
COPY go.mod .
COPY go.sum .
COPY views ./views
COPY --from=builder /Events_Radar_Developement/build/bin/api .


CMD [ "/Events_Radar_Developement/api"] 