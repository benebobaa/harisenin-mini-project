#!/bin/sh

set -e

echo "run db migrations"

source /app/app.env

echo "start app"

exec "$@"