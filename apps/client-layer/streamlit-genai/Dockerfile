FROM langchain/langchain

WORKDIR /app

RUN apt-get update && apt-get install -y \
    build-essential \
    curl \
    software-properties-common \
    python3-dev \
    && rm -rf /var/lib/apt/lists/*

ARG ENVIRONMENT
ENV ENVIRONMENT=$ENVIRONMENT

COPY dist/*.tar.gz .
RUN pip install *.tar.gz
RUN rm -rf *.tar.gz

COPY envs/.env.$ENVIRONMENT .
COPY streamlit_genai/ streamlit_genai/

EXPOSE 8503

HEALTHCHECK CMD curl --fail http://localhost:8503/_stcore/health

ENTRYPOINT ["streamlit", "run", "streamlit_genai/main.py", "--server.port=8503", "--server.address=0.0.0.0"]
