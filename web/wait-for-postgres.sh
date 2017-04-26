#!/bin/bash
# wait-for-postgres.sh

set -e

shift
cmd="$@"

PYTHON=$(cat <<-END
import psycopg2
import time

def wait():

    try:
        conn = psycopg2.connect("dbname='$DB_NAME' "+
                                "user='$DB_USER' "+
                                "host='$DB_SERVICE' "+
                                "port=$DB_PORT "+
                                "password='$DB_PASS'")
    except:
        print("wait")
        time.sleep(1)
        wait()

wait()
END
)
python -c "$PYTHON"

>&2 echo "Postgres is up - executing command"
exec $cmd
