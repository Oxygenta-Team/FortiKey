#!/bin/bash
set -x

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

# Extract values from config.yaml using yq
export POSTGRES_USER=$(yq '.database.user' "$CONFIG_FILE")
export POSTGRES_PASSWORD=$(yq '.database.password' "$CONFIG_FILE")
export POSTGRES_DB=$(yq '.database.dbname' "$CONFIG_FILE")
export POSTGRES_HOST=$(yq '.database.host' "$CONFIG_FILE")
export POSTGRES_PORT=$(yq '.database.port' "$CONFIG_FILE")
export POSTGRES_SSL_MODE=$(yq '.database.sslmode' "$CONFIG_FILE")

# Print the values for debugging
#echo "POSTGRES_USER: $POSTGRES_USER"
#echo "POSTGRES_PASSWORD: $POSTGRES_PASSWORD"
#echo "POSTGRES_DB: $POSTGRES_DB"
#echo "POSTGRES_HOST: $POSTGRES_HOST"
#echo "POSTGRES_PORT: $POSTGRES_PORT"
#echo "POSTGRES_SSL_MODE: $POSTGRES_SSL_MODE"

# Update docker-compose.yml by replacing 'null' with actual values
sed -i "s|POSTGRES_USER: null|POSTGRES_USER: $POSTGRES_USER|" "$DOCKER_COMPOSE_PATH"
sed -i "s|POSTGRES_PASSWORD: null|POSTGRES_PASSWORD: $POSTGRES_PASSWORD|" "$DOCKER_COMPOSE_PATH"
sed -i "s|POSTGRES_DB: null|POSTGRES_DB: $POSTGRES_DB|" "$DOCKER_COMPOSE_PATH"
sed -i "s|POSTGRES_HOST: null|POSTGRES_HOST: $POSTGRES_HOST|" "$DOCKER_COMPOSE_PATH"
sed -i "s|POSTGRES_PORT: null|POSTGRES_PORT: $POSTGRES_PORT|" "$DOCKER_COMPOSE_PATH"
sed -i "s|POSTGRES_SSL_MODE: null|POSTGRES_SSL_MODE: $POSTGRES_SSL_MODE|" "$DOCKER_COMPOSE_PATH"

# Update Makefile by replacing 'null' with actual values
sed -i "s|DB_USER ?= null|DB_USER ?= $POSTGRES_USER|" "$MAKEFILE_PATH"
sed -i "s|DB_PASSWORD ?= null|DB_PASSWORD ?= $POSTGRES_PASSWORD|" "$MAKEFILE_PATH"
sed -i "s|DB_HOST ?= null|DB_HOST ?= $POSTGRES_HOST|" "$MAKEFILE_PATH"
sed -i "s|DB_PORT ?= null|DB_PORT ?= $POSTGRES_PORT|" "$MAKEFILE_PATH"
sed -i "s|DB_NAME ?= null|DB_NAME ?= $POSTGRES_DB|" "$MAKEFILE_PATH"
sed -i "s|SSL_MODE ?= null|SSL_MODE ?= $POSTGRES_SSL_MODE|" "$MAKEFILE_PATH"

echo "POSTGRES_USER: $POSTGRES_USER"
echo "POSTGRES_PASSWORD: $POSTGRES_PASSWORD"
echo "POSTGRES_HOST: $POSTGRES_HOST"
echo "POSTGRES_PORT: $POSTGRES_PORT"
echo "POSTGRES_DB: $POSTGRES_DB"
echo "POSTGRES_SSL_MODE: $POSTGRES_SSL_MODE"

# Run the containers using docker-compose
docker compose -f "$DOCKER_COMPOSE_PATH" up -d

# if you want run with custom paths
# just specify when you run the script
# EXAMPLE
# ./compose-start.sh /path/to/your/docker-compose.yml /path/to/your/Makefile
