from pymongo import MongoClient
from pysd.service_discovery import new_from_env


def get_mongo_client():
    mongo_url = new_from_env().mongo_endpoint()
    return MongoClient(mongo_url)


def drop_database(database_name):
    mongo_client = get_mongo_client()
    mongo_client.drop_database(database_name)
