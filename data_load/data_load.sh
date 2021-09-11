#!/bin/bash
mongoimport --host 127.0.0.1:27017  -u "$MONGO_INITDB_ROOT_USERNAME" -p "$MONGO_INITDB_ROOT_PASSWORD" \
      --authenticationDatabase admin -d titanic -c passengers \
      --type csv --columnsHaveTypes --fields "PassengerId.int32(),Survived.boolean(),Pclass.int32(),Name.string(),Sex.string(),Age.string(),SibSp.int32(),Parch.int32(),Ticket.string(),Fare.double(),Cabin.string(),Embarked.string()" \
      --file /data_load/train.csv  --drop