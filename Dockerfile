FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o pulse .

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder  . /pulse
CMD sh -c "./pulse"