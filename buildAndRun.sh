#!/bin/bash

# Build the Docker image
docker image build -t forum .

# Run the Docker container
docker container run -p 8000:8000 -d forum

# chmod +x buildAndRun.sh
# $ ./buildAndRun.sh
