FROM alpine:latest

WORKDIR /app

RUN apk update && apk add curl

ARG ENVIRONMENT
ENV ENVIRONMENT=$ENVIRONMENT

COPY dist/server .
COPY envs/.env.$ENVIRONMENT .

CMD ["./server"]
