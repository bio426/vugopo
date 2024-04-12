#!/bin/bash

declare -A envs

# Tern
envs["TERN_CONFIG"]="tern.conf"
envs["TERN_MIGRATIONS"]="migrations/"

# Auth
envs["AUTH_JWT_SECRET"]="my-secret"
envs["AUTH_COOKIE_NAME"]="vugopo-token"

# Postgres
envs["PG_HOST"]="localhost"
envs["PG_HOST"]="localhost"
envs["PG_PORT"]="5432"
envs["PG_USER"]="postgres"
envs["PG_PASSWORD"]="password"
envs["PG_DATABASE"]="vugopo"

for env in "${!envs[@]}"; do
  export ${env}=${envs[${env}]}
done

echo "${#envs[@]} variables setted"
