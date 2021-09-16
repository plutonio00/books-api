cp .env.example .env
cd docker

cp .env.example .env
docker-compose build --no-cache
docker-compose up -d

docker exec go-books make
docker exec go-books make db_init
docker exec go-books make exec