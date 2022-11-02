#!/usr/bin/env python3

import os
import json

from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI()
path = "files"


class File(BaseModel):
    filename: str
    subject: str


@app.post("/retrieve")
async def retrieve(file: File):
    file_dict = file.dict()
    file_path = os.path.join(path, file_dict["subject"], file_dict["filename"])
    return open(file_path, "r").read()


@app.get("/directory")
async def directory():
    data = {
        "files": []
    }
    for dirpath, dirnames, filenames in os.walk(path):
        if filenames:
            for f in filenames:
                obj = {
                    "subject": dirpath.split("/")[1],
                    "filename": f
                }
                data["files"].append(obj)

    return data["files"]
