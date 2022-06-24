#!/bin/sh
MIGRATION_NAME="${1:?"missing migration name"}"
migrate create -ext sql -dir ./backend/src/migrations -seq "$MIGRATION_NAME"
