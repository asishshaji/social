docker stop user_db_postgresql
docker rm user_db_postgresql
docker run -p 5432:5432 --name user_db_postgresql -e POSTGRES_PASSWORD=pass -e POSTGRES_USER=postgres -e POSTGRES_DB=tdevs -d postgres
