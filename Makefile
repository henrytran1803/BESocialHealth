mysql:
	docker run --name mysql-socialhealth -e MYSQL_ROOT_PASSWORD=18032002 -p 3306:3306 -d mysql:8.0
dropmysql:
	docker stop mysql-socialhealth && docker rm mysql-socialhealth
createdb:
	docker exec -it mysql-socialhealth mysql -u root -p'18032002' -e "CREATE DATABASE socialhealth;"
dropdb:
	docker exec -it mysql-socialhealth mysql -u root -p'18032002' -e "DROP DATABASE socialhealth;"
migrateup:
	migrate -path db/migration -database "mysql://root:18032002@tcp(127.0.0.1:3306)/socialhealth?charset=utf8mb4&parseTime=True&loc=Local" -verbose up
migratedown:
	migrate -path db/migration -database "mysql://root:18032002@tcp(127.0.0.1:3306)/socialhealth?charset=utf8mb4&parseTime=True&loc=Local" -verbose down -all

.PHONY: mysql dropmysql  createdb dropdb migrateup migrateup
