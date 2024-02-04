#!/bin/sh

set -e


source /app/app.env

echo "start app"

exec "$@"