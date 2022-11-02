#!/usr/bin/env python3

import requests

data = {
    "filename": "test-h1.txt",
    "subject": "math",
    "subject": "science"
}

r = requests.post("http://localhost:8080/homework", json=data)
print(r.text)
print(r.status_code)