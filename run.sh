#!/bin/bash

# running images simultaniously.
docker run -p 8080:8080 code_storage/rest-api & docker run -p 5000:5000 code_storage/svelte