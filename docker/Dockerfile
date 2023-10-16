FROM golang:1.21 as builder
WORKDIR /app

COPY . .

RUN go mod download
CMD make seeder run MODE=production
