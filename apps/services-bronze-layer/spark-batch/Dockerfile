FROM docker.io/bitnami/spark:3.3

ARG USER
ARG USER_ID
ARG GROUP_ID
ARG ENVIRONMENT
ENV ENVIRONMENT=$ENVIRONMENT

COPY dist/*.tar.gz .
RUN pip install *.tar.gz
RUN rm -rf *.tar.gz

COPY envs/.env.${ENVIRONMENT} /app/.env.${ENVIRONMENT}
COPY spark_batch_bronze/ /app/spark_batch_bronze/
COPY tests/ /app/tests/
COPY hack/ /app/hack/

RUN chmod +x /app/hack/check-integration.sh

RUN groupadd -f -g ${GROUP_ID} ${USER} && \
    useradd -m -g ${GROUP_ID} -u ${USER_ID} ${USER} || echo "User already exists"

USER ${USER_ID}:${GROUP_ID}

WORKDIR /app


ENTRYPOINT ["python", "/app/spark_batch_bronze/main.py"]


FROM devimage as prodimage

COPY --from=devimage /app/.env.${ENVIRONMENT} /app/.env.${ENVIRONMENT}
COPY --from=devimage /app/spark_batch_bronze/ /app/spark_batch_bronze/

ENTRYPOINT ["python", "/app/spark_batch_bronze/main.py"]
