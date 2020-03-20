#!/usr/bin/env python3

import requests
import logging
import os
import json
logging.basicConfig(level=logging.DEBUG)

logging.debug(os.path.join(os.getenv("alfred_workflow_cache"), "entries.json"))
r = requests.get(os.environ['shortify_url'] + '/_entries')
response = r.json()
alfred_object = {
    "items": []
}
for item in response:
    alfred_object['items'].append({
        "uid": item['name'],
        "title": item['name'],
        "arg": item['name'],
    })


with open(os.path.join(os.getenv("alfred_workflow_cache"), "entries.json"), 'w+') as file:
    file.write(json.dumps(alfred_object))
