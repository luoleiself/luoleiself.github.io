from langchain_openai.embeddings import OpenAIEmbeddings

"""
langchain_community 社区包不再活跃维护, 推荐使用独立集成包 https://docs.langchain.com/oss/python/integrations/embeddings
"""

embeddings = OpenAIEmbeddings()

print(embeddings.embed_query("你好"))
print(embeddings.embed_documents(["你好", "世界"]))
