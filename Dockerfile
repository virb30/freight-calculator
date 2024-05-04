FROM golang:1.22 AS builder

WORKDIR /app

CMD ["tail", "-f", "/dev/null"]