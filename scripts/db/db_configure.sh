# Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

# bin/bash
set -o errexit
set -o nounset
set -o pipefail

echo "* Generating user name and password"
DB='omxdb'
DB_HOST='timescaledb' # name of the docker container
DB_PORT='5432'
USER='omxuser'
PASSWORD='2ee2e41ec0a7e6441e0038'

echo "* Configure database with generated user and password"

# Delete everything, if exists, to start anew from scratch.
command docker exec -it "$DB_HOST" psql -U postgres -h localhost -c "DROP DATABASE IF EXISTS $DB;"
command docker exec -it "$DB_HOST" psql -U postgres -h localhost -c "DROP USER IF EXISTS $USER;"

# Create generated user with generated password.
command docker exec -it "$DB_HOST" psql -U postgres -h localhost -c "CREATE USER $USER with encrypted password '$PASSWORD';"

# Create new DB linked to the new user created above.
command docker exec -it "$DB_HOST"  psql -U postgres -h localhost -c "CREATE DATABASE $DB WITH OWNER $USER TABLESPACE pg_default;"



echo "* Database configured. Connection details:"
echo "  -- DB Name: $DB"
echo "  -- DB User: $USER"
echo "  -- DB Host: $DB_HOST"
echo "  -- DB Port: $DB_PORT"
