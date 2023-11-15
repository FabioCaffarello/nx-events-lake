# streamlit-genai

## Chat with your Knowledge Graph

This application provides a simple interface for users to interact with a `knowledge graph` structured in Neo4J.
The underlying technology leverages language models and vector embeddings to retrieve relevant information from semantic queries in the ` Neo4J knowledge graph`.

## How to Run the Service

To run your service, follow these steps:

1. **Setup Configuration**:

    - Create a `.env.{ENVIRONMENT}` file based on the example provided in the `envs` folder.

2. **Build The Service**:

```sh
npx nx image client-layer-streamlit-genai --env=<ENVIRONMENT>
```

3. **Run the Application**:

```sh
docker-compose up -d
```

This will launch the application, and you can access it through your web browser.

- Endpoint: http://localhost:8503

### Application Structure

- `main.py`: Contains the Streamlit application code, user interface, and integration with language models and vector stores.
- `chains.py`: Defines functions for loading embedding models and language models.
- `streamlit_genai/utils.py`: Contains a utility class for logging information.


### Embedding Models

The application supports two types of embedding models:

- **Ollama:** Leveraging the OllamaEmbeddings class with the "llama2" model.
- **Sentence Transformer:** Using the SentenceTransformerEmbeddings class with the "all-MiniLM-L6-v2" model.

You can choose the desired embedding model by setting the `EMBEDDING_MODEL` environment variable.

### Language Model

The application uses the ChatOllama language model, configurable by the `LLM` environment variable. Adjust parameters such as temperature, top_k, top_p, and num_ctx in the `load_llm` function.

### Logging

The application uses a simple logging mechanism provided by the `BaseLogger` class in `streamlit_genai/utils.py`. You can extend or replace this logger based on your requirements.

Feel free to explore and customize the code to fit your specific use case. Enjoy chatting with your PDF file!