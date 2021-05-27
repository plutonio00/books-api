cp .env.example .env

cd docker
cp .env.example .env

docker-compose up -d
docker exec go-books make
docker exec go-books goose mysql "books_user:books_pass@tcp(mysql-books:3306)/books_db?parseTime=true" up
docker exec go-books server_run

