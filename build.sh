#!/bin/bash

# building docker images.
docker build -t code_storage/rest-api ./rest-api
docker build -t code_storage/svelte ./svelte

# creating containers.
docker create --name rest-api -p 8080:8080 code_storage/rest-api
docker create --name svelte -p 5000:5000 code_storage/svelte