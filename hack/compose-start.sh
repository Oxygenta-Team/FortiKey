#!/bin/bash
set -x

CONFIG_FILE="../cmd/cipher/config.yaml"

if ! command -v yq &> /dev/null; then
    echo    "Error: pls install yq for your system"
    echo    "sudo apt install yq #for Ubuntu/Debian"
    echo    "sudo dnf install yq #for Fedora/CentOS"
    exit 1
fi

export POSTGRES_USER=$(yq '.database.user' "$CONFIG_FILE")
export POSTGRES_PASSWORD=$(yq '.database.password' "$CONFIG_FILE")
export POSTGRES_DB=$(yq '.database.dbname' "$CONFIG_FILE")
export POSTGRES_HOST=$(yq '.database.host' "$CONFIG_FILE")
export POSTGRES_PORT=$(yq '.database.port' "$CONFIG_FILE")
export POSTGRES_SSL_MODE=$(yq '.database.sslmode' "$CONFIG_FILE")

docker compose up -d