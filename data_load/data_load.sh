#!/bin/bash
mongoimport --host 127.0.0.1:27017  -u "$MONGO_INITDB_ROOT_USERNAME" -p "$MONGO_INITDB_ROOT_PASSWORD" \
      --authenticationDatabase admin -d titanic -c passengers \
      --file /data_load/titanic-100.json  --jsonArray --drop