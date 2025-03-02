#!/bin/bash
set -x

# Get the directory where the script is located
SCRIPT_DIR=$(dirname "$(realpath "$0")")

CONFIG_FILE="$SCRIPT_DIR/../cmd/cipher/config.yaml"
DOCKER_COMPOSE_PATH="$SCRIPT_DIR/../docker-compose.yaml"
MAKEFILE_PATH="$SCRIPT_DIR/../Makefile"

# Check if yq is installed
if ! command -v yq &> /dev/null; then
    echo "Error: Please install yq for your system"
    exit 1
fi

# Check if sed is installed
if ! command -v sed &> /dev/null; then
    echo "Error: Please install sed for your system"
    exit 1
fi

# Set null values for data in docker-compose.yaml
echo "Setting null values in $DOCKER_COMPOSE_PATH..."
sed -i "s|POSTGRES_USER: .*|POSTGRES_USER: null|" "$DOCKER_COMPOSE_PATH"
sed -i "s|POSTGRES_PASSWORD: .*|POSTGRES_PASSWORD: null|" "$DOCKER_COMPOSE_PATH"
sed -i "s|POSTGRES_DB: .*|POSTGRES_DB: null|" "$DOCKER_COMPOSE_PATH"
sed -i "s|POSTGRES_HOST: .*|POSTGRES_HOST: null|" "$DOCKER_COMPOSE_PATH"
sed -i "s|POSTGRES_PORT: .*|POSTGRES_PORT: null|" "$DOCKER_COMPOSE_PATH"
sed -i "s|POSTGRES_SSL_MODE: .*|POSTGRES_SSL_MODE: null|" "$DOCKER_COMPOSE_PATH"

# Clear login data in Makefile
echo "Clearing login data in $MAKEFILE_PATH..."
sed -i "s|DB_USER ?=.*|DB_USER ?= null|" "$MAKEFILE_PATH"
sed -i "s|DB_PASSWORD ?=.*|DB_PASSWORD ?= null|" "$MAKEFILE_PATH"
sed -i "s|DB_HOST ?=.*|DB_HOST ?= null|" "$MAKEFILE_PATH"
sed -i "s|DB_PORT ?=.*|DB_PORT ?= null|" "$MAKEFILE_PATH"
sed -i "s|DB_NAME ?=.*|DB_NAME ?= null|" "$MAKEFILE_PATH"
sed -i "s|SSL_MODE ?=.*|SSL_MODE ?= null|" "$MAKEFILE_PATH"

echo "Login data cleared in configuration files"
