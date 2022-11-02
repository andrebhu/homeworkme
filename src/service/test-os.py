#!/usr/bin/env python3

import os
import json

data = {
    "files": []
}

for dirpath, dirnames, filenames in os.walk("files"):
    if filenames:
        for f in filenames:
            # print(os.path.join(dirpath, f))

            obj = {
                "subject": dirpath.split("/")[1],
                "filename": f
            }
            data["files"].append(obj)
            
print(json.dumps(data))
