#!/bin/sh
DATABASE_URL="${1:?"missing database url"}"
MIGRATION_NAME="${2:?"missing migration target"}"
migrate -path ./backend/src/migrations -database "$DATABASE_URL" -verbose up "$MIGRATION_NAME"
