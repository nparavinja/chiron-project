# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:alpine AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /server

CMD ["/server"]


# FROM alpine

# EXPOSE 8000

# COPY --from=build /server /server

# CMD ["/server"]