#!/bin/sh
set -e

host="$1"
port="$2"
shift 2
cmd="$@"

until mysqladmin ping -h "$host" -P "$port" --silent; do
    echo "Waiting for MySQL to be available at $host:$port..."
    sleep 1
done

echo "MySQL is available, executing command"
exec $cmd
