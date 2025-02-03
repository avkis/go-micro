# base go image
# FROM golang:1.23-alpine AS builder

# RUN mkdir /app

# COPY . /app

# WORKDIR /app

# RUN CGO_ENABLE=0 go build -o brokerApp ./cmd/api

# RUN chmod +x /app/brokerApp

# RUN apt install net-tools

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

# COPY --from=builder /app/brokerApp /app
COPY  brokerApp /app

CMD [ "/app/brokerApp" ]