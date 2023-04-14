import time
from elasticsearch import Elasticsearch
# python3 -m pip install elasticsearch


# Connect to Elasticsearch
es = Elasticsearch([{'host': 'localhost', 'port': 9200, 'scheme': 'http'}])

# Define the query
query = {
    "size": 1,
    "sort": [
        {
            "_script": {
                "script": {
                    "source": "def json = params._source.toString(); return json.length();",
                    "lang": "painless"
                },
                "type": "number", "order": "desc"
            }
        }
    ],
    "_source": ["id"]
}


# Time the query
start_time = time.time()
# Run the query and get the response
response = es.search(index='vendor', body=query)
end_time = time.time()
query_time = end_time - start_time

# Print the query time
print(f"The query took {query_time:.2f} seconds to run.")


# Get the size of the response in bytes
response_size = len(str(response).encode('utf-8'))

# Print the response size
print("Response size in bytes:", response_size)
