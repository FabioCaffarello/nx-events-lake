FROM docker.io/bitnami/spark:3.3
# FROM python:3.10-slim-buster

USER root

WORKDIR /app

ARG ENVIRONMENT
ENV ENVIRONMENT=$ENVIRONMENT

COPY dist/*.tar.gz .
RUN pip install *.tar.gz
RUN rm -rf *.tar.gz

COPY envs/.env.$ENVIRONMENT .
COPY spark_batch_bronze/ spark_batch_bronze/


ENTRYPOINT ["python", "spark_batch_bronze/main.py"]

