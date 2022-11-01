#!/usr/bin/env python3

import requests

data = {
    'filename': 'flag.txt',
    'subject': 'math'
}

r = requests.post("http://localhost:8000/retrieve", json=data)
print(r.text)