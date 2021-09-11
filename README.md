# titanic-go-gin
## An example Titanic passenger REST API using Go and Gin 

Dataset download link: https://www.kaggle.com/c/titanic/data?select=train.csv

1. In the root of the project, create a .env file for docker compose and add environment variables: DB_USER and DB_PASSWORD (eg. DB_USER=mongo  DB_PASSWORD=mongo1234)
2. Start the MongoDB container: docker compose up -d
3. Get shell to MongoDB docker container: docker exec -it mongodb bash
4. From docker shell: chmod +x /data_load/data_load.sh 
5. From docker shell: data_load/data_load.sh  (should import 891 documents)
6. The service application depends on the following environment variables:
   - DB_USER: ie root
   - DB_PASSWORD: ie mongo1243
   - DB_ENDPOINT: default - localhost
   - DB_PORT: default - 27017
   - PORT: default - 8080
   
To access Mongo Shell :
1. mongosh --host 127.0.0.1:27017 -u "$MONGO_INITDB_ROOT_USERNAME" -p "$MONGO_INITDB_ROOT_PASSWORD" --authenticationDatabase admin
2. use titanic
3. db.passengers.find