#!/bin/bash
set -e

BASE_NAME="todo-list-backend"
TAG="dev"

IMAGE_NAME="$BASE_NAME:$TAG"

echo ">>> IMAGE_NAME: $IMAGE_NAME"

echo ">>> Building docker image "
docker build -t $IMAGE_NAME .
