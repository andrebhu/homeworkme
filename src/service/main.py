#!/usr/bin/env python3

import os
from fastapi import Body, FastAPI
from pydantic import BaseModel

app = FastAPI()

class File(BaseModel):
    filename: str
    subject: str

@app.post("/retrieve")
async def retrieve(file: File):
    file_dict = file.dict()
    path = os.path.join(file_dict["subject"], file_dict["filename"])
    return open(path, "r").read()