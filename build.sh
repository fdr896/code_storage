#!/bin/bash

# building docker images.
docker build -t code_storage/rest-api ./rest-api
docker build -t code_storage/svelte ./svelte