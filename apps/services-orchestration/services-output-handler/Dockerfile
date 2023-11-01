FROM scratch
WORKDIR /app
COPY dist/server .

ARG ENVIRONMENT
ENV ENVIRONMENT=$ENVIRONMENT

COPY dist/server .
COPY envs/.env.$ENVIRONMENT .

CMD ["./server"]
