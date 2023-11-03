FROM python:3.10-slim-buster

USER root

WORKDIR /app

ARG ENVIRONMENT
ENV ENVIRONMENT=$ENVIRONMENT

COPY dist/*.tar.gz .
RUN pip install *.tar.gz
RUN rm -rf *.tar.gz

COPY envs/.env.$ENVIRONMENT .
COPY source_watcher/ source_watcher/

ENTRYPOINT ["python", "source_watcher/main.py"]

