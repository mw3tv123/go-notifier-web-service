#syntax=docker/dockerfile:experimental
FROM golang:alpine AS build-env
WORKDIR /app/
COPY . /app/
ENV GO111MODULE=on
RUN CGO_ENABLED=0 go build -o /main ./main.go

FROM alpine:3.13
WORKDIR /app
RUN apk add --no-cache && \
    apk add tzdata
ENV TZ=Asia/Ho_Chi_Minh
COPY --from=build-env /main /app/
COPY config /app/config
ENTRYPOINT ["./main"]
