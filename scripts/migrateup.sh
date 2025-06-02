#!/bin/bash

if [ -f .env ]; then
    source .env
fi


echo "DATABASE_URL is set: $DATABASE_URL"
echo "Current directory: $(pwd)"
ls -la sql/schema



cd sql/schema
goose turso $DATABASE_URL up
