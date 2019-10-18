FROM alpine:latest

RUN mkdir /app
WORKDIR /app
COPY bin /application

CMD ["/application/go-proto-micro", "serve"]