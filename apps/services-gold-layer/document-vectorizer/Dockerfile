FROM langchain/langchain

WORKDIR /app

RUN apt-get update && apt-get install -y \
    build-essential \
    curl \
    software-properties-common \
    && rm -rf /var/lib/apt/lists/*

ARG ENVIRONMENT
ENV ENVIRONMENT=$ENVIRONMENT

COPY dist/*.tar.gz .
RUN pip install *.tar.gz
RUN rm -rf *.tar.gz

COPY envs/.env.$ENVIRONMENT .
COPY document_vectorizer/ document_vectorizer/

ENTRYPOINT ["python", "document_vectorizer/main.py"]