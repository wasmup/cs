import time
import json
from elasticsearch import Elasticsearch
# python3 -m pip install elasticsearch>=7.0.0


# Connect to Elasticsearch
es = Elasticsearch([{'host': 'localhost', 'port': 9200, 'scheme': 'http'}])


# Define your Elasticsearch query
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

response = es.search(index="vendor", body=query, headers={'Content-Type': 'application/json'})
end_time = time.time()
query_time = end_time - start_time

# Get the response size in bytes
response_size = json.dumps(response).encode('utf-8')

# Print the query time, response size, and the time it took to get the response
print(f"The query took {query_time:.2f} seconds to run.")
print(f"The response size is {len(response_size)} bytes.")
print(f"The response took {response['took']} milliseconds to generate.")
