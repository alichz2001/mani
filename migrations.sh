#!/bin/bash

curl -X POST -u "$ARANGO_USERNAME:$ARANGO_PASSWORD" -H "accept: application/json" "$ARANGO_URI"/_api/database -d "{\"name\": \"$ARANGO_DATABASE\"}"

echo

curl -X POST -u "$ARANGO_USERNAME:$ARANGO_PASSWORD" -H 'accept: application/json' "$ARANGO_URI"/_db/"$ARANGO_DATABASE"/_api/collection -d '{"name": "files"}'

echo

curl -X POST -u "$ARANGO_USERNAME:$ARANGO_PASSWORD" -H 'accept: application/json' "$ARANGO_URI"/_db/"$ARANGO_DATABASE"/_api/collection -d '{"name": "users"}'

echo

curl -X POST -u "$ARANGO_USERNAME:$ARANGO_PASSWORD" -H 'accept: application/json' "$ARANGO_URI"/_db/"$ARANGO_DATABASE"/_api/index?collection=files \
-d '{
  "type": "persistent",
  "unique": true,
  "fields": [
    "name"
  ],
  "name":"unique_name_index"
}'

echo

curl -X POST -u "$ARANGO_USERNAME:$ARANGO_PASSWORD" -H 'accept: application/json' "$ARANGO_URI"/_db/"$ARANGO_DATABASE"/_api/index?collection=users \
-d '{
  "type": "persistent",
  "unique": true,
  "fields": [
    "username"
  ],
  "name":"unique_username_index"
}'


exit 0
