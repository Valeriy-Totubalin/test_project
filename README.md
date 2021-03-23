
Для запуска проекта необходимо:
1. создать файл .env по примеру example.env
2. прописать dsn для подключения к базе данных db/dbconf.yml
3. создать пустую базу данных с указанным выше названием
4. выполнить команду goose up (запуск миграций)
5. go run cmd/app/main.go

Swagger http://localhost:8080/swagger/index.html#/items/deleteItem