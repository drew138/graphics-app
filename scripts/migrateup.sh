#!/bin/sh
DATABASE_URL="$1"
migrate -path ./backend/src/migrations -database "$DATABASE_URL" -verbose up
