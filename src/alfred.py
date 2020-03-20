#!/usr/bin/env python3

import json
import os

with open(os.path.join(os.getenv("alfred_workflow_cache"), "entries.json"), 'r') as file:
    print(file.read())
