FROM python:3.10-slim-buster

USER root

WORKDIR /app

ARG ENVIRONMENT
ENV ENVIRONMENT=$ENVIRONMENT

COPY dist/*.tar.gz .
RUN pip install *.tar.gz
RUN rm -rf *.tar.gz

COPY envs/.env.$ENVIRONMENT .
COPY file_downloader/ file_downloader/

ENTRYPOINT ["python", "file_downloader/main.py"]

