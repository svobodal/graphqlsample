#!/bin/sh

curl -X POST http://localhost:7070/query -d '{"query": "{ video(id: 10){name} }"}'