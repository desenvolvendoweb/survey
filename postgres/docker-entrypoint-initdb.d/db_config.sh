#!/bin/env bash

echo "Create USER postgres executing ..."
psql -U postgres -c "CREATE USER $DB_USER PASSWORD '$DB_PASS'"
psql -U postgres -c "CREATE DATABASE $DB_NAME OWNER $DB_USER"
echo "Create USER postgres finished !"
