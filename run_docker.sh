#!/bin/bash

# build docker images.
docker build -t svelte ./svelte
docker build -t rest-api ./rest-api

# run both of images simultaniously.
docker run -p 5000:5000 svelte & docker run -p 8080:8080 rest-api
