FROM python:3.10-slim-buster

USER root

WORKDIR /app

ARG ENVIRONMENT
ENV ENVIRONMENT=$ENVIRONMENT

RUN apt-get update && apt-get install -y ffmpeg

COPY dist/*.tar.gz .
RUN pip install *.tar.gz
RUN rm -rf *.tar.gz

COPY envs/.env.$ENVIRONMENT .
COPY media_transcoder/ media_transcoder/

ENTRYPOINT ["python", "media_transcoder/main.py"]

