cp .env.example .env
cd docker

cp .env.example .env
docker-compose build --no-cache
docker-compose up -d

docker exec go-books make
docker exec go-books goose -dir internal/migration mysql "books_user:books_pass@tcp(mysql-books:3306)/books_db?parseTime=true"  up
docker exec go-books charlatan load ./tests/fixtures -u=books_user -d=books_db -p=books_pass --host=mysql-books
docker exec go-books make exec