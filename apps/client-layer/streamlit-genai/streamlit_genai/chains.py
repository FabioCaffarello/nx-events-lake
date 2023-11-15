from streamlit_genai.utils import BaseLogger
from langchain.chat_models import ChatOllama
from langchain.embeddings import (OllamaEmbeddings,
                                  SentenceTransformerEmbeddings)


def load_embedding_model(embedding_model_name: str, logger=BaseLogger(), config={}):
    """
    Load the embedding model based on the specified name.

    Args:
        embedding_model_name (str): The name of the embedding model.
        logger (BaseLogger, optional): The logger instance. Defaults to BaseLogger().
        config (dict, optional): Additional configuration for loading the model. Defaults to {}.

    Returns:
        Tuple: A tuple containing embeddings and dimension.
    """
    if embedding_model_name == "ollama":
        embeddings = OllamaEmbeddings(
            base_url=config["ollama_base_url"], model="llama2"
        )
        dimension = 4096
        logger.info("Embedding: Using Ollama")
    else:
        embeddings = SentenceTransformerEmbeddings(
            model_name="all-MiniLM-L6-v2", cache_folder="/embedding_model"
        )
        dimension = 384
        logger.info("Embedding: Using SentenceTransformer")
    return embeddings, dimension


def load_llm(llm_name: str, logger=BaseLogger(), config={}):
    logger.info(f"LLM: Using Ollama: {llm_name}")
    return ChatOllama(
        temperature=0,
        base_url=config["ollama_base_url"],
        model=llm_name,
        streaming=True,
        # seed=2,
        top_k=10,  # A higher value (100) will give more diverse answers, while a lower value (10) will be more conservative.
        top_p=0.3,  # Higher value (0.95) will lead to more diverse text, while a lower value (0.5) will generate more focused text.
        num_ctx=3072,  # Sets the size of the context window used to generate the next token.
    )
