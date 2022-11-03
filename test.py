#!/usr/bin/env python3

import re
import requests

# internal tests
r = requests.get("http://localhost:8000/directory")
assert r.status_code == 200
print("OK service /directory")

data = {
    "filename": "test-hw1.txt",
    "subject": "math"
}
r = requests.post("http://localhost:8000/retrieve", json=data)
assert r.status_code == 200
print("OK service /retrieve")

# app tests
r = requests.get("http://localhost:1234")
assert r.status_code == 200
print("OK app /")


'''
curl -X POST -d '{"filename": "test-hw1.txt", "filename": "../../main.py", "subject": "math"}' http:/
/localhost:1234/homework
'''