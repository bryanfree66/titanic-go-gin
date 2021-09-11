# titanic-go-gin
## An example Titanic passenger REST API using go and Gin 

Dataset download link: https://www.kaggle.com/c/titanic/data?select=train.csv

1. In the root of the project, create a .env file and add environment variables: DB_USER and DB_PASSWORD (eg. DB_USER=mongo  DB_PASSWORD=mongo1234)
2. Start the MongoDB container: docker compose up -d
3. Get shell to MongoDB docker container: docker exec -it mongodb bash
4. From docker shell: chmod +x /data_load/data_load.sh 
5. From docker shell: data_load/data_load.sh  (should import 891 documents)
