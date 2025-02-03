# base go image
# FROM golang:1.23-alpine AS builder

# RUN mkdir /app

# COPY . /app

# WORKDIR /app

# RUN CGO_ENABLE=0 go build -o authApp ./cmd/api

# RUN chmod +x /app/authApp

# RUN apt install net-tools

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

# COPY --from=builder /app/authApp /app
COPY authApp /app

CMD [ "/app/authApp" ]