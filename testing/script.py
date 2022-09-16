#!/usr/bin/env python3

import json

d = "{\"test\": 1, \"test\": 2}"

j = json.loads(d)
print(j["test"])

