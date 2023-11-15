
import os
from pydotenv.loader import DotEnvLoader
import streamlit as st
from langchain.chains import RetrievalQA
from langchain.callbacks.base import BaseCallbackHandler
from langchain.vectorstores.neo4j_vector import Neo4jVector
from streamlit.logger import get_logger
from chains import (
    load_embedding_model,
    load_llm,
)

ENVIRONMENT = os.getenv("ENVIRONMENT")
envs = DotEnvLoader(environment=ENVIRONMENT)
url = envs.get_variable("NEO4J_URI")
# Remapping for Langchain Neo4j integration
os.environ["NEO4J_URL"] = url

username = envs.get_variable("NEO4J_USERNAME")
password = envs.get_variable("NEO4J_PASSWORD")
ollama_base_url = envs.get_variable("OLLAMA_BASE_URL")
embedding_model_name = envs.get_variable("EMBEDDING_MODEL")
llm_name = envs.get_variable("LLM")

logger = get_logger(__name__)

embeddings, dimension = load_embedding_model(
    embedding_model_name, config={"ollama_base_url": ollama_base_url}, logger=logger
)

llm = load_llm(llm_name, logger=logger, config={"ollama_base_url": ollama_base_url})

class StreamHandler(BaseCallbackHandler):
    def __init__(self, container, initial_text=""):
        self.container = container
        self.text = initial_text

    def on_llm_new_token(self, token: str, **kwargs) -> None:
        self.text += token
        self.container.markdown(self.text)


def main():
    st.header("📄Chat with your pdf file")

    store = Neo4jVector.from_existing_index(
        embeddings,
        url=url,
        username=username,
        password=password,
        index_name="pdf_enbeddings"
    )

    qa = RetrievalQA.from_chain_type(
        llm=llm, chain_type="stuff", retriever=store.as_retriever()
    )

    # Accept user questions/query
    query = st.text_input("Ask questions about your PDF file")

    if query:
        stream_handler = StreamHandler(st.empty())
        qa.run(query, callbacks=[stream_handler])


if __name__ == "__main__":
    main()
