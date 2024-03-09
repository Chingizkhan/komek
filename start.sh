#!/bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/migrations/ -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"

# https://www.youtube.com/watch?v=jf6sQsz0M1M&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=26