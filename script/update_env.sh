#! /bin/bash

set -e

root_path=$(cd "$(dirname "$0")/.." && pwd -P)
script_path=$(cd "$(dirname "$0")" && pwd -P)

# Nx Dev Environment env vars stay in `.env`

if [ "$ENV" == "development" ]; then
    if [ -z "$NGROK_URL" ]; then
        echo "NGROK_URL is not set"
        exit 1
    fi

    # App environment
    echo "Setting up App environment..."
    cat ${root_path}/env/.env.dev ${root_path}/env/app/.env.dev > ${root_path}/app/.env
    echo "NGROK_API_URL=$NGROK_URL" >> ${root_path}/app/.env

    # API environment 
    echo "Setting up API environment..."
    cat ${root_path}/env/.env.dev ${root_path}/env/api/.env.dev > ${root_path}/.env.dev
fi

if [ "$ENV" == "production" ]; then
    # App environment
    echo "Setting up App environment..."
    cat ${root_path}/env/.env.prod ${root_path}/env/app/.env.prod > ${root_path}/app/.env

    # API environment
    echo "Setting up API environment..."
    cat ${root_path}/env/.env.prod ${root_path}/env/api/.env.prod > ${root_path}/.env.prod
fi
